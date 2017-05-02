package store

import (
	//"gopkg.in/pg.v3"
	"tantan-demo/model"
	//"time"
	"tantan-demo/util"
)


func DeleteAccountRows() error {
	_, err := util.Db.Exec(`delete from account`)
	return err
}

func DeleteRelationRows() error {
	_, err := util.Db.Exec(`delete from relation`)
	return err
}

func QueryOneAccount(account *model.Account) error {
	if _, err := util.Dbs[0].QueryOne(account,`SELECT * FROM account WHERE name = ?`, account.Name); err != nil {
		return err
	}
	return nil
}


func ListUsers() (model.Accounts, error) {
	var accounts model.Accounts
	_, err := util.Dbs[0].Query(&accounts, `SELECT * FROM account`)
	if err != nil {
		panic(err)
	}
	return accounts, err
}

func CreateUsers(account *model.Account) (*model.Account, error) {
	_, err := util.Dbs[0].QueryOne(account, `
		INSERT INTO account (name, type) VALUES (?name, ?type)
		RETURNING id
		`, account)
	if err != nil {
		panic(err)
	}
	return account, err
}

func ListRelations(userId int) (model.Relations, error) {
	var relations model.Relations
	_, err := util.Db.Query(&relations, `SELECT * FROM relation WHERE user_id = ?`, userId)
	if err != nil {
		panic(err)
	}
	return relations, err
}

func UpdateRelations(relation *model.Relation) (*model.Relation, error) {
	_, err := util.Db.Exec(`UPDATE relation SET state = ?state WHERE id = ?id`, relation)
	if err != nil {
		panic(err)
	}
	return relation, err
}

func InsertRelations(relation *model.Relation) (*model.Relation, error) {
	_, err := util.Db.Exec(`INSERT INTO relation (id, user_id, other_user_id, state) VALUES (?id, ?user_id, ?other_user_id, ?state)`, relation)
	if err != nil {
		panic(err)
	}
	return relation, err
}

func QueryOneRelation(relation *model.Relation) error {
	var relations model.Relations
	_, err := util.Db.Query(&relations, `
		SELECT * FROM relation WHERE user_id = ?user_id AND other_user_id = ?other_user_id`, relation)
	if err != nil {
		panic(err)
	}

	if len(relations) == 1 {
		relation.Id = relations[0].Id
		relation.State = relations[0].State
	}

	return nil
}


