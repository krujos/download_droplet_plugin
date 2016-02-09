package table_test

import (
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/krujos/download_droplet_plugin/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Table Suite")
}
