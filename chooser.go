package main

import(
	"math/rand"
	"time"
)

func chooseTalker(members []string) (string) {
	finishedMember := finishedMemberFromStrage()

	notFinishedMember := compare(members, finishedMember)
	if len(notFinishedMember) == 0 {
		println("🍣🍣🍣  ALL members done! start the next round. 🍺🍺🍺")
		notFinishedMember = members
		finishedMember = []string{}
	}

	talker := choosingBy(notFinishedMember)

	finishedMember = append(finishedMember, talker)
	saveFinishedMember(finishedMember)

	return talker
}

func compare(members []string, finishedMember []string) ([]string) {
	for _, val := range finishedMember {
		for key2, val2 := range members {
			if val == val2 {
				members = append(members[:key2], members[key2+1:]...)
			}
		}
	}

	return members
}

func choosingBy(members []string) (string) {
	rand.Seed(time.Now().Unix())
	order := rand.Intn(len(members))

	return members[order]
}
