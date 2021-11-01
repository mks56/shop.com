package main

import (
	"fmt"
	"strings"

	"shop.com/admin_accounts"
	"shop.com/logic_db"
	"shop.com/manager_api"
	"shop.com/par_calc"
	"shop.com/user_api"
)

func main() {
	db := logic_db.InitDb()
	defer db.Close()

	//admin_accounts.RegisterNewAccountApi(db, "NewManager", "98765", "manager")
	//admin_accounts.RegisterNewAccountApi(db, "NewUser", "qwerty", "user")
	//admin_accounts.UpdatePasswordByAccountApi(db, "NewUser", "NEWPASS")
	//admin_accounts.DeleteAccountApi(db, "NewUser12")

	typeManager, result := admin_accounts.AuthorizationByAccountAndPasswordApi(db, "NewManager", "98765")
	fmt.Println(typeManager, "", result)

	//persons := admin_accounts.GetAllAccountsApi(db)
	//fmt.Println(persons)

	if strings.Compare(typeManager, "manager") == 0 && result {
		//manager_api.AddNewGoodsApi(db, "NewGoods1", "Description NewGoods1", 200)
		//manager_api.UpdateCountGoodsByIdAndNameApi(db, 10, "NewGoods1", 273)
		//manager_api.AddCountGoodsByIdAndNameApi(db, 10, "NewGoods1", 42)
		//manager_api.DeleteCountGoodsByIdAndNameApi(db, 10, "NewGoods1", 277)

		//goods := manager_api.GetGoodsByIdAndNameApi(db, 9, "NewGoods1")
		//fmt.Println(goods)

		//manager_api.DeleteGoodsByIdAndNameApi(db, 9, "NewGoods1")

		listGoods := manager_api.GetAllGoodsApi(db)
		fmt.Println(listGoods)

		//manager_api.DeleteGoodsIfCountGoodsEqualZeroApi(db)

		//shoppingCartForManyUser := manager_api.GetAllShoppingCartsApi(db)
		//fmt.Println(shoppingCartForManyUser)
	}

	typeUser, result := admin_accounts.AuthorizationByAccountAndPasswordApi(db, "NewUser", "NEWPASS")
	fmt.Println(typeUser, "", result)

	if strings.Compare(typeUser, "user") == 0 && result {
		//user_api.CreateNewShoppingCartApi(db, "Alice")
		//user_api.AddNewGoodsIntoShoppingCartApi(db, 13, 6, 55)
		//user_api.UpdateCountGoodsInShoppingCartApi(db, 9, 1, 100)
		//user_api.DeleteGoodsByIdShoppingCartAndIdGoodsApi(db, 13, 6)
		shoppingCartForOneUser := user_api.GetAllGoodsFromShoppingCartByIdApi(db, 9)
		fmt.Println(shoppingCartForOneUser)

		par_calc.RunMultipleDbAccessesApi(db, 100)
	}

}
