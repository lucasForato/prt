package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SessionExists(sessionName string) bool {
	cmd := exec.Command("tmux", "has-session", "-t", sessionName)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return false
	}

	output := stdout.String()
	fmt.Println(output)
	return output == "0"
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

type Tmux struct {
	Name  string
	Git   bool
	Terms int
	Cmd   []string
}

func (t Tmux) CreateSession() {
	cmd := "tmux"
	args := []string{"new-session", "-s", t.Name, "-n", "vim", "nvim ."}

	if t.Git {
		args = append(args, ";", "new-window", "-n", "git", "lazygit")
	}

	if t.Terms == 1 {
		if len(t.Cmd) > 0 {
			args = append(args, ";", "new-window", "-n", "terminal", t.Cmd[0])
		} else {
			args = append(args, ";", "new-window", "-n", "terminal")
		}
	}

	if t.Terms == 2 {
		if len(t.Cmd) > 1 {
			args = append(args, ";", "new-window", "-n", "terminal", t.Cmd[0])
		  args = append(args, ";", "splitw", "-h", t.Cmd[1])
		} else {
			args = append(args, ";", "new-window", "-n", "terminal", t.Cmd[0])
		  args = append(args, ";", "splitw", "-h")
		}
	}

	exec := exec.Command(cmd, args...)
	exec.Stdin = os.Stdin
	exec.Stdout = os.Stdout
	exec.Stderr = os.Stderr

	if err := exec.Start(); err != nil {
		log.Fatal(err)
	}

	if err := exec.Wait(); err != nil {
		log.Fatal(err)
	}
}
