//go:build windows
// +build windows

// The syscall package provides a common set of interfaces to system calls on
// Windows and Linux.  All versions offer a set of functions with the same signatures.
// The constant OSName contains the string "windows" or "linux", the same name as the
// build tag for the target system.  The Windows version of the functions all return
// a syscall.EWINDOWS ("not supported by windows") error when called.
//
// This allows source code that is intended to run under Linux and uses system calls to
// at least compile under Windows.  The result can also be run in a limited fashion
// under Windows as long as it uses the OSName constant to avoid calling the syscall
// functions (which would break if called in that environment).  This allows system
// testing under Windows of other parts of the solution that don't use this package.
package portablesyscall

import (
	"io/fs"
	"os"
	"syscall"
)

// OSName contains the name of the target operating system.  It's the same value as the build tag
// for that system ("windows", "linux" or whatever).
const OSName = "windows"

// Stat_t (and therfore Timespec) is defined in syscall under Linux.
// Timespec is a dependency of Stat_t.  This is the definition in Go 1.24.1.
type Timespec struct {
	Sec  int64
	Nsec int64
}

type Stat_t struct {
	Dev       uint64
	Ino       uint64
	Nlink     uint64
	Mode      uint32
	Uid       uint32
	Gid       uint32
	X__pad0   int32
	Rdev      uint64
	Size      int64
	Blksize   int64
	Blocks    int64
	Atim      Timespec
	Mtim      Timespec
	Ctim      Timespec
	X__unused [3]int64
}

// EWINDOWS is defined in syscall under Windows but not under Linux.
const EWINDOWS = syscall.EWINDOWS

// Setuid switches the effective user to the user with the given user ID.  The
// Windows version always returns a syscall.EWINDOWS wrapped in a PathError
// (which is what os.Chown does in the same situation).
func Setuid(targetID int) error {
	return &fs.PathError{Op: "setuid", Err: EWINDOWS}
}

func Stat(f os.File) (*Stat_t, error) {
	return nil, &fs.PathError{Op: "stat", Err: EWINDOWS}
}
