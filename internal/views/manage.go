package view

const ManageBaseForm = `
{
	"text":            "Available QA environments",  # Description Text
	"fallback":        "Well, this is awkward..",  # Error Text
	"color":           "#3AA3E3",
	"attachment_type": "default",
	"callback_id":     "manage.main_selection",
	"actions":         [
		{
			"name":  "create_action",
			"text":  "Add Something",  # Prompt Text
			"type":  "button",
			"value": "create"
		},  # Create
		# TODO : feature/update_thing
		{
			"name":  "remove_action",
			"text":  "Remove Something",  # Prompt Text
			"type":  "button",
			"value": "remove"
		},  # Remove
	]
}
`

