package main

import (
    "fmt"
    "flag"
    "strconv"
    "strings"
    "net/http"
    "log"
    "os"
    "os/exec"
)

func CreateExeHandler(exe string) func(http.ResponseWriter, *http.Request) {
    return (func(w http.ResponseWriter, r *http.Request){
      var err = exec.Command(os.Getenv("SHELL"), "-c", exe).Run()
      if err == nil {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Execute `%s` succeed.", exe)
      } else {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Execute `%s` failed. %s.", exe, err.Error())
      }
    })
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
      info := []string{r.RemoteAddr, r.Proto, r.Method, r.RequestURI}
      log.Println(strings.Join(info, " "))

      handler.ServeHTTP(w, r)
    })
}

func main() {
    var b = flag.String("b", "", "Bind address.")
    var p = flag.Int("p", 8080, "Listen port.")
    var u = flag.String("u", "/", "Handle URL.")
    var e = flag.String("e", "true", "Execute command by $SHELL.")
    flag.Parse()

    var ps = strconv.Itoa(*p)
    log.Printf("Starting web server at %s:%s execute `%s` when handle path '%s'.", *b, ps, *e, *u)

    http.HandleFunc(*u, CreateExeHandler(*e))
    err := http.ListenAndServe(*b + ":" + ps, Log(http.DefaultServeMux))
    if err != nil {
      log.Panic(err)
    }
}
