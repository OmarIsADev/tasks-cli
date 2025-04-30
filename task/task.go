package task

import (
	"errors"
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in_progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          uint16     `json:"id"`
	TaskStatus  TaskStatus `json:"taskStatus"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
}

func NewTask(id uint16, description string) *Task {
	return &Task{
		ID:          id,
		TaskStatus:  TASK_STATUS_TODO,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFile()
	if err != nil {
		return nil
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet!, Try adding using the add command")
		return nil
	}

	filteredTasks := []Task{}

	switch status {
	case "all":
		filteredTasks = tasks

	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.TaskStatus == TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.TaskStatus == TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.TaskStatus == TASK_STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	if len(filteredTasks) == 0 {
		fmt.Printf("No %v tasks yet", status)

		return nil
	}
	for _, task := range filteredTasks {
		fmt.Println(task.ID, ":", task.Description, task.TaskStatus)
	}
	return nil
}

func AddTask(description string) error {
	tasks, err := ReadTasksFile()
	if err != nil {
		return err
	}

	var newTaskID uint16

	if len(tasks) > 0 {
		newTaskID = tasks[len(tasks)-1].ID + 1
	} else {
		newTaskID = 1
	}

	newTask := NewTask(newTaskID, description)

	tasks = append(tasks, *newTask)

	fmt.Println("Task added succesfully! task id: ", newTaskID)

	return WriteTasksToFile(tasks)
}

func DeleteTask(id uint16) error {
	tasks, err := ReadTasksFile()
	if err != nil {
		return err
	}

	var deletedTask Task
	newTasks := []Task{}

	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		} else {
			deletedTask = task
		}
	}

	if deletedTask.ID > 0 {
		return WriteTasksToFile(newTasks)
	}

	return errors.New("no task with this id")
}

func UpdateTaskDescription(id uint16, description string) error {
	tasks, err := ReadTasksFile()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			newTask := task
			newTask.Description = description

			tasks[i] = newTask

			return WriteTasksToFile(tasks)
		}
	}

	return errors.New("no task with this id")
}

func UpdateTaskStatus(id uint16, taskStatus TaskStatus) error {
	tasks, err := ReadTasksFile()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			newTask := task
			newTask.TaskStatus = taskStatus

			tasks[i] = newTask

			return WriteTasksToFile(tasks)
		}
	}

	return errors.New("no task with this id")
}
