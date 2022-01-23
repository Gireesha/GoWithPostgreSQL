package main

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Customer struct {
	customerid int
	FirstName  string
	LastName   string
}

const (
	host     = "localhost"  //hostname(IP Address) of the server
	port     = 5432         //TCP/IP port the server is listening on (by default, 5432)
	user     = "postgres"   //PostgreSQL username to use when connecting to the database
	password = "myPa$$word" //database password
	dbname   = "customer"   //database name
	sslmode  = "disable"    //must be set to disabled unless using SSL
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	// query 1 row from the table customer.customer
	var activeCustomer Customer

	customerSql := "SELECT customerid, firstname, lastname FROM customer.customer WHERE customerid = $1"
	err = db.QueryRow(customerSql, 1).Scan(&activeCustomer.customerid, &activeCustomer.FirstName, &activeCustomer.LastName)
	if err != nil {
		log.Fatal("Failed to fetch first row: ", err)
	}

	fmt.Printf("First Record belongs to %s!\n", activeCustomer.FirstName)

	rows, err := db.Query("SELECT * FROM customer.customer FETCH FIRST 10 ROWS ONLY")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := make([]Customer, 0)

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.customerid, &customer.FirstName, &customer.LastName)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		customers = append(customers, customer)
	}

	for _, customer := range customers {
		fmt.Printf("%d %s %s \n", customer.customerid, customer.FirstName, customer.LastName)
	}

	// update row
	sqlUpdateStmts := `
    UPDATE customer.customer
    SET LastName = $1 WHERE customerid = $2;`
	res, err := db.Exec(sqlUpdateStmts, "Clark", 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Row updated: %v\n", count)

	// delete row
	sqlDelete := `
	  DELETE FROM customer.customer
	  WHERE customerid = $1;`
	res1, err := db.Exec(sqlDelete, 2)
	if err != nil {
		panic(err)
	}
	count, err = res1.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Row deleted: %v\n", count)

	// insert row
	sqlInsertStmt := `
		INSERT INTO customer.customer(
	 firstname, lastname)
		VALUES ($1,$2);`
	resInsert, err := db.Exec(sqlInsertStmt, "Henry", "Ford")
	if err != nil {
		panic(err)
	}
	count, err = resInsert.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("1 row inserted: %v\n", count)

	//calling a SP with parameters
	resSProc, err := db.Exec("CALL customer.insertcustomer('Thomas','Jefferson')")
	if err != nil {
		panic(err)
	}
	count, err = resSProc.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("SP called successfully %v", count)

}
