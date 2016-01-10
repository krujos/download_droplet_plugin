package droplet_test

import (
	"errors"

	. "github.com/krujos/download_droplet_plugin/droplet"

	"github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/cloudfoundry/cli/plugin/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DropletDownloader", func() {

	var fakeCliConnection *fakes.FakeCliConnection
	var downloader *Downloader

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		downloader = &Downloader{Cli: fakeCliConnection}
	})

	Describe("app guids", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{Guid: "1234"}, nil)
		})

		It("Should retreive the a guid", func() {
			Ω(downloader.GetGUID("foo")).To(Equal("1234"))
		})

		It("Should call the plugin service to get the app", func() {
			_, err := downloader.GetGUID("foo")
			Ω(fakeCliConnection.GetAppArgsForCall(0)).Should(Equal("foo"))
			Ω(fakeCliConnection.GetAppCallCount()).To(Equal(1))
			Ω(err).Should(BeNil())
		})
	})

	Describe("GetApp failures", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{},
				errors.New("Bad Things"))
		})

		It("Should reutrn the error that GetApp does", func() {
			_, err := downloader.GetGUID("bar")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Describe("Downloading the droplet", func() {
		tarFileContents := "This is a tar file"
		BeforeEach(func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(
				[]string{tarFileContents}, nil)
		})

		It("Should get back a non null byte array for valid droplet", func() {
			bytes, _ := downloader.GetDroplet("1234")
			Ω(bytes).Should(Equal([]byte(tarFileContents)))
		})
	})

	Describe("Downloading the droplet fails", func() {
		BeforeEach(func() {
			fakeCliConnection.CliCommandWithoutTerminalOutputReturns(
				[]string{""}, errors.New("Busted"))
		})

		It("Should return the download error", func() {
			_, err := downloader.GetDroplet("1234")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Describe("Saveing a droplet to a file", func() {
		It("Should write to a file", func() {

		})
	})

})
