package models

import (
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type PlayerOfTheWeek struct {
	CodPoftw     int         `json:"codPoftw"`
	CodJog       int         `json:"codJog"`
	Dados        null.Float  `json:"dados"`
	InfoDados    string      `json:"infoDados"`
	SemanaInicio string      `json:"semanaInicio"`
	SemanaFim    null.String `json:"semanaFim"`
}

// Valida uma instância de PlayerOfTheWeek
func (p PlayerOfTheWeek) IsValid() (bool, error) {
	// TODO

	return true, nil
}

// Retorna Player of the week da última semana
func (p *PlayerOfTheWeek) GetUltimoPlayerOfTheWeek() (int, error) {
	// TODO

	return http.StatusOK, nil
}
