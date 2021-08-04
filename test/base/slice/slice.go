package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	//test1 := make([]int, 0)
	//for FileOp := 0; FileOp < 8; FileOp++ {
	//	test1 = append(test1, FileOp)
	//}
	//fmt.Println(test1, len(test1), cap(test1))
	//
	//test2 := test1[0:2]					//  0 1 2
	//fmt.Println(test2, len(test2), cap(test2))
	//
	//test3 := test1[0:3]					//  0 1 2 3
	//fmt.Println(test3, len(test3), cap(test3))
	//
	//test2 = append(test2, 99)	//  0 1 2 99
	//fmt.Println(test2, len(test2), cap(test2))

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}

}
