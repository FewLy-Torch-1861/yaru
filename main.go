package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/FewLy-Torch-1861/yaru/color"
	"github.com/FewLy-Torch-1861/yaru/command"
	"github.com/FewLy-Torch-1861/yaru/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) < 2 {
		model, err := tui.InitialModel()
		if err != nil {
			log.Fatal(err)
		}

		p := tea.NewProgram(model)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("%sPlease provide a task description%s\n", color.RED, color.RESET)
			return
		}

		description := os.Args[2]
		err := command.Add(description)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Task added: %s%s%s\n", color.PURPLE, description, color.RESET)

	case "done":
		if len(os.Args) < 3 {
			fmt.Printf("%sPlease provide a task ID%s\n", color.RED, color.RESET)
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		err = command.Done(id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Marked task %s#%d%s as doned!\n", color.PURPLE, id, color.RESET)

	case "list":
		err := command.List()
		if err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Printf("%sUnkown command%s\n", color.RED, color.RESET)
	}
}
