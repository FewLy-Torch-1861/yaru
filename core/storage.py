import json
from core.model import Task, TaskStatus
from warnings import warn
from platformdirs import user_data_dir
from pathlib import Path
from dataclasses import asdict

app_name = "yaru"
data_dir = Path(user_data_dir(app_name))
data_dir.mkdir(parents=True, exist_ok=True)
data_file = data_dir / "data.json"

placeholder_data = {"tasks": []}


def load_tasks():
    try:
        with open(data_file, "r", encoding="UTF-8") as f:
            data = json.load(f)
            tasks_objects = []

            for t in data["tasks"]:
                task = Task(
                    id=t["id"],
                    title=t["title"],
                    status=TaskStatus(t["status"]),
                )
                tasks_objects.append(task)

            return tasks_objects

    except FileNotFoundError:
        with open(data_file, "w", encoding="UTF-8") as f:
            json.dump(placeholder_data, f, ensure_ascii=False, indent=2)
            return []

    except json.JSONDecodeError as e:
        warn(f"JSON corrupted: {e}")
        return []


def save_tasks(tasks: list[Task]):
    data = {"tasks": [asdict(t) for t in tasks]}

    try:
        with open(data_file, "w", encoding="UTF-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)
    except OSError as e:
        warn(f"Error saving tasks: {e}")
