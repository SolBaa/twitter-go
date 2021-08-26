package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/models"
)

func SendTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweets
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error()+" Error en los datos recibidos", 400)
		return
	}

	register := models.Tweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)

	if err != nil {
		http.Error(w, err.Error()+"Error al enviar el tweet intente nuevamente", 400)
		return
	}

	if !status {
		http.Error(w, "Error al enviar el tweet.", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
