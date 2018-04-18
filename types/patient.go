package types

import (
	"strconv"
	"math"
	"errors"
	"database/sql"
	"encoding/json"
)

type Patient struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Surname    NullString `json:"surname"`
	Lastname   NullString `json:"lastname"`
	Nif        string     `json:"nif"`
	BirthDate  NullString `db:"birth_date" json:"birth_date"`
	Age        NullInt64  `json:"age"`
	Gender     string     `json:"gender"`
	Profession NullString `json:"profession"`
	Hobbies    NullString `json:"hobbies"`
	Address    NullString `json:"address"`
	City       string     `json:"city"`
	Phone      NullString `json:"phone"`
	Email      NullString `json:"email"`
	PostalCode NullString `db:"postal_code" json:"postal_code"`
	CreatedAt  NullString `db:"created_at" json:"created_at"`
	UpdatedAt  NullString `db:"updated_at" json:"updated_at"`
	DeletedAt  NullString `db:"deleted_at" json:"deleted_at"`
}

const (
	DNILetters   = "TRWAGMYFPDXBNJZSQVHLCKE"
	ControlCheck = 23
)

func (p *Patient) ValidateDni() (bool, error) {
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

type NullString struct {
	sql.NullString
}

// NullString MarshalJSON interface redefinition
func (r NullString) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.String)
	} else {
		return json.Marshal(nil)
	}
}

type NullInt64 struct {
	sql.NullInt64
}

// NullInt64 MarshalJSON interface redefinition
func (r NullInt64) MarshalJSON() ([]byte, error) {
	if r.Valid {
		return json.Marshal(r.Int64)
	} else {
		return json.Marshal(nil)
	}
}
