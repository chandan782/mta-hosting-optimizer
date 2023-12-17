package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpConfigJSONSerialization(t *testing.T) {
	ipConfig := IpConfig{
		IP:       "127.0.0.1",
		Hostname: "test-host",
		Active:   true,
	}

	jsonData, err := json.Marshal(ipConfig)
	assert.NoError(t, err)

	var deserializedIpConfig IpConfig
	err = json.Unmarshal(jsonData, &deserializedIpConfig)
	assert.NoError(t, err)

	assert.Equal(t, ipConfig, deserializedIpConfig)
}

func TestIpConfigEquality(t *testing.T) {
	ipConfig1 := IpConfig{
		IP:       "127.0.0.1",
		Hostname: "test-host",
		Active:   true,
	}

	ipConfig2 := IpConfig{
		IP:       "127.0.0.1",
		Hostname: "test-host",
		Active:   true,
	}

	assert.Equal(t, ipConfig1, ipConfig2)
}

func TestIpConfigInequality(t *testing.T) {
	ipConfig1 := IpConfig{
		IP:       "127.0.0.1",
		Hostname: "test-host-1",
		Active:   true,
	}

	ipConfig2 := IpConfig{
		IP:       "127.0.0.1",
		Hostname: "test-host-2",
		Active:   true,
	}

	assert.NotEqual(t, ipConfig1, ipConfig2)
}
