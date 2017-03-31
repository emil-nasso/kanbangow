package kanbanflow

// Task -  a kanbanflow task/card
type Task struct {
	ID                   string      `json:"_id"`
	Name                 string      `json:"name"`
	Position             int         `json:"position,omitempty"`
	Description          string      `json:"description"`
	Color                string      `json:"color"`
	ColumnID             string      `json:"columnId"`
	Number               *TaskID     `json:"number,omitempty"`
	ResponsibleUserID    string      `json:"responsibleUserId,omitempty"`
	TotalSecondsSpent    int         `json:"totalSecondsSpent"`
	TotalSecondsEstimate int         `json:"totalSecondsEstimate"`
	SwimlaneID           string      `json:"swimlaneId"`
	DateGrouping         string      `json:"groupingDate,omitempty"`
	Dates                []*TaskDate `json:"dates"`
	SubTasks             []*SubTask  `json:"subTasks"`
	Labels               []*Label    `json:"labels"`
}

// SubTask - A tasks subtask
type SubTask struct {
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}

// Label - a tasks label
type Label struct {
	Name   string `json:"name"`
	Pinned bool   `json:"pinned"`
}

// TaskDate - represent a data on a task, for example a duedate
type TaskDate struct {
	TargetColumnID    string `json:"targetColumnId"`
	Status            string `json:"status"`
	DateType          string `json:"dateType"`
	DueTimestamp      string `json:"dueTimestamp"`
	DueTimestampLocal string `json:"dueTimestampLocal"`
}

// TaskWebhook is the base struct for multiple webhooks
type TaskWebhook struct {
	EventType    string `json:"eventType"`
	UserID       string `json:"userId"`
	UserFullName string `json:"userFullName"`
	Timestamp    string `json:"timestamp"`
	Task         Task   `json:"task"`
}

// CreateTaskWebhook - The data sent from kanbanflow when creating a task
type CreateTaskWebhook struct {
	TaskWebhook
}

// ChangeTaskWebhook - The data sent from kanbanflow when changing a task
type ChangeTaskWebhook struct {
	TaskWebhook
	ChangedProperties []ChangedProperty `json:"changedProperties"`
}

// CommentCreateWebhook - New comment created webhook
type CommentCreateWebhook struct {
	EventType    string       `json:"eventType"`
	UserID       string       `json:"userId"`
	UserFullName string       `json:"userFullName"`
	Timestamp    string       `json:"timestamp"`
	TaskID       string       `json:"taskId"`
	TaskName     string       `json:"taskName"`
	TaskComment  *TaskComment `json:"taskComment"`
}

// TaskComment - a comment published on a task
type TaskComment struct {
	ID               string `json:"_id"`
	Text             string `json:"text"`
	AuthorUserID     string `json:"authorUserId"`
	CreatedTimestamp string `json:"createdTimestamp"`
}

// ChangedProperty represents a property that was changed on a task
type ChangedProperty struct {
	Property string `json:"property"`
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

// TaskID - A kanbanflow task id. It has an optional prefix.
// Example: { "value": 5 }, { "prefix": "BUG-", "value": 5 }
type TaskID struct {
	Value  int    `json:"value"`
	Prefix string `json:"prefix"`
}

// TaskListResponse - The response from a call to the "Get all tasks"-endpoint
// https://kanbanflow.com/api/v1/tasks
type TaskListResponse struct {
	ColumnID     string `json:"columnId"`
	ColumnName   string `json:"columnName"`
	TasksLimited bool   `json:"tasksLimited"`
	SwimlaneID   string `json:"swimlaneId"`
	SwimlaneName string `json:"swimlaneName"`
	Tasks        []Task `json:"tasks"`
}
