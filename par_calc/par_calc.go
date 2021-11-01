package par_calc

import (
	"database/sql"
	"fmt"
	"sync"

	"shop.com/user_api"
)

func RunMultipleDbAccessesApi(db *sql.DB, number_of_requests int) {
	mutex := sync.Mutex{}
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(number_of_requests)

	for i := 0; i < number_of_requests; i++ {

		go func(i int) {
			mutex.Lock()
			shoppingCartForOneUser := user_api.GetAllGoodsFromShoppingCartByIdApi(db, 9)
			fmt.Println(shoppingCartForOneUser)
			mutex.Unlock()
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
}
