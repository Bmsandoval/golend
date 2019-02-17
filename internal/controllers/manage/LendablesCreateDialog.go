package manage

import (
	"fmt"
	"net/http"

	"github.com/nlopes/slack"
)

const lendablesCreateDialogCallback string = "lendables_create_dialog"

func LendablesCreateDialog(requestValues slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", requestValues.DialogSubmissionCallback)
}
