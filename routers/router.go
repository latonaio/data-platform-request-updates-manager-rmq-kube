package routers

import (
	"data-platform-request-updates-manager-rmq-kube/cache"
	"data-platform-request-updates-manager-rmq-kube/config"
	controllersArticleCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/article-creates/all"
	controllersArticleDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/article-doc-creates/header-doc"
	controllersArticleUpdatesCounter "data-platform-request-updates-manager-rmq-kube/controllers/article-updates/counter"
	controllersArticleUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/article-updates/header"
	controllersArticleUpdatesLike "data-platform-request-updates-manager-rmq-kube/controllers/article-updates/like"
	controllersAttendanceCancelsAll "data-platform-request-updates-manager-rmq-kube/controllers/attendance-cancels/all"
	controllersAttendanceCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/attendance-creates/all"
	controllersAuthenticatorCreatesInitialEmailAuth "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/initial-email-auth"
	controllersAuthenticatorCreatesInitialGoogleAccountAuth "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/initial-google-account-auth"
	controllersAuthenticatorCreatesInitialSMSAuth "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/initial-sms-auth"
	controllersAuthenticatorCreatesInstagramAuth "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/instagram-auth"
	controllersAuthenticatorCreatesUser "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-creates/user"
	controllersAuthenticatorUpdatesInstagramAuth "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-updates/instagram-auth"
	controllersAuthenticatorUpdatesUser "data-platform-request-updates-manager-rmq-kube/controllers/authenticator-updates/user"
	controllersBillOfMaterialDetailList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialList "data-platform-request-updates-manager-rmq-kube/controllers/bill-of-material/list"
	controllersBusinessPartnerCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-creates/all"
	controllersBusinessPartnerCreatesRole "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-creates/role"
	controllersBusinessPartnerDocCreatesGeneralDoc "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-doc-creates/general-doc"
	controllersBusinessPartnerUpdatesAddress "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/address"
	controllersBusinessPartnerUpdatesGeneral "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/general"
	controllersBusinessPartnerUpdatesPerson "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person"
	controllersBusinessPartnerUpdatesPersonGoogleAccountAuth "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person-google-account-auth"
	controllersBusinessPartnerUpdatesPersonMobilePhoneAuth "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/person-mobile-phone-auth"
	controllersBusinessPartnerUpdatesRank "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/rank"
	controllersBusinessPartnerUpdatesSNS "data-platform-request-updates-manager-rmq-kube/controllers/business-partner-updates/sns"
	controllersDeliveryDocumentFunctionActualGoodsIssuePosting "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-actual-goods-issue-posting"
	controllersDeliveryDocumentFunctionActualGoodsReceiptPosting "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-actual-goods-receipt-posting"
	controllersDeliveryDocumentFunctionReferFromOrders "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/function-refer-from-orders"
	controllersDeliveryDocumentItemSingleUnit "data-platform-request-updates-manager-rmq-kube/controllers/delivery-document/item-single-unit"
	controllersEventCancelsHeader "data-platform-request-updates-manager-rmq-kube/controllers/event-cancels/header"
	controllersEventCancelsParticipation "data-platform-request-updates-manager-rmq-kube/controllers/event-cancels/participation"
	controllersEventCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/all"
	controllersEventCreatesAttendance "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/attendance"
	controllersEventCreatesParticipation "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/participation"
	controllersEventCreatesPointTransaction "data-platform-request-updates-manager-rmq-kube/controllers/event-creates/point-transaction"
	controllersEventDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/event-doc-creates/header-doc"
	controllersEventUpdatesCounter "data-platform-request-updates-manager-rmq-kube/controllers/event-updates/counter"
	controllersEventUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/event-updates/header"
	controllersEventUpdatesLike "data-platform-request-updates-manager-rmq-kube/controllers/event-updates/like"
	controllersFriendCreatesGeneralFriend "data-platform-request-updates-manager-rmq-kube/controllers/friend-creates/general/friend"
	controllersFriendCreatesGeneralMe "data-platform-request-updates-manager-rmq-kube/controllers/friend-creates/general/me"
	controllersGoogleAccountAuthKey "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/auth-key"
	controllersGoogleAccountAuthInitialAuth "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/initial-auth"
	controllersGoogleAccountAuthUserInfo "data-platform-request-updates-manager-rmq-kube/controllers/google-account-auth/user-info"
	controllersInstagramUserMediaList "data-platform-request-updates-manager-rmq-kube/controllers/instagram/user-media/list"
	controllersMessageCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/message-creates/all"
	controllersOrdersFunctionReferFromQuotations "data-platform-request-updates-manager-rmq-kube/controllers/orders/function-refer-from-quotations"
	controllersOrdersItemSingleUnit "data-platform-request-updates-manager-rmq-kube/controllers/orders/item-single-unit"
	controllersParticipationCancelsAll "data-platform-request-updates-manager-rmq-kube/controllers/participation-cancels/all"
	controllersParticipationCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/participation-creates/all"
	controllersParticipationDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/participation-doc-creates/header-doc"
	controllersPointBalanceCreatesPointBalance "data-platform-request-updates-manager-rmq-kube/controllers/point-balance-creates/point-balance"
	controllersPointBalanceUpdatesByShop "data-platform-request-updates-manager-rmq-kube/controllers/point-balance-updates/by-shop"
	controllersPointBalanceUpdatesPointBalance "data-platform-request-updates-manager-rmq-kube/controllers/point-balance-updates/point-balance"
	controllersPointTransactionCancelsHeader "data-platform-request-updates-manager-rmq-kube/controllers/point-transaction-cancels/header"
	controllersPointTransactionCreatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/point-transaction-creates/header"
	controllersPostCreatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/post-creates/header"
	controllersPostCreatesHeadersWithInstagramMedia "data-platform-request-updates-manager-rmq-kube/controllers/post-creates/headers-with-instagram-media"
	controllersPostDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/post-doc-creates/header-doc"
	controllersPostUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/post-updates/header"
	controllersProductionOrderItemOperation "data-platform-request-updates-manager-rmq-kube/controllers/production-order-confirmation/item-operation"
	controllersQRCodeGeneratesBusinessPartner "data-platform-request-updates-manager-rmq-kube/controllers/qr-code-generates/business-partner"
	controllersQRCodeGeneratesEvent "data-platform-request-updates-manager-rmq-kube/controllers/qr-code-generates/event"
	controllersQRCodeGeneratesParticipation "data-platform-request-updates-manager-rmq-kube/controllers/qr-code-generates/participation"
	controllersQRCodeGeneratesSite "data-platform-request-updates-manager-rmq-kube/controllers/qr-code-generates/site"
	controllersQuestionnaireCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/questionnaire-creates/all"
	controllersShopCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/shop-creates/all"
	controllersShopUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/shop-updates/header"
	controllersShopDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/shop-doc-creates/header-doc"
	controllersSiteCreatesAll "data-platform-request-updates-manager-rmq-kube/controllers/site-creates/all"
	controllersSiteDocCreatesHeaderDoc "data-platform-request-updates-manager-rmq-kube/controllers/site-doc-creates/header-doc"
	controllersSiteUpdatesCounter "data-platform-request-updates-manager-rmq-kube/controllers/site-updates/counter"
	controllersSiteUpdatesHeader "data-platform-request-updates-manager-rmq-kube/controllers/site-updates/header"
	controllersSiteUpdatesLike "data-platform-request-updates-manager-rmq-kube/controllers/site-updates/like"
	controllersSMSAuthToken "data-platform-request-updates-manager-rmq-kube/controllers/sms-auth/sms-auth-token"
	controllersSMSAuthTokenNotificationViaAWS "data-platform-request-updates-manager-rmq-kube/controllers/sms-auth/sms-auth-token-notification-via-aws"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"net/http"
)

