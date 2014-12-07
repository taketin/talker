package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSelect,
    commandRewind,
	commandForward,
}

var selectFlags = []cli.Flag{
	cli.BoolFlag{Name: "dry, d", Usage: "Dry run mode"},
	cli.BoolFlag{Name: "random, r", Usage: "Without storage, choose at random"},
}

var commandSelect = cli.Command{
	Name:  "select",
	Usage: `
	Choose talker. The options are as follows:

		-d (--dry) : dry run mode

		-r (--random) : random choose mode (not use storage)
`,
	Description: `
`,
	Action: doSelect,
	Flags: selectFlags,
}

var commandRewind = cli.Command{
	Name:  "rewind",
	Usage: "Remove the last of the talker from strage.",
	Description: `
`,
	Action: doRewind,
}

var commandForward = cli.Command{
	Name:  "forward",
	Usage: "Add the given talker in storage.",
	Description: `
`,
	Action: doForward,
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
	isDryrun := c.Bool("dry")
	isRandom := c.Bool("random")

	members := strings.Split(Configs.Talker.Members, ",")

	talker := ""
	isLast := false

	if isRandom {
		talker = choosingBy(members)
	} else {
		talker, isLast = chooseTalker(members, isDryrun)
	}

	fmt.Printf("%s \n", talker)

	if (isLast) {
		println("🍣🍣🍣  ALL members done! start the next round. 🍺🍺🍺")
	}
}

func doRewind(c *cli.Context) {
	finishedMember := finishedMemberFromStrage()
	removeMember := finishedMember[len(finishedMember)-1]
	saveFinishedMember(finishedMember[:len(finishedMember)-1])
	fmt.Printf("'%s' remove from storage.\n", removeMember)
}

func doForward(c *cli.Context) {
	name := ""
	if len(c.Args()) > 0 {
		name = c.Args()[0]
	} else {
		fmt.Printf("Please input forward name.\n")
		os.Exit(1)
	}

	finishedMember := finishedMemberFromStrage()
	finishedMember = append(finishedMember, name)
	saveFinishedMember(finishedMember)
	fmt.Printf("'%s' add to storage.\n", name)
}
