package main

import (
    "fmt"
    "flag"
    "strings"
    "net/http"
    "log"
    "os"
    "os/exec"
)

func getShell() string {
  var shell string = os.Getenv("SHELL")
  if shell == "" {
    shell = "/bin/sh"
  }
  return shell
}

func CreateExeHandler(exe string) func(http.ResponseWriter, *http.Request) {
    return (func(w http.ResponseWriter, r *http.Request){
      var err = exec.Command(getShell(), "-c", exe).Run()
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
    var (
      bind string
      portNum int
      url string
      exe string
    )

    flag.StringVar(&bind, "b", "", "Bind address.")
    flag.IntVar(&portNum, "p", 8080, "Listen port.")
    flag.StringVar(&url, "u", "/", "Handle URL.")
    flag.StringVar(&exe, "e", "true", "Execute command by $SHELL.")
    flag.Parse()

    var portStr = fmt.Sprint(portNum)
    log.Printf("Starting web server at %s:%s execute `%s` when handle path '%s'.", bind, portStr, exe, url)

    http.HandleFunc(url, CreateExeHandler(exe))
    e := http.ListenAndServe(bind + ":" + portStr, Log(http.DefaultServeMux))
    if e != nil {
      log.Panic(e)
    }
}
