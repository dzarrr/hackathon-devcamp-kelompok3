package internal

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	// "strings"
	"time"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// GetBookByID a function to get a single book given it's ID
func (h *Handler) GetSavingByUserID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}
	query := fmt.Sprintf("SELECT id, user_id, balance, target, start_date, end_date FROM savings WHERE user_id = %d", userID)
	rows, err := h.DB.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	var savings []Saving

	for rows.Next() {
		saving := Saving{}
		err := rows.Scan(
			&saving.ID,
			&saving.UserID,
			&saving.Balance,
			&saving.Target,
			&saving.StartDate,
			&saving.EndDate,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		savings = append(savings, saving)
	}

	bytes, err := json.Marshal(savings)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, bytes, http.StatusOK)
}


// InsertBook a function to insert book to DB
func (h *Handler) InsertSaving(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
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
	var saving Saving
	err = json.Unmarshal(body, &saving)
	if err != nil {
		log.Printf("[Saving][InsertSaving] Unmarshal error%+v\n", err)
		return
	}

	// executing insert query
	query := fmt.Sprintf(`INSERT INTO savings (user_id, balance, target, start_date, end_date) VALUES (%d, %d, %d, '%s', '%s')`, saving.UserID, saving.Balance, saving.Target, saving.StartDate, saving.EndDate)
	log.Printf(">> query: %s\n", query)
	// _, err = h.DB.Query(query)
	_, err = h.DB.Query(`INSERT INTO savings (user_id, balance, target, start_date, end_date) VALUES ($1, $2, $3, $4, $5)`, saving.UserID, saving.Balance, saving.Target, saving.StartDate, saving.EndDate)
	
	if err != nil {
		log.Printf("[Saving][InsertSaving]%+v", err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Insert saving success!"
	}
	`), http.StatusOK)
}

// EditBook a function to change book data in DB, with given params
func (h *Handler) EditSaving(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}
	
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
	var saving Saving
	err = json.Unmarshal(body, &saving)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println(saving);
	// layout := "2006-01-02T15:04:05-0700";
	// query := fmt.Sprintf("UPDATE savings SET balance = %d, target = %d, start_date = '%s', end_date = '%s' WHERE user_id = %d", saving.Balance, saving.Target, time.Parse(layout, saving.StartDate), time.Parse(layout, saving.EndDate), userID)
	// _, err = h.DB.Query(query)
	_, err = h.DB.Query(`UPDATE savings SET balance = $1, target = $2, start_date = $3, end_date = $4 WHERE user_id = $5`, saving.Balance, saving.Target, saving.StartDate, saving.EndDate, userID)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Update book success!"
	}
	`), http.StatusOK)
}

func (h *Handler) AddBalance(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}
	
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
	var saving Saving
	err = json.Unmarshal(body, &saving)
	if err != nil {
		log.Println(err)
		return
	}

	add_balance := saving.Balance
	if (add_balance <= 0) {
		log.Println("Add zero or negative balance")
		renderJSON(w, []byte(`
			message: "Failed to add balance. Zero or negative balance detected."
		`), http.StatusBadRequest)
		return
	}

	var end_date, new_end_date time.Time
	var target, balance, new_balance int
	err = h.DB.QueryRow(`SELECT balance, target, end_date FROM savings where user_id = $1`, userID).Scan(&balance, &target, &end_date)
	if err != nil {
		log.Fatal(err)
	}
	
	days_left := (end_date.Sub(time.Now()).Seconds()) / (24*3600)
	daily_pay  := float64(target-balance)/days_left
	
	new_balance = balance + add_balance
	new_days_left := float64(target-new_balance)/daily_pay
	new_end_date = time.Now().AddDate(0, 0, int(new_days_left))

	_, err = h.DB.Query(`UPDATE savings SET balance = $1, end_date = $2 WHERE user_id = $3`, new_balance, new_end_date, userID)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Update balance success!",
		balance: `+ strconv.Itoa(new_balance) +`,
		daily_pay: `+ fmt.Sprintf("%.2f", daily_pay) +`,
		end_date: `+ new_end_date.String() +`
	}
	`), http.StatusOK)
}

func (h *Handler) EditEndDate(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}
	
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
	var saving Saving
	err = json.Unmarshal(body, &saving)
	if err != nil {
		log.Println(err)
		return
	}

	var end_date, new_end_date time.Time
	var target, balance int
	new_end_date = saving.EndDate
	err = h.DB.QueryRow(`SELECT balance, target, end_date FROM savings where user_id = $1`, userID).Scan(&balance, &target, &end_date)
	if err != nil {
		log.Fatal(err)
	}
	days_left := (end_date.Sub(time.Now()).Seconds()) / (24*3600)
	daily_pay  := float64(target-balance)/days_left
	
	_, err = h.DB.Query(`UPDATE savings SET end_date = $1 WHERE user_id = $2`, new_end_date, userID)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Update End Date success!"
		daily_pay `+ fmt.Sprintf("%.2f", daily_pay) +`,
	}
	`), http.StatusOK)
}

// DeleteBookByID a function to remove book data from DB, given bookID
func (h *Handler) DeleteSavingByUserID(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userID, err := strconv.ParseInt(param.ByName("userID"), 10, 64)
	if err != nil {
		log.Printf("[internal][GetUserByID] invalid user_id :%+v\n", err)
		return
	}

	query := fmt.Sprintf("DELETE FROM savings WHERE user_id = %d", userID)
	_, err = h.DB.Exec(query)
	if err != nil {
		log.Println(err)
		return
	}

	renderJSON(w, []byte(`
	{
		status: "success",
		message: "Delete book success!"
	}
	`), http.StatusOK)
}
