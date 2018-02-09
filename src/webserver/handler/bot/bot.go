package bot

import (
	"encoding/json"
	"math/rand"
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
	intent, confidence, _ := bot.GetIntent(args.NormalizedText)

	// generate message by confidence
	levelConfidence := "notsure"
	if confidence >= 0.6 {
		levelConfidence = "confident"
	} else if confidence >= 0.3 {
		levelConfidence = "doubt"
	}

	// get message
	msg := msgConf[levelConfidence]
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(msg))

	msgResp := msg[index]

	intent = intentGreeting

	// convert intent into assistant if the answer is confidence
	data := []map[string]interface{}{}
	if levelConfidence != "notsure" {
		switch intent {
		case intentAssistant:
			data, err = handleAssistant(args.NormalizedText, sess.ID)
		case intentGrade:
			data, err = handleGrade(args.NormalizedText, sess.ID)
		case intentAssignment:
			data, err = handleAssignment(args.NormalizedText, sess.ID)
		case intentInformation:
			data, err = handleInformation(args.NormalizedText, sess.ID)
		case intentSchedule:
			data, err = handleSchedule(args.NormalizedText, sess.ID)
		case intentGreeting:
			intent = intentOther
			msgResp = handleGreeting(sess.Name)
		case intentAboutBot:
			intent = intentOther
			msgResp = handleAboutBot()
		case intentAboutStudent:
			intent = intentOther
			msgResp = handleAboutStudent(sess.Name)
		case intentAboutCreator:
			intent = intentOther
			msgResp = handleAboutCreator()
		case intentKidding:
			intent = intentOther
			msgResp = handleKidding()
		case intentUnknown:
			intent = intentOther
		default:
			break
		}
	}

	// intent and entity
	respData := map[string]interface{}{
		"intent": intent,
		"entity": data,
		"text":   msgResp,
	}

	// log message into database
	jsn, _ := json.Marshal(respData)
	jsnStr := string(jsn)

	tx, err := conn.DB.Beginx()
	if err != nil {
		return
	}

	_, err = log.Insert(args.Text, sess.ID, bot.StatusUser, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	lastInsertID, err := log.Insert(jsnStr, sess.ID, bot.StatusBot, tx)
	if err != nil {
		tx.Rollback()
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	tx.Commit()

	// prepare for response
	respData["id"] = lastInsertID
	resp := messageResponse{
		Status:    bot.StatusBot,
		Text:      args.Text,
		TimeStamp: time.Now().Unix(),
		Response:  respData,
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}

func LoadHistoryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sess := r.Context().Value("User").(*auth.User)
	params := loadHistoryParams{
		id: r.FormValue("id"),
	}

	args, err := params.validate()
	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusBadRequest).
			AddError("Invalid Request"))
		return
	}

	var log []bot.Log
	if args.id.Valid {
		log, err = bot.LoadByID(args.id.Int64, sess.ID)
	} else {
		log, err = bot.LoadByUserID(sess.ID)
	}

	if err != nil {
		template.RenderJSONResponse(w, new(template.Response).
			SetCode(http.StatusInternalServerError))
		return
	}

	resp := []map[string]interface{}{}
	for _, val := range log {
		if val.Status == bot.StatusUser {
			resp = append(resp, map[string]interface{}{
				"status": bot.StatusUser,
				"time":   val.CreatedAt.Unix(),
				"message": map[string]interface{}{
					"id":   val.ID,
					"text": val.Message,
				},
			})
			continue
		}

		jsnMap := map[string]interface{}{}
		json.Unmarshal([]byte(val.Message), &jsnMap)
		jsnMap["id"] = val.ID
		resp = append(resp, map[string]interface{}{
			"status":  bot.StatusBot,
			"time":    val.CreatedAt.Unix(),
			"message": jsnMap,
		})
	}

	template.RenderJSONResponse(w, new(template.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return
}
