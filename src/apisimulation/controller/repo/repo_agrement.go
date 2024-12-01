package repo

import (

	//modelApp "simulation/src/apisimulation/model"
	"log"
	conn "simulation/src/apisimulation/controller/connection"
	modelApp "simulation/src/apisimulation/model"
)

// GetListaAgrement retrieves a list of agreements for a given user ID
func GetListaAgrement(userId int) ([]modelApp.ListAgreement, error) {
	// Connect to DB
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return nil, errConn
	}
	defer db.Close() // Ensure the database connection is closed after the function returns

	query := `
	SELECT "AGREMENT", "AMOUNT_INSTALMENT", "PRODUCT", "INSTALMENT", "DUE_DATE"
	FROM public."AGREMENT"
	WHERE "USER_ID" = $1
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing

	var agreements []modelApp.ListAgreement
	for rows.Next() {
		var ag modelApp.ListAgreement
		if err := rows.Scan(&ag.Agrement, &ag.AmountInstalment, &ag.Product, &ag.Instalment, &ag.DueDate); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		agreements = append(agreements, ag)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating rows: %v", err)
		return nil, err
	}

	return agreements, nil
}