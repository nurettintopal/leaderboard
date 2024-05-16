package leaderboard

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println("initial test")

	var got = "a"
	var wanted = "b"

	if got == wanted {
		t.Errorf("got %q, wanted %q", got, wanted)
	}

}
