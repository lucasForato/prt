package utils

import (
	"bytes"
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
}

func SwitchSession(name string) {
	cmd := exec.Command("tmux", "switch-client", "-t", name)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

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
}

func UnsetSession() {
	cmd := exec.Command("unset", "TMUX")

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func AttachSession(name string) {
	cmd := exec.Command("tmux", "attach-session", "-t", name)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
