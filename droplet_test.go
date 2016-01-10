package main_test

import (
	"errors"

	"github.com/cloudfoundry/cli/plugin/fakes"
	"github.com/cloudfoundry/cli/plugin/models"
	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Droplet", func() {
	var fakeCliConnection *fakes.FakeCliConnection
	var droplet *Droplet
	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		droplet = &Droplet{Cli: fakeCliConnection}
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

	// Describe("Downloading the droplet", func() {
	// 	It("Should call download on the downloader", func() {
	// 		Fail("NYI")
	// 	})
	// })

})
