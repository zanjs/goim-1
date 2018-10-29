package connect

import (
	"goim/logic/db"
	"goim/public/ctx"
)

func Context() *ctx.Context {
	return ctx.NewContext(db.Factoty.GetSession())
}
