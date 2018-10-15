package main

import (
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (s *server) createUser(user *User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	sql := `
    INSERT INTO users(username, password)
    VALUES($1, $2)
    RETURNING id, username
  `
	err = s.db.QueryRow(sql, user.Username, string(password)).Scan(&user.ID, &user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) showUser(user *User) error {
	sql := `
		SELECT password, id, username
		FROM users
		WHERE username=$1;
	`
	var storedPassword string

	err := s.db.QueryRow(
		sql,
		user.Username).
		Scan(
			&storedPassword,
			&user.ID,
			&user.Username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(storedPassword),
		[]byte(user.Password))
	if err != nil {
		return err
	}

	return nil
}
