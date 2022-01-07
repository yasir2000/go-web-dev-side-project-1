package cms

import (
	"database/sql"

	//Use the Postgres SQL driver
	_ "github.com/lib/pq"
)

var store = newDB()

type PgStore struct {
	DB *sql.DB
}

// a new instance of DB
func newDB() *PgStore {
	db, err := sql.Open("postgres", "user=postgres dbname=yasir2000 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &PgStore{
		DB: db,
	}
}

// takes page id, grab value to pointer struct from db that fulfill grabbed id, then produce page construct
func GetPage(id string) (*Page, error) {
	var p Page
	//Scan copies the columns from the matched row into the values pointed at by dest.
	//Scan for details. If more than one row matches the query,
	//Scan uses the first row and discards the rest. If no row matches the query, Scan returns ErrNoRows.
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id= $1", id).Scan(&p.ID, &p.Title, &p.Content)

	return &p, err
}

// GetPages is a new function that allows us to get every page from our database.
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// logic
	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}

//return values of a function dont need names, only types are enough
func CreatePage(p *Page) (int, error) {
	var id int
	// .Scan function assigns a value aand returns id from query to id variable
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}
