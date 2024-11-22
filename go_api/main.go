package main

import (
	"goapi/constants"
	"goapi/handler"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/generate/monster/beast", handler.RandomBeastMonsterHandler)

	http.HandleFunc("/generate/monster/human", handler.RandomHumanMonsterHandler)

	log.Println("Server running on http://127.0.0.1:" + strconv.Itoa(constants.APP_PORT))
	http.ListenAndServe((":" + strconv.Itoa(constants.APP_PORT)), nil)
}
