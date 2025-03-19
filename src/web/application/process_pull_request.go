package application

import (
	"encoding/json"
	"fmt"
	"log"

	"webhock-disc.com/w/src/web/domain"
)

func ProcessPullRequestEvent(rawData []byte)string{
	var eventPayload domain.PullRequestEvent
	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al deserializar payload")
		return "ERROR"
	}

	log.Printf("Evento de pull request recibido: %s", eventPayload.Action)

	// base := eventPayload.PullRequest.Base.Ref
	baseBranch := eventPayload.PullRequest.Base.Ref
	headBranch := eventPayload.PullRequest.Head.Ref
	titulo := eventPayload.PullRequest.Title
	repoFullName := eventPayload.Repository.FullName
	user := eventPayload.PullRequest.User.Login
	urlPullRequest := eventPayload.PullRequest.HTMLURL

	var msg string

	switch eventPayload.Action {
	case "opened":
		msg = fmt.Sprintf("**Nuevo PR abierto en** `%s`\n **Título:** %s\n **Autor:** %s\n **De:** `%s` → `%s`\n [Ver Pull Request](%s)", 
			repoFullName, titulo, user, urlPullRequest, baseBranch, headBranch)
	case "closed":
		msg = fmt.Sprintf("**PR cerrado en** `%s`\n **Título:** %s\n **Autor:** %s\n **De:** `%s` → `%s`\n [Ver Pull Request](%s)", 
			repoFullName, titulo, user, urlPullRequest, baseBranch, headBranch)
	case "merged":
		msg = fmt.Sprintf("**PR MERGEADO en** `%s`\n **Título:** %s\n **Autor:** %s\n **De:** `%s` → `%s`\n [Ver Pull Request](%s)", 
			repoFullName, titulo, user, urlPullRequest, baseBranch, headBranch)
	default:
		msg = fmt.Sprintf("**PR actualizado en** `%s`\n **Título:** %s\n **Autor:** %s\n [Ver Pull Request](%s)", 
			repoFullName, titulo, user, urlPullRequest)
	}
	return msg
}