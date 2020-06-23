package main
import (
	"log"
	"net/http"
	"flag"
	"path/filepath"
	"sync"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)
// templ represent a single template
type templateHandler struct {
	once		sync.Once
	filename	string
	templ		*template.Template
}
// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		t.once.Do(func() {
			t.templ = template.Must(template.ParseFiles(filespath.Join("templates", t.filename)))
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

	r := newRoom()
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFUnc("/auth/", loginHandler)
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //parse the flags
	//setup gomniauth
	gomniuath.SetSecurityKey("PUT YOUR AUTH KEY HERE")
	gomniauth.WithProviders(
		facebook.New("key", "secret",
			"http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret",
			"http://localhost:8080/auth/callback/github"),
		google.New("key", "secret",
			"http://localhost:8080/auth/callback/google"),
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
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
