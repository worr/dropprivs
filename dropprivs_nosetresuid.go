// +build netbsd darwin solaris
// +build !linux,!openbsd,!freebsd,!dragonflybsd

package dropprivs

import "syscall"

// Drops all elevated privileges permanently
func Dropprivs(uid, gid int) error {
	// Drop extra groups
	if err := syscall.Setgroups([]int{gid}); err != nil {
		return err
	}

	if err := syscall.Setgid(gid); err != nil {
		return err
	}

	if err := syscall.Setegid(gid); err != nil {
		return err
	}

	if err := syscall.Setuid(uid); err != nil {
		return err
	}

	if err := syscall.Seteuid(uid); err != nil {
		return err
	}

	return nil
}
