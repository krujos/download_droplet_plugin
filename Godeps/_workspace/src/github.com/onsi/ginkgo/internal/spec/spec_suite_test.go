package spec_test

import (
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}
