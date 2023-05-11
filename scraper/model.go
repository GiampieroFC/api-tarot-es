package scraper

import (
	"database/sql"
	"fmt"
	"log"
)

type modelCard struct {
	Id                        string  `json:"id"`
	Arcano                    *string `json:"arcano"`
	Número                    int16   `json:"numero"`
	Palo                      *string `json:"palo"`
	Nombre                    *string `json:"nombre"`
	Descripción               *string `json:"descripcion"`
	Significado_al_derecho    *string `json:"significado_al_derecho"`
	Significado_al_revés      *string `json:"significado_al_revés"`
	Interpretación_al_revés   *string `json:"interpretación_al_revés"`
	Interpretación_al_derecho *string `json:"interpretación_al_derecho"`
	Imagen                    *string `json:"imagen"`
}

type ModelCards []modelCard

type modeler interface {
	CreateTable(*sql.DB)
	FillTabla(db *sql.DB, data *modelCard)
}

type cards struct {
	modeler modeler
}

func NewCards() *cards {
	var m modeler
	return &cards{m}
}

func (c *cards) CreateTable(db *sql.DB) {
	stmt, err := db.Prepare(psql_table)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migración de 'cartas' ejecutada correctamente")
}

func (c *cards) FillTabla(db *sql.DB, data *modelCard) {
	stmt, err := db.Prepare(psql_insert_row)
	if err != nil {
		log.Fatal("FillTabla() prepare ", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		"DEFAULT",
		data.Arcano,
		intToNull(data.Número),
		stringToNull(*data.Palo),
		data.Nombre,
		data.Descripción,
		data.Significado_al_derecho,
		data.Significado_al_revés,
		data.Interpretación_al_revés,
		data.Interpretación_al_derecho,
		data.Imagen,
	).Scan()
	if err != nil {
		log.Fatal("FillTabla() QueryRow ", err)
	}

	fmt.Println("inserción de 'carta' ejecutada correctamente")

}
