package route

import (
	"identify/backend/bootstrap"
	cc "identify/backend/controllers/card"
)

func Configure(b *bootstrap.Bootstrapper) {
	v1Card := b.App.Group("/identify/card")
	{
		cc.NewCard().Router(v1Card)
		if b.Config.App.Mode == "debug" {
			cc.NewTest().Router(v1Card)
		}
	}
}
