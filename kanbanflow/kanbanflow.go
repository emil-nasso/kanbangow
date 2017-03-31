package kanbanflow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// KanbanflowAPIKey can be created on kanbanflows website
var KanbanflowAPIKey string

// GetTasksByColumnAndSwimlane returns a list of tasks in the specified column/swimlane
func GetTasksByColumnAndSwimlane(columnID, swimlaneID string) []Task {
	var taskResponse []TaskListResponse
	url := fmt.Sprintf(
		"https://kanbanflow.com/api/v1/tasks?swimlaneId=%s&columnId=%s",
		swimlaneID,
		columnID,
	)
	doRequest(url, &taskResponse)
	if len(taskResponse) != 1 {
		panic("Unexpected format when listing tasks by column and swimlan. Should be list with one item.")
	}
	tasks := taskResponse[0].Tasks
	return tasks
}

func getTaskByID(taskID string) *Task {
	url := fmt.Sprintf(
		"https://kanbanflow.com/api/v1/tasks/%s",
		taskID,
	)
	task := &Task{}
	doRequest(url, task)
	return task
}

// Returns the response data as a string if nil is provides if the
// response argument is nil
func doRequest(url string, response interface{}) string {
	client := &http.Client{}
	var err error
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.SetBasicAuth("apiToken", KanbanflowAPIKey)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if response == nil {
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			panic(readErr)
		}
		return string(body)
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		panic(err)
	}

	return ""
}
