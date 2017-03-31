package main

import (
	"fmt"
	"time"

	kanbanflow "github.com/emil-nasso/kanbangow/kanbanflow"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	config()
	kanbanflow.KanbanflowAPIKey = kanbanflowAPIKey
}

func main() {
	r := gin.New()
	r.GET("/bug-free-days", bugFreeDays)
	r.Run()
}

func bugFreeDays(c *gin.Context) {
	bugInProgressColumns := append(preDevColumns, devColumns...)
	bugInProgressColumns = append(bugInProgressColumns, postDevColumns...)
	inProgressBugs := 0
	for _, columnID := range bugInProgressColumns {
		tasks := kanbanflow.GetTasksByColumnAndSwimlane(columnID, bugSwimlaneID)
		inProgressBugs += len(tasks)
	}

	lastCompletedTaskDate := "0000-00-00"

	for _, columnID := range []string{doneColumnID} {
		tasks := kanbanflow.GetTasksByColumnAndSwimlane(columnID, bugSwimlaneID)
		//Check if there are any tasks in the column and if it has date grouping
		if len(tasks) > 0 && len(tasks[0].DateGrouping) > 0 {
			lastCompletedTaskDate = tasks[0].DateGrouping
			break
		}
	}

	dateFormat := "2006-01-02"
	t, _ := time.Parse(dateFormat, lastCompletedTaskDate)
	bugFreeDuration := time.Now().Sub(t).Hours() / 24

	fmt.Println(lastCompletedTaskDate)

	c.JSON(200, gin.H{
		"in-progress-bugs-count":          inProgressBugs,
		"last-solved-bug-completion-date": lastCompletedTaskDate,
		"days-since-last-completed-bug":   bugFreeDuration,
	})
}
