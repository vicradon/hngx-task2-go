package personcrud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapis/personcrudapi/database"

	"github.com/gorilla/mux"
)

type ResponseData struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Person struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreatePersonRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func getIdParam(r *http.Request) int64 {
	vars := mux.Vars(r)
	userID := vars["id"]

	intValue, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return intValue
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM persons")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := ResponseData{
		Message: "Successfully fetched persons",
		Status:  "success",
		Data:    rows,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func CreateNewPerson(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	var requestBody CreatePersonRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	insertQuery := "INSERT INTO persons (first_name, last_name, email) VALUES (?, ?, ?)"
	lastInsertID, err := db.ExecAndGetLastInsertID(insertQuery, requestBody.FirstName, requestBody.LastName, requestBody.Email)
	if err != nil {
		http.Error(w, "Error inserting into the database", http.StatusBadRequest)
		return
	}

	createdPerson := Person{
		Id:        lastInsertID,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
	}

	responseData := ResponseData{
		Message: "Successfully created a person",
		Status:  "success",
		Data:    createdPerson,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetPersonDetails(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	id := getIdParam(r)

	getQuery := "SELECT id, first_name, last_name, email FROM persons WHERE id = ?"

	rows, err := db.Query(getQuery, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := ResponseData{
		Message: "Successfully fetched person details",
		Status:  "success",
		Data:    rows[0],
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func UpdatePersonDetails(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	id := getIdParam(r)
	var requestBody CreatePersonRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	updateQuery := "UPDATE persons SET first_name = ?, last_name = ?, email = ? WHERE id = ?"
	if err := db.Exec(updateQuery, requestBody.FirstName, requestBody.LastName, requestBody.Email, id); err != nil {
		http.Error(w, "Error inserting into the database", http.StatusBadRequest)
		return
	}

	updatedPerson := Person{
		Id:        id,
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
	}

	responseData := ResponseData{
		Message: "Successfully updated person details",
		Status:  "success",
		Data:    updatedPerson,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	id := getIdParam(r)

	deleteQuery := "DELETE FROM persons WHERE id = ?"
	if err := db.Exec(deleteQuery, id); err != nil {
		http.Error(w, "Error deleting person", http.StatusBadRequest)
		return
	}

	responseData := ResponseData{
		Message: "Successfully deleted a person",
		Status:  "success",
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
