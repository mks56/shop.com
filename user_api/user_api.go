package user_api

import (
	"database/sql"
	"math"

	"shop.com/err"
	goods_tables "shop.com/logic_db/goods_table"
	shopping_cart_tables "shop.com/logic_db/shopping_cart_tables"
	"shop.com/manager_api"
	struct_data "shop.com/struct_data"
)

func CreateNewShoppingCartApi(db *sql.DB, account string) {
	result := shopping_cart_tables.CreateNewShoppingCart(db, account)
	if result != nil {
		err.PrintError(result)
	}
}

func AddNewGoodsIntoShoppingCartApi(db *sql.DB, id_shopping_cart, id_goods, count_goods uint32) {
	var result error
	var name string
	name, result = goods_tables.GetNameGoodsById(db, id_goods)
	if result != nil {
		err.PrintError(result)
	}

	canPutCountGoodsIntoShoppingCart := manager_api.DeleteCountGoodsByIdAndNameApi(db, id_goods, name, count_goods)

	result = shopping_cart_tables.AddNewGoodsIntoShoppingCart(db, id_shopping_cart, id_goods, canPutCountGoodsIntoShoppingCart)
	if result != nil {
		err.PrintError(result)
	}
}

func UpdateCountGoodsInShoppingCartApi(db *sql.DB, id_shopping_cart, id_goods, count_goods uint32) {
	var result error

	var countGoodsInStock uint32
	countGoodsInStock, result = goods_tables.GetCountGoodsById(db, id_goods)
	if result != nil {
		err.PrintError(result)
	}

	var countGoodsInShoppingCart uint32
	countGoodsInShoppingCart, result = shopping_cart_tables.GetCountGoodsById(db, id_shopping_cart)
	if result != nil {
		err.PrintError(result)
	}

	tempCount := int32(int32(count_goods) - int32(countGoodsInShoppingCart))

	var countForUpdate uint32
	countForUpdate = count_goods

	var name string
	name, result = goods_tables.GetNameGoodsById(db, id_goods)
	if result != nil {
		err.PrintError(result)
	}

	if tempCount > 0 {
		countForUpdate = manager_api.DeleteCountGoodsByIdAndNameApi(db, id_goods, name, uint32(tempCount))
		countForUpdate = countForUpdate + countGoodsInShoppingCart
		if countGoodsInStock > uint32(math.Abs(float64(tempCount))) {
			countForUpdate = count_goods
		}

	} else if tempCount < 0 {
		manager_api.AddCountGoodsByIdAndNameApi(db, id_goods, name, uint32(math.Abs(float64(tempCount))))
	}

	result = shopping_cart_tables.UpdateCountGoodsInShoppingCartBy(db, id_shopping_cart, id_goods, countForUpdate)
	if result != nil {
		err.PrintError(result)
	}

}

func DeleteGoodsByIdShoppingCartAndIdGoodsApi(db *sql.DB, id_shopping_cart, id_goods uint32) {
	var countGoodsInShoppingCart uint32
	var result error
	countGoodsInShoppingCart, result = shopping_cart_tables.GetCountGoodsById(db, id_shopping_cart)
	if result != nil {
		err.PrintError(result)
	}

	result = shopping_cart_tables.DeleteGoodsByIdShoppingCartAndIdGoods(db, id_shopping_cart, id_goods)
	if result != nil {
		err.PrintError(result)
	}

	var name string
	name, result = goods_tables.GetNameGoodsById(db, id_goods)
	if result != nil {
		err.PrintError(result)
	}

	manager_api.AddCountGoodsByIdAndNameApi(db, id_goods, name, countGoodsInShoppingCart)

}

func GetAllGoodsFromShoppingCartByIdApi(db *sql.DB, id_shopping_cart uint32) struct_data.ShoppingCartForOneUser {
	listShoppingCart, result := shopping_cart_tables.GetAllGoodsFromShoppingCartById(db, id_shopping_cart)
	if result != nil {
		err.PrintError(result)
	}
	return struct_data.FilterShoppingCartForOneUser(listShoppingCart)
}
