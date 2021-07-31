package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Moldy-Community/moldy/core/terminal"
	"github.com/Moldy-Community/moldy/utils/colors"
	"github.com/Moldy-Community/moldy/utils/functions"
	"github.com/spf13/cobra"
)

var (
	newTodoFlg, listFlg, selectDoneFlg, editTodoFlg, deleteTodoFlg bool
)

var filename string = "Moldy.todo.json"

type todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type allData []todo

var todoCmd = &cobra.Command{
	Use:     "todo",
	Short:   "Sort all of your tasks in a project",
	Long:    "Sort all of your task in a project.\nComplete it and if you want you can make commit in base of it: The most sorted and easier way to do commits",
	Aliases: []string{"todos", "td", "to-do", "task", "tasks", "ts"},
	Example: "moldy todo new",
	Run: func(cmd *cobra.Command, args []string) {
		var dataInFile allData
		if newTodoFlg {
			if !ExistsFile(filename) {
				CreateFile()
				dataInFile = append(dataInFile, CreateData(0))
				dataBytes, err := json.Marshal(dataInFile)
				functions.CheckErrors(err, "2", "Error saving the new to-do", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
				err = ioutil.WriteFile(filename, dataBytes, 0644)
				functions.CheckErrors(err, "2", "Error saving the new to-do", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
				colors.Success("The new todo was created successfully")
				return
			}

			file, err := ioutil.ReadFile(filename)
			functions.CheckErrors(err, "2", "Error reading the todo file", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			json.Unmarshal(file, &dataInFile)
			data := CreateData(len(dataInFile))
			dataInFile = append(dataInFile, data)
			dataBytes, err := json.Marshal(dataInFile)
			functions.CheckErrors(err, "2", "Error saving the new to-do", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			functions.CheckErrors(err, "2", "Error saving the new to-do", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			colors.Success("The new todo was created successfully")
		}
		file, err := ioutil.ReadFile(filename)
		functions.CheckErrors(err, "2", "Error reading the todo file", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
		json.Unmarshal(file, &dataInFile)
		if listFlg && !selectDoneFlg {
			for _, value := range dataInFile {
				fmt.Printf("%v. %v: %v | DONE %v\n", value.Id+1, value.Title, value.Description, value.Done)
			}
		}

		var titles []string
		for _, value := range dataInFile {
			titles = append(titles, fmt.Sprintf("%v. %v | DONE: %v", value.Id+1, value.Title, value.Done))
		}

		if selectDoneFlg {
			selected := terminal.SelectPrompt("Select a task to mark it done", titles)
			idSelected, err := strconv.Atoi(strings.Split(selected, "")[0])
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			idSelected += -1
			for i := 0; i < len(dataInFile); i++ {
				if dataInFile[i].Id == idSelected {
					dataInFile[i].Done = !dataInFile[i].Done
				}
			}
			dataBytes, err := json.Marshal(dataInFile)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
		}

		if editTodoFlg {
			selected := terminal.SelectPrompt("Select some task to edit", titles)
			idSelected, err := strconv.Atoi(strings.Split(selected, "")[0])
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			idSelected += -1
			for i := 0; i < len(dataInFile); i++ {
				if dataInFile[i].Id == idSelected {
					colors.Info("If some value is correct only press enter in this camp and the value will be the same\n")
					dataInFile[i].Title = terminal.BasicPrompt("Title", dataInFile[i].Title)
					dataInFile[i].Description = terminal.BasicPrompt("Description", dataInFile[i].Description)
				}
			}
			dataBytes, err := json.Marshal(dataInFile)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
		}

		if deleteTodoFlg {
			selected := terminal.SelectPrompt("Select some task to delete", titles)
			idSelected, err := strconv.Atoi(strings.Split(selected, "")[0])
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			idSelected += -1
			dataInFile = append(dataInFile[:idSelected], dataInFile[idSelected+1:]...)
			for i := idSelected; i < len(dataInFile); i++ {
				dataInFile[i].Id += -1
			}
			dataBytes, err := json.Marshal(dataInFile)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
			err = ioutil.WriteFile(filename, dataBytes, 0644)
			functions.CheckErrors(err, "2", "Error changing the done value", "Try again and if the problem persist leave the issue in github.com/Moldy-community/moldy")
		}
	},
}

func ExistsFile(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile() {
	_, err := os.Create(filename)
	functions.CheckErrors(err, "2", "error creating the Moldy.todo.json", "Try again and leave the issue in github.com/Moldy-community/moldy")
}

func CreateData(id int) todo {
	data := todo{
		Id:          id,
		Title:       terminal.BasicPrompt("Title", ""),
		Description: terminal.BasicPrompt("Description", ""),
		Done:        false,
	}
	return data
}

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.Flags().BoolVarP(&newTodoFlg, "new", "n", false, "Create a new todo")
	todoCmd.Flags().BoolVarP(&listFlg, "list", "l", false, "List all todo's")
	todoCmd.Flags().BoolVarP(&selectDoneFlg, "select", "s", false, "Select a task to mark it done or undone")
	todoCmd.Flags().BoolVarP(&editTodoFlg, "edit", "e", false, "Change values of a todo")
	todoCmd.Flags().BoolVarP(&deleteTodoFlg, "delete", "d", false, "Delete a todo")
}
