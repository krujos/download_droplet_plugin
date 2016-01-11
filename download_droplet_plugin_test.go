package main_test

import (
	"os/exec"

	"github.com/cloudfoundry/cli/testhelpers/rpc_server"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"
	. "github.com/krujos/download_droplet_plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

const pluginPath = "./download-droplet-plugin"

var _ = Describe("DownloadDropletPlugin", func() {

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
