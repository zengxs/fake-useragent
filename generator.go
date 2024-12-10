package fakeuseragent

import (
	_ "embed"
	"errors"
	"math/rand"
	"time"
)

//go:embed database.yml
var defaultUserAgentDBText string

type UserAgentGenerator struct {
	database   []UserAgent
	randSource rand.Source
}

func NewUserAgentGenerator(opts ...OptionSetter) (*UserAgentGenerator, error) {
	options, err := NewOptions(opts...)
	if err != nil {
		return nil, errors.Join(ErrNewOptions, err)
	}

	// filter user agents by tags
	var filteredUserAgents []UserAgent
	if len(options.TagFilters) > 0 {
		for _, ua := range options.Database {
			if options.FilterMode == FilterModeInclude {
				for _, tag := range options.TagFilters {
					if ua.HasTag(tag) {
						filteredUserAgents = append(filteredUserAgents, ua)
						break
					}
				}
			} else {
				// FilterModeExclude
				hasTag := false
				for _, tag := range options.TagFilters {
					if ua.HasTag(tag) {
						hasTag = true
						break
					}
				}
				if !hasTag {
					filteredUserAgents = append(filteredUserAgents, ua)
				}
			}
		}
	} else {
		filteredUserAgents = options.Database
	}

	gen := &UserAgentGenerator{
		database:   filteredUserAgents,
		randSource: rand.NewSource(time.Now().UnixNano()),
	}

	return gen, nil
}

func (gen *UserAgentGenerator) Random() UserAgent {
	randIndex := rand.New(gen.randSource).Intn(len(gen.database))
	return gen.database[randIndex]
}
