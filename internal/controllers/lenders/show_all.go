package lenders

import (
	"fmt"
	"golend/internal/models"
	"golend/pkg/db"
	"net/http"

	"github.com/gorilla/mux"
)

func ShowAll(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["TeamId"]
	var lendr = models.Lender{TeamId: teamId}
	db.DB.Table("lenders").Find(&lendr)
	result := db.DB.Find(&lendr)
	fmt.Println(result)
}
