package dice

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Register adds a new Die type to the database.
//
// TODO: Implement Register
func Register(d Die) error {
	return nil
}

// NewDatabase creates a new database to work with.
//
// TODO: implement NewDatabase
func NewDatabase() (db r :=sql.DB, err error) {
	connStr := "user=postgres dbname=postgres sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
            return
        }
}

func main() {

	// age := 21
	// rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	// …
}
