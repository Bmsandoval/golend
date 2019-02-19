package database

import (
	"golend/internal/models"
	"golend/pkg/db"
	"net/http"

	//_ "github.com/go-sql-driver/mysql"

)

func Migrate(_ http.ResponseWriter, _ *http.Request) {
	db.DB.Debug().AutoMigrate(&models.Lender{})
	db.DB.Debug().AutoMigrate(&models.Grouper{})
	db.DB.Debug().AutoMigrate(&models.Lendable{})
}

func Revert(_ http.ResponseWriter, _ *http.Request) {
	db.DB.Debug().DropTableIfExists(&models.Lendable{})
	db.DB.Debug().DropTableIfExists(&models.Grouper{})
	db.DB.Debug().DropTableIfExists(&models.Lender{})
}
