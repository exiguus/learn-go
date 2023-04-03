package network_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/exiguus/gotesting/network"
	"github.com/stretchr/testify/assert"
)

func TestReadHosts(t *testing.T) {
	t.Parallel()
	filesystem := fstest.MapFS{
		"hosts": {
			Data: []byte("127.0.0.1 localhost\n"),
		},
	}
	content, err := network.ReadHosts(filesystem)
	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
}

func TestReadFileFromHost(t *testing.T) {
	t.Parallel()
	tempDir, err := os.MkdirTemp("", "test-hostfile--")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	hostfile := filepath.Join(tempDir, "hosts")
	err = os.WriteFile(hostfile, []byte("127.0.0.1 localhost\n"), 0644)
	assert.NoError(t, err)

	content, err := network.ReadHostsFromFile(hostfile)
	assert.Equal(t, "127.0.0.1 localhost\n", content)
	assert.NoError(t, err)
} // or use https://pkg.go.dev/github.com/spf13/afero
