package models

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Dupla struct {
	CodDupla    int         `json:"codDupla"`
	Nome        string      `json:"nome"`
	Logo        null.String `json:"logo"`
	Roster      Roster      `json:"roster"`
	Elo         float64     `json:"elo"`
	DataCriacao string      `json:"dataCriacao"`
	PaisOrg     string      `json:"paisOrg"`
}

type Duplas []Dupla

type Roster struct {
	CodRoster    int         `json:"codRoster"`
	CodDupla     int         `json:"codDupla"`
	CodIGL       int         `json:"codIGL"`
	CodJog2      int         `json:"codJog2"`
	CodJogBench  null.Int    `json:"codJogBench"`
	CodCoach     null.Int    `json:"codCoach"`
	InicioRoster string      `json:"inicioRoster"`
	FimRoster    null.String `json:"fimRoster"`
}

// Valida uma instância de Dupla
func (d Dupla) IsValid() (bool, error) {
	// TODD

	return true, nil
}

// Valida uma instância de Roster
func (r Roster) IsValid() (bool, error) {
	// TODO

	return true, nil
}

// Pega o ranking das duplas atualizado
func (d *Duplas) GetRankingDuplas() (int, error) {
	query := "SELECT * FROM Duplas ORDER BY elo DESC;"

	rows, err := E.DB.Query(query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao retornar ranking do banco de dados.")
	}

	for rows.Next() {
		var dupla Dupla
		err := rows.Scan(&dupla.CodDupla, &dupla.Nome, &dupla.Logo, &dupla.Elo,
			&dupla.DataCriacao, &dupla.PaisOrg)
		if err != nil {
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao retornar ranking do banco de dados.")
		}

		*d = append(*d, dupla)
	}

	defer rows.Close()

	return http.StatusOK, nil
}

// Pega todos os dados de uma dupla
func (d *Dupla) GetDuplaByID(codDupla int) (int, error) {
	// TODO

	return http.StatusOK, nil
}
