package fakeuseragent

import "github.com/goccy/go-yaml"

type FilterMode int

const (
	FilterModeInclude FilterMode = iota
	FilterModeExclude
)

type Options struct {
	// Database is full list of user agents
	Database []UserAgent

	// TagFilters is a list of tags to filter user agents
	TagFilters []UserAgentTag

	// FilterMode is a mode to filter user agents
	FilterMode FilterMode
}

type OptionSetter func(*Options)

// NewOptions creates a new options with default values and user options
func NewOptions(opts ...OptionSetter) (*Options, error) {
	// get default user agents
	var allUserAgents []UserAgent
	err := yaml.Unmarshal([]byte(defaultUserAgentDBText), &allUserAgents)
	if err != nil {
		return nil, err
	}

	// set default options and apply user options
	options := &Options{
		Database:   allUserAgents,
		TagFilters: nil, // no filter
		FilterMode: FilterModeInclude,
	}
	for _, opt := range opts {
		opt(options)
	}

	if len(options.Database) == 0 {
		return nil, ErrDatabaseIsEmpty
	}

	return options, nil
}

// WithDatabase sets your own database of user agents to randomize
func WithDatabase(database []UserAgent) OptionSetter {
	return func(o *Options) {
		o.Database = database
	}
}

// WithTagFilters sets tags to filter user agents
func WithTagFilters(tags ...UserAgentTag) OptionSetter {
	return func(o *Options) {
		o.TagFilters = tags
	}
}

// WithFilterMode sets a mode to filter user agents
func WithFilterMode(mode FilterMode) OptionSetter {
	return func(o *Options) {
		o.FilterMode = mode
	}
}
