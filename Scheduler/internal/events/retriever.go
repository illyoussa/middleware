package events

import (
	"fmt"
	"io"
	"net/http"
)

// FetchICalData récupère les données iCal depuis une URL donnée.
func FetchICalData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la récupération des données : %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("échec de la récupération des données : statut %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du corps de la réponse : %w", err)
	}

	return string(body), nil
}
