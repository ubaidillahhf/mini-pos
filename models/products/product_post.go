package products

type ProductPost struct {
	Id          int    `json:"id" gorm:"primaryKey;index"`
	MerchantId  int    `json:"merchant_id" gorm:"not null;index"`
	Sku         string `json:"sku" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Image       string `json:"image" `
	Remark      string `json:"remark"`
	Description string `json:"description"`
}
