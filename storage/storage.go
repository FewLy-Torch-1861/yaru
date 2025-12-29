package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/FewLy-Torch-1861/yaru/task"
)

func getFilePath() (string, error) {
	dataDir := os.Getenv("XDG_DATA_HOME")

	if dataDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		dataDir = filepath.Join(home, ".local", "share")
	}

	return filepath.Join(dataDir, "yaru", "tasks.json"), nil
}

func Load() ([]task.Task, error) {
	path, err := getFilePath()
	if err != nil {
		return nil, err
	}

	taskFile, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	var taskData []task.Task
	err = json.Unmarshal(taskFile, &taskData)
	if err != nil {
		return nil, err
	}

	return taskData, nil
}

func Save(tasks []task.Task) error {
	path, err := getFilePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}