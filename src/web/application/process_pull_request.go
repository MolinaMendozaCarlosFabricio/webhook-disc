package application

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"webhock-disc.com/w/src/web/domain"
)

func ProcessPullRequestEvent(rawData []byte) (string, string) {
	var eventPayload domain.PullRequestEvent
	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al deserializar payload")
		return "ERROR", ""
	}

	log.Printf("Evento de pull request recibido: %s", eventPayload.Action)

	// base := eventPayload.PullRequest.Base.Ref
	titulo := eventPayload.PullRequest.Title
	repoFullName := eventPayload.Repository.FullName
	user := eventPayload.PullRequest.User.Login
	urlPullRequest := eventPayload.PullRequest.HTMLURL

	msg := fmt.Sprintf("Nuevo PR en **%s**\n **%s**\n %s\n %s", repoFullName, titulo, user, urlPullRequest)
	return msg, os.Getenv("DISCORD_DEV_WEBHOOK_URL")
}