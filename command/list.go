package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/FewLy-Torch-1861/yaru/color"
	"github.com/FewLy-Torch-1861/yaru/storage"
)

func List() error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Printf("%sYou have nothing to do!%s\n", color.PURPLE, color.RESET)
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, t := range tasks {
		var doneIcon string

		if t.Done == true {
			doneIcon = "✅"
		} else {
			doneIcon = "❌"
		}

		fmt.Fprintf(
			w,
			"%s#%d%s\t[%s] %s%s%s\t%s%s(%s)%s\n",
			color.DIM,
			t.ID,
			color.RESET,
			doneIcon,
			color.BOLD,
			t.Description,
			color.RESET,
			color.ITALIC,
			color.DIM,
			t.CreatedAt.Format("06-01-02"),
			color.RESET,
		)
	}

	return w.Flush()
}
