package infrastructure

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"webhock-disc.com/w/src/web/application"
)

func HandlePullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("Evento recibido: %s (ID: %s)", eventType, deliveryID)

	rawData, err := ctx.GetRawData()
	if err != nil {
		log.Println("Error al leer el cuerpo de la solicitud")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
		return
	}

	var msg, webhookURL string
	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "pong"})
		return
	case "pull_request":
		msg, webhookURL = application.ProcessPullRequestEvent(rawData)
	case "workflow_run":
		msg, webhookURL = application.ProcessWorkflowRunEvent(rawData)
	default:
		log.Println("Evento no soportado:", eventType)
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento ignorado"})
		return
	}

	if msg == "ERROR" || webhookURL == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error procesando evento"})
		return
	}

	statusCode := postDiscord(msg, webhookURL)
	if statusCode == 200 {
		ctx.JSON(http.StatusOK, gin.H{"success": "Evento procesado con éxito"})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo al enviar mensaje a Discord"})
	}
}
