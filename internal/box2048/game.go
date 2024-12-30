package box2048

import "net/http"

type Game struct {
	board []int
	score int
}

func Get2048(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/root/internal/box2048/index.html")
}
