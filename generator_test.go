package fakeuseragent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fakeuseragent "github.com/zengxs/fake-useragent"
)

func TestGenerate(t *testing.T) {
	gen, err := fakeuseragent.NewUserAgentGenerator()
	assert.NoError(t, err)

	ua := gen.Random()
	assert.NotEmpty(t, ua)
}

func TestGenerateFilteredInclude(t *testing.T) {
	gen, err := fakeuseragent.NewUserAgentGenerator(
		fakeuseragent.WithTagFilters(fakeuseragent.TagChrome),
		fakeuseragent.WithFilterMode(fakeuseragent.FilterModeInclude),
	)
	assert.NoError(t, err)

	// repeat 100 times
	for i := 0; i < 100; i++ {
		ua := gen.Random()
		assert.NotEmpty(t, ua)
		assert.True(t, ua.HasTag(fakeuseragent.TagChrome))
	}
}

func TestGenerateFilteredExclude(t *testing.T) {
	gen, err := fakeuseragent.NewUserAgentGenerator(
		fakeuseragent.WithTagFilters(fakeuseragent.TagChrome),
		fakeuseragent.WithFilterMode(fakeuseragent.FilterModeExclude),
	)
	assert.NoError(t, err)

	// repeat 100 times
	for i := 0; i < 100; i++ {
		ua := gen.Random()
		assert.NotEmpty(t, ua)
		assert.False(t, ua.HasTag(fakeuseragent.TagChrome))
	}
}
