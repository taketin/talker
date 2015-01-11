package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandSelect,
	commandRewind,
	commandForward,
	commandMembers,
	commandFinishedMembers,
	commandCurrent,
}

var selectFlags = []cli.Flag{
	cli.BoolFlag{Name: "dry, d", Usage: "Dry run mode"},
	cli.BoolFlag{Name: "random, r", Usage: "Without storage, choose at random"},
}

var commandSelect = cli.Command{
	Name: "select",
	Usage: `
	Choose talker. The options are as follows:

		-d (--dry) : dry run mode

		-r (--random) : random choose mode (not use storage)
`,
	Description: `
`,
	Action: doSelect,
	Flags:  selectFlags,
}

var commandRewind = cli.Command{
	Name:  "rewind",
	Usage: "Remove the last of the talker from strage.",
	Description: `
`,
	Action: doRewind,
}

var commandMembers = cli.Command{
	Name:  "members",
	Usage: "Show all members.",
	Description: `
`,
	Action: doMembers,
}

var commandFinishedMembers = cli.Command{
	Name:  "finished",
	Usage: "Shows a list of selected users from strage.",
	Description: `
`,
	Action: doFinishedMembers,
}

var commandForward = cli.Command{
	Name:  "forward",
	Usage: "Add the given talker in storage.",
	Description: `
`,
	Action: doForward,
}

var commandCurrent = cli.Command{
	Name:  "current",
	Usage: "Show current talkers.",
	Description: `
`,
	Action: doCurrent,
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

	if isLast {
		println("🍣🍣🍣  All members has been completed! start the next round. 🍺🍺🍺\n")
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

func doMembers(c *cli.Context) {
	members := strings.Split(Configs.Talker.Members, ",")
	for _, member := range members {
		fmt.Printf("%s\n", member)
	}
}

func doFinishedMembers(c *cli.Context) {
	finishedMembers := finishedMemberFromStrage()
	for _, finishedMember := range finishedMembers {
		fmt.Printf("%s\n", finishedMember)
	}
}

func doCurrent(c *cli.Context) {
	currentNum := 2

	finishedMembers := finishedMemberFromStrage()
	currents := finishedMembers[len(finishedMembers) - currentNum:]
	for _, current := range currents {
		fmt.Printf("%s\n", current)
	}
}
