package goods_table

import (
	"database/sql"

	"shop.com/err"
	struct_data "shop.com/struct_data"
)

func InsertNewGoods(db *sql.DB, goods struct_data.Goods) error {
	insertGoods := `insert into "GoodsTab"("name", "description", "count_good") values($1, $2, $3)`
	var result error
	_, result = db.Exec(insertGoods, goods.Name, goods.Description,
		goods.Count)
	return result
}

func UpdateCountGoodsByIdAndName(db *sql.DB, id uint32, name string, count uint32) error {
	updateCountGoods := `update "GoodsTab" set "count_good"=$1 where "id_good"=$2 and "name"=$3`
	_, result := db.Exec(updateCountGoods, count, id, name)
	return result
}

func AddCountGoodsByIdAndName(db *sql.DB, id uint32, name string, adding_count uint32) error {
	updateCountGoods := `update "GoodsTab" set "count_good"="count_good"+$1 where "id_good"=$2 and "name"=$3`
	_, result := db.Exec(updateCountGoods, adding_count, id, name)
	return result
}

func DeleteCountGoodsByIdAndName(db *sql.DB, id uint32, name string, deleting_count uint32) error {
	updateCountGoods := `update "GoodsTab" set "count_good"="count_good"-$1 where "id_good"=$2 and "name"=$3`
	_, result := db.Exec(updateCountGoods, deleting_count, id, name)
	return result
}

func DeleteGoodsByIdAndName(db *sql.DB, id uint32, name string) error {
	deleteGoods := `delete from "GoodsTab" where "id_good"=$1 and "name"=$2`
	_, result := db.Exec(deleteGoods, id, name)
	return result
}

func DeleteGoodsIfCountGoodsEqualZero(db *sql.DB) error {
	deleteGoods := `delete from "GoodsTab" where "count_good"=0`
	_, result := db.Exec(deleteGoods)
	return result
}

func GetDataGoodsByIdAndName(db *sql.DB, id uint32, name string) (struct_data.Goods, error) {
	getRowData := `SELECT "count_good", "description" FROM "GoodsTab" where "id_good"=$1 and "name"=$2`
	row := db.QueryRow(getRowData, id, name)
	var good struct_data.Goods
	result := row.Scan(&good.Count, &good.Description)
	good.Id = id
	good.Name = name
	return good, result
}

func GetNameGoodsById(db *sql.DB, id uint32) (string, error) {
	getRowData := `SELECT "name" FROM "GoodsTab" where "id_good"=$1`
	row := db.QueryRow(getRowData, id)
	var name string
	result := row.Scan(&name)
	return name, result
}

func GetCountGoodsById(db *sql.DB, id uint32) (uint32, error) {
	getRowData := `SELECT "count_good" FROM "GoodsTab" where "id_good"=$1`
	row := db.QueryRow(getRowData, id)
	var countGoods uint32
	result := row.Scan(&countGoods)
	return countGoods, result
}

func GetAllGoods(db *sql.DB) ([]struct_data.Goods, error) {
	rows, result := db.Query(`SELECT "id_good", "name", "description", "count_good" FROM "GoodsTab"`)
	err.CheckError(result)

	defer rows.Close()

	listGoods := []struct_data.Goods{}
	for rows.Next() {
		var goods struct_data.Goods
		result = rows.Scan(&goods.Id, &goods.Name, &goods.Description, &goods.Count)
		err.CheckError(result)

		listGoods = append(listGoods, goods)
	}

	err.CheckError(result)
	return listGoods, result
}
