package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func SessionExists(sessionName string) bool {
	cmd := exec.Command("tmux", "list-sessions")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	output := stdout.String()
	return strings.Contains(output, sessionName)
}

func InTmuxSession() bool {
	env := os.Getenv("TMUX")
	if len(env) > 0 {
		return true
	}
	return false
}

func KillSession(name string) {
	cmd := exec.Command("tmux", "kill-session", "-t", name)

  err := cmd.Start() 
  if err != nil {
    log.Fatal(err)
  }

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session killed.")
}

func KillCurrSession() {
	cmd := exec.Command("tmux", "kill-session")

  err := cmd.Start()
  if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println("Current session killed.")
}
