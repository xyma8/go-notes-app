package notes

import (
	"log"
	"net/http"

	"github.com/xyma8/go-notes-app/pkg/handler"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /api/notes/{noteID}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		noteID := r.PathValue("noteID")
		noteJSON, err := handler.GetNote(noteID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(noteJSON)
	})

	//v1 := http.NewServeMux()
	//v1.Handle("/api/", http.StripPrefix("/api", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}
