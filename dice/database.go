package dice

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"sync"

	_ "github.com/lib/pq" // Postgres Driver for sql package
)

// DB is the standard database.
var DB *Database

func init() {
	var err error
	DB = new(Database)
	DB.Mut = sync.Mutex{}
	if err = DB.Connect(); err != nil {
		panic(err)
	}
}

// Database is a postgresql database.
type Database struct {
	DB  *sql.DB
	Mut sync.Mutex
}

//IsInit returns whether or not the database has been initialized.
func (db *Database) IsInit() bool {
	return db.DB != nil
}

// Init initializes the database, if it hasn't already been initialized.
func (db *Database) Init() error {
	if !db.IsInit() {
		var err error
		if db, err = NewDatabase(); err != nil {
			return err
		}
		// db = database
	}
	return nil
}

// NewTableTemplate contains the Postgresql code for creating a roll table.
func NewTableTemplate() *template.Template {
	t := template.New("NewTableTemplate").Funcs(map[string]interface{}{
		// add simply sums the inputs.
		"add": func(a, b int) int { return a + b },
		// last returns whether or not the current index is the last in the slice.
		"last": func(i int, d Dice) bool { return i+1 != len(d) },
	})
	return template.Must(t.Parse(newTableTemplate))
}

// newTableTemplate holds the template for creating a new roll table in sql.
const newTableTemplate = `{{$Dice := .Dice}}WITH{{range $idx, $elem := .Dice}} t_{{add $idx 1}}(d{{add $idx 1}}_side) AS (
	SELECT side FROM dice.sides WHERE die_id=(SELECT die_id FROM dice.dice_types WHERE name='{{.Name}}')
){{if last $idx $Dice}},{{end}}{{end}}
SELECT
    row_number() over (ORDER BY {{range $idx, $elem := .Dice}}d{{add $idx 1}}_side{{if last $idx $Dice}},{{end}} {{end}}ASC) as roll_id,
    {{range $idx, $elem := .Dice}}t_{{add $idx 1}}.d{{add $idx 1}}_side{{if last $idx $Dice}},
    {{end}}{{end}}
INTO dice.t_{{.Name}}
FROM{{range $idx, $elem := .Dice}}{{if $idx}}
CROSS JOIN{{end}} t_{{add $idx 1}}{{end}};`

// NewTable creates a new roll table for n dice of type d.
//
// TODO: Make NewTable private and generate tables only when called upon.
func (db *Database) NewTable(d Die, n int) error {
	// TODO: Update NewTable to check for registered dice.
	if err := db.RegisterDieType(d); err != nil {
		return err
	}

	var args = struct {
		Name string
		Dice
	}{
		fmt.Sprintf("%dd%d", n, len(d.Sides())),
		Copy(d, n),
	}
	var query bytes.Buffer
	if err := NewTableTemplate().Execute(&query, args); err != nil {
		return err
	}
	if _, err := db.DB.Exec(query.String()); err != nil {
		return err
	}
	return nil
}

// RegisterDieType adds a new Die type to the database.
func (db *Database) RegisterDieType(d Die) error {
	if err := db.Init(); err != nil {
		return err
	}

	name := d.Name()

	// Proccess our die name to make it more SQL friendly.
	shortName := strings.ToLower(strings.Replace(name, " ", "_", -1))

	// Removes all special characters.
	regexp.MustCompile(`[^_A-Z0-9]`).ReplaceAllString(shortName, "")

	// Add the die type to the die_types table.
	sides := d.Sides()
	if _, err := db.DB.Exec("INSERT INTO dice.dice_types (name, short, sides) "+
		"VALUES ($1, $2, $3);", name, shortName, len(sides)); err != nil {
		return err
	}

	// Add the sides to the sides table.
	var query strings.Builder
	query.WriteString(`
INSERT INTO dice.sides (die_id, side, value)
SELECT (SELECT die_id from dice.dice_types WHERE name=$1), * FROM (VALUES`)
	for i, s := range sides {
		query.WriteString(fmt.Sprintf("(%d,%d)", i+1, s.Value()))
		if i+1 != len(sides) {
			query.WriteString(", ")
		}
	}
	query.WriteString(`) as t;`)
	_, err := db.DB.Exec(query.String(), name)
	return err
}

