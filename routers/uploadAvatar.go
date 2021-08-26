package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("avatar")
	ext := strings.Split(handler.Filename, ".")[1]
	arch := "uploads/avatar/" + IDUser + "." + ext
	f, err := os.OpenFile(arch, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{}

	user.Avatar = IDUser + "." + ext

	status, err := db.ChangeRegister(user, IDUser)

	if err != nil || !status {
		http.Error(w, "Error al al guardar el avatar en la DB "+err.Error(), http.StatusBadRequest)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
