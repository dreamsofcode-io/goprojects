package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/dreamsofcode-io/scrape-me/internal/middleware"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Pages     []Page     `yaml:"pages"`
	Status    []Status   `yaml:"statuses"`
	Redirects []Redirect `yaml:"redirects"`
}

type Link struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Redirect struct {
	Path string `yaml:"path"`
	To   string `yaml:"to"`
}

type Page struct {
	Path  string `yaml:"path"`
	Title string `yaml:"title"`
	Links []Link `yaml:"links"`
	Extra *Link  `yaml:"extra"`
}

type Status struct {
	Path       string `yaml:"path"`
	StatusCode int    `yaml:"status"`
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func handleError(tmpl *template.Template, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		if wrapped.statusCode >= 400 {
			tmpl.ExecuteTemplate(w, "error.html", map[string]string{
				"ErrorMessage": http.StatusText(wrapped.statusCode),
				"Status":       strconv.Itoa(wrapped.statusCode),
			})
		}
	})
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	files := http.FileServer(http.Dir("./static"))

	r := http.NewServeMux()
	r.Handle("GET /static/", http.StripPrefix("/static", files))

	tmpl := template.Must(template.New("").ParseGlob("./templates/*"))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	cfgFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalln("failed to read config file")
	}

	var cfg Config
	if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
		log.Fatalln("failed to decode file", err)
	}

	for _, page := range cfg.Pages {
		r.HandleFunc(
			fmt.Sprintf("GET %s", page.Path),
			func(w http.ResponseWriter, r *http.Request) {
				tmpl.ExecuteTemplate(w, "index.html", page)
			},
		)
	}

	for _, status := range cfg.Status {
		r.HandleFunc(
			fmt.Sprintf("GET %s", status.Path),
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(status.StatusCode)
			},
		)
	}

	for _, redirect := range cfg.Redirects {
		r.HandleFunc(
			fmt.Sprintf("GET %s", redirect.Path),
			func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, redirect.To, http.StatusFound)
			},
		)
	}

	s := http.Server{
		Addr: ":8080",
		Handler: middleware.NoCache(
			middleware.Logging(logger, handleError(tmpl, r)),
		),
	}

	logger.Info("Server starting", slog.String("port", "8080"))

	s.ListenAndServe()
}
