package routers

import (
	"data-platform-request-updates-manager-rmq-kube/cache"
	"data-platform-request-updates-manager-rmq-kube/config"
	controllersBillOfMaterialDetailList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/list"
	controllersProductionOrderItemOperation "data-platform-request-updates-manager-rmq-kube/controllers/production-order-confirmation/item-operation"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()

	redisTokenCacheKeyPrefix := "tokens"

	_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, &redisTokenCacheKeyPrefix)
	//redisTokenCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1)

	billOfMaterialListController := &controllersBillOfMaterialList.BillOfMaterialListController{
		CustomLogger: l,
	}

	billOfMaterialDetailListController := &controllersBillOfMaterialDetailList.BillOfMaterialDetailListController{
		CustomLogger: l,
	}

	productionOrderConfirmationItemOperationController := &controllersProductionOrderItemOperation.ProductionOrderConfirmationItemOperationController{
		CustomLogger: l,
	}

	billOfMaterial := beego.NewNamespace(
		"/bill-of-material",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/header/updates", billOfMaterialListController),
		beego.NSRouter("/item/updates", billOfMaterialDetailListController),
	)

	productionOrderConfirmation := beego.NewNamespace(
		"/production-order-confirmation",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/header/updates", productionOrderConfirmationItemOperationController),
	)

	beego.AddNamespace(
		billOfMaterial,
		productionOrderConfirmation,
	)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"}, // allow method only POST
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		_ = ctx.Input.Header("Authorization")
		//jwtToken := ctx.Input.Header("Authorization")
		//
		//trimJwtToken := strings.TrimPrefix(jwtToken, "Bearer ")
		//
		//token, err := redisTokenCache.GetRaw(trimJwtToken)
		//
		//fmt.Sprintf("token: %v", token)
		//
		//if err == nil {
		//	return
		//}

		//services.VerifyToken(ctx, l, jwtToken)
	})

	//	beego.AddNamespace(billOfMaterial)

	//beego.Router("/:aPIServiceName/:aPIType", &controllers.APIModuleRuntimesController{})
	//beego.Router("/register", &controllers.RegisterController{})
	//beego.Router("/login", &controllers.LoginController{})
}
