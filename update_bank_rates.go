package main

import "fmt"

func updateBankRates() {
	updateTinkoffRates()
}

func saveBankRates(models []Rate) {

	//db, err := sql.Open(getDatabaseSecrets())
	//tx, err := db.Begin()
	//defer tx.Rollback()
	//defer db.Close()
	//
	//if err != nil {
	//	log.Error(err)
	//	return
	//}

	for _, model := range models {
		fmt.Println(model)
		//tx.Exec("INSERT INTO rates (model, company, price) VALUES ('iPhone X', $1, $2)",
		//	"Apple", 72000)
	} /*&& categoryCondition*/

	//tx.Commit()
}
