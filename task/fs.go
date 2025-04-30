package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

// Gets json file path
func tasksFilePath() string {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current working directory: ", err)
		return ""
	}

	return path.Join(cwd, "tasks.json")
}

// Lists all tasks from json file
func ReadTasksFile() ([]Task, error) {
	filePath := tasksFilePath()
	log.SetPrefix("[Reading tasks] ")
	log.SetFlags(0)

	// Checks for file stats
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Println("Couldn't read save file, Generating new one..")

		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		// Handle creating error
		if err != nil {
			log.Println("Error creating save file: ", err)
			return nil, err
		}

		defer file.Close()
		log.Println("Save file created succesfully!")

		return []Task{}, nil
	}

	// Opening file if exists
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file: ", err)
		return nil, err
	}

	defer file.Close()

	// Initialize new tasks slice
	tasks := []Task{}

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		log.Println("Error decoding saved tasks:", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	log.SetPrefix("[Saving tasks]")
	log.SetFlags(0)

	filePath := tasksFilePath()

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}

	defer file.Close()

	// Encode file
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		log.Println("Error encoding file:", err)
		return err
	}

	return nil
}
