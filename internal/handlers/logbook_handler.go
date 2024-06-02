package handlers

import (
	"github.com/beatrizrdgs/literalog/internal/services"
)

type LogbookHandler struct {
	svc *services.LogbookService
}

func NewLogbookHandler(svc *services.LogbookService) *LogbookHandler {
	return &LogbookHandler{svc: svc}
}

// func (h *LogbookHandler) Add(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	req := new(models.LogbookRequest)
// 	json.NewDecoder(r.Body).Decode(req)
// 	logbook := models.NewLogbook(*req)
// 	err := h.svc.Add(ctx, logbook)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(logbook)
// }
