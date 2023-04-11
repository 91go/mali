package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/app"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/goods"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/life"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	DailyApiGroup   life.ApiGroup
	GoodsApiGroup   goods.ApiGroup
	LifeApiGroup    life.ApiGroup
	RssApiGroup     rss.ApiGroup
	AppApiGroup     app.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
