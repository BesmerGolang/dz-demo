package verify

import (
	"dz3/configs"
	"dz3/internal/verify/email"
	"dz3/internal/verify/storage"
	"encoding/json"
	"net/http"
)

type EmailRequest struct {
	Email string `json:"email"`
}
type Handler struct {
	Config *configs.EmailConfig
}

func (h *Handler) SendHandler(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Email == "" {
		http.Error(w, "Invalid request: missing email", http.StatusBadRequest)
		return
	}

	token, err := storage.GenerateToken()
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}
	if err := storage.SaveToken(token, req.Email); err != nil {
		http.Error(w, "failed to save data", http.StatusInternalServerError)
		return
	}
	err = email.SendVerificationEmail(req.Email, token, h.Config)
	if err != nil {
		http.Error(w, "fail;ed to send email", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("verefication email sent"))
}
func (h *Handler) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash") //вытаскиваем значение из URL по имени параметра\
	if hash == "" {
		http.Error(w, "missing hash", http.StatusBadRequest)
		return
	}
	isValid := storage.CheckToken(hash)
	if !isValid {
		http.Error(w, "invalid token", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email successfully verified !"))
}
