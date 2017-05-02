package util

import (
	"math/rand"
	"fmt"
	"time"
)

func RelationIdGen(userId int, otherUserId int) string {
	rand.Seed(time.Now().UnixNano())

	return fmt.Sprintf("%d@%d@%d@%d",
		time.Now().Unix(), userId, otherUserId, rand.Intn(1000))
}
