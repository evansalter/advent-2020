package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
		return
	}
	cmd := args[0]
	var arg string
	if len(args) > 1 {
		arg = args[1]
	}

	switch cmd {
	case "new":
		newDay(arg)
	default:
		help()
	}
}

func help() {
	fmt.Print(`
Command must be one of:
	new <day number>
`)
}

func handleError(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
	help()
}

func newDay(day string) {
	if day == "" {
		help()
		return
	}

	dirs := []string{
		day,
		day + "/part1",
		day + "/part2",
	}
	for _, d := range dirs {
		if err := os.Mkdir(d, os.ModePerm); err != nil {
			handleError(err)
			return
		}
	}
	if f, err := os.Create(day + "/input.txt"); err != nil {
		handleError(err)
		return
	} else {
		if err := f.Close(); err != nil {
			handleError(err)
			return
		}
	}

	pathsToCopy := []string{
		"/part1/part1.go",
		"/part2/part2.go",
	}
	for _, p := range pathsToCopy {
		cmd := exec.Command("cp", "template"+p, day+p)
		if err := cmd.Run(); err != nil {
			handleError(err)
			return
		}
	}

	t, err := template.ParseFiles("template/main.template.txt")
	if err != nil {
		handleError(err)
		return
	}

	f, err := os.Create(day + "/main.go")
	if err != nil {
		handleError(err)
		return
	}

	if err := t.Execute(f, struct{ Day string }{Day: day}); err != nil {
		handleError(err)
		return
	}

	if err := f.Close(); err != nil {
		handleError(err)
		return
	}

	if err := exec.Command("git", "add", day).Run(); err != nil {
		handleError(err)
		return
	}
}
