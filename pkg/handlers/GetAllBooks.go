package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

// var shops []string
// queryErr := ssr.Db.Table("shopify_shops").Pluck("domain", &shops).Error
// if queryErr != nil {
// 	return nil, queryErr
// }

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	//var shopDomain []string
	// fetch all shop domain

	//h.DB.Table("books").Pluck("shop_domain", &shopDomain)
	//h.DB.Select("shop_domain", "desc").Find(&books)
	//h.Db.Select("domain", "config").First(&shop)
	// if result := h.DB.Find(&books); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	var count int64
	h.DB.Table("books").Count(&count)
	PageSize := 4
	if count > 0 {
		iterate := int(count) / PageSize
		if int(count)%PageSize == 0 {
			iterate = iterate - 1
		}
		page := 1
		for i := 0; i <= int(iterate); i++ {

			temp := []models.Book{}
			offset := (page - 1) * PageSize
			queryErr := h.DB.Offset(offset).Limit(PageSize).Select("shop_domain", "desc").Find(&temp).Error
			if queryErr != nil {
				log.Println(queryErr)
			}
			page++
			books = append(books, temp...)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
