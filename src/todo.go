package main

import "time"

type Todo struct {
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due_time"`
}

type Todos []Todo