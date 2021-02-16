package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func writeTask(id int, text string) {

	dt := time.Now()

	data := Task{
		Id:     id,
		Date:   dt.Format("01-02-2006 15:04:05"),
		Text:   text,
		Done:   false,
		Status: "En cours",
	}

	jsonString, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("tasks_store/task"+strconv.Itoa(id)+".json", jsonString, os.ModePerm)
	if err != nil {
		panic(err)
	}

}
