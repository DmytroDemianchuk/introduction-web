package horoshop

import (
	"strconv"

	"github.com/Easy-Job-Developer/catalog_plus/domain/records/models"
)

type Product struct {
	Article                *string          `json:"article,omitempty"`
	Title                  *LocalizedString `json:"title,omitempty"`
	H1Title                *LocalizedString `json:"h1_title,omitempty"`
	SEOTitle               *LocalizedString `json:"seo_title,omitempty"`
	Description            *LocalizedString `json:"description,omitempty"`
	ShortDescription       *LocalizedString `json:"short_description,omitempty"`
	SEODescription         *LocalizedString `json:"seo_description,omitempty"`
	MarketplaceDescription *LocalizedString `json:"marketplace_description,omitempty"`
	GTIN                   *string          `json:"gtin,omitempty"`
	MPN                    *string          `json:"mpn,omitempty"`
	Popularity             *int             `json:"popularity,omitempty"`
	GuaranteeShop          *string          `json:"guarantee_shop,omitempty"`
	GuaranteeLength        *int             `json:"guarantee_length,omitempty"`
	Parent                 *ParentID        `json:"parent,omitempty"`
	Presence               *LocalizedString `json:"presence,omitempty"`
	Price                  *float64         `json:"price,omitempty"`
	PriceOld               *float64         `json:"price_old,omitempty"`
	Row                    int              `json:"excel_row,omitempty"`
}
type ParentID struct {
	ID    *int    `json:"id"`
	Value *string `json:"value"`
}

type LocalizedString struct {
	RU *string `json:"ru,omitempty"`
	UA *string `json:"ua,omitempty"`
}

type ProductSheetsMap struct {
	Sheet      string     `json:"sheetName"`
	ProductMap ProductMap `json:"product"`
}

type ProductMap struct {
	Article                *HeaderColumnIndex `json:"article,omitempty"`
	Title                  LocalizedStringMap `json:"title,omitempty"`
	H1Title                LocalizedStringMap `json:"h1_title,omitempty"`
	SEOTitle               LocalizedStringMap `json:"seo_title,omitempty"`
	Description            LocalizedStringMap `json:"description,omitempty"`
	ShortDescription       LocalizedStringMap `json:"short_description,omitempty"`
	SEODescription         LocalizedStringMap `json:"seo_description,omitempty"`
	MarketplaceDescription LocalizedStringMap `json:"marketplace_description,omitempty"`
	GTIN                   *HeaderColumnIndex `json:"gtin,omitempty"`
	MPN                    *HeaderColumnIndex `json:"mpn,omitempty"`
	Popularity             *HeaderColumnIndex `json:"popularity,omitempty"`
	GuaranteeShop          *HeaderColumnIndex `json:"guarantee_shop,omitempty"`
	GuaranteeLength        *HeaderColumnIndex `json:"guarantee_length,omitempty"`
	Parent                 ParentIDMap        `json:"parent,omitempty"`
	Presence               LocalizedStringMap `json:"presence,omitempty"`
	Price                  *HeaderColumnIndex `json:"price,omitempty"`
	PriceOld               *HeaderColumnIndex `json:"price_old,omitempty"`
}
type ParentIDMap struct {
	ID    *HeaderColumnIndex `json:"id"`
	Value *HeaderColumnIndex `json:"value"`
}

type LocalizedStringMap struct {
	RU *HeaderColumnIndex `json:"ru,omitempty"`
	UA *HeaderColumnIndex `json:"ua,omitempty"`
}

type HeaderColumnIndex struct {
	Column models.Column `json:"column,omitempty"`
}

func (productMap *ProductMap) InitProductFromTable(rows [][]string, productRows []int) ([]Product, []Product) {
	var products []Product
	var productsFailed []Product

	getStrValue := func(hci *HeaderColumnIndex, rowInd int) *string {
		if hci == nil {
			return nil
		}
		colIndex := hci.Column.Index.Column
		if colIndex < 0 || colIndex >= len(rows[rowInd]) {
			return nil
		}
		content := rows[rowInd][colIndex]
		if content == "" {
			return nil
		}
		return &content
	}
	getIntValue := func(hci *HeaderColumnIndex, rowInd int) *int {
		if hci == nil {
			return nil
		}
		colIndex := hci.Column.Index.Column
		if colIndex < 0 || colIndex >= len(rows[rowInd]) {
			return nil
		}
		content := rows[rowInd][colIndex]
		if content == "" {
			return nil
		}
		val, _ := strconv.Atoi(content)
		return &val
	}

	getFloatValue := func(hci *HeaderColumnIndex, rowInd int) *float64 {
		if hci == nil {
			return nil
		}
		colIndex := hci.Column.Index.Column
		if colIndex < 0 || colIndex >= len(rows[rowInd]) {
			return nil
		}
		content := rows[rowInd][colIndex]
		if content == "" {
			return nil
		}
		cleanContent := ""
		for _, char := range content {
			if char >= '0' && char <= '9' || char == '.' {
				cleanContent += string(char)
			} else {
				break
			}
		}

		val, _ := strconv.ParseFloat(cleanContent, 64)
		return &val
	}

	for _, r := range productRows {
		product := Product{
			Article: getStrValue(productMap.Article, r-1),
			Title: &LocalizedString{
				RU: getStrValue(productMap.Title.RU, r-1),
				UA: getStrValue(productMap.Title.UA, r-1),
			},
			H1Title: &LocalizedString{
				RU: getStrValue(productMap.H1Title.RU, r-1),
				UA: getStrValue(productMap.H1Title.UA, r-1),
			},
			SEOTitle: &LocalizedString{
				RU: getStrValue(productMap.SEOTitle.RU, r-1),
				UA: getStrValue(productMap.SEOTitle.UA, r-1),
			},
			Description: &LocalizedString{
				RU: getStrValue(productMap.Description.RU, r-1),
				UA: getStrValue(productMap.Description.UA, r-1),
			},
			ShortDescription: &LocalizedString{
				RU: getStrValue(productMap.ShortDescription.RU, r-1),
				UA: getStrValue(productMap.ShortDescription.UA, r-1),
			},
			SEODescription: &LocalizedString{
				RU: getStrValue(productMap.SEODescription.RU, r-1),
				UA: getStrValue(productMap.SEODescription.UA, r-1),
			},
			MarketplaceDescription: &LocalizedString{
				RU: getStrValue(productMap.MarketplaceDescription.RU, r-1),
				UA: getStrValue(productMap.MarketplaceDescription.UA, r-1),
			},
			GTIN:            getStrValue(productMap.GTIN, r-1),
			MPN:             getStrValue(productMap.MPN, r-1),
			Popularity:      getIntValue(productMap.Popularity, r-1),
			GuaranteeShop:   getStrValue(productMap.GuaranteeShop, r-1),
			GuaranteeLength: getIntValue(productMap.GuaranteeLength, r-1),
			Parent: &ParentID{
				ID:    getIntValue(productMap.Parent.ID, r-1),
				Value: getStrValue(productMap.Parent.Value, r-1),
			},
			Presence: &LocalizedString{
				RU: getStrValue(productMap.Presence.RU, r-1),
				UA: getStrValue(productMap.Presence.UA, r-1),
			},
			Price:    getFloatValue(productMap.Price, r-1),
			PriceOld: getFloatValue(productMap.PriceOld, r-1),
			Row:      r,
		}

		if product.Article != nil {
			products = append(products, product)
		} else {
			productsFailed = append(productsFailed, product)
		}

	}
	return products, productsFailed

}
