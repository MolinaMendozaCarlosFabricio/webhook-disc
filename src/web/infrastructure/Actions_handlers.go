package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"webhock-disc.com/w/src/web/application"
)

func HandleGithubActionEvent(ctx *gin.Context){
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("Evento recibido: %s (ID: %s)", eventType, deliveryID)

	if eventType != "workflow_run" {
		ctx.JSON(http.StatusContinue, gin.H{"Message": "Evento no correspondiente a una Action"})
		return
	}

	rawData, err := ctx.GetRawData()
	if err != nil {
		log.Println("Error al leer el cuerpo de la solicitud")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer datos"})
		return
	}

	msg := application.ProcessWorkflowRunEvent(rawData)

	webhookURL := os.Getenv("DISCORD_TEST_WEBHOOK_URL")

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