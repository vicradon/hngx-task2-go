package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(dbPath string) (*Database, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Exec(query string, args ...interface{}) error {
	_, err := d.db.Exec(query, args...)
	return err
}

type RowJSON map[string]interface{}

func (d *Database) Query(query string, args ...interface{}) ([]RowJSON, error) {
	rows, err := d.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get column names from the rows.
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Create a slice to store the result rows.
	result := make([]RowJSON, 0)

	// Create a buffer to scan values into.
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	// Iterate over rows and scan values into the buffer.
	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		// Create a map to store row data.
		row := make(RowJSON)

		// Iterate over the columns and map values to column names.
		for i, colName := range columns {
			val := *(values[i].(*interface{}))
			row[colName] = val
		}

		// Append the row to the result slice.
		result = append(result, row)
	}

	// Check for any errors during iteration.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
