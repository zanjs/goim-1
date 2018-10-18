package service

import (
	"goim/logic/db"
	"goim/logic/entity"
	"goim/public/context"
	"log"
	"testing"
)

var ctx = context.NewContext(db.Factoty.GetSession())

func TestFriendService_Add(t *testing.T) {
	add := entity.FriendAdd{
		UserId:      1,
		UserLabel:   "alber",
		Friend:      2,
		FriendLabel: "h",
	}
	err := FriendService.Add(ctx, add)
	if err != nil {
		log.Println(err)
	}
}

func TestFriendService_Delete(t *testing.T) {
	err := FriendService.Delete(ctx, 1, 2)
	if err != nil {
		log.Println(err)
	}
}
