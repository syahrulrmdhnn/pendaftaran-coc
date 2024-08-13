package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/config"
	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/models"
)

func AmbilHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nama := vars["nama"]
	kunci := vars["kunci"]

	n := os.Getenv("APP_NAMA")
	k := os.Getenv("APP_KUNCI")

	if nama != n || kunci != k {
		http.Error(w, `{"message":"Tidak memiliki akses!"}`, http.StatusUnauthorized)
		return
	}

	var pendaftar []models.Pendaftar
	config.DB.Find(&pendaftar)

	response := `{"message":"success","data":{`
	for i, p := range pendaftar {
		if i > 0 {
			response += ","
		}
		response += fmt.Sprintf(`"%d":{"nama_lengkap":"%s","email":"%s","telepon":"%s","bukti_transfer":"/static/%s"}`, p.ID, p.NamaLengkap, p.Email, p.NoTelp, p.BuktiTransfer)
	}
	response += `}}`
	w.Write([]byte(response))
}