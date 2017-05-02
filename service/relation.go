package service

import (
	"tantan-demo/util"
	"tantan-demo/model"
	//"tantan-demo/store"
	"tantan-demo/store/sharding"
)

func UpdateRelation(userId int, otherUserId int, relation *model.Relation)  {

	//construct relation
	relation.Id = util.RelationIdGen(userId, otherUserId)
	relation.UserId = userId
	relation.OtherUserId = otherUserId

	//construct opponent relation
	opponentRelation := model.Relation{
		UserId: relation.OtherUserId,
		OtherUserId: relation.UserId,
		State: "",
	}

	sharding.QueryOneRelation(&opponentRelation)
	//store.QueryOneRelation(&opponentRelation)
	if opponentRelation.State == "liked" && relation.State == "liked" {
		// update state = matched, then insert self relation
		relation.State = "matched"
		sharding.InsertRelations(relation)
		//store.InsertRelations(relation)

		// update opponent info, then insert opponent's relation
		opponentRelation.State = "matched"
		sharding.UpdateRelations(&opponentRelation)
		//store.UpdateRelations(&opponentRelation)
	} else {
		//insert self relation
		sharding.InsertRelations(relation)
		//store.InsertRelations(relation)
	}
}
