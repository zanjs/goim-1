package connect

import (
	"goim/logic/db"
	"goim/public/context"
)

func Context() *context.Context {
	return context.NewContext(db.Factoty.GetSession())
}
