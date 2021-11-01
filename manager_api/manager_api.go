package manager_api

import (
	"database/sql"
	"fmt"
	"math"

	"shop.com/err"
	goods_table "shop.com/logic_db/goods_table"
	"shop.com/logic_db/shopping_cart_tables"
	"shop.com/msg"
	struct_data "shop.com/struct_data"
)

func AddNewGoodsApi(db *sql.DB, name, description string, count uint32) {
	newGoods := struct_data.Goods{Name: name, Description: description,
		Count: count}
	result := goods_table.InsertNewGoods(db, newGoods)
	if result != nil {
		err.PrintError(result)
	}
}

func UpdateCountGoodsByIdAndNameApi(db *sql.DB, id uint32, name string, count uint32) {
	result := goods_table.UpdateCountGoodsByIdAndName(db, id, name, count)
	if result != nil {
		err.PrintError(result)
	}
}

func AddCountGoodsByIdAndNameApi(db *sql.DB, id uint32, name string, adding_count uint32) {
	result := goods_table.AddCountGoodsByIdAndName(db, id, name, adding_count)
	if result != nil {
		err.PrintError(result)
	}
}

func DeleteCountGoodsByIdAndNameApi(db *sql.DB, id uint32, name string, deleting_count uint32) uint32 {
	var result error
	var goods struct_data.Goods
	goods, result = goods_table.GetDataGoodsByIdAndName(db, id, name)
	if result != nil {
		err.PrintError(result)
	}

	tempCount := int32(goods.Count - deleting_count)
	if (tempCount) < 0.0 {
		cannotCountDeleting := uint32(math.Abs(float64(tempCount)))
		msgString := fmt.Sprintf("You want to pick up the goods in the amount of %d, but there are only %d such goods in stock. It will be delete from the stock of goods in the amount of %d. And %d goods are out of stock",
			deleting_count, goods.Count, goods.Count, cannotCountDeleting)
		msg.PrintMsgBy(msgString)
	} else {
		goods.Count = deleting_count
	}

	result = goods_table.DeleteCountGoodsByIdAndName(db, id, name, goods.Count)
	if result != nil {
		err.PrintError(result)
	}
	return goods.Count
}

func GetGoodsByIdAndNameApi(db *sql.DB, id uint32, name string) struct_data.Goods {
	goods, result := goods_table.GetDataGoodsByIdAndName(db, id, name)
	if result != nil {
		err.PrintError(result)
	}
	return goods
}

func DeleteGoodsByIdAndNameApi(db *sql.DB, id uint32, name string) {
	result := goods_table.DeleteGoodsByIdAndName(db, id, name)
	if result != nil {
		err.PrintError(result)
	}
}

func DeleteGoodsIfCountGoodsEqualZeroApi(db *sql.DB) {
	result := goods_table.DeleteGoodsIfCountGoodsEqualZero(db)
	if result != nil {
		err.PrintError(result)
	}
}

func GetAllGoodsApi(db *sql.DB) []struct_data.Goods {
	listGoods, result := goods_table.GetAllGoods(db)
	if result != nil {
		err.PrintError(result)
	}
	return listGoods
}

func GetAllShoppingCartsApi(db *sql.DB) []struct_data.ShoppingCartForOneUser {
	listShoppingCarts, result := shopping_cart_tables.GetAllShoppingCarts(db)
	if result != nil {
		err.PrintError(result)
	}

	shoppingCartForManyUser := struct_data.FilterShoppingCartsForManyUser(listShoppingCarts)
	return shoppingCartForManyUser
}
