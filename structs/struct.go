package structs

type Items struct {
	ItemId      int    `gorm:"primary_key;auto_increment" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
	Orders      Orders `gorm:"foreignkey:OrderId"`
}

type Orders struct {
	OrderID      int    `gorm:"primary_key;auto_increment" json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
}

type ItemRequestCreate struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ItemRequestUpdate struct {
	LineItemId  int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type RequestCreate struct {
	OrderedAt    string              `json:"orderedAt"`
	CustomerName string              `json:"customerName"`
	Items        []ItemRequestCreate `json:"items"`
}

type RequestUpdate struct {
	CustomerName string              `json:"customerName"`
	OrderedAt    string              `json:"orderedAt"`
	Items        []ItemRequestUpdate `json:"items"`
}

type ResponseItem struct {
	ItemId      int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ResponseOrder struct {
	OrderID      int            `json:"order_id"`
	CustomerName string         `json:"customer_name"`
	OrderedAt    string         `json:"ordered_at"`
	Items        []ResponseItem `json:"items"`
}
