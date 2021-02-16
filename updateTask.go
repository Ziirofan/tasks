package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func updateTask(id string, txt string) {
	t := readTask(id)
	t.Text = txt

	jsonString, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("tasks_store/task"+id+".json", jsonString, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func doneTask(id string) {
	t := readTask(id)
	t.Done = true
	t.Status = "complete"

	jsonString, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("tasks_store/task"+id+".json", jsonString, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
