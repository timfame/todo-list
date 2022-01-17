package http

import (
	"github.com/gorilla/mux"
	"github.com/timfame/todo-list/internal/service"
	"html/template"
	"log"
	"net/http"
)

type server struct {
	router  *mux.Router
	service service.IService
	tmpl    *template.Template
}

func NewServer(service service.IService) (*server, error) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println("Parse files error:", err)
		return nil, err
	}
	s := &server{
		router:  mux.NewRouter(),
		service: service,
		tmpl:    tmpl,
	}
	s.initRoutes()
	return s, nil
}

func (s *server) initRoutes() {
	public := s.router.PathPrefix("").Subrouter()
	public.HandleFunc("/", s.index).Methods(http.MethodGet)
	public.Handle("/static/", http.StripPrefix("/static/",  http.FileServer(http.Dir("static"))))

	api := s.router.PathPrefix("/api").Subrouter()
	api.Use(s.indexRedirect)
	api.HandleFunc("/add-list", s.addList).Methods(http.MethodPost)
	api.HandleFunc("/remove-list", s.removeList).Methods(http.MethodPost)
	api.HandleFunc("/add-task", s.addTask).Methods(http.MethodPost)
	api.HandleFunc("/mark-task-done", s.markTaskDone).Methods(http.MethodPost)
}

func (s *server) Run() error {
	httpServer := &http.Server{
		Handler: s.router,
		Addr:    "localhost:8008",
	}
	return httpServer.ListenAndServe()
}

func (s *server) indexRedirect(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})
}
