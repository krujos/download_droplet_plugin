package droplet_test

import (
	"errors"
	"os"

	cliFakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/krujos/download_droplet_plugin/droplet"
	"github.com/krujos/download_droplet_plugin/droplet/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DropletDownloader", func() {

	var fakeCliConnection *cliFakes.FakeCliConnection
	var downloader *Downloader

	BeforeEach(func() {
		fakeCliConnection = &cliFakes.FakeCliConnection{}
		downloader = &Downloader{Cli: fakeCliConnection}
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
		tarFileContents := []byte("This is the droplet")
		var fileWriter *fakes.FakeFileWriter
		var downloader *Downloader

		BeforeEach(func() {
			fileWriter = new(fakes.FakeFileWriter)
			downloader = &Downloader{Writer: fileWriter}
		})

		It("Should write the droplet to a file", func() {
			downloader.SaveDropletToFile("/tmp", tarFileContents)
			dir, data, mode := fileWriter.WriteFileArgsForCall(0)
			Ω(dir).Should(Equal("/tmp"))
			Ω(data).Should(Equal(tarFileContents))
			Ω(mode).Should(Equal(os.FileMode(0644)))
		})

		It("should return the fs error if it has one", func() {
			fileWriter.WriteFileReturns(errors.New("Broke!!!"))
			Ω(downloader.SaveDropletToFile("/tmp", tarFileContents)).ShouldNot(BeNil())
		})
	})

})
