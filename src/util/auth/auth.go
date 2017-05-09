package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/melodiez14/meiko/src/module/user"
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
)

var (
	c                  Config
	errSessionNotlogin = errors.New("SessionNotLogin")
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
		userData := &user.User{}
		cookie, err := r.Cookie(c.SessionKey)
		if err == nil {
			userData, _ = getUserInfo(cookie.Value)
		}

		r = r.WithContext(context.WithValue(r.Context(), "User", userData))
		h(w, r, ps)
	}
}

func getUserInfo(session string) (*user.User, error) {

	session = strings.Trim(session, " ")
	client := conn.Redis.Get()
	defer client.Close()

	key := sessionPrefix + session
	jsd, err := redis.String(client.Do("GET", key))

	if err != nil && err != redis.ErrNil {
		return nil, err
	}

	res := &user.User{}
	err = json.Unmarshal([]byte(jsd), res)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errSessionNotlogin
	}

	return res, nil
}
