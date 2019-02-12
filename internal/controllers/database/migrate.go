package database

import (
	"lendr/internal/models/grouper"
	"lendr/internal/models/lender"
	"lendr/pkg/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"lendr/internal/models/lendable"
)

func Migrate(_ http.ResponseWriter, _ *http.Request) {
	db.DB.Debug().AutoMigrate(&lender.Lender{})
	db.DB.Debug().AutoMigrate(&grouper.Grouper{})
	db.DB.Debug().AutoMigrate(&lendable.Lendable{})
}

func Revert(_ http.ResponseWriter, _ *http.Request) {
	db.DB.Debug().DropTableIfExists(&lendable.Lendable{})
	db.DB.Debug().DropTableIfExists(&grouper.Grouper{})
	db.DB.Debug().DropTableIfExists(&lender.Lender{})
}
