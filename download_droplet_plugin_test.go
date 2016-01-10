package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("DownloadDropletPlugin", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var downloadDropletPlugin *DownloadDropletCmd

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		downloadDropletPlugin = &DownloadDropletCmd{Cli: fakeCliConnection}
	})

})
