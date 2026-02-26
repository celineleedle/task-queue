package queue

type Stats struct {
	NumTasks      int `json:"total_tasks"`
	NumPending    int `json:"num_pending"`
	NumProcessing int `json:"num_processing"`
	NumCompleted  int `json:"num_completed"`
	NumFailed     int `json:"num_failed"`
}
