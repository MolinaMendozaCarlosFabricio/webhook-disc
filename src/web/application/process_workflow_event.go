package application

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"webhock-disc.com/w/src/web/domain"
)

func ProcessWorkflowRunEvent(rawData []byte) (string, string) {
	var eventPayload struct {
		Action      string `json:"action"`
		WorkflowRun struct {
			Name       string `json:"name"`
			Conclusion string `json:"conclusion"`
			HTMLURL    string `json:"html_url"`
		} `json:"workflow_run"`
		Repository domain.Repository `json:"repository"`
	}

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al deserializar payload de workflow_run")
		return "ERROR", ""
	}

	status := "Fallido"
	if eventPayload.WorkflowRun.Conclusion == "success" {
		status = "Exitoso"
	}

	msg := fmt.Sprintf("**Workflow:** %s\n📂 Repositorio: %s\n🚀 Estado: %s\n🔗 [Ver detalles](%s)",
		eventPayload.WorkflowRun.Name, eventPayload.Repository.FullName, status, eventPayload.WorkflowRun.HTMLURL)

	return msg, os.Getenv("DISCORD_TEST_WEBHOOK_URL")
}
