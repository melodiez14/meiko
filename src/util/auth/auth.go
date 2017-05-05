// package auth

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// 	"github.com/melodiez14/meiko/src/util/env"
// 	"github.com/tokopedia/reputation/src/utils/external"
// 	"github.com/tokopedia/reputation/src/utils/template"
// )

// var SessionKey = map[string]string{
// 	"production":  "_SID_Tokopedia_",
// 	"staging":     "_SID_Tokopedia_Coba_",
// 	"alpha":       "_SID_Tokopedia_Alpha_",
// 	"development": "_SID_Tokopedia_",
// }

// // MustAuthorize you must provide the Bearer token on header if you're using this middleware
// func MustAuthorize(h httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 		env := env.Get()
// 		sessionKey := SessionKey[env]
// 		cookie, err := r.Cookie(sessionKey)
// 		if err != nil {
// 			template.RenderResponse(w, r, new(template.Response).
// 				SetCode(http.StatusForbidden).
// 				AddError("Invalid Session"))
// 			return
// 		}

// 		userData, err := getUserInfo(cookie.Value)
// 		if err != nil {
// 			template.RenderResponse(w, r, new(template.Response).
// 				SetCode(http.StatusForbidden).
// 				AddError("Invalid Session"))
// 			return
// 		}
// 		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		r = r.WithContext(context.WithValue(r.Context(), "User", userData))
// 		r = r.WithContext(context.WithValue(r.Context(), "Language", userData.Lang))

// 		h(w, r, ps)
// 	}
// }

// // OptionalAuthorize you don't really have to pass the Bearer token if using this middleware
// func OptionalAuthorize(h httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 		userData := &User{}
// 		env := env.Get()
// 		sessionKey := SessionKey[env]
// 		cookie, err := r.Cookie(sessionKey)
// 		if err == nil {
// 			userData, _ = getUserInfo(cookie.Value)
// 		}
// 		r = r.WithContext(context.WithValue(r.Context(), "User", userData))
// 		r = r.WithContext(context.WithValue(r.Context(), "Language", userData.Lang))
// 		h(w, r, ps)
// 	}
// }

// // UserThumbnail to get user image uri using imagerouter
// func UserThumbnail(UserID int64) string {
// 	if UserID == 0 {
// 		return fmt.Sprintf("")
// 	}
// 	imgPath := external.URL.ImageRouter + "/image/v1"

// 	return fmt.Sprintf("%s/u/%d/user_thumbnail/desktop", imgPath, UserID)
// }

// // ShopThumbnail to get shop image uri using imagerouter
// func ShopThumbnail(ShopID int64) string {
// 	if ShopID == 0 {
// 		return fmt.Sprintf("")
// 	}

// 	imgPath := external.URL.ImageRouter + "/image/v1"

// 	return fmt.Sprintf("%s/s/%d/shop_xs_thumbnail/desktop", imgPath, ShopID)
// }

// func getUserInfo(cookieValue string) (*User, error) {
// 	userSession, err := SHelper.GetUser(cookieValue)
// 	if err != nil {
// 		return &User{}, err
// 	}
// 	userData := &User{
// 		UserID:      userSession.UserID,
// 		Email:       userSession.Email,
// 		FullName:    userSession.FullName,
// 		Lang:        userSession.Lang,
// 		Status:      userSession.Status,
// 		ShopID:      userSession.ShopID,
// 		ShopThumb:   ShopThumbnail(userSession.ShopID),
// 		UserThumb:   UserThumbnail(userSession.UserID),
// 		AccessToken: userSession.AccessToken,
// 	}

// 	return userData, nil
// }

// func GetUserDataFromRequest(r *http.Request) (*User, error) {
// 	if r.Context().Value("User") != nil {
// 		userData := r.Context().Value("User").(*User)
// 		if userData == nil {
// 			return nil, fmt.Errorf("UserData not found")
// 		}

// 		return userData, nil
// 	}
// 	return nil, fmt.Errorf("Cannot get user data")
// }
