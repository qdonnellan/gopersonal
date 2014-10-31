package essays

import (
    "fmt"
    "net/http"
)

// Handle the essays! Also, need to write a better doc comment...
func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello from the Essay folder!!!")
}