package main

import (
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

// CLI information
const (
	AppName = "clip"
	Author  = "mingrammer"
	Version = "0.0.2"
	Usage   = "Store and copy the frequently used text to clipboard"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Author = Author
	app.Version = Version
	app.Usage = Usage
	app.Commands = buildCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func homeDir() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	return home
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
