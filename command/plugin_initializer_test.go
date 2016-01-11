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
	)
	BeforeEach(func() {
		initiliazer = &DownloadDropletCmdInitiliazer{}
		fakeCliConnection = &cli_fakes.FakeCliConnection{}
	})

	It("Should set the Droplet", func() {
		cmd := NewDownloadDropletCmd(initiliazer)
		initiliazer.InitializePlugin(cmd, fakeCliConnection)
		Ω(cmd.Drop).NotTo(BeNil())
	})

	//It should set the downloader on the droplet

})
