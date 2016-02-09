package plugin_test

import (
	"path/filepath"

	"github.com/cloudfoundry/cli/testhelpers/plugin_builder"
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	plugin_builder.BuildTestBinary(filepath.Join("..", "fixtures", "plugins"), "test_1")
	RunSpecs(t, "Plugin Suite")
}
