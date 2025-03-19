package domain

type ActionPayload struct {
	Action      string `json:"action"`
	WorkflowRun WorkflowRun `json:"workflow_run"`
	Repository  Repository `json:"repository"`
}