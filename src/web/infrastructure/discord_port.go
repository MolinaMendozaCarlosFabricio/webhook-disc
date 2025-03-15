package infrastructure

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func postDiscord(msg, webhookURL string) int {
	if webhookURL == "" {
		log.Println("Error: el link del webhook no está configurado")
		return 500
	}

	payload := map[string]string{"content": msg}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error al serializar JSON")
		return 500
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al enviar mensaje a Discord")
		return 500
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Printf("Error al enviar mensaje, código: %d", resp.StatusCode)
		return 400
	}

	return 200
}
