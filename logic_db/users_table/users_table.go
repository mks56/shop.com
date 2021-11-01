package users_table

import (
	"database/sql"

	"shop.com/err"
	struct_data "shop.com/struct_data"
)

func InsertNewUser(db *sql.DB, person struct_data.Person) error {
	insertUser := `insert into "UsersTab"("account", "password", "type") values($1, $2, $3)`
	var result error
	_, result = db.Exec(insertUser, person.Account, person.Password,
		person.TypePerson)
	return result
}

func UpdatePasswordByAccount(db *sql.DB, account, newPassword string) error {
	updatePassword := `update "UsersTab" set "password"=$1 where "account"=$2`
	_, result := db.Exec(updatePassword, newPassword, account)
	return result
}

func DeleteAccount(db *sql.DB, account string) error {
	deleteAccount := `delete from "UsersTab" where account=$1`
	_, result := db.Exec(deleteAccount, account)
	return result
}

func GetAllAccounts(db *sql.DB) ([]struct_data.Person, error) {
	rows, result := db.Query(`SELECT "account", "password", "type" FROM "UsersTab"`)
	err.CheckError(result)

	defer rows.Close()

	persons := []struct_data.Person{}
	for rows.Next() {
		var person struct_data.Person
		result = rows.Scan(&person.Account, &person.Password, &person.TypePerson)
		err.CheckError(result)

		persons = append(persons, person)
	}

	err.CheckError(result)
	return persons, result
}

func GetDataByAccount(db *sql.DB, account string) (struct_data.Person, error) {
	getRowData := `SELECT "account", "password", "type" FROM "UsersTab" where account=$1`
	row := db.QueryRow(getRowData, account)
	var person struct_data.Person
	result := row.Scan(&person.Account, &person.Password, &person.TypePerson)

	return person, result
}
