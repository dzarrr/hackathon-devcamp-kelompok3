package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// GetUserByID a method to get user given userID params in URL
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}

	query := fmt.Sprintf("SELECT id, name, email FROM users WHERE id = %d", userID)
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var users []User

	for rows.Next() {
		user := User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, user)
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, bytes, http.StatusOK)
}

// InsertUser a function to insert user data (id, name, email) to DB
func (h *Handler) InsertUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	// read json body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse json body
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		return
	}

	// executing insert query
	query := fmt.Sprintf("INSERT INTO users (id, name, email) VALUES (%d, '%s', '%s')", user.ID, user.Name, user.Email)
	_, err = h.DB.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Insert user success!"
	}
	`), http.StatusOK)
}

// EditUserByID a function to change user data (name) in DB with given params (id, name)
func (h *Handler) EditUserByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := param.ByName("userID")
	// read json body
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		renderJSON(w, []byte(`
			message: "Failed to read body"
		`), http.StatusBadRequest)
		return
	}

	// parse json body
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println(err)
		return
	}

	query := fmt.Sprintf("UPDATE users SET name = '%s' WHERE id = %s", user.Name, userID)
	_, err = h.DB.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Update user success!"
	}
	`), http.StatusOK)
}

// DeleteUserByID a function to remove user data from DB given the userID
func (h *Handler) DeleteUserByID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID := param.ByName("userID")

	query := fmt.Sprintf("DELETE FROM users WHERE id = %s", userID)
	_, err := h.DB.Exec(query)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Delete user success!"
	}
	`), http.StatusOK)
}
