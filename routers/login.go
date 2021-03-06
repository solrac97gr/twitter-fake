package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/solrac97gr/twitter-fake/database"
	"github.com/solrac97gr/twitter-fake/jwt"
	"github.com/solrac97gr/twitter-fake/models"
)

/*Login : user try to login in the app*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuer or password invalid"+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := database.TryLogin(user.Email, user.Password)
	if !exist {
		http.Error(w, "Usuer or password invalid", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error ocurred"+err.Error(), 400)
		return
	}
	res := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
