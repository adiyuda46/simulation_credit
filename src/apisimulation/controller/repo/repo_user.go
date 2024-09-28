package repo

import (

	//modelApp "simulation/src/apisimulation/model"
	"log"
	conn "simulation/src/apisimulation/controller/connection"
	"simulation/src/apisimulation/controller/utils"
	modelApp "simulation/src/apisimulation/model"

)

func RegisterRepository(name, password, email, phone string) error {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn)
		return errConn
	}
	newPassword, _ := utils.HashPassword(password)

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

	_, err := db.Exec(query, name, newPassword, email, phone)
	if err != nil {
		log.Printf("Execution failed: %v", err)
		return err
	}
	return nil
}

func RegisterCheck(email, phone string) (int, error) {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return 0, errConn
	}
	var result int
	query := `SELECT COUNT("ID") FROM public."USER" where "EMAIL" = $1 or "PHONE_NUMBER" = $2`
	row := db.QueryRow(query, email, phone).Scan(&result)
	if row != nil {
		log.Printf("Execution failed: %v", row)
		return 0, row
	}
	return result, nil
}

func CheckPhoneNumber(phone string) (string, error) {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return "", errConn
	}
	query := `SELECT "PASSWORD" FROM public."USER" WHERE "PHONE_NUMBER" = $1`

	var result string
	row := db.QueryRow(query, phone).Scan(&result)
	if row != nil {
		log.Printf("Execution failed: %v", row)
		return "", row
	}
	return result, nil
}
func Supabase() (string,error) {
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return "", errConn
	}
	var result string
	query := `SELECT "nama" FROM public."user"`
	row := db.QueryRow(query).Scan(&result)
	if row != nil {
		log.Printf("Execution failed: %v", row)
		return "", row
	}
	return result, nil
}

func GetUserbyPhone(phone string) (*modelApp.User, error) {
    // Connect to DB
    db, errConn := conn.ConnectToDatabase()
    if errConn != nil {
        log.Printf("Database connection failed: %v", errConn) // Log error
        return nil, errConn // Kembalikan nil jika terjadi kesalahan
    }

    query := `select "ID","NAME","PASSWORD","EMAIL" from public."USER" where "PHONE_NUMBER" = $1`

    var user modelApp.User
    row := db.QueryRow(query, phone).Scan(&user.Id, &user.Name, &user.Password, &user.Email)
    
    if row != nil {
        log.Printf("Query failed: %v", row)
        return nil, row
    }

    return &user, nil // Kembalikan pointer ke pengguna yang ditemukan
}

func GetAgrement(userId string)()  {
	
}