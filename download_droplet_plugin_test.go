package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/cloudfoundry/cli/plugin/models"
	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DownloadDropletPlugin", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var downloadDropletPlugin *DownloadDropletCmd

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		downloadDropletPlugin = &DownloadDropletCmd{Cli: fakeCliConnection}
	})

	Describe("app guids", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{Guid: "1234"}, nil)
		})

		It("Should retreive the a guid", func() {
			Ω(downloadDropletPlugin.GetGUID("foo")).To(Equal("1234"))
		})

		It("Should call the plugin service to get the app", func() {
			downloadDropletPlugin.GetGUID("foo")
			Ω(fakeCliConnection.GetAppArgsForCall(0)).Should(Equal("foo"))
			Ω(fakeCliConnection.GetAppCallCount()).To(Equal(1))
		})
	})
})
