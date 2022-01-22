package main
import "net/http"

func main() {

	http.HandleFunc("/", Home)
	http.ListenAndServe(":3000", nil)

}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Olá Mundo, só que agora do k8s</h1>"))
}