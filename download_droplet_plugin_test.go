package main_test

import (
	"os/exec"

	"github.com/cloudfoundry/cli/plugin/fakes"
	io_helpers "github.com/cloudfoundry/cli/testhelpers/io"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"

	"github.com/cloudfoundry/cli/testhelpers/rpc_server"

	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
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

	Describe("Run - usage and other things that exit", func() {
		var (
			rpcHandlers  *fake_rpc_handlers.FakeHandlers
			ts           *test_rpc_server.TestServer
			err          error
			pathToPlugin string
			args         []string
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

		})

		AfterEach(func() {
			ts.Stop()
		})

		It("exits 1 with bad arguments", func() {
			command := exec.Command(pathToPlugin, args...)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
		})

		It("prints the usage & exits 1 when called with less than three arguments", func() {
			command := exec.Command(pathToPlugin, args...)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))
			Eventually(session.Out).Should(gbytes.Say("cf download-droplet APP_NAME PATH"))
		})
	})
})
