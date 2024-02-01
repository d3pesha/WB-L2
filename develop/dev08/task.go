package main

import (
	"bufio"
	"errors"
	"fmt"
	gops "github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func processCommand(s string) {
	cmd := strings.Split(s, " ")
	switch cmd[0] {
	case "cd":
		cd(cmd)
	case "pwd":
		pwd()
	case "echo":
		echo(cmd)
	case "kill":
		kill(cmd)
	case "ps":
		ps()
	case "exec":
		execute(cmd)
	case "fork":
		forkCommand(cmd)
	case "\\q":
		os.Exit(0)
	}
}

func cd(args []string) {
	if len(args) == 1 {
		homeDir := os.Getenv("HOME")
		_ = os.Chdir(homeDir)
	} else if len(args) == 2 {
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Printf("no such file or directory: %s ", args[1])
		}
	}

}

func pwd() {
	current, _ := os.Getwd()
	fmt.Println(current)
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func kill(args []string) {
	id, _ := strconv.Atoi(args[1])
	p, err := os.FindProcess(id)
	if err != nil {
		fmt.Printf("kill %s failed: no such process", args[1])
	}

	err = p.Kill()
	if err != nil {
		fmt.Printf("kill %s failed: %s", args[1], err.Error())
	}
}

func ps() {
	processes, _ := gops.Processes()
	for _, p := range processes {
		fmt.Printf("name: %s \tpid: %d\n", p.Executable(), p.Pid())
	}
}

func execute(args []string) {
	if len(args) == 0 {
		fmt.Println("no command")
	}

	cmd := exec.Command(args[1], args[2:]...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
	fmt.Println(string(stdout))
}

func forkCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("fork: enter process to fork")
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("fork: couldn't get pwd")
	}

	cmd, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("fork: couldn't find path for %v: %w", args[0], err)
	}
	if cmd == "" {
		return "", fmt.Errorf("fork: couldn't find path for %v", args[0])
	}

	args[0] = cmd

	process, err := os.StartProcess(args[0], args, &os.ProcAttr{
		Dir:   pwd,
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})

	if err != nil {
		return "", fmt.Errorf("fork: couldn't fork: %w", err)
	}
	return fmt.Sprintf("Forked process with PID: %d", process.Pid), nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		processCommand(sc.Text())
	}
}
