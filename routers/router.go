package routers

import (
	"data-platform-request-updates-manager-rmq-kube/cache"
	"data-platform-request-updates-manager-rmq-kube/config"
	controllersAuthenticatorCreatesUser "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/user"
	controllersAuthenticatorUpdatesUser "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-updates/user"
	controllersBillOfMaterialDetailList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/list"
	controllersBusinessPartnerCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-creates/all"
	controllersBusinessPartnerCreatesRole "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-creates/role"
	controllersBusinessPartnerUpdatesAddress "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/address"
	controllersBusinessPartnerUpdatesPerson "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person"
	controllersBusinessPartnerUpdatesPersonGoogleAccountAuth "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person-google-account-auth"
	controllersBusinessPartnerUpdatesPersonMobilePhoneAuth "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person-mobile-phone-auth"
	controllersBusinessPartnerUpdatesRank "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/rank"
	controllersDeliveryDocumentFunctionActualGoodsIssuePosting "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-actual-goods-issue-posting"
	controllersDeliveryDocumentFunctionActualGoodsReceiptPosting "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-actual-goods-receipt-posting"
	controllersDeliveryDocumentFunctionReferFromOrders "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-refer-from-orders"
	controllersDeliveryDocumentItemSingleUnit "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/item-single-unit"
	controllersEventCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/all"
	controllersEventCreatesPointTransaction "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/point-transaction"
	controllersEventDocCreatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/event-doc-creates/header-doc"
	controllersEventUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/event-updates/header"
	controllersGoogleAccountAuthKey "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/auth-key"
	controllersGoogleAccountAuthInitialAuth "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/initial-auth"
	controllersGoogleAccountAuthUserInfo "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/user-info"
	controllersOrdersFunctionReferFromQuotations "data-platform-request-updates-manager-rmq-kube/controllers/orders/function-refer-from-quotations"
	controllersOrdersItemSingleUnit "data-platform-request-updates-manager-rmq-kube/controllers/orders/item-single-unit"
	controllersPointBalanceUpdatesPointBalance "data-platform-request-updates-manager-rmq-kube/controllers/point-balance-updates/point-balance"
	controllersPointTransactionCreatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/point-transaction-creates/header"
	controllersProductionOrderItemOperation "data-platform-request-updates-manager-rmq-kube/controllers/production-order-confirmation/item-operation"
	controllersSiteCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/site-creates/all"
	controllersSiteDocCreatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/site-doc-creates/header-doc"
	controllersSiteUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/site-updates/header"
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

	authenticatorCreatesUserController := &controllersAuthenticatorCreatesUser.AuthenticatorCreatesUserController{
		CustomLogger: l,
	}

	authenticatorUpdatesUserController := &controllersAuthenticatorUpdatesUser.AuthenticatorUpdatesUserController{
		CustomLogger: l,
	}

	businessPartnerCreatesAllController := &controllersBusinessPartnerCreatesAll.BusinessPartnerCreatesAllController{
		CustomLogger: l,
	}

	businessPartnerCreatesRoleController := &controllersBusinessPartnerCreatesRole.BusinessPartnerCreatesRoleController{
		CustomLogger: l,
	}

	businessPartnerUpdatesPersonController := &controllersBusinessPartnerUpdatesPerson.BusinessPartnerUpdatesPersonController{
		CustomLogger: l,
	}

	businessPartnerUpdatesAddressController := &controllersBusinessPartnerUpdatesAddress.BusinessPartnerUpdatesAddressController{
		CustomLogger: l,
	}

	businessPartnerUpdatesRankController := &controllersBusinessPartnerUpdatesRank.BusinessPartnerUpdatesRankController{
		CustomLogger: l,
	}

	businessPartnerUpdatesPersonMobilePhoneAuthController := &controllersBusinessPartnerUpdatesPersonMobilePhoneAuth.BusinessPartnerUpdatesPersonMobilePhoneAuthController{
		CustomLogger: l,
	}

	businessPartnerUpdatesPersonGoogleAccountAuthController := &controllersBusinessPartnerUpdatesPersonGoogleAccountAuth.BusinessPartnerUpdatesPersonGoogleAccountAuthController{
		CustomLogger: l,
	}

	eventCreatesAllController := &controllersEventCreatesAll.EventCreatesAllController{
		CustomLogger: l,
	}

	eventCreatesPointTransactionController := &controllersEventCreatesPointTransaction.EventCreatesPointTransactionController{
		CustomLogger: l,
	}

	eventDocCreatesHeaderController := &controllersEventDocCreatesHeader.EventDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	eventUpdatesHeaderController := &controllersEventUpdatesHeader.EventUpdatesHeaderController{
		CustomLogger: l,
	}

	siteCreatesAllController := &controllersSiteCreatesAll.SiteCreatesAllController{
		CustomLogger: l,
	}

	siteDocCreatesHeaderController := &controllersSiteDocCreatesHeader.SiteDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	siteUpdatesHeaderController := &controllersSiteUpdatesHeader.SiteUpdatesHeaderController{
		CustomLogger: l,
	}

	pointTransactionCreatesHeaderController := &controllersPointTransactionCreatesHeader.PointTransactionCreatesHeaderController{
		CustomLogger: l,
	}

	pointBalanceUpdatesPointBalanceController := &controllersPointBalanceUpdatesPointBalance.PointBalanceUpdatesPointBalanceController{
		CustomLogger: l,
	}

	ordersItemSingleUnitController := &controllersOrdersItemSingleUnit.OrdersItemSingleUnitController{
		CustomLogger: l,
	}

	ordersFunctionReferFromQuotationsController := &controllersOrdersFunctionReferFromQuotations.OrdersFunctionReferFromQuotationsController{
		CustomLogger: l,
	}

	deliveryDocumentItemSingleUnitController := &controllersDeliveryDocumentItemSingleUnit.DeliveryDocumentItemSingleUnitController{
		CustomLogger: l,
	}

	deliveryDocumentFunctionActualGoodsIssuePostingController := &controllersDeliveryDocumentFunctionActualGoodsIssuePosting.DeliveryDocumentFunctionActualGoodsIssuePostingController{
		CustomLogger: l,
	}

	deliveryDocumentFunctionActualGoodsReceiptPostingController := &controllersDeliveryDocumentFunctionActualGoodsReceiptPosting.DeliveryDocumentFunctionActualGoodsReceiptPostingController{
		CustomLogger: l,
	}

	deliveryDocumentFunctionReferFromOrdersController := &controllersDeliveryDocumentFunctionReferFromOrders.DeliveryDocumentFunctionReferFromOrdersController{
		CustomLogger: l,
	}

	billOfMaterialListController := &controllersBillOfMaterialList.BillOfMaterialListController{
		CustomLogger: l,
	}

	billOfMaterialDetailListController := &controllersBillOfMaterialDetailList.BillOfMaterialDetailListController{
		CustomLogger: l,
	}

	productionOrderConfirmationItemOperationController := &controllersProductionOrderItemOperation.ProductionOrderConfirmationItemOperationController{
		CustomLogger: l,
	}

	googleAccountAuthInitialAuthController := &controllersGoogleAccountAuthInitialAuth.GoogleAccountAuthInitialAuthController{
		CustomLogger: l,
	}

	googleAccountAuthKeyController := &controllersGoogleAccountAuthKey.GoogleAccountAuthKeyController{
		CustomLogger: l,
	}

	googleAccountAuthUserInfoController := &controllersGoogleAccountAuthUserInfo.GoogleAccountAuthUserInfoController{
		CustomLogger: l,
	}

	authenticator := beego.NewNamespace(
		"/authenticator",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/user", authenticatorCreatesUserController),
		beego.NSRouter("/updates/user", authenticatorUpdatesUserController),
	)

	businessPartner := beego.NewNamespace(
		"/business-partner",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", businessPartnerCreatesAllController),
		beego.NSRouter("/creates/role", businessPartnerCreatesRoleController),
		beego.NSRouter("/updates/person", businessPartnerUpdatesPersonController),
		beego.NSRouter("/updates/address", businessPartnerUpdatesAddressController),
		beego.NSRouter("/updates/rank", businessPartnerUpdatesRankController),
		beego.NSRouter("/updates/person-mobile-phone-auth", businessPartnerUpdatesPersonMobilePhoneAuthController),
		beego.NSRouter("/updates/person-google-account-auth", businessPartnerUpdatesPersonGoogleAccountAuthController),
	)

	event := beego.NewNamespace(
		"/event",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", eventCreatesAllController),
		beego.NSRouter("/creates/point-transaction", eventCreatesPointTransactionController),
		beego.NSRouter("/updates/header", eventUpdatesHeaderController),
	)

	eventDoc := beego.NewNamespace(
		"/event-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", eventDocCreatesHeaderController),
	)

	site := beego.NewNamespace(
		"/site",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", siteCreatesAllController),
		beego.NSRouter("/updates/header", siteUpdatesHeaderController),
	)

	siteDoc := beego.NewNamespace(
		"/site-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", siteDocCreatesHeaderController),
	)

	pointTransaction := beego.NewNamespace(
		"/point-transaction",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header", pointTransactionCreatesHeaderController),
	)

	pointBalance := beego.NewNamespace(
		"/point-balance",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/updates/point-balance", pointBalanceUpdatesPointBalanceController),
	)

	orders := beego.NewNamespace(
		"/orders",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/item/updates", ordersItemSingleUnitController),
		beego.NSRouter("/function-refer-from-quotations/creates", ordersFunctionReferFromQuotationsController),
	)

	deliveryDocument := beego.NewNamespace(
		"/delivery-document",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/item/updates", deliveryDocumentItemSingleUnitController),
		beego.NSRouter("/function-actual-goods-issue-posting/updates", deliveryDocumentFunctionActualGoodsIssuePostingController),
		beego.NSRouter("/function-actual-goods-receipt-posting/updates", deliveryDocumentFunctionActualGoodsReceiptPostingController),
		beego.NSRouter("/function-refer-from-orders/creates", deliveryDocumentFunctionReferFromOrdersController),
	)

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

	googleAccountInitialAuth := beego.NewNamespace(
		"/google-account-initial-auth",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/initial-auth", googleAccountAuthInitialAuthController),
		beego.NSRouter("/auth-key", googleAccountAuthKeyController),
		beego.NSRouter("/user-info", googleAccountAuthUserInfoController),
	)

	beego.AddNamespace(
		authenticator,
		businessPartner,
		event,
		eventDoc,
		site,
		siteDoc,
		pointTransaction,
		pointBalance,
		orders,
		deliveryDocument,
		billOfMaterial,
		productionOrderConfirmation,
		googleAccountInitialAuth,
	)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"}, // allow method only POST
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
