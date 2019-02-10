package controllers

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"lendr/internal/models"
)

func MigrateLendables(_ http.ResponseWriter, _ *http.Request) {
	models.CreateLendablesTable()
}

func RevertLendables(_ http.ResponseWriter, _ *http.Request) {
	models.DeleteLendablesTable()
}
