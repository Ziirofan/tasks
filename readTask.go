package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cheynewallace/tabby"
	"github.com/i582/cfmt"
)

func readTask(id string) Task {

	jsonFile, err := os.Open("tasks_store/task" + id + ".json")
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var task Task

	json.Unmarshal(byteValue, &task)

	return task
}

func displayTask(task Task) {

	cfmt.Printf(`
    {{                             }}::bgGreen
    {{       Tasks Ongoing         }}::bgGreen|#ffffff
    {{                             }}::bgGreen
	`)

	t := tabby.New()
	t.AddHeader("Id", "Text", "Date", "Status")
	t.AddLine(strconv.Itoa(task.Id), task.Text, task.Date, task.Status)
	t.Print()
}
