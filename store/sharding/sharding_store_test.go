package sharding

import (
	"testing"
	"tantan-demo/util"
	"tantan-demo/model"
)

func TestListRelations(t *testing.T) {
	util.InitDbs()

	ListRelations(23)
}

func TestInsertRelations(t *testing.T) {

	util.InitDbs()

	var relation model.Relation
	relation.Id = util.RelationIdGen(1, 2)
	relation.UserId = 1
	relation.OtherUserId = 2
	InsertRelations(&relation)

	relation.Id = util.RelationIdGen(2, 3)
	relation.UserId = 2
	relation.OtherUserId = 3
	InsertRelations(&relation)
}
