package htmlcontentservice

import (
	"github.com/akosgarai/go_akos_httpd/game"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Htmlpage interface {
	Getpage() string
}

type Game interface {
	Render() string
}

type Service struct {
	addr   string
	ln     net.Listener
	logger *log.Logger

	htmlpage Htmlpage
	game     Game
}

// New returns an uninitialized HTTP service.
func New(addr string, htmlpage Htmlpage) *Service {
	return &Service{
		addr:     addr,
		htmlpage: htmlpage,
		logger:   log.New(os.Stderr, "[httpContent] ", log.LstdFlags),
	}
}

// Start starts the service.
func (s *Service) Start() error {
	server := http.Server{
		Handler: s,
	}

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.ln = ln

	http.Handle("/page", s)

	go func() {
		err := server.Serve(s.ln)
		if err != nil {
			s.logger.Fatalf("HTTP serve: %s", err)
		}
	}()

	return nil
}

// Close closes the service.
func (s *Service) Close() error {
	s.ln.Close()
	return nil
}

// ServeHTTP allows Service to serve HTTP requests.
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Println(r.Method, r.URL)
	if strings.HasPrefix(r.URL.Path, "/page") {
		s.handleKeyRequest(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) handleKeyRequest(w http.ResponseWriter, r *http.Request) {
	getKey := func() string {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			return "default"
		}
		return parts[2]
	}
	getRowNum := func() int {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 {
			return 3
		}
		i, err := strconv.Atoi(parts[3])
		if err == nil {
			return i
		} else {
			return 3
		}
	}
	getColNum := func() int {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 5 {
			return 3
		}
		i, err := strconv.Atoi(parts[4])
		if err == nil {
			return i
		} else {
			return 3
		}
	}

	switch r.Method {
	case "GET":
		k := getKey()
		if k == "default" {
			io.WriteString(w, string(s.htmlpage.Getpage()))
			return
		}
		if k == "game" {
			s.game = game.New(getRowNum(), getColNum())
			io.WriteString(w, string(s.game.Render()))
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// Addr returns the address on which the Service is listening
func (s *Service) Addr() net.Addr {
	return s.ln.Addr()
}
