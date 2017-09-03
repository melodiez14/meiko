package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/webserver/template"
)

type (
	Config struct {
		SessionKey string `json:"sessionkey"`
	}
)

const (
	sessionPrefix = "session:"
	character     = "!QAZ@WSX#EDC$RFV%TGB^YHN&UJM*IK<(OL>)P:?_{+}|1qaz2wsx3edc4rfv5tgb6yhn7ujm8ik,9ol.0p-[=]"
)

var (
	c                  Config
	errSessionNotlogin = errors.New("SessionNotLogin")
	charMaxIndex       = len(character)
)

func Init(cfg Config) {
	c = cfg
}

// MustAuthorize you must provide the Bearer token on header if you're using this middleware
func MustAuthorize(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookie, err := r.Cookie(c.SessionKey)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusForbidden).
				AddError("Invalid Session"))
			return
		}

		userData, err := getUserInfo(cookie.Value)
		if err != nil {
			template.RenderJSONResponse(w, new(template.Response).
				SetCode(http.StatusForbidden).
				AddError("Invalid Session"))
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		r = r.WithContext(context.WithValue(r.Context(), "User", userData))

		h(w, r, ps)
	}
}

// OptionalAuthorize you don't really have to pass the Bearer token if using this middleware
func OptionalAuthorize(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var userData *User
		cookie, err := r.Cookie(c.SessionKey)
		if err == nil {
			userData, _ = getUserInfo(cookie.Value)
		}

		r = r.WithContext(context.WithValue(r.Context(), "User", userData))
		h(w, r, ps)
	}
}

func getUserInfo(session string) (*User, error) {

	session = strings.Trim(session, " ")
	client := conn.Redis.Get()
	defer client.Close()

	key := sessionPrefix + session
	jsd, err := redis.String(client.Do("GET", key))

	if err != nil && err != redis.ErrNil {
		return nil, err
	}

	res := &User{}
	err = json.Unmarshal([]byte(jsd), res)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errSessionNotlogin
	}

	return res, nil
}

func (u User) SetSession(w http.ResponseWriter) error {

	var cookie string
	for i := 0; i < 32; i++ {
		cookie = cookie + string(character[rand.Intn(charMaxIndex)])
	}

	key := sessionPrefix + cookie
	data, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	client := conn.Redis.Get()
	_, err = redis.String(client.Do("SET", key, data))
	if err != nil {
		return fmt.Errorf("Failed to set session to Redis")
	}

	http.SetCookie(w, &http.Cookie{
		Name:    c.SessionKey,
		Expires: time.Now().AddDate(0, 1, 0),
		Value:   cookie,
		Path:    "/",
	})

	return nil
}
