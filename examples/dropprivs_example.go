package main

// This should be built, and then made setuid root to test

import (
	"fmt"
	"github.com/worr/dropprivs"
	"os"
)

func main() {
	uid := os.Getuid()
	gid := os.Getgid()

	euid := os.Geteuid()
	egid := os.Getegid()

	fmt.Printf("%v %v %v %v\n", uid, gid, euid, egid)

	if err := dropprivs.Dropprivs(4710, 4710); err != nil {
		fmt.Errorf("%s\n", err)
	}

	fmt.Printf("%v %v %v %v\n", os.Getuid(), os.Getgid(), os.Geteuid(), os.Getegid())

	if err := dropprivs.Dropprivs(euid, egid); err != nil {
		fmt.Errorf("%s\n", err)
	}

	fmt.Printf("%v %v %v %v\n", os.Getuid(), os.Getgid(), os.Geteuid(), os.Getegid())
}
