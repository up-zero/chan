package test

import (
	"fmt"
	"os"
	"os/user"
	"testing"
)

func TestUser(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u.HomeDir)
}

func TestStop(t *testing.T) {
	p, err := os.FindProcess(2648)
	if err != nil {
		t.Fatal(err)
	}
	err = p.Kill()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("KILL SUCCESS")
}
