package manage

import (
	"net/http"
	"strings"

	"github.com/nlopes/slack"
)

// ************************************
// Entry point for management commands
// ************************************
func CallbackHandler(callback slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(callback.CallbackID, "."+baseActionSelectCallback) {
		BaseActionSelect(callback, w, r)
	} else if strings.HasSuffix(callback.CallbackID, "."+lendablesCreateDialogCallback) {
		CreateNewLendable(callback, w, r)
	} else if strings.HasSuffix(callback.CallbackID, "."+lendablesCreateCallback) {
		lendablesCreateView(callback, w, r)
	}
	return
}
