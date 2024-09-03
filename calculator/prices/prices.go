package prices

import (
	"fmt"

	"example.com/file_operations"
)

type TaxedPrice struct {
	IOManager   file_operations.FileManager `json:"-"`
	TaxRate     float64                     `json:"tax_rate"`
	Prices      []float64                   `json:"prices"`
	TaxedPrices map[string]string           `json:"taxed_prices"`
}

func (taxedPrice TaxedPrice) Process() {

	taxedPrice.LoadSavedData()

	resultMap := make(map[string]string)

	for _, price := range taxedPrice.Prices {
		priceAfterTax := price * (1 + taxedPrice.TaxRate)
		resultMap[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceAfterTax)
	}

	taxedPrice.TaxedPrices = resultMap

	taxedPrice.IOManager.SavePriceData(taxedPrice)
}

func (taxedPrice *TaxedPrice) LoadSavedData() {
	prices, err := taxedPrice.IOManager.ReadSavedPriceData()

	if err != nil {
		fmt.Println(err)
		return
	}

	taxedPrice.Prices = prices
}

func New(prices []float64, taxRate float64, fm file_operations.FileManager) *TaxedPrice {

	return &TaxedPrice{
		IOManager: fm,
		Prices:    prices,
		TaxRate:   taxRate,
	}
}
