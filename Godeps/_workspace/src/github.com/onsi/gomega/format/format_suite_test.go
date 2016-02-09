package format_test

import (
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}
