package declarations

type Task struct {
	ID          string `json:"id"`
	TaskName    string `json:"taskName"`
	TaskDetails string `json:"taskDetails"`
	CreatedDate string `json:"createdDate"`
}
