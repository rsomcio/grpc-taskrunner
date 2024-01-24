package main

import (
	"fmt"

	"github.com/rsomcio/grpc-taskrunner/cmd"
	"github.com/rsomcio/grpc-taskrunner/queue"
)

type Task struct {
	Name string
	Cmd  string
	Args []string
}

func main() {

	q := queue.New()
	q.Push(&Task{Name: "test1", Cmd: "./bin/test.sh", Args: []string{""}})
    q.Push(&Task{Name: "test2", Cmd: "curl", Args: []string{"-s", "ifconfig.me"}})
	q.Push(&Task{Name: "test3"})
	q.Push(&Task{Name: "test4"})
	q.Push(&Task{Name: "test5"})
	q.Push(&Task{Name: "test6"})

	for q.IsEmpty() == false {
		task := q.Pop()
		values := task.(*Task)
		if values.Cmd == "" {
			continue
		}
		fmt.Printf("Name: %s Cmd: %+v\n", values.Name, values.Cmd)
		cmd.Run(values.Cmd, values.Args)
	}
}
