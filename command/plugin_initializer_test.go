package command_test

import (
	cli_fakes "github.com/cloudfoundry/cli/plugin/fakes"
	. "github.com/krujos/download_droplet_plugin/command"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PluginInitializer", func() {
	var (
		initiliazer       *DownloadDropletCmdInitiliazer
		fakeCliConnection *cli_fakes.FakeCliConnection
		cmd               *DownloadDropletCmd
	)
	BeforeEach(func() {
		initiliazer = &DownloadDropletCmdInitiliazer{}
		fakeCliConnection = &cli_fakes.FakeCliConnection{}
		cmd = NewDownloadDropletCmd(initiliazer)
		initiliazer.InitializePlugin(cmd, fakeCliConnection)
	})

	It("Should set the Droplet", func() {
		Ω(cmd.Drop).NotTo(BeNil())
	})

	It("Should set the downloader on the droplet", func() {
		Ω(cmd.Drop.GetDownloader()).ShouldNot(BeNil())
	})
})
