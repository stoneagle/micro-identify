package route

import (
	"identify/backend/bootstrap"
	"identify/backend/common"
	cc "identify/backend/controllers/card"
)

func Configure(b *bootstrap.Bootstrapper) {
	v1Card := b.App.Group("/v1/card")
	{
		engine := common.GetEngine(b.Config.Card.Database.Name)
		cc.NewCard(engine).Router(v1Card)
		if b.Config.App.Mode == "debug" {
			cc.NewTest(engine).Router(v1Card)
		}
	}
}
