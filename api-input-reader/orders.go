package apiInputReader

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type OrdersSDC struct {
	ConnectionKey    string       `json:"connection_key"`
	Result           bool         `json:"result"`
	RedisKey         string       `json:"redis_key"`
	Filepath         string       `json:"filepath"`
	APIStatusCode    int          `json:"api_status_code"`
	RuntimeSessionID string       `json:"runtime_session_id"`
	BusinessPartner  *int         `json:"business_partner"`
	ServiceLabel     string       `json:"service_label"`
	APIType          string       `json:"api_type"`
	Header           OrdersHeader `json:"Orders"`
	APISchema        string       `json:"api_schema"`
	Accepter         []string     `json:"accepter"`
	Deleted          bool         `json:"deleted"`
}

type OrdersHeader struct {
	OrderID                          int              `json:"OrderID"`
	OrderDate                        *string          `json:"OrderDate"`
	OrderType                        *string          `json:"OrderType"`
	OrderStatus                      *string          `json:"OrderStatus"`
	SupplyChainRelationshipID        *int             `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int             `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int             `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            *int             `json:"Buyer"`
	Seller                           *int             `json:"Seller"`
	BillToParty                      *int             `json:"BillToParty"`
	BillFromParty                    *int             `json:"BillFromParty"`
	BillToCountry                    *string          `json:"BillToCountry"`
	BillFromCountry                  *string          `json:"BillFromCountry"`
	Payer                            *int             `json:"Payer"`
	Payee                            *int             `json:"Payee"`
	ContractType                     *string          `json:"ContractType"`
	OrderValidityStartDate           *string          `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string          `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string          `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string          `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   *float32         `json:"TotalNetAmount"`
	TotalTaxAmount                   *float32         `json:"TotalTaxAmount"`
	TotalGrossAmount                 *float32         `json:"TotalGrossAmount"`
	HeaderDeliveryStatus             *string          `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus              *string          `json:"HeaderBillingStatus"`
	HeaderDocReferenceStatus         *string          `json:"HeaderDocReferenceStatus"`
	TransactionCurrency              *string          `json:"TransactionCurrency"`
	PricingDate                      *string          `json:"PricingDate"`
	PriceDetnExchangeRate            *float32         `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            *string          `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime            *string          `json:"RequestedDeliveryTime"`
	HeaderCompleteDeliveryIsDefined  *bool            `json:"HeaderCompleteDeliveryIsDefined"`
	Incoterms                        *string          `json:"Incoterms"`
	PaymentTerms                     *string          `json:"PaymentTerms"`
	PaymentMethod                    *string          `json:"PaymentMethod"`
	Contract                         *int             `json:"Contract"`
	ContractItem                     *int             `json:"ContractItem"`
	ProductionVersion                *int             `json:"ProductionVersion"`
	ProductionVersionItem            *int             `json:"ProductionVersionItem"`
	ProductionOrder                  *int             `json:"ProductionOrder"`
	ProductionOrderItem              *int             `json:"ProductionOrderItem"`
	Operations                       *int             `json:"Operations"`
	OperationsItem                   *int             `json:"OperationsItem"`
	OperationID                      *int             `json:"OperationID"`
	ReferenceDocument                *int             `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int             `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           *string          `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32         `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              *string          `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool            `json:"IsExportImport"`
	HeaderText                       *string          `json:"HeaderText"`
	HeaderBlockStatus                *bool            `json:"HeaderBlockStatus"`
	HeaderDeliveryBlockStatus        *bool            `json:"HeaderDeliveryBlockStatus"`
	HeaderBillingBlockStatus         *bool            `json:"HeaderBillingBlockStatus"`
	ExternalReferenceDocument        *string          `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain        *string          `json:"CertificateAuthorityChain"`
	UsageControlChain                *string          `json:"UsageControlChain"`
	CreationDate                     *string          `json:"CreationDate"`
	CreationTime                     *string          `json:"CreationTime"`
	LastChangeDate                   *string          `json:"LastChangeDate"`
	LastChangeTime                   *string          `json:"LastChangeTime"`
	IsCancelled                      *bool            `json:"IsCancelled"`
	IsMarkedForDeletion              *bool            `json:"IsMarkedForDeletion"`
	Item                             *[]OrdersItem    `json:"Item"`
	Partner                          *[]OrdersPartner `json:"Partner"`
	Address                          *[]OrdersAddress `json:"Address"`
}

