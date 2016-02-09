package droplet_test

import (
	"errors"
	"net/http"
	"os"

	cliFakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/krujos/download_droplet_plugin/droplet"
	"github.com/krujos/download_droplet_plugin/droplet/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("DropletDownloader", func() {

	var fakeCliConnection *cliFakes.FakeCliConnection
	var downloader *CFDownloader
	var server *ghttp.Server
	tarFileContents := "This is a tar file"

	BeforeEach(func() {
		fakeCliConnection = &cliFakes.FakeCliConnection{}
		downloader = &CFDownloader{Cli: fakeCliConnection}
		server = ghttp.NewServer()
		fakeCliConnection.AccessTokenReturns("bearer 1234", nil)
		fakeCliConnection.ApiEndpointReturns(server.URL(), nil)
		fakeCliConnection.IsSSLDisabledReturns(true, nil)
		server.AppendHandlers(
			ghttp.VerifyRequest("GET", "/v2/apps/1234/droplet/download"),
			ghttp.VerifyHeader(
				http.Header{
					"Authorization": []string{"bearer 1234"},
				}),
			ghttp.RespondWith(http.StatusOK, []byte(tarFileContents)),
		)
	})

	AfterEach(func() {
		server.Close()
	})

	// Describe("Downloading the droplet", func() {
	// 	It("Should get back a non null byte array for valid droplet", func() {
	// 		bytes, _ := downloader.GetDroplet("1234")
	// 		Ω(bytes).Should(Equal([]byte(tarFileContents)))
	// 	})
	// })

	// Describe("Downloading the droplet fails", func() {
	// 	BeforeEach(func() {
	// 		fakeCliConnection.CliCommandWithoutTerminalOutputReturns(
	// 			[]string{""}, errors.New("Busted"))
	// 	})
	//
	// 	It("Should return the download error", func() {
	// 		_, err := downloader.GetDroplet("1234")
	// 		Ω(err).ShouldNot(BeNil())
	// 	})
	// })

	Describe("Saveing a droplet to a file", func() {
		tarFileContents := []byte("This is the droplet")
		var fileWriter *fakes.FakeFileWriter
		var downloader *CFDownloader

		BeforeEach(func() {
			fileWriter = new(fakes.FakeFileWriter)
			downloader = &CFDownloader{Writer: fileWriter}
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
