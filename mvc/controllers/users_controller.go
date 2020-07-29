package controllers

import (
	"encoding/json"
	"github.com/uwaifo/introserver/mvc/services"
	"github.com/uwaifo/introserver/mvc/utils"
	"net/http"
	"strconv"
)

//call
//curl localhost:8800/users?user_id=123

func GetUser(rw http.ResponseWriter, req *http.Request) {

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		rw.WriteHeader(apiErr.StatusCode)
		rw.Write(jsonValue)
		return

	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue,_ := json.Marshal(apiErr)
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(jsonValue))
		return
	}

	jsonValue,_ := json.Marshal(user)
	rw.Write(jsonValue)
	//rw.Write([]byte("User is me "))

}
