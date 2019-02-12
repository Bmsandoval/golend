package lenders

import (
	"fmt"
	"github.com/gorilla/mux"
	"lendr/internal/models/lender"
	"lendr/pkg/db"
	"net/http"
)

func ShowAll(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamId := vars["TeamId"]
	var lendr = lender.Lender{TeamId: teamId}
	db.DB.Table("lenders").Find(&lendr)
	result := db.DB.Find(&lendr)
	fmt.Println(result)
}
