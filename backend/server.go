package notes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xyma8/go-notes-app/pkg/handler"
	models "github.com/xyma8/go-notes-app/pkg/models"
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

	router.HandleFunc("OPTIONS /api/notes/add", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("POST /api/notes/add", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		log.Printf("api/notes/add")

		var noteDto models.NoteDto
		err := json.NewDecoder(r.Body).Decode(&noteDto)
		if err != nil {
			http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
			return
		}

		noteUUID, err := handler.AddNote(noteDto)

		if err != nil {
			log.Printf("/notes/add " + err.Error())
			http.Error(w, "Ошибка при добавлении новой заметки", http.StatusInternalServerError)
			return
		}

		w.Write(noteUUID)
	})

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
