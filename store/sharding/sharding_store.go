package sharding

import (
	"tantan-demo/util"
	"tantan-demo/model"
	"fmt"
)

func ListRelations(userId int) (model.Relations, error) {
	db, tableName := util.CaculateDbAndTable(userId)
	var relations model.Relations
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?", tableName)
	_, err := db.Query(&relations, query, userId)
	if err != nil {
		panic(err)
	}
	return relations, err
}

func UpdateRelations(relation *model.Relation) (*model.Relation, error) {
	db, tableName := util.CaculateDbAndTable(relation.UserId)
	query := fmt.Sprintf("UPDATE %s SET state = ?state WHERE id = ?id", tableName)
	_, err := db.Exec(query, relation)
	if err != nil {
		panic(err)
	}
	return relation, err
}

func InsertRelations(relation *model.Relation) (*model.Relation, error) {
	db, tableName := util.CaculateDbAndTable(relation.UserId)
	query := fmt.Sprintf("INSERT INTO %s (id, user_id, other_user_id, state) VALUES (?id, ?user_id, ?other_user_id, ?state)", tableName)
	_, err := db.Exec(query, relation)
	if err != nil {
		panic(err)
	}
	return relation, err
}

func QueryOneRelation(relation *model.Relation) error {
	db, tableName := util.CaculateDbAndTable(relation.UserId)
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?user_id AND other_user_id = ?other_user_id", tableName)
	var relations model.Relations
	_, err := db.Query(&relations, query, relation)
	if err != nil {
		panic(err)
	}

	if len(relations) == 1 {
		relation.Id = relations[0].Id
		relation.State = relations[0].State
	}

	return nil
}

