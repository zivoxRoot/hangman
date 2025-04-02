package config

// import (
// 	"os"
// 	"errors"
// 	"io/fs"
// 	"io/ioutil"
// 	"strings"
// )
//
// type Configurator struct {
// 	FilePath string
// 	WordListPath string
// }
//
// type Configuration struct {
// 	PrintBanner bool
// 	PrintHints bool
// 	UseColors bool
// }
//
// func NewConfigurator(folder string) *Configurator {
// 	configurator := Configurator{
// 		FilePath: folder,
// 	}
//
// 	return &configurator
// }
//
// func (c *Configurator) checkFolderExists() (bool, error) {
// 	path := c.FilePath
// 	_, err := os.Stat(path)
//
// 	if err == nil {
// 		return true, nil
// 	}
//
// 	if errors.Is(err, fs.ErrNotExist) {
// 		return false, nil
// 	}
//
// 	return false, err
// }
//
// func (c *Configurator) NewConfiguration() error {
// 	_, err := c.checkFolderExists()
// 	if err != nil {
// 		return err
// 	}
//
// 	// Generate a default configuration.
// 	config := Configuration{
// 		PrintBanner: true,
// 		PrintHints: true,
// 		UseColors: true,
// 	}
//
// 	// Read the configuration file.
// 	content, err := c.readConfig()
// 	if err != nil {
// 		return err
// 	}
//
// 	// Split the content in a slice
// 	contentSlice := strings.Split(content, " ")
//
// 	for _, i := range contentSlice {
// 		if i == "PrintBanner" {
// 			config.PrintBanner = []contentSlice{i + 2} // NOTE: ERRORS HERE !!!!!!!
// 			continue
// 		}
// 	}
// }
//
// func (c *Configurator) readConfig() (string, error) {
// 	content, err := ioutil.ReadFile(c.FilePath)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return string(content), nil
// }
