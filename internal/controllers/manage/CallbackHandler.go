package manage

import (
	"github.com/nlopes/slack"
	"net/http"
	"strings"
)

// ************************************
// Entry point for management commands
// ************************************
func CallbackHandler(callback slack.InteractionCallback, w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(callback.CallbackID, "."+baseActionSelectCallback) {
		BaseActionSelect(callback, w, r)
	} else if strings.HasSuffix(callback.CallbackID, "."+lendablesCreateDialogCallback) {
		LendablesCreateDialog(callback, w, r)
	}
	return
}


