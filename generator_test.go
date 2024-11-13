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
