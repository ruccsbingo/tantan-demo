package controller

import (
	"encoding/json"
	"net/http"
	"tantan-demo/store"
	"tantan-demo/model"
	"github.com/gorilla/mux"
	"strconv"
	"tantan-demo/util"
	"tantan-demo/service"
	"tantan-demo/store/sharding"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	accounts, err := store.ListUsers()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func ParseAccountData(r *http.Request) (*model.Account, error) {
	var account model.Account
	if body, err := util.CheckInput(r); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}
	return &account, nil
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	var account *model.Account
	var err error
	if account, err = ParseAccountData(r); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	store.CreateUsers(account)

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(account); err != nil {
		panic(err)
	}
}

func ListRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	vars := mux.Vars(r)
	var userId int
	var err error
	if userId, err = strconv.Atoi(vars["user_id"]); err != nil {
		panic(err)
	}

	//relations, err := store.ListRelations(userId)
	relations, err := sharding.ListRelations(userId)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		panic(err)
	}
}

func ParseRelationData(r *http.Request) (*model.Relation, error) {
	var relation model.Relation
	if body, err := util.CheckInput(r); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &relation); err != nil {
		return nil, err
	}
	return &relation, nil
}

func UpdateRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	//parse body
	var relation *model.Relation
	var err error
	if relation, err = ParseRelationData(r); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//parse path
	var userId int
	var otherUserId int
	if userId, otherUserId, err = util.ParsePath(r); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	service.UpdateRelation(userId, otherUserId, relation)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(relation); err != nil {
		panic(err)
	}
}

