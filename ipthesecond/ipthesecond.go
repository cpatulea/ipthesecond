package ipthesecond

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  w.Header().Set("Cache-Control", "private, max-age=0, no-cache")

  fmt.Fprint(w, r.RemoteAddr + "\n")
}
