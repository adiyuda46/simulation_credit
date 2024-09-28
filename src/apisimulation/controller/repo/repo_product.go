package repo

import (
	"errors"
	"log"
	conn "simulation/src/apisimulation/controller/connection"
	modelApp "simulation/src/apisimulation/model"
)


func GetAllLob()([]modelApp.AllLob,error)  {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return []modelApp.AllLob{}, errConn
	}

	query := `SELECT * FROM public."LOB"`
	rows ,err := db.Query(query)
	if err != nil {
		return nil, errors.New("Failed to execute the query: " + err.Error())
	}

	var results []modelApp.AllLob
	for rows.Next() {
		var result modelApp.AllLob
		err := rows.Scan(&result.Id,&result.LobName,&result.Desc)
		if err != nil {
			return nil, errors.New("Failed to scan row: " + err.Error())
		}
		results = append(results, result)
	}
	
	return results,nil
}

func GetLobByid(id int)(string,error)  {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return "", errConn
	}
	query := `select "LOB_NAME" from public."LOB" where "ID" = $1`
	var result string
	row := db.QueryRow(query,id).Scan(&result)
	if row != nil {
		log.Printf("Execution failed: %v", row)
		return "", row
	}
	return result , nil
}
