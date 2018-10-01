package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesteban19/edcomments/models"
)

func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error al convertir el mensjae: %s", err)
	}

	w.WriteHeader(m.Code)
	w.Write(j)
}
