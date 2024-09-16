package repo

import (

	//modelApp "simulation/src/apisimulation/model"
	"log"
	conn "simulation/src/apisimulation/controller/connection"
)

func RegisterRepository(name, password, email, phone string) error {
    // Connect to DB
    db, errConn := conn.ConnectToDatabase()
    if errConn != nil {
        log.Printf("Database connection failed: %v", errConn) 
        return errConn
    }

    query := `INSERT INTO public."USER"
(
    "NAME",
    "PASSWORD",
    "EMAIL",
    "DTM_CRT",
    "DTM_UPD",
    "PHONE_NUMBER"
)
VALUES (
    $1,
    $2,
    $3,
    NOW(),
    NOW(),
    $4)`

    _, err := db.Exec(query, name, password, email, phone)
    if err != nil {
        log.Printf("Execution failed: %v", err) 
        return err 
    }
    return nil
}

func RegisterCheck(email,phone string)(int,error)  {
	// Connect to DB
    db, errConn := conn.ConnectToDatabase()
    if errConn != nil {
        log.Printf("Database connection failed: %v", errConn) // Log error
        return 0,errConn
    }
var result int
	query := `SELECT COUNT("ID") FROM public."USER"
where "EMAIL" = $1 or "PHONE_NUMBER" = $2`
 row := db.QueryRow(query,email,phone).Scan(&result)
 if row != nil {
	log.Printf("Execution failed: %v", row) 
	return 0,row 
 }
 return result,nil
}