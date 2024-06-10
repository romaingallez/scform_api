package models

import (
	"strconv"
)

type Matiere struct {
	NomMatiere string
	Notes      []Note
}

type Note struct {
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Type    string `json:"type"`
	Note    string `json:"note"`
	NoteMax string `json:"noteMax"`
	Coeef   string `json:"coeef"`
}

// NoteInt represents the note struct with Note and NoteMax as integers
type NoteFloat struct {
	Name    string  `json:"name"`
	Desc    string  `json:"desc"`
	Type    string  `json:"type"`
	Note    float64 `json:"note"`
	NoteMax float64 `json:"noteMax"`
	Coeef   string  `json:"coeef"`
}

func (n *Note) ConvertToFloat() (*NoteFloat, error) {
	// test if note as . or ,

	note, err := strconv.ParseFloat(n.Note, 64)
	if err != nil {
		return nil, err
	}
	noteMax, err := strconv.ParseFloat(n.NoteMax, 64)
	if err != nil {
		return nil, err
	}
	return &NoteFloat{
		Name:    n.Name,
		Desc:    n.Desc,
		Type:    n.Type,
		Note:    note,
		NoteMax: noteMax,
		Coeef:   n.Coeef,
	}, nil
}
