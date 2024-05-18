package notes

import (
	"log"
	"net/http"
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
	router.HandleFunc("GET /notes/{noteID}", func(w http.ResponseWriter, r *http.Request) {
		noteID := r.PathValue("noteID")
		w.Write([]byte("Note ID: " + noteID))
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/", http.StripPrefix("/api", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started %s", s.addr)
	return server.ListenAndServe()
}

/*
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
*/
