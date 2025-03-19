package infrastructure

import (
	"log"
	"net/http"
	"os"

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

	var msg string
	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "pong"})
		return
	case "pull_request":
		msg = application.ProcessPullRequestEvent(rawData)
	case "push":
		msg = application.ProcessPullRequestEvent(rawData)
	default:
		log.Println("Evento no soportado:", eventType)
		ctx.JSON(http.StatusOK, gin.H{"status": "Evento ignorado"})
		return
	}

	webhookURL := os.Getenv("DISCORD_DEV_WEBHOOK_URL")

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
