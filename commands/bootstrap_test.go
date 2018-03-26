package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBootstrapList(t *testing.T) {
	assert := assert.New(t)

	d := NewDaemon(t).Start()
	defer d.ShutdownSuccess()

	bs := d.RunSuccess("bootstrap list")

	assert.Equal("[TODO]\n", bs.ReadStdout())
}