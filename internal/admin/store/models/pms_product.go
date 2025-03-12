package models

type PmsProduct struct {
	Model
	ProductSn                 string                      `json:"product_sn" gorm:"product_sn"`
	BrandId                   int64                       `json:"brand_id" gorm:"brand_id"`
	BrandName                 string                      `json:"brand_name" gorm:"brand_name"`
	ProductCategoryId         int64                       `json:"product_category_id" gorm:"product_category_id"`
	ProductCategoryName       string                      `json:"product_category_name" gorm:"product_category_name"`
	Pic                       string                      `json:"pic" gorm:"pic"`
	Name                      string                      `json:"name" gorm:"name"`
	SubTitle                  string                      `json:"sub_title" gorm:"sub_title"`
	KeyWord                   string                      `json:"key_word" gorm:"key_word"`
	Price                     string                      `json:"price" gorm:"price"`
	Sale                      int                         `json:"sale" gorm:"sale"`
	NewStatus                 int                         `json:"new_status" gorm:"new_status"`
	RecommendStatus           int                         `json:"recommand_status" gorm:"recommand_status"`
	Stock                     int                         `json:"stock" gorm:"stock"`
	PromotionType             int                         `json:"promotion_type" gorm:"promotion_type"`
	Sort                      int                         `json:"sort" gorm:"sort"`
	ProductAttributeValueList []*PmsProductAttributeValue `json:"attr_value_list" gorm:"foreignKey:ProductId"`
}

func (*PmsProduct) TableName() string {
	return "pms_product"
}
