package main_test

import (
	"errors"
	"os/exec"

	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	"github.com/cloudfoundry/cli/testhelpers/rpc_server"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"
	. "github.com/krujos/download_droplet_plugin"
	fake_droplet "github.com/krujos/download_droplet_plugin/droplet/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const pluginPath = "./download-droplet-plugin"

var _ = Describe("DownloadDropletPlugin", func() {

	var (
		fakeCliConnection     *fakes.FakeCliConnection
		downloadDropletPlugin *DownloadDropletCmd
		goodArgs              []string
		uninstallArgs         []string
		fakeDroplet           *fake_droplet.FakeDroplet
	)

	BeforeEach(func() {
		fakeCliConnection = &fakes.FakeCliConnection{}
		fakeDroplet = new(fake_droplet.FakeDroplet)
		downloadDropletPlugin = &DownloadDropletCmd{Drop: fakeDroplet}
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

	Describe("Run - integration tests", func() {
		var (
			rpcHandlers  *fake_rpc_handlers.FakeHandlers
			ts           *test_rpc_server.TestServer
			err          error
			pathToPlugin string
			args         []string
			badArgs      []string
		)

		BeforeSuite(func() {
			var err error
			pathToPlugin, err = gexec.Build("github.com/krujos/download_droplet_plugin")
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterSuite(func() {
			gexec.CleanupBuildArtifacts()
		})

		BeforeEach(func() {
			rpcHandlers = &fake_rpc_handlers.FakeHandlers{}
			ts, err = test_rpc_server.NewTestRpcServer(rpcHandlers)
			Ω(err).NotTo(HaveOccurred())

			err = ts.Start()
			Ω(err).NotTo(HaveOccurred())

			rpcHandlers.CallCoreCommandStub = func(_ []string, retVal *bool) error {
				*retVal = true
				return nil
			}

			rpcHandlers.GetOutputAndResetStub = func(_ bool, retVal *[]string) error {
				*retVal = []string{"{}"}
				return nil
			}
			args = []string{ts.Port(), "download-droplet"}
			badArgs = []string{ts.Port(), "garbage", "foo", "/path"}

		})

		AfterEach(func() {
			ts.Stop()
		})

		It("prints the usage & exits 1 when called with less than three arguments", func() {
			command := exec.Command(pathToPlugin, args...)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Out).Should(gbytes.Say("cf download-droplet APP_NAME PATH"))
		})

		It("prints usage and exits 1 if the first arg is not download-droplet", func() {
			command := exec.Command(pathToPlugin, badArgs...)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Out).Should(gbytes.Say("cf download-droplet APP_NAME PATH"))
		})
	})
})
