# Task Tracker CLI Tool

https://roadmap.sh/projects/task-tracker
A simple CLI tool for managing tasks.

## Build & run

Linux:
```bash
go build . & ./tasks-cli
```
Windows: 
```bash
go build . ; ./tasks-cli.exe 
```

## Features

### List Tasks

* List all tasks
* Filter tasks by status (todo, in-progress, done)

### Add Tasks

* Add new tasks with description

### Update Tasks

* Update task description
* Update task status

### Delete Tasks

* Delete tasks by id

## Usage

### List Tasks

* `task-cli list`
* `task-cli list todo`
* `task-cli list in-progress`
* `task-cli list done`

### Add Tasks

* `task-cli add <description>`

### Update Tasks

* `task-cli update description <id> <new description>`
* `task-cli update status <id> <new status>`
* `task-cli update status <id> <new status code>`
	+ `1` - done
	+ `2` - in-progress
	+ `3` - todo

### Delete Tasks

* `task-cli delete <id>`