type HealthCheckController struct {
	beego.Controller
}

func (c *HealthCheckController) Get() {
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = map[string]string{"status": "healthy"}
	c.ServeJSON()
}

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()

	redisTokenCacheKeyPrefix := "tokens"

	_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, &redisTokenCacheKeyPrefix)
	//redisTokenCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1)

	authenticatorCreatesUserController := &controllersAuthenticatorCreatesUser.AuthenticatorCreatesUserController{
		CustomLogger: l,
	}

	authenticatorCreatesInitialEmailAuthController := &controllersAuthenticatorCreatesInitialEmailAuth.AuthenticatorCreatesInitialEmailAuthController{
		CustomLogger: l,
	}

	authenticatorCreatesInitialSMSAuthController := &controllersAuthenticatorCreatesInitialSMSAuth.AuthenticatorCreatesInitialSMSAuthController{
		CustomLogger: l,
	}

	authenticatorCreatesInitialGoogleAccountAuthController := &controllersAuthenticatorCreatesInitialGoogleAccountAuth.AuthenticatorCreatesInitialGoogleAccountAuthController{
		CustomLogger: l,
	}

	authenticatorCreatesInstagramAuthController := &controllersAuthenticatorCreatesInstagramAuth.AuthenticatorCreatesInstagramAuthController{
		CustomLogger: l,
	}

	authenticatorUpdatesUserController := &controllersAuthenticatorUpdatesUser.AuthenticatorUpdatesUserController{
		CustomLogger: l,
	}

	authenticatorUpdatesInstagramAuthController := &controllersAuthenticatorUpdatesInstagramAuth.AuthenticatorUpdatesInstagramAuthController{
		CustomLogger: l,
	}

	businessPartnerCreatesAllController := &controllersBusinessPartnerCreatesAll.BusinessPartnerCreatesAllController{
		CustomLogger: l,
	}

	businessPartnerCreatesRoleController := &controllersBusinessPartnerCreatesRole.BusinessPartnerCreatesRoleController{
		CustomLogger: l,
	}

	businessPartnerDocCreatesGeneralDocController := &controllersBusinessPartnerDocCreatesGeneralDoc.BusinessPartnerDocCreatesGeneralDocController{
		CustomLogger: l,
	}

	businessPartnerUpdatesGeneralController := &controllersBusinessPartnerUpdatesGeneral.BusinessPartnerUpdatesGeneralController{
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

	businessPartnerUpdatesSNSController := &controllersBusinessPartnerUpdatesSNS.BusinessPartnerUpdatesSNSController{
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

	eventCreatesParticipationController := &controllersEventCreatesParticipation.EventCreatesParticipationController{
		CustomLogger: l,
	}

	eventCreatesAttendanceController := &controllersEventCreatesAttendance.EventCreatesAttendanceController{
		CustomLogger: l,
	}

	eventCreatesPointTransactionController := &controllersEventCreatesPointTransaction.EventCreatesPointTransactionController{
		CustomLogger: l,
	}

	eventDocCreatesHeaderDocController := &controllersEventDocCreatesHeaderDoc.EventDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	eventUpdatesHeaderController := &controllersEventUpdatesHeader.EventUpdatesHeaderController{
		CustomLogger: l,
	}

	eventUpdatesCounterController := &controllersEventUpdatesCounter.EventUpdatesCounterController{
		CustomLogger: l,
	}

	eventUpdatesLikeController := &controllersEventUpdatesLike.EventUpdatesLikeController{
		CustomLogger: l,
	}

	eventCancelsHeaderController := &controllersEventCancelsHeader.EventCancelsHeaderController{
		CustomLogger: l,
	}

	eventCancelsParticipationController := &controllersEventCancelsParticipation.EventCancelsParticipationController{
		CustomLogger: l,
	}

	friendCreatesGeneralMeController := &controllersFriendCreatesGeneralMe.FriendCreatesGeneralMeController{
		CustomLogger: l,
	}

	friendCreatesGeneralFriendController := &controllersFriendCreatesGeneralFriend.FriendCreatesGeneralFriendController{
		CustomLogger: l,
	}

	articleCreatesAllController := &controllersArticleCreatesAll.ArticleCreatesAllController{
		CustomLogger: l,
	}

	articleDocCreatesHeaderDocController := &controllersArticleDocCreatesHeaderDoc.ArticleDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	articleUpdatesHeaderController := &controllersArticleUpdatesHeader.ArticleUpdatesHeaderController{
		CustomLogger: l,
	}

	articleUpdatesCounterController := &controllersArticleUpdatesCounter.ArticleUpdatesCounterController{
		CustomLogger: l,
	}

	articleUpdatesLikeController := &controllersArticleUpdatesLike.ArticleUpdatesLikeController{
		CustomLogger: l,
	}

	siteCreatesAllController := &controllersSiteCreatesAll.SiteCreatesAllController{
		CustomLogger: l,
	}

	siteDocCreatesHeaderDocController := &controllersSiteDocCreatesHeaderDoc.SiteDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	siteUpdatesHeaderController := &controllersSiteUpdatesHeader.SiteUpdatesHeaderController{
		CustomLogger: l,
	}

	siteUpdatesCounterController := &controllersSiteUpdatesCounter.SiteUpdatesCounterController{
		CustomLogger: l,
	}

	siteUpdatesLikeController := &controllersSiteUpdatesLike.SiteUpdatesLikeController{
		CustomLogger: l,
	}

	shopCreatesAllController := &controllersShopCreatesAll.ShopCreatesAllController{
		CustomLogger: l,
	}

	shopDocCreatesHeaderDocController := &controllersShopDocCreatesHeaderDoc.ShopDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	shopUpdatesHeaderController := &controllersShopUpdatesHeader.ShopUpdatesHeaderController{
		CustomLogger: l,
	}

	postCreatesHeaderController := &controllersPostCreatesHeader.PostCreatesHeaderController{
		CustomLogger: l,
	}

	postCreatesHeadersWithInstagramMediaController := &controllersPostCreatesHeadersWithInstagramMedia.PostCreatesHeadersWithInstagramMediaController{
		CustomLogger: l,
	}

	postDocCreatesHeaderDocController := &controllersPostDocCreatesHeaderDoc.PostDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	postUpdatesHeaderController := &controllersPostUpdatesHeader.PostUpdatesHeaderController{
		CustomLogger: l,
	}

	pointTransactionCreatesHeaderController := &controllersPointTransactionCreatesHeader.PointTransactionCreatesHeaderController{
		CustomLogger: l,
	}

	pointTransactionCancelsHeaderController := &controllersPointTransactionCancelsHeader.PointTransactionCancelsHeaderController{
		CustomLogger: l,
	}

	pointBalanceUpdatesPointBalanceController := &controllersPointBalanceUpdatesPointBalance.PointBalanceUpdatesPointBalanceController{
		CustomLogger: l,
	}

	pointBalanceUpdatesByShopController := &controllersPointBalanceUpdatesByShop.PointBalanceUpdatesByShopController{
		CustomLogger: l,
	}

	pointBalanceCreatesPointBalanceController := &controllersPointBalanceCreatesPointBalance.PointBalanceCreatesPointBalanceController{
		CustomLogger: l,
	}

	participationCreatesAllController := &controllersParticipationCreatesAll.ParticipationCreatesAllController{
		CustomLogger: l,
	}
	participationDocCreatesHeaderDocController := &controllersParticipationDocCreatesHeaderDoc.ParticipationDocCreatesHeaderDocController{
		CustomLogger: l,
	}

	participationCancelsAllController := &controllersParticipationCancelsAll.ParticipationCancelsAllController{
		CustomLogger: l,
	}

	attendanceCreatesAllController := &controllersAttendanceCreatesAll.AttendanceCreatesAllController{
		CustomLogger: l,
	}

	attendanceCancelsAllController := &controllersAttendanceCancelsAll.AttendanceCancelsAllController{
		CustomLogger: l,
	}

	messageCreatesAllController := &controllersMessageCreatesAll.MessageCreatesAllController{
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

	instagramUserMediaListController := &controllersInstagramUserMediaList.InstagramUserMediaListController{
		CustomLogger: l,
	}

	questionnaireCreatesAllController := &controllersQuestionnaireCreatesAll.QuestionnaireCreatesAllController{
		CustomLogger: l,
	}

	sMSAuthTokenController := &controllersSMSAuthToken.SMSAuthTokenController{
		CustomLogger: l,
	}

	sMSAuthTokenNotificationViaAWSController := &controllersSMSAuthTokenNotificationViaAWS.SMSAuthTokenNotificationViaAWSController{
		CustomLogger: l,
	}

	qRCodeGeneratesBusinessPartnerController := &controllersQRCodeGeneratesBusinessPartner.QRCodeGeneratesBusinessPartnerController{
		CustomLogger: l,
	}

	qRCodeGeneratesEventController := &controllersQRCodeGeneratesEvent.QRCodeGeneratesEventController{
		CustomLogger: l,
	}

	qRCodeGeneratesSiteController := &controllersQRCodeGeneratesSite.QRCodeGeneratesSiteController{
		CustomLogger: l,
	}

	qRCodeGeneratesParticipationController := &controllersQRCodeGeneratesParticipation.QRCodeGeneratesParticipationController{
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

	authenticator := beego.NewNamespace(
		"/authenticator",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/user", authenticatorCreatesUserController),
		beego.NSRouter("/creates/initial-email-auth", authenticatorCreatesInitialEmailAuthController),
		beego.NSRouter("/creates/initial-sms-auth", authenticatorCreatesInitialSMSAuthController),
		beego.NSRouter("/creates/initial-google-account-auth", authenticatorCreatesInitialGoogleAccountAuthController),
		beego.NSRouter("/creates/instagram-auth", authenticatorCreatesInstagramAuthController),
		beego.NSRouter("/updates/user", authenticatorUpdatesUserController),
		beego.NSRouter("/updates/instagram-auth", authenticatorUpdatesInstagramAuthController),
	)

	businessPartner := beego.NewNamespace(
		"/business-partner",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", businessPartnerCreatesAllController),
		beego.NSRouter("/creates/role", businessPartnerCreatesRoleController),
		beego.NSRouter("/updates/general", businessPartnerUpdatesGeneralController),
		beego.NSRouter("/updates/person", businessPartnerUpdatesPersonController),
		beego.NSRouter("/updates/address", businessPartnerUpdatesAddressController),
		beego.NSRouter("/updates/rank", businessPartnerUpdatesRankController),
		beego.NSRouter("/updates/sns", businessPartnerUpdatesSNSController),
		beego.NSRouter("/updates/person-mobile-phone-auth", businessPartnerUpdatesPersonMobilePhoneAuthController),
		beego.NSRouter("/updates/person-google-account-auth", businessPartnerUpdatesPersonGoogleAccountAuthController),
	)

	businessPartnerDoc := beego.NewNamespace(
		"/business-partner-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/general-doc", businessPartnerDocCreatesGeneralDocController),
	)

	event := beego.NewNamespace(
		"/event",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", eventCreatesAllController),
		beego.NSRouter("/creates/participation", eventCreatesParticipationController),
		beego.NSRouter("/creates/attendance", eventCreatesAttendanceController),
		beego.NSRouter("/creates/point-transaction", eventCreatesPointTransactionController),
		beego.NSRouter("/updates/header", eventUpdatesHeaderController),
		beego.NSRouter("/updates/counter", eventUpdatesCounterController),
		beego.NSRouter("/updates/like", eventUpdatesLikeController),
		beego.NSRouter("/cancels/header", eventCancelsHeaderController),
		beego.NSRouter("/cancels/participation", eventCancelsParticipationController),
	)

	eventDoc := beego.NewNamespace(
		"/event-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", eventDocCreatesHeaderDocController),
	)

	article := beego.NewNamespace(
		"/article",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", articleCreatesAllController),
		beego.NSRouter("/updates/header", articleUpdatesHeaderController),
		beego.NSRouter("/updates/counter", articleUpdatesCounterController),
		beego.NSRouter("/updates/like", articleUpdatesLikeController),
	)

	articleDoc := beego.NewNamespace(
		"/article-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", articleDocCreatesHeaderDocController),
	)

	site := beego.NewNamespace(
		"/site",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", siteCreatesAllController),
		beego.NSRouter("/updates/header", siteUpdatesHeaderController),
		beego.NSRouter("/updates/counter", siteUpdatesCounterController),
		beego.NSRouter("/updates/like", siteUpdatesLikeController),
	)

	siteDoc := beego.NewNamespace(
		"/site-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", siteDocCreatesHeaderDocController),
	)

	shop := beego.NewNamespace(
		"/shop",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", shopCreatesAllController),
		beego.NSRouter("/updates/header", shopUpdatesHeaderController),
	)

	shopDoc := beego.NewNamespace(
		"/shop-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", shopDocCreatesHeaderDocController),
	)

	friend := beego.NewNamespace(
		"/friend",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/general/me", friendCreatesGeneralMeController),
		beego.NSRouter("/creates/general/friend", friendCreatesGeneralFriendController),
	)

	post := beego.NewNamespace(
		"/post",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header", postCreatesHeaderController),
		beego.NSRouter("/creates/headers-with-instagram-media", postCreatesHeadersWithInstagramMediaController),
		beego.NSRouter("/updates/header", postUpdatesHeaderController),
	)

	postDoc := beego.NewNamespace(
		"/post-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", postDocCreatesHeaderDocController),
	)

	pointTransaction := beego.NewNamespace(
		"/point-transaction",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header", pointTransactionCreatesHeaderController),
		beego.NSRouter("/cancels/header", pointTransactionCancelsHeaderController),
	)

	pointBalance := beego.NewNamespace(
		"/point-balance",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/point-balance", pointBalanceCreatesPointBalanceController),
		beego.NSRouter("/updates/point-balance", pointBalanceUpdatesPointBalanceController),
		beego.NSRouter("/updates/by-shop", pointBalanceUpdatesByShopController),
	)

	participation := beego.NewNamespace(
		"/participation",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", participationCreatesAllController),
		beego.NSRouter("/cancels/all", participationCancelsAllController),
	)

	participationDoc := beego.NewNamespace(
		"/participation-doc",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/header-doc", participationDocCreatesHeaderDocController),
	)

	attendance := beego.NewNamespace(
		"/attendance",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", attendanceCreatesAllController),
		beego.NSRouter("/cancels/all", attendanceCancelsAllController),
	)

	message := beego.NewNamespace(
		"/message",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", messageCreatesAllController),
	)

	googleAccountInitialAuth := beego.NewNamespace(
		"/google-account-initial-auth",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/initial-auth", googleAccountAuthInitialAuthController),
		beego.NSRouter("/auth-key", googleAccountAuthKeyController),
		beego.NSRouter("/user-info", googleAccountAuthUserInfoController),
	)

	questionnaire := beego.NewNamespace(
		"/questionnaire",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/creates/all", questionnaireCreatesAllController),
	)

	instagram := beego.NewNamespace(
		"/instagram",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/user-media/list", instagramUserMediaListController),
	)

	sMSAuth := beego.NewNamespace(
		"/sms-auth",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/sms-auth-token", sMSAuthTokenController),
		beego.NSRouter("/sms-auth-token-notification-via-aws", sMSAuthTokenNotificationViaAWSController),
	)

	qRCode := beego.NewNamespace(
		"/qr-code",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/generates/business-partner", qRCodeGeneratesBusinessPartnerController),
		beego.NSRouter("/generates/event", qRCodeGeneratesEventController),
		beego.NSRouter("/generates/site", qRCodeGeneratesSiteController),
		beego.NSRouter("/generates/participation", qRCodeGeneratesParticipationController),
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

	beego.AddNamespace(
		beego.NewNamespace("/api/updates").
			Namespace(
				authenticator,
				businessPartner,
				businessPartnerDoc,
				event,
				eventDoc,
				article,
				articleDoc,
				site,
				siteDoc,
				shop,
				shopDoc,
				friend,
				post,
				postDoc,
				pointTransaction,
				pointBalance,
				participation,
				participationDoc,
				attendance,
				message,
				googleAccountInitialAuth,
				instagram,
				questionnaire,
				sMSAuth,
				qRCode,
				orders,
				deliveryDocument,
				billOfMaterial,
				productionOrderConfirmation,
			),
	)

	beego.Router("/", &HealthCheckController{})

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
