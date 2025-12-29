package command

import (
	"github.com/FewLy-Torch-1861/yaru/storage"
)

func Done(ID int) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == ID {
			tasks[i].Done = true
			break
		}
	}

	return storage.Save(tasks)
}
