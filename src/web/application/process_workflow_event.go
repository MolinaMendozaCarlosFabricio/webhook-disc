package application

import (
	"encoding/json"
	"fmt"
	"log"

	"webhock-disc.com/w/src/web/domain"
)

func ProcessWorkflowRunEvent(rawData []byte)string{
	var eventPayload domain.ActionPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al deserializar payload de workflow_run")
		return "ERROR"
	}

	status := "En progreso"
	switch eventPayload.WorkflowRun.Conclusion {
	case "success":
		status = "Exitoso"
	case "failure":
		status = "Fallido"
	case "cancelled":
		status = "Cancelado"
	case "skipped":
		status = "Saltado"
	case "timed_out":
		status = "Tiempo agotado"
	}

	msg := fmt.Sprintf("**Workflow:** %s\n Repositorio: %s\n Estado: %s\n [Ver detalles](%s)",
		eventPayload.WorkflowRun.Name, eventPayload.Repository.FullName, status, eventPayload.WorkflowRun.HTMLURL)

	return msg
}
