package handler

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	si "github.com/ut-issl/wings/internal/core/service/interfaces"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, mux *chi.Mux) *http.Server {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			l, err := net.Listen("tcp", ":8080")
			if err != nil {
				log.Printf("Failed to listen on port 8080: %v", err)
				return err
			}
			go func() {
				log.Println("Starting HTTP server on :8080")
				if err := http.Serve(l, mux); err != nil {
					log.Printf("HTTP server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping HTTP server")
			return nil
		},
	})
	return &http.Server{
		Addr: ":8080",
	}
}

func Route(cs si.ICmdService) *chi.Mux {
	api := chi.NewRouter()

	api.Get("/queue", func(w http.ResponseWriter, r *http.Request) {
		// 実装が必要
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not implemented yet",
		})
	})

	api.Route("/api/operations/{ID}", func(r chi.Router) {
		// get all Cmds
		r.Get("/cmd", func(w http.ResponseWriter, r *http.Request) {
			operationID := chi.URLParam(r, "ID")
			if operationID == "" {
				http.Error(w, "Operation ID is required", http.StatusBadRequest)
				return
			}

			cmds, err := cs.GetAllCmds(operationID)
			if err != nil {
				log.Printf("Error getting commands: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(cmds); err != nil {
				log.Printf("Error encoding response: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		})
	})

	return api
}
