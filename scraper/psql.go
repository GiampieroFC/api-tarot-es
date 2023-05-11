package scraper

const (
	psql_table = `CREATE TABLE IF NOT EXISTS cartas(
		id UUID DEFAULT gen_random_uuid() NOT NULL,
		arcano VARCHAR(30),
		número INT,
		palo VARCHAR(30),
		nombre VARCHAR(30),
		descripción TEXT,
		significado_al_derecho TEXT,
		significado_al_revés       TEXT,
		interpretación_al_revés    TEXT,
		interpretación_al_derecho  TEXT,
		imagen TEXT,
		CONSTRAINT cartas_id_pk PRIMARY KEY (id)
		);`

	psql_insert_row = `INSERT INTO cartas(
		id,
		arcano,
		número,
		palo,
		nombre,
		descripción,
		significado_al_derecho,
		significado_al_revés,
		interpretación_al_revés,
		interpretación_al_derecho,
		imagen) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11) RETURNING cartas.nombre;`

	psql_get_all       = `SELECT * FROM cartas ORDER BY número;`
	psql_get_by_number = `SELECT * FROM cartas WHERE número = $1;`
	psql_get_menor     = `SELECT * FROM cartas WHERE arcano = 'Menor' ORDER BY número;`
	psql_get_mayor     = `SELECT * FROM cartas WHERE arcano = 'Mayor' ORDER BY número;`
	psql_get_by_palo   = `SELECT * FROM cartas WHERE palo = $1 ORDER BY número;`
	psql_get_random    = `SELECT * FROM cartas ORDER BY RANDOM() LIMIT $1;`
)
