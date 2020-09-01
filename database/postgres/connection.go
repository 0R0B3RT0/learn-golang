package main

func main() {
	Connect()
}

func Connect() {
	dbsource := "postgres://postgres:password@example.com/testdb?sslmode=verify-full"
	conn, err := sqlx.Connect("postgres", dbsource)
	if err != nil {
		fmt.Printf("error open db: %v\n", err)
		return
	}
}
