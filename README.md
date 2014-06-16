# dropprivs

## Why dropprivs?

Face it &mdash; sometimes you need to start a service with elevated
privileges. Even if you want to do something benign like drop yourself
immediately into a chroot, you need to elevate privileges first.

Doing this is non-obvious though, as well as platform-specific. Linux,
FreeBSD, OpenBSD and HP-UX include
the handy setresuid(2) and retresgid(2) functions to facilitate setting the
real, effective and *saved* uid and gid, respectively.

NetBSD, OS X and Solaris...do not.
