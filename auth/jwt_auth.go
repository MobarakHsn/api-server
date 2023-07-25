package auth

import (
	"api-server/data"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var mySigningKey = []byte("gimme")

func GenerateToken(cred data.Credential) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = cred.Username
	claims["password"] = cred.Password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var cred data.Credential
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	correctPassword, ok := data.CredentialList[cred.Username]
	if !ok || cred.Password != correctPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenString, err := GenerateToken(cred)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(tokenString))
}

func AuthMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Invalid token!", http.StatusUnauthorized)
			return
		}
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Invalid token!", http.StatusUnauthorized)
			return
		}
	}
}
