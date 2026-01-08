package portablesyscall

import (
	"testing"
)

// TestSetuidUnderWindows checks Setuid under Windows
func TestSetuidUnderWindows(t *testing.T) {

	if OSName == "windows" {
		err := Setuid(1)

		if OSName == "windows" {
			if err.Error() != "setuid : not supported by windows" {
				t.Errorf("want \"setuid : not supported by windows\" got %v", err.Error())
			}
		}
	}
}
