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
	Rank        int         `json:"rank"`
	Elo         float64     `json:"elo"`
	DataCriacao string      `json:"dataCriacao"`
	PaisOrg     string      `json:"paisOrg"`
	Ativo       bool        `json:"ativo"`
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

type Rosters []Roster

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
	query := "SELECT * FROM Duplas WHERE ativo = TRUE ORDER BY elo DESC LIMIT 5;"

	rows, err := E.DB.Query(query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao retornar ranking do banco de dados.")
	}

	i := 1
	for rows.Next() {
		var dupla Dupla
		dupla.Rank = i

		err := rows.Scan(&dupla.CodDupla, &dupla.Nome, &dupla.Logo, &dupla.Elo,
			&dupla.DataCriacao, &dupla.PaisOrg, &dupla.Ativo)
		if err != nil {
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao retornar ranking do banco de dados.")
		}

		*d = append(*d, dupla)
		i++
	}

	defer rows.Close()

	return http.StatusOK, nil
}

// Pega todos os dados de uma dupla
func (d *Dupla) GetDuplaByID(codDupla int) (int, error) {
	queryDupla := "SELECT * FROM Duplas WHERE cod_dupla = $1;"

	row := E.DB.QueryRow(queryDupla, codDupla)
	err := row.Scan(&d.CodDupla, &d.Nome, &d.Logo, &d.Elo, &d.DataCriacao,
		&d.PaisOrg)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound,
			fmt.Errorf("Dupla não existente ou não encontrada.")
	}

	queryRoster := `
		SELECT * FROM DuplasRoster
		WHERE cod_dupla = ?
		ORDER BY inicio_roster DESC
		LIMIT 1;`
	row = E.DB.QueryRow(queryRoster, codDupla)
	err = row.Scan(&d.Roster.CodRoster, &d.Roster.CodDupla, &d.Roster.CodIGL,
		&d.Roster.CodJog2, &d.Roster.CodJogBench, &d.Roster.CodCoach,
		&d.Roster.InicioRoster, &d.Roster.FimRoster)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound,
			fmt.Errorf("Dupla não possui jogadores em atividade.")
	}

	return http.StatusOK, nil
}

// Pega todas as Rosters de uma dupla
func (r *Rosters) GetAllRosters(codDupla int) (int, error) {
	// TODO

	return http.StatusOK, nil
}
