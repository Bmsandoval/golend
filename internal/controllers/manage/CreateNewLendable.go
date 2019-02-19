package manage

import (
	"github.com/nlopes/slack"
	"golend/internal/models"
	"log"
	"net/http"
	"strconv"
)

const lendablesCreateDialogCallback string = "lendables_create_dialog"

func CreateNewLendable(requestValues slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	hashState := requestValues.State
	lendr, err := models.SearchLendersAsAdmin(requestValues.Team.ID, requestValues.User.ID)
	if err != nil {
		log.Fatal(err.Error())
	}
	if ! lendr.ValidateHash(hashState) {
		w.WriteHeader(http.StatusForbidden)
	}

	var grouperId uint
	group := requestValues.Submission["groupStatus"]
	name := requestValues.Submission["lendableName"]
	if "isGroup" == group {
		grouper := models.MakeNewGrouper(lendr.ID, name)
		grouperId = grouper.ID
	} else {
		// TODO - probably shouldn't ignore this error
		groupStatus, _ := strconv.ParseUint(requestValues.Submission["groupStatus"], 10, 32)
		grouperId = uint(groupStatus)
	}
	models.MakeNewLendable(grouperId, name)
}
