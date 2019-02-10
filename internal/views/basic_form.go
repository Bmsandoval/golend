package view

const TestForm = `
{
  Text:       "Which beer do you want? :beer:",
  Color:      "#f9a41b",
  CallbackID: "beer",
  Actions: {
    {
      Name: "actionSelect",
      Type: "select",
      Options: {
        {
          Text:  "Asahi Super Dry",
          Value: "Asahi Super Dry",
        },
        {
          Text:  "Kirin Lager Beer",
          Value: "Kirin Lager Beer",
        },
        {
          Text:  "Sapporo Black Label",
          Value: "Sapporo Black Label",
        },
        {
          Text:  "Suntory Malt’s",
          Value: "Suntory Malts",
        },
        {
          Text:  "Yona Yona Ale",
          Value: "Yona Yona Ale",
        },
      },
    },
    {
      Name:  "actionCancel",
      Text:  "Cancel",
      Type:  "button",
      Style: "danger",
    },
  },
}
`

const TestFormTwo = `
[
	{
		"text":            "prompt text",
		"fallback":        "well, this is awkward..",  # error text
		"color":           "#3aa3e3",
		"attachment_type": "default",
		"callback_id":     "callback id",
		"actions": [
			{
				"name":    "action name",
				"text":    "prompt text",  # prompt text
				"type":    "select",
				"options": [
          {
            "text":  "option text",
            "value": "option value"
          }
        ]
			}
		]
	}
]
`

const TestFormThree = `
{"response_type":"{{.Count}}","text":"{{.Material}}"}
`
