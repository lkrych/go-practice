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

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
	yaml "gopkg.in/yaml.v2"
)

//templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

type config struct {
	GoogleClient string `yaml:"GOOGLE_CLIENT_ID"`
	GoogleSecret string `yaml:"GOOGLE_SECRET"`
	GithubClient string `yaml:"GITHUB_CLIENT_ID"`
	GithubSecret string `yaml:"GITHUB_SECRET"`
}

//serveHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
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
		google.New(c.GoogleClient, c.GoogleSecret, "http://localhost:8080/auth/callback/google"),
		github.New(c.GithubClient, c.GithubSecret, "http://localhost:8080/auth/callback/github"),
	)

	r := newRoom(UseGravatar)
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
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
	yamlFile, err := ioutil.ReadFile("secrets.yml")

	if err != nil {
		log.Printf("Error reading YAML: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
