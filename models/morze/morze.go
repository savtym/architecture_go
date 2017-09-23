package morze

import (
		"fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, decoder(r.URL.Path));
}

func decoder(buffer string) string {
    fmt.Println(buffer);
    return "morze";
}
