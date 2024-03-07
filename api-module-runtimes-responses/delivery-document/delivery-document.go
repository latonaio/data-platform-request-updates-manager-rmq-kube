package apiModuleRuntimesResponsesDeliveryDocument

type DeliveryDocumentRes struct {
	Message DeliveryDocument `json:"message,omitempty"`
}

type DeliveryDocument struct {
	Header *[]Header `json:"Header,omitempty"`
	Item   *[]Item   `json:"Item,omitempty"`
}

type Header struct {
	DeliveryDocument                       int      `json:"DeliveryDocument"`
	DeliveryDocumentDate                   string   `json:"DeliveryDocumentDate"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int     `json:"SupplyChainRelationshipPaymentID"`
	SupplyChainRelationshipFreightID       *int     `json:"SupplyChainRelationshipFreightID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	BillToParty                            *int     `json:"BillToParty"`
	BillFromParty                          *int     `json:"BillFromParty"`
	BillToCountry                          *string  `json:"BillToCountry"`
	BillFromCountry                        *string  `json:"BillFromCountry"`
	Payer                                  *int     `json:"Payer"`
	Payee                                  *int     `json:"Payee"`
	FreightPartner                         *int     `json:"FreightPartner"`
	IsExportImport                         *bool    `json:"IsExportImport"`
	DeliverToPlantTimeZone                 *string  `json:"DeliverToPlantTimeZone"`
	DeliverFromPlantTimeZone               *string  `json:"DeliverFromPlantTimeZone"`
	ReferenceDocument                      *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                  *int     `json:"ReferenceDocumentItem"`
	OrderID                                *int     `json:"OrderID"`
	OrderItem                              *int     `json:"OrderItem"`
	Contract                               *int     `json:"Contract"`
	ContractItem                           *int     `json:"ContractItem"`
	Project                                *int     `json:"Project"`
	WBSElement                             *int     `json:"WBSElement"`
	ProductionVersion                      *int     `json:"ProductionVersion"`
	ProductionVersionItem                  *int     `json:"ProductionVersionItem"`
	ProductionOrder                        *int     `json:"ProductionOrder"`
	ProductionOrderItem                    *int     `json:"ProductionOrderItem"`
	Operations                             *int     `json:"Operations"`
	OperationsItem                         *int     `json:"OperationsItem"`
	OperationID                            *int     `json:"OperationID"`
	BillOfMaterial                         *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                     *int     `json:"BillOfMaterialItem"`
	ContractType                           *string  `json:"ContractType"`
	OrderValidityStartDate                 *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string  `json:"OrderValidityEndDate"`
	PlannedGoodsIssueDate                  string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                  string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                string   `json:"PlannedGoodsReceiptTime"`
	FreightOrder                           *int     `json:"FreightOrder"`
	InvoiceDocumentDate                    *string  `json:"InvoiceDocumentDate"`
	HeaderCompleteDeliveryIsDefined        *bool    `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus                   *string  `json:"HeaderDeliveryStatus"`
	GoodsIssueOrReceiptSlipNumber          *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus                    *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus                *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockStatus               *bool    `json:"HeaderBillingBlockStatus"`
	HeaderGrossWeight                      *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight                        *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit                       *string  `json:"HeaderWeightUnit"`
	Incoterms                              *string  `json:"Incoterms"`
	TransactionCurrency                    *string  `json:"TransactionCurrency"`
	HeaderDeliveryBlockStatus              *bool    `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus               *bool    `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus             *bool    `json:"HeaderReceivingBlockStatus"`
	ExternalReferenceDocument              *string  `json:"ExternalReferenceDocument"`
	CertificateAuthorityChain              *string  `json:"CertificateAuthorityChain"`
	UsageControlChain                      *string  `json:"UsageControlChain"`
	CreationDate                           string   `json:"CreationDate"`
	CreationTime                           string   `json:"CreationTime"`
	LastChangeDate                         string   `json:"LastChangeDate"`
	LastChangeTime                         string   `json:"LastChangeTime"`
	IsCancelled                            *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                    *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	DeliveryDocument                              int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                          int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory                  string   `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID                     int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       int      `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	SupplyChainRelationshipBillingID              *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID              *int     `json:"SupplyChainRelationshipPaymentID"`
	SupplyChainRelationshipFreightID              *int     `json:"SupplyChainRelationshipFreightID"`
	Buyer                                         int      `json:"Buyer"`
	Seller                                        int      `json:"Seller"`
	DeliverToParty                                int      `json:"DeliverToParty"`
	DeliverFromParty                              int      `json:"DeliverFromParty"`
	DeliverToPlant                                string   `json:"DeliverToPlant"`
	DeliverFromPlant                              string   `json:"DeliverFromPlant"`
	BillToParty                                   *int     `json:"BillToParty"`
	BillFromParty                                 *int     `json:"BillFromParty"`
	BillToCountry                                 *string  `json:"BillToCountry"`
	BillFromCountry                               *string  `json:"BillFromCountry"`
	Payer                                         *int     `json:"Payer"`
	Payee                                         *int     `json:"Payee"`
	FreightPartner                                *int     `json:"FreightPartner"`
	Product                                       string   `json:"Product"`
	SizeOrDimensionText                           *string  `json:"SizeOrDimensionText"`
	ProductStandardID                             *string  `json:"ProductStandardID"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	ProductSpecification                          *string  `json:"ProductSpecification"`
	MarkingOfMaterial                             *string  `json:"MarkingOfMaterial"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	OriginalQuantityInBaseUnit                    float32  `json:"OriginalQuantityInBaseUnit"`
	OriginalQuantityInDeliveryUnit                float32  `json:"OriginalQuantityInDeliveryUnit"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	ProductIsBatchManagedInDeliverToPlant         *bool    `json:"ProductIsBatchManagedInDeliverToPlant"`
	BatchMgmtPolicyInDeliverToPlant               *string  `json:"BatchMgmtPolicyInDeliverToPlant"`
	DeliverToPlantBatch                           *string  `json:"DeliverToPlantBatch"`
	DeliverToPlantBatchValidityStartDate          *string  `json:"DeliverToPlantBatchValidityStartDate"`
	DeliverToPlantBatchValidityStartTime          *string  `json:"DeliverToPlantBatchValidityStartTime"`
	DeliverToPlantBatchValidityEndDate            *string  `json:"DeliverToPlantBatchValidityEndDate"`
	DeliverToPlantBatchValidityEndTime            *string  `json:"DeliverToPlantBatchValidityEndTime"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	ProductIsBatchManagedInDeliverFromPlant       *bool    `json:"ProductIsBatchManagedInDeliverFromPlant"`
	BatchMgmtPolicyInDeliverFromPlant             *string  `json:"BatchMgmtPolicyInDeliverFromPlant"`
	DeliverFromPlantBatch                         *string  `json:"DeliverFromPlantBatch"`
	DeliverFromPlantBatchValidityStartDate        *string  `json:"DeliverFromPlantBatchValidityStartDate"`
	DeliverFromPlantBatchValidityStartTime        *string  `json:"DeliverFromPlantBatchValidityStartTime"`
	DeliverFromPlantBatchValidityEndDate          *string  `json:"DeliverFromPlantBatchValidityEndDate"`
	DeliverFromPlantBatchValidityEndTime          *string  `json:"DeliverFromPlantBatchValidityEndTime"`
	StockConfirmationBusinessPartner              int      `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        string   `json:"StockConfirmationPlant"`
	ProductIsBatchManagedInStockConfirmationPlant *bool    `json:"ProductIsBatchManagedInStockConfirmationPlant"`
	BatchMgmtPolicyInStockConfirmationPlant       *string  `json:"BatchMgmtPolicyInStockConfirmationPlant"`
	StockConfirmationPlantBatch                   *string  `json:"StockConfirmationPlantBatch"`
	StockConfirmationPlantBatchValidityStartDate  *string  `json:"StockConfirmationPlantBatchValidityStartDate"`
	StockConfirmationPlantBatchValidityStartTime  *string  `json:"StockConfirmationPlantBatchValidityStartTime"`
	StockConfirmationPlantBatchValidityEndDate    *string  `json:"StockConfirmationPlantBatchValidityEndDate"`
	StockConfirmationPlantBatchValidityEndTime    *string  `json:"StockConfirmationPlantBatchValidityEndTime"`
	StockConfirmationPolicy                       *string  `json:"StockConfirmationPolicy"`
	StockConfirmationStatus                       *string  `json:"StockConfirmationStatus"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation                *string  `json:"ProductionPlantStorageLocation"`
	ProductIsBatchManagedInProductionPlant        *bool    `json:"ProductIsBatchManagedInProductionPlant"`
	BatchMgmtPolicyInProductionPlant              *string  `json:"BatchMgmtPolicyInProductionPlant"`
	ProductionPlantBatch                          *string  `json:"ProductionPlantBatch"`
	ProductionPlantBatchValidityStartDate         *string  `json:"ProductionPlantBatchValidityStartDate"`
	ProductionPlantBatchValidityStartTime         *string  `json:"ProductionPlantBatchValidityStartTime"`
	ProductionPlantBatchValidityEndDate           *string  `json:"ProductionPlantBatchValidityEndDate"`
	ProductionPlantBatchValidityEndTime           *string  `json:"ProductionPlantBatchValidityEndTime"`
	InspectionPlantBusinessPartner                *int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                               *string  `json:"InspectionPlant"`
	InspectionPlan                                *int     `json:"InspectionPlan"`
	InspectionLot                                 *int     `json:"InspectionLot"`
	DeliveryDocumentItemText                      string   `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer               string   `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller              string   `json:"DeliveryDocumentItemTextBySeller"`
	PlannedGoodsIssueDate                         string   `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime                         string   `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate                       string   `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime                       string   `json:"PlannedGoodsReceiptTime"`
	PlannedGoodsIssueQuantity                     float32  `json:"PlannedGoodsIssueQuantity"`
	PlannedGoodsIssueQtyInBaseUnit                float32  `json:"PlannedGoodsIssueQtyInBaseUnit"`
	PlannedGoodsReceiptQuantity                   float32  `json:"PlannedGoodsReceiptQuantity"`
	PlannedGoodsReceiptQtyInBaseUnit              float32  `json:"PlannedGoodsReceiptQtyInBaseUnit"`
	ActualGoodsIssueDate                          *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                          *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                        *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                        *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                      *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit                 *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity                    *float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit               *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	QuantityPerPackage                            *float32 `json:"QuantityPerPackage"`
	ItemBillingStatus                             *string  `json:"ItemBillingStatus"`
	ItemCompleteDeliveryIsDefined                 *bool    `json:"ItemCompleteDeliveryIsDefined"`
	ProductWeightUnit                             *string  `json:"ProductWeightUnit"`
	ProductNetWeight                              *float32 `json:"ProductNetWeight"`
	ItemWeightUnit                                *string  `json:"ItemWeightUnit"`
	ItemNetWeight                                 *float32 `json:"ItemNetWeight"`
	ItemGrossWeight                               *float32 `json:"ItemGrossWeight"`
	InternalCapacityQuantity                      *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit                  *string  `json:"InternalCapacityQuantityUnit"`
	ItemIsBillingRelevant                         *bool    `json:"ItemIsBillingRelevant"`
	NetAmount                                     *float32 `json:"NetAmount"`
	TaxAmount                                     *float32 `json:"TaxAmount"`
	GrossAmount                                   *float32 `json:"GrossAmount"`
	OrderID                                       *int     `json:"OrderID"`
	OrderItem                                     *int     `json:"OrderItem"`
	Contract                                      *int     `json:"Contract"`
	ContractItem                                  *int     `json:"ContractItem"`
	ProductionVersion                             *int     `json:"ProductionVersion"`
	ProductionVersionItem                         *int     `json:"ProductionVersionItem"`
	ProductionOrder                               *int     `json:"ProductionOrder"`
	ProductionOrderItem                           *int     `json:"ProductionOrderItem"`
	BillOfMaterial                                *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                            *int     `json:"BillOfMaterialItem"`
	Operations                                    *int     `json:"Operations"`
	OperationsItem                                *int     `json:"OperationsItem"`
	OperationID                                   *int     `json:"OperationID"`
	OrderType                                     *string  `json:"OrderType"`
	ContractType                                  *string  `json:"ContractType"`
	OrderValidityStartDate                        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                          *string  `json:"OrderValidityEndDate"`
	PaymentTerms                                  *string  `json:"PaymentTerms"`
	DueCalculationBaseDate                        *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                                *string  `json:"PaymentDueDate"`
	NetPaymentDays                                *int     `json:"NetPaymentDays"`
	PaymentMethod                                 *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                        *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                          *string  `json:"InvoicePeriodEndDate"`
	ConfirmedDeliveryDate                         *string  `json:"ConfirmedDeliveryDate"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	ReferenceDocument                             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                         *int     `json:"ReferenceDocumentItem"`
	TransactionTaxClassification                  string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry         string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry       string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                      string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                        string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup                 string   `json:"ProductAccountAssignmentGroup"`
	TaxCode                                       *string  `json:"TaxCode"`
	TaxRate                                       *float32 `json:"TaxRate"`
	CountryOfOrigin                               *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                       *string  `json:"CountryOfOriginLanguage"`
	Equipment                                     *int     `json:"Equipment"`
	FreightOrder                                  *int     `json:"FreightOrder"`
	ItemDeliveryBlockStatus                       *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                        *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                      *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                        *bool    `json:"ItemBillingBlockStatus"`
	ExternalReferenceDocument                     *string  `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem                 *string  `json:"ExternalReferenceDocumentItem"`
	CreationDate                                  string   `json:"CreationDate"`
	CreationTime                                  string   `json:"CreationTime"`
	LastChangeDate                                string   `json:"LastChangeDate"`
	LastChangeTime                                string   `json:"LastChangeTime"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}
