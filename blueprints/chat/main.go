package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"trace"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
)

//templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type config struct {
	googleClient string `yaml:"GOOGLE_CLIENT_ID"`
	googleSecret string `yaml:"GOOGLE_SECRET"`
	githubClient string `yaml:"GITHUB_CLIENT_ID"`
	githubSecret string `yaml:"GOOGLE_SECRET"`
}

//serveHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	//parse secrets
	var c config
	c.getConfig()
	//setup gomniauth
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(c.googleClient, c.googleSecret, "http://localhost:8080/auth/callback/google"),
		github.New(c.githubClient, c.githubSecret, "http://localhost:8080/auth/callback/github"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	fmt.Printf("Server listening on port %v \n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

//room is run on separate goroutine so that the chatting operations occur in the background, allowing our main
//goroutine to run the web server.

func (c *config) getConfig() *config {
	yaml, err := ioutil.ReadFile("secrets.yml")
	if err != nil {
		log.Printf("Error reading YAML: %v", err)
	}
	err = yaml.Unmarshal(yaml, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
