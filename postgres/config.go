package main
import (
	"fmt"
	// "database/sql"
	// _ "github.com/lib/pq"
	 "github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
	// "log"
)
type config struct {
	name string
	user string
	pass string
	host string
	port int
	dbname string
	sslmode string
}
func NewDBConn(con config) (db *pg.DB) {
	address := fmt.Sprintf("%s:%s", con.host, con.port)
	options := &pg.Options{
		User: con.user,
		Password: con.pass,
		Addr: address,
		Database: con.name,
		PoolSize: 50,
	}
	db = pg.Connect(options)
	if db == nil{
		fmt.Println("cannot connect")
	}
	return db
}
type Post struct {
	ID int	`pg:"id,pk"`
	Name string
	Email string
}
func InsertDB(pg *pg.DB, post Post) error {
	_, error := pg.Model(&post).Insert()
	return error
}
func main() {
	configs := config{
		name: "postgres",
		user: "postgres",
		pass: "postgres",
		host: "127.0.0.1",
		port: 5432,
		dbname: "postgres",
		sslmode: "disable" }
	

	// _, err := sql.Open(configs.name, fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=%s", configs.user, configs.pass, configs.host, configs.port, configs.dbname, configs.sslmode))
	db := NewDBConn(configs)
	fmt.Println(db)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("The connection to db is success")
	// }
	p := Post{
		ID: 5,
		Name: "vijay",
		Email: "a@gmai.com" }
	defer db.Close()
	err := InsertDB(db, p)
	if err!=nil {
		fmt.Println(err)
	}
}