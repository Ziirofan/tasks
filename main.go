package main

// Task Main struct of task
type Task struct {
	Id     int    `json:"Id"`
	Date   string `json:"Date"`
	Text   string `json:"Text"`
	Done   bool   `json:"Done"`
	Status string `json:"Status"`
}

func main() {
	commandTask()
}
