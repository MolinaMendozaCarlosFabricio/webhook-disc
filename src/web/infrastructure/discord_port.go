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
		return http.StatusInternalServerError
	}

	payload := map[string]string{"content": msg}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al serializar JSON: %v", err)
		return http.StatusInternalServerError
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error al enviar mensaje a Discord: %v", err)
		return http.StatusInternalServerError
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error al cerrar el cuerpo de la respuesta: %v", err)
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		return http.StatusOK
	default:
		log.Printf("Error al enviar mensaje, código: %d, respuesta: %s", resp.StatusCode, resp.Status)
		return http.StatusBadRequest
	}
}
