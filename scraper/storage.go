package scraper

// Conect to
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

type dBaser interface {
	Connect()
	GetPool() *sql.DB
	GetAll() (ModelCards, error)
}

type DataBase struct {
	dbaser dBaser
}

func NewDB() *DataBase {
	var d dBaser
	return &DataBase{d}
}

func (*DataBase) Connect() {

	connectionString := os.Getenv("CONNECTION_STRING")

	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// GetPool return a unique instance of db
func (*DataBase) GetPool() *sql.DB {
	return db
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowProduct(s scanner) (modelCard, error) {
	card := modelCard{}

	// palo := stringToNull(card.Palo)

	err := s.Scan(
		&card.Id,
		&card.Arcano,
		&card.Número,
		&card.Palo,
		&card.Nombre,
		&card.Descripción,
		&card.Significado_al_derecho,
		&card.Significado_al_revés,
		&card.Interpretación_al_revés,
		&card.Interpretación_al_derecho,
		&card.Imagen,
	)
	if err != nil {
		return modelCard{}, err
	}

	return card, nil
}

func (*DataBase) GetAll() ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_all)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_all): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query(): %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		return nil, err
	}

	return cards, nil
}

func (*DataBase) GetMenor() ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_menor)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_menor): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query() (psql_get_menor): %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		return nil, err
	}

	return cards, nil
}

func (*DataBase) GetRandom(r string) ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_random)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_random): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(r)
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query() (psql_get_random): %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		return nil, err
	}

	return cards, nil
}

func (*DataBase) GetMayor() ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_mayor)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_mayor): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query() (psql_get_mayor): %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		return nil, err
	}

	return cards, nil
}

func (*DataBase) GetByNumber(numero string) ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_by_number)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_by_number): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(numero)
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query() GetByNumber: %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		log.Fatalf("rows.Err() GetByNumber: %v", err)
		return nil, err
	}

	return cards, nil
}

func (*DataBase) GetByPalo(palo string) ([]modelCard, error) {
	stmt, err := db.Prepare(psql_get_by_palo)
	if err != nil {
		stmt.Close()
		log.Fatalf("db.Prepare(psql_get_by_palo): %v", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(palo)
	if err != nil {
		stmt.Close()
		log.Fatalf("stmt.Query(palo) GetByPalo: %v", err)
		return nil, err
	}
	defer rows.Close()

	var cards ModelCards

	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			stmt.Close()
			return nil, err
		}
		cards = append(cards, m)
	}

	if err := rows.Err(); err != nil {
		stmt.Close()
		log.Fatalf("rows.Err() GetByPalo: %v", err)
		return nil, err
	}

	return cards, nil
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func intToNull(i16 int16) sql.NullInt16 {
	null := sql.NullInt16{Int16: i16}
	if null.Int16 != 0 {
		null.Valid = true
	}
	return null
}
