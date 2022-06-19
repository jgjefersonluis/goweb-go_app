package main 

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hi everyone, i like %s!", r.URL.Path[1:])
	//www.seudominio.com/qualquercoisa
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

