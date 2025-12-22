import dataclasses as dc
from enum import StrEnum, auto


class TaskStatus(StrEnum):
    TODO = auto()
    IN_PROGRESS = auto()
    CANCELLED = auto()
    ARCHIVE = auto()
    DONE = auto()


@dc.dataclass
class Task:
    id: int
    title: str
    status: TaskStatus = TaskStatus.TODO

    @property
    def is_done(self) -> bool:
        return self.status in {
            TaskStatus.DONE,
            TaskStatus.ARCHIVE,
            TaskStatus.CANCELLED,
        }
