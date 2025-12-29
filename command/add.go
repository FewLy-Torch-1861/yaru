package command

import (
	"time"

	"github.com/FewLy-Torch-1861/yaru/storage"
	"github.com/FewLy-Torch-1861/yaru/task"
)

func Add(description string) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newID := maxID + 1

	newTask := task.Task{
		ID:          newID,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	return storage.Save(tasks)
}
