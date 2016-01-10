package main_test

import (
	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"

	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const pluginPath = "./download-droplet-plugin"

var _ = Describe("DownloadDropletPlugin", func() {

	var (
		// rpcHandlers           *fake_rpc_handlers.FakeHandlers
		// ts                    *test_rpc_server.TestServer
		// err                   error
		fakeCliConnection     *fakes.FakeCliConnection
		downloadDropletPlugin *DownloadDropletCmd
	)

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		downloadDropletPlugin = &DownloadDropletCmd{}
	})

	Describe("GetMetadata", func() {

		It("Returns metadata", func() {
			md := downloadDropletPlugin.GetMetadata()
			Ω(md).NotTo(BeNil())
		})

		It("Has a help message", func() {
			md := downloadDropletPlugin.GetMetadata()
			Ω(md.Commands[0].HelpText).NotTo(BeNil())
		})
	})

	Describe("Run", func() {
		It("pritns an informative message when downloading the droplet", func() {
			output := io_helpers.CaptureOutput(func() {
				downloadDropletPlugin.Run(fakeCliConnection, []string{"download-droplet", "theApp", "/tmp"})
			})
			Ω(output[0]).To(Equal("Saving theApp's droplet to /tmp"))
		})

		It("throws an error if the first arg is not download-droplet", func() {
			output := io_helpers.CaptureOutput(func() {
				downloadDropletPlugin.Run(fakeCliConnection, []string{"garbage", "foo", "/path"})
			})
			Ω(output[0]).To(ContainSubstring("unknown command"))
		})
	})

	// Describe("Run - usage and other things that exit", func() {
	// 	BeforeSuite(func() {
	// 		var err error
	// 		pathToPlugin, err = gexec.Build("github.com/krujos/download_droplet_plugin")
	// 		Ω(err).ShouldNot(HaveOccurred())
	// 	})
	//
	// 	AfterSuite(func() {
	// 		gexec.CleanupBuildArtifacts()
	// 	})
	//
	// 	It("prints the usage & exits 1 when called wiht less than three arguments", func() {
	//
	// 		Eventually(session).Should(gexec.Exit(1))
	//
	// 		output := io_helpers.CaptureOutput(func() {
	// 			downloadDropletPlugin.Run(fakeCliConnection, []string{"download-droplet"})
	// 		})
	// 		Ω(output[0]).To(Equal("NAME:"))
	// 		Ω(output[3]).To(Equal("USAGE:"))
	// 		Ω(output[5]).To(Equal("cf download-droplet APP_NAME PATH"))
	// 	})
	// })

})
