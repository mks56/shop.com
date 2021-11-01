package shopping_cart_tables

import (
	"database/sql"

	"shop.com/err"
	struct_data "shop.com/struct_data"
)

func CreateNewShoppingCart(db *sql.DB, account string) error {
	createNewShoppingCart := `insert into "ShoppingCartUserTab"("account") SELECT "account" FROM "UsersTab" WHERE "type"='user' AND "account"=$1`
	var result error
	_, result = db.Exec(createNewShoppingCart, account)
	return result
}

func AddNewGoodsIntoShoppingCart(db *sql.DB, id_shopping_cart, id_goods, count_goods uint32) error {
	insertGood := `insert into "ShoppingCartTab"("id_shop_cart", "id_good", "count_shop_goods") values($1, $2, $3)`
	var result error
	_, result = db.Exec(insertGood, id_shopping_cart, id_goods, count_goods)
	return result
}

func UpdateCountGoodsInShoppingCartBy(db *sql.DB, id_shopping_cart, id_goods, count_goods uint32) error {
	updateCountGoods := `update "ShoppingCartTab" set "count_shop_goods"=$1 where "id_shop_cart"=$2 and "id_good"=$3`
	_, result := db.Exec(updateCountGoods, count_goods, id_shopping_cart, id_goods)
	return result
}

func GetCountGoodsById(db *sql.DB, id uint32) (uint32, error) {
	getRowData := `SELECT "count_shop_goods" FROM "ShoppingCartTab" where "id_shop_cart"=$1`
	row := db.QueryRow(getRowData, id)
	var countGoods uint32
	result := row.Scan(&countGoods)
	return countGoods, result
}

func DeleteGoodsByIdShoppingCartAndIdGoods(db *sql.DB, id_shopping_cart, id_goods uint32) error {
	deleteGoods := `delete from "ShoppingCartTab" where "id_shop_cart"=$1 and "id_good"=$2`
	_, result := db.Exec(deleteGoods, id_shopping_cart, id_goods)
	return result
}

func GetAllGoodsFromShoppingCartById(db *sql.DB, id_shopping_cart uint32) ([]struct_data.ShoppingCart, error) {
	getGoods := `SELECT S.id_shop_cart, U.account, G.id_good, G.name, G.description, S.count_shop_goods FROM "ShoppingCartTab" AS S, "GoodsTab" AS G, "ShoppingCartUserTab" AS U WHERE S.id_shop_cart=U.id_shop_cart_user AND S.id_shop_cart=$1 AND S.id_good=G.id_good`
	rows, result := db.Query(getGoods, id_shopping_cart)
	err.CheckError(result)

	defer rows.Close()

	listShoppingCart := []struct_data.ShoppingCart{}
	for rows.Next() {
		var shoppingCart struct_data.ShoppingCart
		result = rows.Scan(&shoppingCart.Id, &shoppingCart.NameAccount, &shoppingCart.Goods.Id, &shoppingCart.Goods.Name, &shoppingCart.Goods.Description, &shoppingCart.Goods.Count)
		err.CheckError(result)

		listShoppingCart = append(listShoppingCart, shoppingCart)
	}

	return listShoppingCart, result
}

func GetAllShoppingCarts(db *sql.DB) ([]struct_data.ShoppingCart, error) {
	getAllShoppingCarts := `SELECT S.id_shop_cart, U.account, G.id_good, G.name, G.description, S.count_shop_goods FROM "ShoppingCartTab" AS S, "GoodsTab" AS G, "ShoppingCartUserTab" AS U WHERE S.id_shop_cart=U.id_shop_cart_user AND S.id_good=G.id_good ORDER BY S.id_shop_cart ASC`
	rows, result := db.Query(getAllShoppingCarts)
	err.CheckError(result)

	defer rows.Close()

	listShoppingCart := []struct_data.ShoppingCart{}
	for rows.Next() {
		var shoppingCart struct_data.ShoppingCart
		result = rows.Scan(&shoppingCart.Id, &shoppingCart.NameAccount, &shoppingCart.Goods.Id, &shoppingCart.Goods.Name, &shoppingCart.Goods.Description, &shoppingCart.Goods.Count)
		err.CheckError(result)

		listShoppingCart = append(listShoppingCart, shoppingCart)
	}

	return listShoppingCart, result
}
