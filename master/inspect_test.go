package master

import (
	"path"
	"sync"
	"testing"

	"github.com/countstarlight/homo/sdk/homo-go"
	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	info := infoStats{
		Inspect:  homo.Inspect{},
		services: nil,
		file:     path.Join("testdata", "application_docker.stats"),
		RWMutex:  sync.RWMutex{},
	}
	sss := map[string]map[string]interface{}{}
	b := info.LoadStats(&sss)
	assert.Equal(t, true, b)
}
