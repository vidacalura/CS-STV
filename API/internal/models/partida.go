package models

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Partida struct {
	CodPrtd  int               `json:"codPrtd"`
	CodEvnt  int               `json:"codEvnt"`
	NomeEvnt string            `json:"evento"`
	TimeCasa DadosPartidaDupla `json:"timeCasa"`
	TimeFora DadosPartidaDupla `json:"timeFora"`
	BestOf   int               `json:"bestOf"`
	Mapas    string            `json:"mapas"`
	CodMVP   null.Int          `json:"codMvp"`
	NomeMVP  string            `json:"mvp"`
	DataJogo string            `json:"dataJogo"`
}

type Partidas []Partida

type DadosPartidaDupla struct {
	CodDupla int         `json:"codDupla"`
	Nome     string      `json:"nome"`
	LogoURL  null.String `json:"logoURL"`
	Pontos   int         `json:"pontos"`
}

// Valida uma instância de Partida
func (p Partida) IsValid() (bool, error) {
	// TODO

	return true, nil
}

// Valida uma instância de DadosPartidaDupla
func (d DadosPartidaDupla) IsValid() (bool, error) {
	// TODO

	return true, nil
}

// Pega as 8 partidas mais recentemente cadastradas no sistema
func (partidas *Partidas) GetPartidasRecentes() (int, error) {
	query := `
		SELECT * FROM Partidas
		ORDER BY data DESC
		LIMIT 8;`

	rows, err := E.DB.Query(query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
	}

	for rows.Next() {
		var p Partida
		err := rows.Scan(&p.CodPrtd, &p.CodEvnt, &p.TimeCasa.CodDupla, &p.TimeFora.CodDupla,
			&p.BestOf, &p.TimeCasa.Pontos, &p.TimeFora.Pontos, &p.Mapas, &p.CodMVP,
			&p.DataJogo)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		queryDupla := "SELECT nome, logo FROM Duplas WHERE cod_dupla = $1;"
		row := E.DB.QueryRow(queryDupla, p.TimeCasa.CodDupla)
		if err := row.Scan(&p.TimeCasa.Nome, &p.TimeCasa.LogoURL); err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		row = E.DB.QueryRow(queryDupla, p.TimeFora.CodDupla)
		if err := row.Scan(&p.TimeFora.Nome, &p.TimeFora.LogoURL); err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		*partidas = append(*partidas, p)
	}

	return http.StatusOK, nil
}

// Pega todas as partidas de um time
func (partidas *Partidas) GetPartidasByTimeID(codTime int) (int, error) {
	query := `
		SELECT * FROM Partidas
		WHERE time_casa = $1
		OR time_fora = $1;`

	rows, err := E.DB.Query(query, codTime)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
	}

	for rows.Next() {
		var p Partida
		err := rows.Scan(&p.CodPrtd, &p.CodEvnt, &p.TimeCasa.CodDupla, &p.TimeFora.CodDupla,
			&p.BestOf, &p.TimeCasa.Pontos, &p.TimeFora.Pontos, &p.Mapas, &p.CodMVP,
			&p.DataJogo)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		queryDupla := "SELECT nome, logo FROM Duplas WHERE cod_dupla = $1;"
		row := E.DB.QueryRow(queryDupla, p.TimeCasa.CodDupla)
		if err := row.Scan(&p.TimeCasa.Nome, &p.TimeCasa.LogoURL); err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		row = E.DB.QueryRow(queryDupla, p.TimeFora.CodDupla)
		if err := row.Scan(&p.TimeFora.Nome, &p.TimeFora.LogoURL); err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber partidas recentes do banco de dados.")
		}

		*partidas = append(*partidas, p)
	}


	return http.StatusOK, nil
}
