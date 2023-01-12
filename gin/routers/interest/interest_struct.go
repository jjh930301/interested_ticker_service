package interest

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type interestUserModel struct {
	ID       uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"`
}

type interestTickerChartModel struct {
	TickerSymbol string          `gorm:"column:ticker_symbol;type:varchar(12);not null" json:"-"`
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close"`
	Percent      string          `gorm:"column:percent;type:float;not null" json:"percent"`
	Volume       string          `gorm:"column:volume;type:bigint;not null" json:"volume"`
}

type interestList struct {
	UserId       uuid.UUID `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User         interestUserModel
	TickerSymbol string    `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol,omitempty"` //index
	DateTime     time.Time `json:"date_time"`
	Type         int       `json:"type"`
	Name         string    `json:"name"`
	Close        string    `json:"close"`
	Percent      string    `json:"percent"`
	Volume       string    `json:"volume"`
}

type InterestListResponse struct {
	Ticker interestList             `json:"ticker"`
	Recent interestTickerChartModel `json:"recent"`
}

type SetInterestBody struct {
	Code string `json:"code"`
}