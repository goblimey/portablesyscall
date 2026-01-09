//go:build linux
// +build linux

// The portablesyscall package provides a common set of interfaces to system calls on
// Windows and Linux.  All versions offer a set of functions with the same signatures.
// The constant OSName contains the string "windows" or "linux", the same name as the
// build tag for the target system.  The Windows version of the functions all return
// a syscall.EWINDOWS error when called.  This allows source code that is intended to
// run under Linux and uses system calls to at least compile under Windows.  The
// result can also be run in a limited fashion under Windows as long as it uses the
// OSName constant to avoid calling the syscall functions (which would break if
// called in that environment).
//
// A process can only use the features of this package if it's running as root on
// a Linux system because the underlying functionality only exists on that system and
// only root can use it.
package portablesyscall

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// OSName contains the name of the target operating system.  It's the same value as the build tag
// for that system ("windows", "linux" or whatever).
const OSName = "linux"

type Timespec syscall.Timespec
type Stat_t syscall.Stat_t

// EWINDOWS is defined in the Windows syscall.  Used to create errors.  The
// value is the one defined in Go 1.24.1.
const EWINDOWS = 536871042

// Setuid switches to the user with the given user ID or returns an error.
func Setuid(targetID int) error {
	return unix.Setuid(targetID)
}
