package fakeuseragent

import (
	_ "embed"
	"math/rand"
	"time"

	"github.com/goccy/go-yaml"
)

//go:embed database.yml
var defaultUserAgentDBText string

type UserAgentGenerator struct {
	database   []UserAgent
	randSource rand.Source
}

type Option struct {
	TagFilters []UserAgentTag
}

var DefaultOptions = []Option{
	{
		TagFilters: []UserAgentTag{
			TagDesktop,
			TagWindows,
		},
	},
}

func NewUserAgentGenerator(opts ...Option) (*UserAgentGenerator, error) {
	if len(opts) == 0 {
		opts = DefaultOptions
	}

	var allUserAgents []UserAgent
	err := yaml.Unmarshal([]byte(defaultUserAgentDBText), &allUserAgents)
	if err != nil {
		return nil, err
	}

	// filter user agents by tags
	for _, opt := range opts {
		if len(opt.TagFilters) == 0 {
			continue
		}

		var filteredUserAgents []UserAgent
		for _, ua := range allUserAgents {
			for _, tag := range opt.TagFilters {
				if ua.HasTag(tag) {
					filteredUserAgents = append(filteredUserAgents, ua)
					break
				}
			}
		}

		allUserAgents = filteredUserAgents
	}

	gen := &UserAgentGenerator{
		database:   allUserAgents,
		randSource: rand.NewSource(time.Now().UnixNano()),
	}

	return gen, nil
}

func (gen *UserAgentGenerator) Random() UserAgent {
	randIndex := rand.New(gen.randSource).Intn(len(gen.database))
	return gen.database[randIndex]
}
