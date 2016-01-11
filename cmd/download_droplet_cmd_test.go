package cmd_test

import (
	"errors"

	cli_fakes "github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	. "github.com/krujos/download_droplet_plugin/cmd"
	fake_droplet "github.com/krujos/download_droplet_plugin/droplet/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DownloadDropletCmd", func() {
	var (
		fakeCliConnection     *cli_fakes.FakeCliConnection
		downloadDropletPlugin *DownloadDropletCmd
		goodArgs              []string
		uninstallArgs         []string
		fakeDroplet           *fake_droplet.FakeDroplet
	)

	BeforeEach(func() {
		fakeCliConnection = &cli_fakes.FakeCliConnection{}
		fakeDroplet = new(fake_droplet.FakeDroplet)
		downloadDropletPlugin = &DownloadDropletCmd{}
		downloadDropletPlugin.Drop = fakeDroplet
		goodArgs = []string{"download-droplet", "theApp", "/tmp"}
		uninstallArgs = []string{"CLI-MESSAGE-UNINSTALL"}
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
		Context("Messages", func() {
			It("pritns an informative message when downloading the droplet", func() {
				output := io_helpers.CaptureOutput(func() {
					downloadDropletPlugin.Run(fakeCliConnection, goodArgs)
				})
				Ω(output[0]).To(Equal("Saving theApp's droplet to /tmp"))
			})
		})

		Context("initializer complication", func() {
			It("Should call the initializer during run", func() {
				Fail("NYI")
			})
		})

		Context("Saving a droplet", func() {
			It("should call save dropplet with the right arguments", func() {
				downloadDropletPlugin.Run(fakeCliConnection, goodArgs)
				name, path := fakeDroplet.SaveDropletArgsForCall(0)
				Ω(name).Should(Equal("theApp"))
				Ω(path).Should(Equal("/tmp"))
				Ω(fakeDroplet.SaveDropletCallCount()).To(Equal(1))
			})

			It("Should print an error message if an err is returned", func() {
				fakeDroplet.SaveDropletReturns(errors.New("This is an error"))
				output := io_helpers.CaptureOutput(func() {
					downloadDropletPlugin.Run(fakeCliConnection, goodArgs)
				})
				Ω(output[1]).To(Equal("This is an error"))
			})

			It("Should print an uninstall message", func() {
				output := io_helpers.CaptureOutput(func() {
					downloadDropletPlugin.Run(fakeCliConnection, uninstallArgs)
				})
				Ω(output[0]).To(Equal("Thanks for using droplet downloader!"))
			})
		})
	})
})
