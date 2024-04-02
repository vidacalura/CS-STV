package models

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/guregu/null.v3"
)

type Noticia struct {
	CodNotc   int         `json:"codNotc"`
	Titulo    string      `json:"titulo"`
	Subtitulo null.String `json:"subtitulo"`
	Noticia   string      `json:"noticia"`
	Data      string      `json:"data"`
}

type FeedNoticias []Noticia

// Valida uma instância de Noticia
func (n Noticia) IsValid() (bool, error) {
	// TODO
	
	return true, nil
}

// Pega as 10 notícias mais recentes
func (f *FeedNoticias) GetFeedNoticias() (int, error) {
	query := `
		SELECT cod_notc, titulo, subtitulo, data FROM Noticias
		ORDER BY data DESC
		LIMIT 10;`

	rows, err := E.DB.Query(query)
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError,
			fmt.Errorf("Erro ao receber feed de notícias.")
	}

	for rows.Next() {
		var n Noticia
		err := rows.Scan(&n.CodNotc, &n.Titulo, &n.Subtitulo, &n.Data)
		if err != nil {
			log.Println(err)
			return http.StatusInternalServerError,
				fmt.Errorf("Erro ao receber feed de notícias.")
		}

		*f = append(*f, n)
	}

	return http.StatusOK, nil
}
