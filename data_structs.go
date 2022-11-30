package shipdis

import (
	"net/http"
	"net/url"
)

type RequestPayload struct {
	URL         string
	Method      string // Use http method constants
	Headers     http.Header
	QueryParams url.Values
	Payload     []byte
}

type DeleteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BillTo and ShipTo
type Address struct {
	Name        string `json:"name,omitempty"`
	Company     string `json:"company,omitempty"`
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	Street3     string `json:"street3,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postalCode,omitempty"`
	Country     string `json:"country,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Residential bool   `json:"residential,omitempty"`
}

type Weight struct {
	Value float64 `json:"value,omitempty"`
	Units string  `json:"units,omitempty"`
}

type Option struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ShipStationItem struct {
	LineItemKey       string   `json:"lineItemKey,omitempty"`
	Sku               string   `json:"sku,omitempty"`
	Name              string   `json:"name,omitempty"`
	ImageURL          string   `json:"imageUrl,omitempty"`
	Weight            *Weight  `json:"weight,omitempty"`
	Quantity          int      `json:"quantity,omitempty"`
	UnitPrice         float64  `json:"unitPrice,omitempty"`
	ShippingAmount    float64  `json:"shippingAmount,omitempty"`
	WarehouseLocation string   `json:"warehouseLocation,omitempty"`
	Options           []Option `json:"options,omitempty"`
	ProductID         int      `json:"productId,omitempty"`
	FulfillmentSku    string   `json:"fulfillmentSku,omitempty"`
	Adjustment        bool     `json:"adjustment,omitempty"`
	Upc               string   `json:"upc,omitempty"`
}

type AdvancedOptions struct {
	Source string `json:"source,omitempty"`
}

type ShipStationOrder struct {
	OrderId         int64             `json:"orderId,omitempty"`
	OrderNumber     string            `json:"orderNumber,omitempty"`
	OrderDate       string            `json:"orderDate,omitempty"`
	PaymentDate     string            `json:"paymentDate,omitempty"`
	ShipByDate      string            `json:"shipByDate,omitempty"`
	OrderStatus     string            `json:"orderStatus,omitempty"`
	CustomerEmail   string            `json:"customerEmail,omitempty"`
	BillTo          *Address          `json:"billTo,omitempty"`
	ShipTo          *Address          `json:"shipTo,omitempty"`
	Items           []ShipStationItem `json:"items,omitempty"`
	AmountPaid      float64           `json:"amountPaid,omitempty"`
	TaxAmount       float64           `json:"taxAmount,omitempty"`
	InternalNotes   string            `json:"internalNotes,omitempty"`
	AdvancedOptions AdvancedOptions   `json:"advancedOptions,omitempty"`
	TagIds          []int             `json:"tagIds,omitempty"`
}

type ShipStationOrderList struct {
	Orders []ShipStationOrder `json:"orders"`
	Total  int                `json:"total"`
	Page   int                `json:"page"`
	Pages  int                `json:"pages"`
}

type ShipStationResponse struct {
	HasErrors bool `json:"hasErrors,omitempty"`
	Results   []struct {
		OrderID      int    `json:"orderId,omitempty"`
		OrderNumber  string `json:"orderNumber,omitempty"`
		OrderKey     string `json:"orderKey,omitempty"`
		Success      bool   `json:"success,omitempty"`
		ErrorMessage string `json:"errorMessage,omitempty"`
	} `json:"results,omitempty"`
}

type ShipStationMarkAsShipped struct {
	OrderID            int64  `json:"orderId,omitempty"`
	CarrierCode        string `json:"carrierCode,omitempty"`
	ShipDate           string `json:"shipDate,omitempty"`
	TrackingNumber     string `json:"trackingNumber,omitempty"`
	NotifyCustomer     bool   `json:"notifyCustomer,omitempty"`
	NotifySalesChannel bool   `json:"notifySalesChannel,omitempty"`
}

// Shipments

type ShipmentList struct {
	Shipments []Shipment `json:"shipments"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Pages     int        `json:"pages"`
}

type Shipment struct {
	ShipmentID          int              `json:"shipmentId"`
	OrderID             int              `json:"orderId"`
	OrderKey            string           `json:"orderKey"`
	UserID              string           `json:"userId"`
	OrderNumber         string           `json:"orderNumber"`
	CreateDate          string           `json:"createDate"`
	ShipDate            string           `json:"shipDate"`
	ShipmentCost        float64          `json:"shipmentCost"`
	InsuranceCost       int              `json:"insuranceCost"`
	TrackingNumber      string           `json:"trackingNumber"`
	IsReturnLabel       bool             `json:"isReturnLabel"`
	BatchNumber         string           `json:"batchNumber"`
	CarrierCode         string           `json:"carrierCode"`
	ServiceCode         string           `json:"serviceCode"`
	PackageCode         string           `json:"packageCode"`
	Confirmation        string           `json:"confirmation"`
	WarehouseID         int              `json:"warehouseId"`
	Voided              bool             `json:"voided"`
	VoidDate            string           `json:"voidDate"`
	MarketplaceNotified bool             `json:"marketplaceNotified"`
	NotifyErrorMessage  interface{}      `json:"notifyErrorMessage"`
	ShipTo              ShipTo           `json:"shipTo"`
	Weight              Weight           `json:"weight"`
	Dimensions          interface{}      `json:"dimensions"`
	InsuranceOptions    InsuranceOptions `json:"insuranceOptions"`
	AdvancedOptions     AdvancedOptions  `json:"advancedOptions"`
	ShipmentItems       []ShipmentItems  `json:"shipmentItems"`
	LabelData           interface{}      `json:"labelData"`
	FormData            interface{}      `json:"formData"`
}
type ShipTo struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Street3     string `json:"street3"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Residential bool   `json:"residential"`
}
type InsuranceOptions struct {
	Provider       string `json:"provider"`
	InsureShipment bool   `json:"insureShipment"`
	InsuredValue   int    `json:"insuredValue"`
}
type ShipmentItems struct {
	OrderItemID       int    `json:"orderItemId"`
	LineItemKey       string `json:"lineItemKey"`
	Sku               string `json:"sku"`
	Name              string `json:"name"`
	ImageURL          string `json:"imageUrl"`
	Weight            Weight `json:"weight"`
	Quantity          int    `json:"quantity"`
	UnitPrice         int    `json:"unitPrice"`
	WarehouseLocation string `json:"warehouseLocation"`
	Options           Option `json:"options"`
	ProductID         int    `json:"productId"`
	FulfillmentSku    string `json:"fulfillmentSku"`
}

type ShipstationHoldItem struct {
	OrderID       int64  `json:"orderId"`
	HoldUntilDate string `json:"holdUntilDate"`
}

type ShipstationMarkShippedResponse struct {
	OrderID     int64  `json:"orderId"`
	OrderNumber string `json:"orderNumber"`
}

type ShipstationHoldResponse struct {
	Success  bool   `json:"success"`
	Messages string `json:"message"`
}
