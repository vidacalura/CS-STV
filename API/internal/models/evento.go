package models

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Evento struct {
	CodEvnt int         `json:"codEvnt"`
	Nome    string      `json:"evento"`
	Inicio  string      `json:"inicio"`
	Fim     null.String `json:"fim"`
}

type Eventos []Evento

// Valida uma instância de Evento
func (e Evento) IsValid() (bool, error) {
	// TODO

	return true, nil
}

// Pega os eventos mais recentes em ordem
func (e *Eventos) GetEventosRecentes() (int, error) {
	query := `
		SELECT * FROM Eventos
		ORDER BY inicio DESC;`

	rows, err := E.DB.Query(query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao receber eventos recentes do banco de dados.")
	}

	for rows.Next() {
		var evnt Evento
		err := rows.Scan(&evnt.CodEvnt, &evnt.Nome, &evnt.Inicio, &evnt.Fim)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber eventos recentes do banco de dados.")
		}

		*e = append(*e, evnt)
	}

	return http.StatusOK, nil
}
