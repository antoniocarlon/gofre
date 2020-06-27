package gofre

import (
	"log"
	"strings"
)

type runFunction func(task *Task, c chan bool)
type checkFunction func(task *Task) bool

// Task is the main Gofre Task struct
type Task struct {
	ID        string
	Params    map[string]interface{}
	DependsOn map[string]Task
	Execute   runFunction
	Check     checkFunction
}

func (task *Task) getID() string {
	if (*task).ID != "" {
		return (*task).ID
	}
	return "noname"
}

func (task *Task) getParams() map[string]interface{} {
	return (*task).Params
}

func (task *Task) setParams(params map[string]interface{}) {
	(*task).Params = params
}

func (task *Task) doRun() bool {
	log.Println("Executing task " + (*task).getID())

	c := make(chan bool)
	go (*task).Execute(task, c)

	ok := false

	if <-c {
		message := []string{"The task", (*task).getID(), "run well \u2705"}

		if (*task).Check != nil {
			if (*task).Check(task) {
				message = append(message, "and the check went well \u2705  :)")
				ok = true
			} else {
				message = append(message, "but the check didn't go well \u274C  :(")
				ok = false
			}
		} else {
			message = append(message, "but there's no check function \u2754  :|")
			ok = true
		}

		log.Println(strings.Join(message, " "))
	} else {
		log.Println("The task " + (*task).getID() + " didn't run well \u274C  :(")
		ok = false
	}

	return ok
}

func (task *Task) isAlreadyExecuted() bool {
	return (*task).Check != nil && (*task).Check(task)
}

// Run is the receiver function that executes the Gofre task
func (task *Task) Run() bool {
	if (*task).isAlreadyExecuted() {
		log.Println("The task " + (*task).ID + " is already executed, skipping \u2705")
		return true
	}

	if (*task).DependsOn != nil {
		for _, childTask := range (*task).DependsOn {
			if ok := childTask.Run(); !ok {
				return false
			}
		}
	}

	return (*task).doRun()
}
