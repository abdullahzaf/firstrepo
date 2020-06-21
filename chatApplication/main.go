package main
import (
	"log"
	"net/http"
	"flag"
)
// temp1 represent a single template
type templateHandler struct {
	once		sync.once
	filename	string
	temp1		*template.Template
}
// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filespath.Join("templates", t.filename)))
}}
t.temp1.Execute(w, r)
}
func main() {

	// root
	http.Handle("/", &templateHandler{filename: "chat.html"})
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //parse the flags
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	//get the room going
	go r.run()
	//start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat!
				</body>
			</html>
		))
	})
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
