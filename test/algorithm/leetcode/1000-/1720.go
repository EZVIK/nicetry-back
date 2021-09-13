package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

func main() {

	//encoded := []int{1, 2, 3}
	//fmt.Println(decode(encoded, 1))

	//a1 := 2 ^ 1
	//a2 := a1 & 0
	//
	//fmt.Println(a1, a2)
	time.Sleep(time.Millisecond * 660)
	fmt.Println("Connecting server...")
	time.Sleep(time.Millisecond * 360)
	fmt.Println("photo will be saved in /usr/local/sexy/photo/")

	fmt.Println("Starting download...")
	initTime := 1000
	for {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		a := r1.Intn(initTime)
		time.Sleep(time.Millisecond * time.Duration(a))
		initTime -= a

		if initTime <= 1 {
			initTime += a * 2
		}

		name := GetSnowflakeId()
		fmt.Println(name, ".jpg is saved.")
	}
}

func GetSnowflakeId() string {
	//default node id eq 1,this can modify to different serverId node
	node, _ := snowflake.NewNode(10)
	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}
