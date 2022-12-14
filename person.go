package main

import (
	"context"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonParams struct {
	FirstName string
	LastName  string
}

// CreatePerson creates a person
func CreatePerson(ctx context.Context, params PersonParams) error {
	const sql = `INSERT INTO hack.person (first_name, last_name) VALUES ($1, $2)`
	_, err := db.ExecContext(ctx, sql, params.FirstName, params.LastName)
	return err
}

// UpdatePerson updates a person
func UpdatePerson(ctx context.Context, id int, params PersonParams) error {
	var sql string = fmt.Sprintf("UPDATE hack.person SET first_name = '%s', last_name = '%s' WHERE id = %d",
		params.FirstName,
		params.LastName,
		id)
	_, err := db.ExecContext(ctx, sql)
	return err
}

// ListPeople returns all people
func ListPeople(ctx context.Context) ([]Person, error) {
	const sql = `SELECT id, first_name, last_name FROM hack.person ORDER BY id`
	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(&i.ID, &i.FirstName, &i.LastName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
