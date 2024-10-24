package task

import (
	"github.com/robfig/cron/v3"
	"golang-senior-learn/live/lesson1/cmd"
	"strings"
)

type Task struct {
	Spec    string `mapstructure:"Spec"`
	Commend string `mapstructure:"Commend"`
	Open    bool   `mapstructure:"Open"`
}

func (t Task) Run() {
	args := strings.Split(t.Commend, " ")
	cmd.RootCmd.SetArgs(args)
	cmd.RootCmd.Execute()
}

func AddTask(spec string, commend string, open bool) {
	tasks = append(tasks, Task{
		Spec:    spec,
		Commend: commend,
		Open:    open,
	})
}

var tasks []Task

func Start() {
	c := cron.New(cron.WithSeconds())
	for _, task := range tasks {
		if task.Open {
			_, err := c.AddJob(task.Spec, task)
			if err != nil {
				panic(err)
			}
		}
	}
	c.Run()
}
