package droplet

import (
	"os"

	"github.com/cloudfoundry/cli/plugin"
)

//CFDownloader real implementation to download droplets.
type CFDownloader struct {
	Cli    plugin.CliConnection
	Writer FileWriter
}

//Downloader interaface for implementing downloaders.
type Downloader interface {
	GetDroplet(guid string) ([]byte, error)
	SaveDropletToFile(filePath string, data []byte) error
}

//FileWriter test shim for writing to a file.
type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

//GetDroplet from CF
func (downloader *CFDownloader) GetDroplet(guid string) ([]byte, error) {
	downloadURL := "/v2/apps/" + guid + "/droplet/download"
	droplet, err := downloader.Cli.CliCommandWithoutTerminalOutput("curl", downloadURL)
	return []byte(droplet[0]), err
}

//SaveDropletToFile writes a downloaded droplet to file
func (downloader *CFDownloader) SaveDropletToFile(filePath string, data []byte) error {
	return downloader.Writer.WriteFile(filePath, data, 0644)
}
