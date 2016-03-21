package main

import (
    "fmt"
    "flag"
    "strconv"
    "strings"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    info := []string{r.RemoteAddr, r.Proto, r.Method, r.RequestURI}
    log.Println(strings.Join(info, " "))

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
}

func main() {
    var b = flag.String("b", "", "Bind address.")
    var p = flag.Int("p", 8080, "Listen port.")
    var u = flag.String("u", "/", "Handle URL.")
    flag.Parse()

    var ps = strconv.Itoa(*p)
    log.Printf("Starting web server at %s:%s handle url %s", *b, ps, *u)

    http.HandleFunc(*u, handler)
    http.ListenAndServe(*b + ":" + ps, nil)
}
