package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSelect,
    commandRewind,
    commandDryrun,
    commandRandom,
    
}


var commandSelect = cli.Command{
	Name:  "select",
	Usage: "",
	Description: `
`,
	Action: doSelect,
}

var commandRewind = cli.Command{
	Name:  "rewind",
	Usage: "",
	Description: `
`,
	Action: doRewind,
}

var commandDryrun = cli.Command{
	Name:  "dryrun",
	Usage: "",
	Description: `
`,
	Action: doDryrun,
}

var commandRandom = cli.Command{
	Name:  "random",
	Usage: "",
	Description: `
`,
	Action: doRandom,
}


func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func doSelect(c *cli.Context) {
}

func doRewind(c *cli.Context) {
}

func doDryrun(c *cli.Context) {
}

func doRandom(c *cli.Context) {
}