// PSQL is a string of Postgresql code.
type PSQL string

// NewDatabasePSQL is the postgreSQL code that creates a new database.
const NewDatabasePSQL PSQL = `DROP SCHEMA dice CASCADE;
CREATE SCHEMA dice;
SET search_path = dice;

CREATE TABLE dice_types (
	die_id SERIAL PRIMARY KEY
	,name TEXT NOT NULL
	,short text NOT NULL
	,sides INT NOT NULL CHECK (sides > 0)
);

CREATE TABLE sides (
	side_id SERIAL PRIMARY KEY
	,die_id INT REFERENCES dice_types(die_id) NOT NULL
	,side INT CHECK (side > 0)
	,value INT NOT NULL DEFAULT 0
);`

func (db *Database) Connect() error {
	connStr := "user=postgres dbname=postgres sslmode=disable"

	var err error
	db.DB, err = sql.Open("postgres", connStr)
	return err
}

// NewDatabase creates a new database to work with.
func NewDatabase() (db *Database, err error) {

	db = new(Database)
	db.Mut = sync.Mutex{}

	if err = db.Connect(); err != nil {
		return db, err
	}

	_, err = db.DB.Exec(string(NewDatabasePSQL))
	return db, err
}

// Roll enumerates all possible rolls of d, using fn to decide when to reroll,
// rerolling up to "rerolls" times.
func (d *Database) Roll(dice *Dice, fn Reroll, rerolls int) Table {

	// full tracks all rolls.
	full := `SELECT * FROM t_1d6`
	/*
		// next tracks the next table in the chain.
		// next := full

		// Keep re-rolling until all our rolls are perfect or we run out of rerolls.
			for i := 0; i < rerolls && next != ""; i++ {
				next = func(next string) string {
					next = strings.TrimPrefix(next, " UNION ALL ")
					rows, err := d.DB.Query(next)
					if err != nil {
						log.Println(err)
						// panic(err)
					}
					defer rows.Close()
					next = ""

					// var keep int
					for rows.Next() {
						var roll_id, side int
						if err := rows.Scan(&roll_id, &side); err != nil {
							log.Println(err)
							// panic(err)
						}

						dice.Roll()
						re := fn(*dice)
						if re[0] {
							next += " UNION ALL SELECT * FROM t_1d6"
						}
					}
					return next
				}(next)
				full += next
			}
	*/

	full = `
	WITH b as (
		SELECT value, side FROM dice.sides where die_id=1
	),
	total as ( SELECT count(*) FROM dice.t_1d6 )
	SELECT a.*, b.value, count(*), round(count(*)::NUMERIC / (select * FROM total), 4) as p FROM
	dice.t_1d6 as a
	INNER JOIN b ON (b.side=a.d1_side)
	GROUP BY value, d1_side, roll_id
	ORDER BY roll_id;
`
	// rows, err := d.DB.Query(full)
	rows, err := d.DB.Query(full)
	if err != nil {
		panic(err)
	}

	out := Table{Dice: *dice}
	for rows.Next() {
		var rollID, d1Side, value, count int
		var p float32
		if err := rows.Scan(&rollID, &d1Side, &value, &count, &p); err != nil {
			panic(err)
		}
		out.Rows = append(out.Rows, Row{Sides: []int{d1Side}, Count: count})
	}

	return out
}

// Table contains the results of a Dice Enumeration.
type Table struct {
	Dice []Die // The dice used in this table.
	Rows []Row // The rolls, which may or may not be grouped.
}

// Row holds a row or grouped set of rows in a DiceEnum table.
type Row struct {
	// Sides holds The Up side of each Die in Table.Die
	Sides []int
	// Count holds the number of times this roll shows up in the table.
	// Usually once, but can be more if results have been grouped.
	Count int
}
