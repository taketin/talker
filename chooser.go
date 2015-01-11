package main

import (
	"math/rand"
	"time"
)

func chooseTalker(members []string, save bool) (talker string, isLast bool) {
	finishedMember := finishedMemberFromStrage()
	notFinishedMember := getDiff(members, finishedMember)
	isLast = false

	if len(notFinishedMember) == 0 {
		notFinishedMember = members
		finishedMember = []string{}
	}

	talker = choosingBy(notFinishedMember)
	finishedMember = append(finishedMember, talker)

	if !save {
		saveFinishedMember(finishedMember)
	}

	if len(finishedMember) == len(members) {
		isLast = true
	}

	return
}

func getDiff(members []string, finishedMember []string) []string {
	for _, val := range finishedMember {
		for key2, val2 := range members {
			if val == val2 {
				members = append(members[:key2], members[key2+1:]...)
			}
		}
	}

	return members
}

func choosingBy(members []string) string {
	rand.Seed(time.Now().Unix())
	order := rand.Intn(len(members))

	return members[order]
}
