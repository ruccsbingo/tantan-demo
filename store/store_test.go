package store

import (
	"testing"
	"tantan-demo/model"
	"fmt"
	"tantan-demo/util"
)

func TestCreateUsers(t *testing.T) {
	util.Init()
	DeleteAccountRows()
	var account *model.Account = new(model.Account)
	account.Name = "sun"
	CreateUsers(account)
	fmt.Println(account.Id)
}

func TestQueryOneAccount(t *testing.T) {
	util.Init()
	DeleteAccountRows()
	account := &model.Account{Name: "dan"}
	another := *account

	CreateUsers(account)

	QueryOneAccount(account)

	if another.Name != account.Name {
		t.Error("QueryOneAccount is fail")
	}
}

func TestListUsers(t *testing.T) {
	util.Init()
	DeleteAccountRows()
	arr := []model.Account{
		{Name: "zhang"},
		{Name: "bing"},
	}

	for _, account := range arr {
		CreateUsers(&account)
	}

	accounts, _ := ListUsers()

	for i, account := range accounts {
		if account.Name != arr[i].Name {
			t.Error("i=", i, "account=", account, "arr[i]=", arr[i])
		}
	}
}