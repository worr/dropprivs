// +build linux openbsd freebsd dragonflybsd
// +build !darwin,!netbsd,!solaris

package dropprivs

import "syscall"

func Dropprivs(uid, gid int) error {
	// Drop extra groups
	if err := syscall.Setgroups([]int{gid}); err != nil {
		return err
	}

	if err := syscall.Setresgid(gid, gid, gid); err != nil {
		return err
	}

	if err := syscall.Setresuid(uid, uid, uid); err != nil {
		return err
	}

	return nil
}
