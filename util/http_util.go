package util

import (
	"io/ioutil"
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func ParsePath(r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	var userId int
	var otherUserId int
	var err error
	if userId, err = strconv.Atoi(vars["user_id"]); err != nil {
		return 0, 0, err
	}
	if otherUserId, err = strconv.Atoi(vars["other_user_id"]); err != nil {
		return 0, 0, err
	}
	return userId, otherUserId, nil
}


func CheckInput(r *http.Request) ([]byte, error)  {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, err
	}
	if err := r.Body.Close(); err != nil {
		return nil, err
	}

	return body, nil
}
