package types

import (
	"strconv"
	"math"
	"errors"
	"database/sql"
)

type Patient struct {
	Id         int
	Name       string
	Surname    sql.NullString
	Lastname   sql.NullString
	Nif        string
	BirthDate  []uint8        `db:"birth_date"`
	Age        sql.NullInt64
	Gender     string
	Profession sql.NullString
	Hobbies    sql.NullString
	Address    sql.NullString
	City       string
	Phone      sql.NullString
	Email      sql.NullString
	PostalCode sql.NullString `db:"postal_code"`
	CreatedAt  sql.NullString `db:"created_at"`
	UpdatedAt  sql.NullString `db:"updated_at"`
	DeletedAt  sql.NullString `db:"deleted_at"`
}

const (
	DNILetters   = "TRWAGMYFPDXBNJZSQVHLCKE"
	ControlCheck = 23
)

func (p *Patient) ValidateDni() (bool, error){
	dni := p.Nif
	numbers, err := strconv.Atoi(dni[:len(dni)-1])
	if err != nil {
		return false, err
	}

	n := float64(numbers)
	mod := math.Mod(n, ControlCheck)

	modString := strconv.FormatFloat(mod, 'f', -1, 64)
	index, err := strconv.Atoi(modString)
	if err != nil {
		return false, err
	}

	if string(DNILetters[index]) != string(dni[len(dni)-1]) {
		return false, errors.New("")
	}

	return true, nil
}
