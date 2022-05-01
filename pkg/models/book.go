package models

import (
	"time"

	"gorm.io/datatypes"
)

type Book struct {
	ID                         int               `json:"id" gorm:"primaryKey"`
	ShopDomain                 string            `json:"shop"`
	ShopId                     int64             `json:"shop_id"`
	Name                       string            `json:"name"`
	ShopOwner                  string            `json:"shop_owner"`
	Email                      string            `json:"email"`
	CustomerSupportEmail       string            `json:"customer_support_email"`
	Address1                   string            `json:"address1"`
	Address2                   string            `json:"address2"`
	City                       string            `json:"city"`
	State                      string            `json:"state"`
	CountryCode                string            `json:"country_code"`
	CountryName                string            `json:"country_name"`
	Pincode                    string            `json:"pincode"`
	Host                       string            `json:"host"`
	Domain                     string            `json:"domain" gorm:"index"`
	Phone                      string            `json:"phone"`
	WidgetAccessScope          string            `json:"widget_access_scope"`
	WidgetAccessToken          string            `json:"widget_access_token"`
	WidgetInstallationStatus   string            `json:"widget_installation_status"`
	IsWidgetEnabled            bool              `json:"is_widget_enabled"`
	PaymentsAccessScope        string            `json:"payments_access_scope"`
	PaymentsAccessToken        string            `json:"payments_access_token"`
	PaymentsInstallationStatus string            `json:"payments_installation_status"`
	IsPaymentsEnabled          bool              `json:"is_payments_enabled"`
	MetadataId                 int64             `json:"metadata_id"`
	WidgetConfigs              datatypes.JSON    `json:"widget_configs"`
	ClientId                   string            `json:"client_id"`
	ClientSecret               string            `json:"client_secret"`
	CreatedAt                  time.Time         `json:"created_at"`
	UpdatedAt                  time.Time         `json:"updated_at"`
	Config                     datatypes.JSONMap `json:"config" sql:"type:jsonb"`
}
