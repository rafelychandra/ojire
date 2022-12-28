package entity

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type (
	Product struct {
		Id        uint64    `json:"id,omitempty" gorm:"column:id"`
		Name      string    `json:"name,omitempty" gorm:"column:name"`
		SKU       string    `json:"sku,omitempty" gorm:"column:sku"`
		Quantity  uint64    `json:"quantity,omitempty" gorm:"column:quantity"`
		CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
		UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	}

	RUserProduct struct {
		Id        uint64    `json:"id,omitempty" gorm:"column:id"`
		UserId    uint64    `json:"userId,omitempty" gorm:"column:userId"`
		ProductId uint64    `json:"productId,omitempty" gorm:"column:productId"`
		CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
		UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	}
)

func GenerateSKU(quantity uint64) (string, error) {
	now := time.Now().Format("20060102150405")
	max := big.NewInt(9999)

	totalStr := strconv.Itoa(int(quantity))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-SKU%s%s", totalStr, n.String(), now), nil
}
