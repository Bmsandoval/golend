package examples

import (
	"bytes"
	"net/http"

	_ "github.com/kr/pretty"
	"golend/internal/templates"
)

func TestTemplatized(w http.ResponseWriter, r *http.Request) {
	//	t, err := template.New("test").Delims("<<",">>").Parse(`
	//{
	//	"title":        "What is your thing?",
	//	"submit_label": "Submit",
	//	"callback_id":  "manage.create_thing",
	//	"elements":     [
	//		{
	//			"type":  "text",
	//			"label": "Name",
	//			"name":  "thing_name"
	//		},
	//		{
	//			"label":   "Related",
	//			"type":    "select",
	//			"name":    "relates_to",
	//			"options": [
	//<< range $key, $value := .Count >>
	//				{
	//					"label": "Parent (probably the one you want)",
	//					"value": "<< $key >>"
	//				},
	//<< end >>
	//			]
	//		},
	//		{
	//			"label":   "Control Access",
	//			"type":    "select",
	//			"name":    "control_access",
	//			"options": [
	//<< range $key, $value := .Count >>
	//				{
	//					"label": "Parent (probably the one you want)",
	//					"value": "<< $key >>"
	//				},
	//<< end >>
	//			]
	//		}
	//	]
	//}
	//`)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//
	//type Inventory struct {
	//	Count   []int
	//}
	//sweaters := Inventory{}
	//
	////err = t.Execute(w, sweaters) //, config)
	//err = t.Execute(os.Stdout, sweaters) //, config)
	//if err != nil {
	//	log.Print("execute: ", err)
	//	return
	//}
	//
	//

	buffer := new(bytes.Buffer)
	template.LendableCreationDialog([]int{1, 2, 3, 4}, buffer)
}
