package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"lendr/internal/models"
	"lendr/pkg/db"
)

func ShowAllLendrs(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lendrId := vars["LendrId"]
	var lendr = models.Lendr{LendrId: lendrId}
	result := db.DB.Find(&lendr)
	fmt.Println(result)
}

func RegisterLendr(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lendrId := vars["LendrId"]
	models.AddNewLendr(lendrId)
	fmt.Println("Registered: " + lendrId)
}
func UnregisterLendr(_ http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lendrId := vars["LendrId"]
	models.RemoveExistingLendr(lendrId)
	models.RemoveLendablesByLendr(lendrId)
	fmt.Println("Unregistered: " + lendrId)
}

func MigrateLendrs(_ http.ResponseWriter, _ *http.Request) {
	models.CreateLendrsTable()
}

func RevertLendrs(_ http.ResponseWriter, _ *http.Request) {
	models.DeleteLendrsTable()
}
