package models

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type PlayerOfTheWeek struct {
	CodPoftw     int         `json:"codPoftw"`
	SteamId      int         `json:"steamId"`
	Nome         string      `json:"nome"`
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

// Retorna o Player of the week da última semana
func (p *PlayerOfTheWeek) GetUltimoPlayerOfTheWeek() (int, error) {
	query := `	
		SELECT *
		FROM PlayerOfTheWeek
		ORDER BY cod_poftw DESC
		LIMIT 1;`

	row := E.DB.QueryRow(query)
	err := row.Scan(&p.CodPoftw, &p.CodJog, &p.Dados, &p.InfoDados, &p.SemanaInicio,
		&p.SemanaFim, &p.Nome)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao receber Jogador da semana do banco de dados.")
	}

	return http.StatusOK, nil
}
