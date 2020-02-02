package engine

import (
	"testing"

	"github.com/aiicy/aiicy/sdk/aiicy-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	static := []string{"name=aiicy", "org=linux", aiicy.EnvKeyServiceToken + "=key"}
	dyn := map[string]string{
		"repo":    "github",
		"project": "aiicy",
	}
	envs := GenerateInstanceEnv(t.Name(), static, dyn)
	assert.Contains(t, envs, "name=aiicy")
	assert.Contains(t, envs, "org=linux")
	assert.Contains(t, envs, "repo=github")
	assert.Contains(t, envs, "project=aiicy")
	assert.Contains(t, envs, aiicy.EnvKeyServiceInstanceName+"="+t.Name())
	assert.NotContains(t, envs, aiicy.EnvKeyServiceToken+"=key")
}
