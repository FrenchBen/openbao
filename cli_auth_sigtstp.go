//go:build unix

package jwtauth

import (
	"os"
	"syscall"
)

// authHalts are the signals we want to interrupt our auth callback on.
// SIGTSTP is omitted for Windows.
var authHalts = []os.Signal{os.Interrupt, os.Kill, syscall.SIGTSTP}
