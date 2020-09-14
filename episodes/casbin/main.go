package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/casbin/casbin/v2"
	"casbin/authorization"
	"casbin/model"
)

// REFERENCE: https://github.com/zupzup/casbin-http-role-example
func main() {

	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcer("./auth_model.conf", "./policy.csv")
	if err != nil { log.Fatal(err) }

	// setup session store
	sessManager := scs.New()
	sessManager.Lifetime = 3*time.Minute

	// setup users
	users := createUsers()

	// setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler(sessManager, users))
	mux.HandleFunc("/logout", logoutHandler(sessManager))
	mux.HandleFunc("/member/current", currentMemberHandler(sessManager))
	mux.HandleFunc("/member/role", memberRoleHandler(sessManager))
	mux.HandleFunc("/admin/stuff", adminHandler())

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", sessManager.LoadAndSave(authorization.Authorizer(authEnforcer, sessManager, users)(mux))))

}

func createUsers() model.Users {
	users := model.Users{}
	users = append(users, model.User{ID: 1, Name: "Admin", Role: "admin"})
	users = append(users, model.User{ID: 2, Name: "Sabine", Role: "member"})
	users = append(users, model.User{ID: 3, Name: "Sepp", Role: "memeber"})
	return users
}

func loginHandler(sessManager *scs.SessionManager, users model.Users) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			writeError(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}
		// setup session
		if err := sessManager.RenewToken(r.Context()); err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}
		sessManager.Put(r.Context(), "userID", user.ID)
		sessManager.Put(r.Context(), "role", user.Role)
		writeSuccess("SUCCESS", w)
	})
}

func logoutHandler(sessManager *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := sessManager.Destroy(r.Context()); err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}
		writeSuccess("SUCCESS", w)
	})
}

func currentMemberHandler(sessManager *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid:= sessManager.GetInt(r.Context(), "userID")

		writeSuccess(fmt.Sprintf("User with ID: %d", uid), w)
	})
}

func memberRoleHandler(sessManager *scs.SessionManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := sessManager.GetString(r.Context(), "role")

		writeSuccess(fmt.Sprintf("User with Role: %s", role), w)
	})
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("I'm an Admin!", w)
	})
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
