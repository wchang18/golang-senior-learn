package cron_learn

import (
	"github.com/robfig/cron/v3"
	"golang-senior-learn/chapter1/cobra_learn"
	"strings"
)

type Task struct {
	Spec    string `mapstructure:"Spec"`
	Commend string `mapstructure:"Commend"`
	Open    bool   `mapstructure:"Open"`
}

func (t Task) Run() {
	cmd := strings.Split(t.Commend, " ")
	cobra_learn.RootCmd.SetArgs(cmd)
	cobra_learn.RootCmd.Execute()
}

func Start() {
	c := cron.New(cron.WithSeconds())
	for _, task := range config.Scheduler {
		if task.Open {
			_, err := c.AddJob(task.Spec, task)
			if err != nil {
				panic(err)
			}
		}
	}
	c.Run()
}
