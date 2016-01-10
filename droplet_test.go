package main_test

import (
	"errors"

	cliFakes "github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/cloudfoundry/cli/plugin/models"
	. "github.com/krujos/download_droplet_plugin"
	"github.com/krujos/download_droplet_plugin/droplet/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Droplet", func() {
	var fakeCliConnection *cliFakes.FakeCliConnection
	var droplet *Droplet
	var fakeDownloader *fakes.FakeDownloader

	BeforeEach(func() {
		fakeCliConnection = &cliFakes.FakeCliConnection{}
		fakeDownloader = &fakes.FakeDownloader{}
		droplet = &Droplet{
			Cli:        fakeCliConnection,
			Downloader: fakeDownloader,
		}
	})

	Describe("Getting the app details from cf", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{Guid: "1234"}, nil)
		})

		It("Should call the plugin service to get the app", func() {
			err := droplet.SaveDroplet("foo", "/tmp")
			Ω(fakeCliConnection.GetAppArgsForCall(0)).Should(Equal("foo"))
			Ω(fakeCliConnection.GetAppCallCount()).To(Equal(1))
			Ω(err).Should(BeNil())
		})
	})

	Describe("Getting app details from cf failure scinerieos", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{},
				errors.New("Bad Things"))
		})

		It("Should reutrn the error that GetApp does", func() {
			err := droplet.SaveDroplet("bar", "/tmp")
			Ω(err).ShouldNot(BeNil())
		})
	})

	Describe("Saving the droplet", func() {
		BeforeEach(func() {
			fakeCliConnection.GetAppReturns(plugin_models.GetAppModel{Guid: "1234"}, nil)
		})

		It("Should call download on the downloader", func() {
			droplet.SaveDroplet("bar", "/tmp")
			Ω(fakeDownloader.SaveDropletToFileCallCount()).Should(Equal(1))
		})

		It("Should return an error if the download fails", func() {
			fakeDownloader.SaveDropletToFileReturns(errors.New("Failed to save"))
			err := droplet.SaveDroplet("bar", "/tmp")
			Ω(fakeDownloader.SaveDropletToFileCallCount()).Should(Equal(1))
			Ω(err).NotTo(BeNil())
		})
	})

	Describe("Downloading the droplet", func() {
		It("Should download the droplet", func() {
			droplet.SaveDroplet("bar", "/tmp")
			Ω(fakeDownloader.GetDropletCallCount()).Should(Equal(1))
		})

		It("Should return an error if downloading the droplet fails", func() {
			droplet.SaveDroplet("bar", "/tmp")
			fakeDownloader.GetDropletReturns(nil, errors.New("Download failed"))
			err := droplet.SaveDroplet("bar", "/tmp")
			Ω(err).NotTo(BeNil())
		})

	})

})
