package admin_accounts

import (
	"database/sql"
	"strings"

	"shop.com/err"
	"shop.com/logic_db/users_table"
	"shop.com/msg"
	struct_data "shop.com/struct_data"
)

func RegisterNewAccountApi(db *sql.DB, account, password, typePerson string) {
	newPerson := struct_data.Person{Account: account, Password: password, TypePerson: typePerson}
	result := users_table.InsertNewUser(db, newPerson)
	if result != nil {
		msg.MsgAboutIncorrectAccount()
		err.PrintError(result)
	}
}

func UpdatePasswordByAccountApi(db *sql.DB, account, password string) {
	result := users_table.UpdatePasswordByAccount(db, account, password)
	if result != nil {
		err.PrintError(result)
	}
}

func DeleteAccountApi(db *sql.DB, account string) {
	result := users_table.DeleteAccount(db, account)
	if result != nil {
		err.PrintError(result)
	}
}

func AuthorizationByAccountAndPasswordApi(db *sql.DB, account, password string) (string, bool) {
	localResult := 0
	person, result := users_table.GetDataByAccount(db, account)
	if result != nil {
		localResult = -1
		msg.MsgAboutDoesNotExistAccount()
		err.PrintError(result)
	} else {
		if strings.Compare(strings.TrimRight(account, "\n"), strings.TrimRight(person.Account, "\n")) != 0 {
			localResult = -1
			msg.MsgAboutIncorrectGotAccount()
		}

		if strings.Compare(password, person.Password) != 0 {
			localResult = -1
			msg.MsgAboutIncorrectPassword()
		}

		if localResult == 0 {
			msg.MsgAboutSuccessfulAccountAuthorization()
			return person.TypePerson, true
		}
	}
	return "", false
}

func GetAllAccountsApi(db *sql.DB) []struct_data.Person {
	persons, result := users_table.GetAllAccounts(db)
	if result != nil {
		err.PrintError(result)
	}
	return persons
}
