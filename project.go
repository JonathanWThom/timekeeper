package main

// Project lives on the projects table
type Project struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Code string `json:"code" db:"code"`
}
