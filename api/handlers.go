package api

import (
	"encoding/json"
	"net/http"
)

// GetRecords
func GetRecords(w http.ResponseWriter, r *http.Request) {
	db := DBconn()

	rows, err := db.Query("SELECT * FROM records;")
	if err != nil {
		return
	}

	results := []Record{}
	setRecords := Record{}

	i := 0
	for rows.Next() {
		var id, date, start, end, breakTime, description string
		rows.Scan(&id, &date, &start, &end, &breakTime, &description)

		i++

		setRecords.ID = id
		setRecords.Date = date
		setRecords.Start = start
		setRecords.End = end
		setRecords.BreakTime = breakTime
		setRecords.Description = description
		results = append(results, setRecords)
	}

	json.NewEncoder(w).Encode(results)
}

// GetRecord
func GetRecord(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r) // Get params
	// // Loop through books and find with id
	// for _, item := range records {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }

	// json.NewEncoder(w).Encode(&Record{})
}

// CreateRecord
func CreateRecord(w http.ResponseWriter, r *http.Request) {

}

// DeleteRecord
func DeleteRecord(w http.ResponseWriter, r *http.Request) {

}

// UpdateRecord
func UpdateRecord(w http.ResponseWriter, r *http.Request) {

}
