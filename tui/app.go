package tui

import (
	"fmt"
	"strings"

	"github.com/FewLy-Torch-1861/yaru/color"
	"github.com/FewLy-Torch-1861/yaru/storage"
	"github.com/FewLy-Torch-1861/yaru/task"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tasks  []task.Task
	cursor int
}

func InitialModel() (Model, error) {
	tasks, err := storage.Load()
	if err != nil {
		return Model{}, err
	}

	return Model{
		tasks:  tasks,
		cursor: 0,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%sYaru Todo List%s\n\n", color.PURPLE, color.RESET)

	for i, t := range m.tasks {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		doneIcon := "❌"
		if t.Done {
			doneIcon = "✅"
		}

		date := t.CreatedAt.Format("06-01-02")

		fmt.Fprintf(&b, "%s%s#%d%s [%s] %s%s%s \t%s%s(%s)%s\n",
			cursor,
			color.DIM,
			t.ID,
			color.RESET,
			doneIcon,
			color.BOLD,
			t.Description,
			color.RESET,
			color.ITALIC,
			color.DIM,
			date,
			color.RESET,
		)
	}

	b.WriteString(fmt.Sprintf("\n%sPress q to quit.%s\n", color.DIM, color.RESET))
	return b.String()
}