type OrdersItem struct {
	OrderID                                       int                        `json:"OrderID"`
	OrderItem                                     int                        `json:"OrderItem"`
	OrderItemCategory                             *string                    `json:"OrderItemCategory"`
	OrderStatus                                   *string                    `json:"OrderStatus"`
	SupplyChainRelationshipID                     *int                       `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int                       `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int                       `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int                       `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int                       `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                         *int                       `json:"Buyer"`
	Seller                                        *int                       `json:"Seller"`
	DeliverToParty                                *int                       `json:"DeliverToParty"`
	DeliverFromParty                              *int                       `json:"DeliverFromParty"`
	DeliverToPlant                                *string                    `json:"DeliverToPlant"`
	DeliverFromPlant                              *string                    `json:"DeliverFromPlant"`
	OrderItemText                                 *string                    `json:"OrderItemText"`
	OrderItemTextByBuyer                          *string                    `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                         *string                    `json:"OrderItemTextBySeller"`
	Product                                       *string                    `json:"Product"`
	SizeOrDimensionText                           *string                    `json:"SizeOrDimensionText"`
	ProductStandardID                             *string                    `json:"ProductStandardID"`
	ProductGroup                                  *string                    `json:"ProductGroup"`
	ProductSpecification                          *string                    `json:"ProductSpecification"`
	MarkingOfMaterial                             *string                    `json:"MarkingOfMaterial"`
	BaseUnit                                      *string                    `json:"BaseUnit"`
	DeliveryUnit                                  *string                    `json:"DeliveryUnit"`
	ProductionVersion                             *int                       `json:"ProductionVersion"`
	ProductionVersionItem                         *int                       `json:"ProductionVersionItem"`
	BillOfMaterial                                *int                       `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int                       `json:"BillOfMaterialItem"`
	ProductionOrder                               *int                       `json:"ProductionOrder"`
	ProductionOrderItem                           *int                       `json:"ProductionOrderItem"`
	Operations                                    *int                       `json:"Operations"`
	OperationsItem                                *int                       `json:"OperationsItem"`
	OperationID                                   *int                       `json:"OperationID"`
	PricingDate                                   *string                    `json:"PricingDate"`
	PriceDetnExchangeRate                         *float32                   `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                         *string                    `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                         *string                    `json:"RequestedDeliveryTime"`
	DeliverToPlantTimeZone                        *string                    `json:"DeliverToPlantTimeZone"`
	DeliverToPlantStorageLocation                 *string                    `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool                      `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string                    `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string                    `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string                    `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string                    `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string                    `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string                    `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantTimeZone                      *string                    `json:"DeliverFromPlantTimeZone"`
	DeliverFromPlantStorageLocation               *string                    `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool                      `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string                    `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string                    `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string                    `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string                    `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string                    `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string                    `json:"DeliverFromPlantBatchValidityEndTime"`
	StockConfirmationBusinessPartner              *int                       `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string                    `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                *string                    `json:"StockConfirmationPlantTimeZone"`
	ProductIsBatchManagedInStockConfirmationPlant *bool                      `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string                    `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string                    `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string                    `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string                    `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string                    `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string                    `json:"StockConfirmationPlantBatchValidityEndTime"`
	ServicesRenderingDate                         *string                    `json:"ServicesRenderingDate"`
	OrderQuantityInBaseUnit                       *float32                   `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit                   *float32                   `json:"OrderQuantityInDeliveryUnit"`
	QuantityPerPackage                            *float32                   `json:"QuantityPerPackage"`
	StockConfirmationPolicy                       *string                    `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string                    `json:"StockConfirmationStatus"`
	ConfirmedOrderQuantityInBaseUnit              *float32                   `json:"ConfirmedOrderQuantityInBaseUnit"`
	ProductWeightUnit                             *string                    `json:"ProductWeightUnit"`
	ProductNetWeight                              *float32                   `json:"ProductNetWeight"`
	ItemNetWeight                                 *float32                   `json:"ItemNetWeight"`
	ProductGrossWeight                            *float32                   `json:"ProductGrossWeight"`
	ItemGrossWeight                               *float32                   `json:"ItemGrossWeight"`
	InternalCapacityQuantity                      *float32                   `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string                    `json:"InternalCapacityQuantityUnit"`
	NetAmount                                     *float32                   `json:"NetAmount"`
	TaxAmount                                     *float32                   `json:"TaxAmount"`
	GrossAmount                                   *float32                   `json:"GrossAmount"`
	InvoiceDocumentDate                           *string                    `json:"InvoiceDocumentDate"`
	ProductionPlantBusinessPartner                *int                       `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string                    `json:"ProductionPlant"`
	ProductionPlantTimeZone                       *string                    `json:"ProductionPlantTimeZone"`
	ProductionPlantStorageLocation                *string                    `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool                      `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string                    `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string                    `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string                    `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string                    `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string                    `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string                    `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlantBusinessPartner                *int                       `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                               *string                    `json:"InspectionPlant"`
	InspectionPlan                                *int                       `json:"InspectionPlan"`
	InspectionLot                                 *int                       `json:"InspectionLot"`
	Incoterms                                     *string                    `json:"Incoterms"`
	TransactionTaxClassification                  *string                    `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         *string                    `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       *string                    `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      *string                    `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        *string                    `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 *string                    `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                                  *string                    `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string                    `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string                    `json:"PaymentDueDate"`
	NetPaymentDays                                *int                       `json:"NetPaymentDays"`
	PaymentMethod                                 *string                    `json:"PaymentMethod"`
	Contract                                      *int                       `json:"Contract"`
	ContractItem                                  *int                       `json:"ContractItem"`
	Project                                       *int                       `json:"Project"`
	WBSElement                                    *int                       `json:"WBSElement"`
	AccountingExchangeRate                        *float32                   `json:"AccountingExchangeRate"`
	ReferenceDocument                             *int                       `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int                       `json:"ReferenceDocumentItem"`
	ItemCompleteDeliveryIsDefined                 *bool                      `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus                            *string                    `json:"ItemDeliveryStatus"`
	IssuingStatus                                 *string                    `json:"IssuingStatus"`
	ReceivingStatus                               *string                    `json:"ReceivingStatus"`
	ItemBillingStatus                             *string                    `json:"ItemBillingStatus"`
	TaxCode                                       *string                    `json:"TaxCode"`
	TaxRate                                       *float32                   `json:"TaxRate"`
	CountryOfOrigin                               *string                    `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string                    `json:"CountryOfOriginLanguage"`
	Equipment                                     *int                       `json:"Equipment"`
	FreightAgreement                              *int                       `json:"FreightAgreement"`
	FreightAgreementItem                          *int                       `json:"FreightAgreementItem"`
	ItemBlockStatus                               *bool                      `json:"ItemBlockStatus"`
	ItemDeliveryBlockStatus                       *bool                      `json:"ItemDeliveryBlockStatus"`
	ItemBillingBlockStatus                        *bool                      `json:"ItemBillingBlockStatus"`
	ExternalReferenceDocument                     *string                    `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                 *string                    `json:"ExternalReferenceDocumentItem"`
	CreationDate                                  *string                    `json:"CreationDate"`
	CreationTime                                  *string                    `json:"CreationTime"`
	LastChangeDate                                *string                    `json:"LastChangeDate"`
	LastChangeTime                                *string                    `json:"LastChangeTime"`
	IsCancelled                                   *bool                      `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool                      `json:"IsMarkedForDeletion"`
	ItemPricingElement                            []OrdersItemPricingElement `json:"ItemPricingElement"`
	ItemScheduleLine                              []OrdersItemScheduleLine   `json:"ItemScheduleLine"`
}

type OrdersItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  *int     `json:"SupplyChainRelationshipID"`
	Buyer                      *int     `json:"Buyer"`
	Seller                     *int     `json:"Seller"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *int     `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
	CreationDate               *string  `json:"CreationDate"`
	CreationTime               *string  `json:"CreationTime"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	LastChangeTime             *string  `json:"LastChangeTime"`
	IsCancelled                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
}

type OrdersItemScheduleLine struct {
	OrderID                                         int      `json:"OrderID"`
	OrderItem                                       int      `json:"OrderItem"`
	ScheduleLine                                    int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID         *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         *string  `json:"Product"`
	StockConfirmationBussinessPartner               *int     `json:"StockConfirmationBussinessPartner"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone                  *string  `json:"StockConfirmationPlantTimeZone"`
	StockConfirmationPlantBatch                     *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate    *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityEndDate      *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	RequestedDeliveryDate                           *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                           *string  `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                           *string  `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryTime                           *string  `json:"ConfirmedDeliveryDate"`
	ScheduleLineOrderQuantityInBaseUnit             *float32 `json:"ScheduleLineOrderQuantityInBaseUnit"`
	OriginalOrderQuantityInBaseUnit                 *float32 `json:"OriginalOrderQuantityInBaseUnit"`
	ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedOrderQuantityByPDTAvailCheckInBaseUnit"`
	DeliveredQuantityInBaseUnit                     *float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit                   *float32 `json:"UndeliveredQuantityInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit                 *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	StockIsFullyConfirmed                           *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                                   *string  `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus             *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	ExternalReferenceDocument                       *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                   *string  `json:"ExternalReferenceDocumentItem"`
	ExternalReferenceDocumentItemScheduleLine       *string  `json:"ExternalReferenceDocumentItemScheduleLine"`
	CreationDate                                    *string  `json:"CreationDate"`
	CreationTime                                    *string  `json:"CreationTime"`
	LastChangeDate                                  *string  `json:"LastChangeDate"`
	LastChangeTime                                  *string  `json:"LastChangeTime"`
	IsCancelled                                     *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type OrdersAddress struct {
	OrderID     int     `json:"OrderID"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}

func OrdersInputRead(
	controller *beego.Controller,
) OrdersSDC {
	var requestBody OrdersSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
