package repo

import (
	//"errors"
	"log"
	conn "simulation/src/apisimulation/controller/connection"
	modelApp "simulation/src/apisimulation/model"
)

func GetCatCar() ([]string, error) {
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return nil, errConn
	}

	query := `SELECT DISTINCT "P"."CATEGORY"
FROM public."CAR" AS "c"
INNER JOIN public."PRODUCT" AS "P" ON "c"."CATEGORY" = "P"."CATEGORY";`

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			log.Printf("Scan failed: %v", err)
			return nil, err
		}
		result = append(result, category)
	}

	return result, nil
}

func GetProductCar(category string) ([]string, error) {
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return nil, errConn
	}
	query := `SELECT DISTINCT "c"."PRODUCT_NAME"
FROM public."CAR" AS "c"
INNER JOIN public."PRODUCT" AS "P" ON "c"."CATEGORY" = "P"."CATEGORY"
WHERE "P"."CATEGORY" = $1`

	rows, err := db.Query(query, category)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, err
	}

	var results []string
	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err != nil {
			log.Printf("Scan failed: %v", err)
			return nil, err
		}
		results = append(results, result)

	}

	return results, nil

}

func GetPriceCar(req modelApp.Price) (string, error) {
	db, errConn := conn.ConnectToDatabase()
	if errConn != nil {
		log.Printf("Database connection failed: %v", errConn) // Log error
		return "", errConn
	}
	query := `SELECT "c"."PRICE"
FROM public."CAR" AS "c"
INNER JOIN public."PRODUCT" AS "P" ON "c"."CATEGORY" = "P"."CATEGORY"
WHERE "P"."CATEGORY" = $1 AND "c"."PRODUCT_NAME" = $2`

	var result string
	row := db.QueryRow(query, req.Category, req.ProductName).Scan(&result)
	if row != nil {
		log.Printf("Query failed: %v", row)
		return "", row
	}
	return result, nil
}
