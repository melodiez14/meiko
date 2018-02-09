package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/webserver/template"
	onesignal "github.com/tbalthazar/onesignal-go"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := onesignal.NewClient(nil)
	client.AppKey = "9f1b7d96-d6d8-4e3f-9d2b-9978f5f2b5a1"
	client.UserKey = "OWQ0MTFkZjYtNjdlOC00N2Y2LWFmN2YtN2IxMDdkOGNhYjEw"
	apps, res, err := client.Apps.List()
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(apps, res)
	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetMessage("Hello from meiko. Have a nice day! :)"))
	return
}
