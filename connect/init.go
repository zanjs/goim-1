package connect

import (
	"goim/logic/db"
	"goim/public/context"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Context() *context.Context {
	return context.NewContext(db.Factoty.GetSession())
}
