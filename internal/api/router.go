package api

import (
	"encoding/json"
	"net/http"

	"github.com/sarch-3/stream-to-display/internal/domain"
	"github.com/sarch-3/stream-to-display/internal/service"
)

type Handler struct {
	playerService *service.PlayerService
}

func NewHandler(playerService *service.PlayerService) *Handler {
	return &Handler{playerService: playerService}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /api/video/add", h.AddVideo)
	// mux.HandleFunc("POST /api/video/playback", h.Playback)
	// mux.HandleFunc("POST /api/video/seek", h.Seek)
	// mux.HandleFunc("POST /api/video/skip", h.Skip)
	// mux.HandleFunc("POST /api/video/volume", h.Volume)
	// mux.HandleFunc("POST /api/video/speed", h.Speed)
	// mux.HandleFunc("POST /api/video/clear-queue", h.ClearQueue)

	return mux
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) AddVideo(w http.ResponseWriter, r *http.Request) {
	var req domain.AddVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res := h.playerService.AddVideo(req.Source)
	writeJSON(w, http.StatusOK, res)
}

// func (h *Handler) Playback(w http.ResponseWriter, r *http.Request) {
// 	writeJSON(w, http.StatusOK, h.playerService.Playback())
// }

// func (h *Handler) Seek(w http.ResponseWriter, r *http.Request) {
// 	var req domain.SeekRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	delta := req.Delta
// 	if delta == 0 && req.Seconds != 0 {
// 		delta = req.Seconds
// 	}

// 	writeJSON(w, http.StatusOK, h.playerService.Seek(delta))
// }

// func (h *Handler) Skip(w http.ResponseWriter, r *http.Request) {
// 	writeJSON(w, http.StatusOK, h.playerService.SkipVideo())
// }

// func (h *Handler) Volume(w http.ResponseWriter, r *http.Request) {
// 	var req domain.VolumeRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, h.playerService.AdjustVolume(req.Delta))
// }

// func (h *Handler) Speed(w http.ResponseWriter, r *http.Request) {
// 	var req domain.SpeedRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, h.playerService.AdjustSpeed(req.Delta))
// }

// func (h *Handler) ClearQueue(w http.ResponseWriter, r *http.Request) {
// 	writeJSON(w, http.StatusOK, h.playerService.ClearQueue())
// }

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
