package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := New("../../configs/fastid.yml")
	require.NoError(t, err)
	assert.Equal(t, cfg.HTTP.Listen, ":8000")
}

func TestConfigError(t *testing.T) {
	_, err := New("invalid.yml")
	require.Error(t, err)
	assert.EqualError(t, err, "open invalid.yml: no such file or directory")
}
