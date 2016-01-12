package main

import "github.com/krujos/download_droplet_plugin/command"

func main() {
	initer := command.NewDownloadDropletCmdInitiliazer()
	command.NewDownloadDropletCmd(initer).Start()
}
