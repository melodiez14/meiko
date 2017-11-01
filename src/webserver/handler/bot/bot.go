package bot

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/bot"
	"github.com/melodiez14/meiko/src/module/log"
	"github.com/melodiez14/meiko/src/util/auth"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

func BotHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := messageParams{
		Text: r.FormValue("text"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError(err.Error()))
		return
	}

	// get text intent
	intent, err := getIntent(args.NormalizedText)

	// convert intent into assistant
	var data interface{}
	switch intent {
	case intentAssistant:
		data, err = handleAssistant(args.NormalizedText, sess.ID)
	case intentGrade:
		break
	case intentAssignment:
		break
	case intentInformation:
		break
	case intentSchedule:
		break
	default:
		break
	}
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusOK).
			SetMessage(err.Error()))
		return
	}

	// intent and entity
	respData := map[string]interface{}{
		"intent": intent,
		"entity": data,
	}

	// prepare for response
	resp := messageResponse{
		Status:    bot.StatusBot,
		Text:      args.Text,
		TimeStamp: time.Now().Unix(),
		Response:  respData,
	}

	// log message into database
	go func() {

		jsn, _ := json.Marshal(respData)
		jsnStr := string(jsn)

		tx, err := conn.DB.Beginx()
		if err != nil {
			return
		}

		err = log.Insert(args.Text, sess.ID, bot.StatusUser, tx)
		if err != nil {
			return
		}

		err = log.Insert(jsnStr, sess.ID, bot.StatusBot, tx)
		if err != nil {
			return
		}
		tx.Commit()
	}()

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

func LoadHistoryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := loadHistoryParams{
		Time:     r.FormValue("time"),
		Position: r.FormValue("position"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Bad Request"))
		return
	}

	log, err := bot.LoadByTime(args.Time, args.IsAfter, sess.ID)
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	var resp []map[string]interface{}
	for _, val := range log {
		if val.Status == bot.StatusUser {
			resp = append(resp, map[string]interface{}{
				"status": bot.StatusUser,
				"time":   val.CreatedAt.Unix(),
				"response": map[string]interface{}{
					"text": val.Message,
				},
			})
			continue
		}

		jsnMap := map[string]interface{}{}
		json.Unmarshal([]byte(val.Message), &jsnMap)
		resp = append(resp, map[string]interface{}{
			"status":   bot.StatusBot,
			"time":     val.CreatedAt.Unix(),
			"response": jsnMap,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
