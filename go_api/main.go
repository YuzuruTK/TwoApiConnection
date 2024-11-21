package main

import (
	"goapi/constants"
	"goapi/handler"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/get-monster/beast", handler.RandomBeastMonsterHandler)

	http.HandleFunc("/get-monster/human", handler.RandomHumanMonsterHandler)

	http.ListenAndServe((":" + strconv.Itoa(constants.APP_PORT)), nil)
}
