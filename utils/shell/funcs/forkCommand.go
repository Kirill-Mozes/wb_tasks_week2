package funcs

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func forkCommand(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("shell: fork command need amount of processes")
	}
	fork, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	children := []int{}
	var builder strings.Builder
	pid := os.Getpid()
	ppid := os.Getppid()
	builder.WriteString(
		fmt.Sprintf("pid: %d, ppid: %d, forks: %d\n", pid, ppid, fork),
	)
	if _, isChild := os.LookupEnv("CHILD_ID"); !isChild {
		for i := 0; i < fork; i++ {
			args := append(os.Args, fmt.Sprintf("#child_%d_of_%d", i, os.Getpid()))
			childENV := []string{
				fmt.Sprintf("CHILD_ID=%d", i),
			}
			pwd, err := os.Getwd()
			if err != nil {
				return "", err
			}
			childPID, _ := syscall.ForkExec(args[0], args, &syscall.ProcAttr{
				Dir: pwd,
				Env: append(os.Environ(), childENV...),
				Sys: &syscall.SysProcAttr{
					Setsid: true,
				},
				Files: []uintptr{0, 1, 2}, // print message to the same pty
			})
			builder.WriteString(
				fmt.Sprintf("parent %d fork %d\n", pid, childPID),
			)
			if childPID != 0 {
				children = append(children, childPID)
			}
		}
		// print children
		builder.WriteString(
			fmt.Sprintf("parent: PID=%d children=%v", pid, children),
		)
		if len(children) == 0 && fork != 0 {
			return "", errors.New("shell: fork no child avaliable, exit")
		}

		// set env
		for _, childID := range children {
			if c := os.Getenv("CHILDREN"); c != "" {
				os.Setenv("CHILDREN", fmt.Sprintf("%s,%d", c, childID))
			} else {
				os.Setenv("CHILDREN", fmt.Sprintf("%d", childID))
			}
		}
	}
	return "", nil
}
