package manage

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
)

const lendablesCreateDialogCallback string = "lendables_create_dialog"
func LendablesCreateDialog(requestValues slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", requestValues.DialogSubmissionCallback)
}
