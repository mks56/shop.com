package struct_data

type ShoppingCart struct {
	Id          uint32
	NameAccount string
	Goods       Goods
}

type ShoppingCartForOneUser struct {
	Id          uint32
	NameAccount string
	Goods       []Goods
}

func FilterShoppingCartForOneUser(shopping_cart []ShoppingCart) ShoppingCartForOneUser {
	var tempShoppingCartForOneUser ShoppingCartForOneUser
	tempShoppingCartForOneUser.Id = shopping_cart[0].Id
	tempShoppingCartForOneUser.NameAccount = shopping_cart[0].NameAccount
	var tempGoods []Goods
	for _, element := range shopping_cart {
		tempGoods = append(tempGoods, element.Goods)
	}
	tempShoppingCartForOneUser.Goods = tempGoods
	return tempShoppingCartForOneUser
}

func FilterShoppingCartsForManyUser(shopping_carts []ShoppingCart) []ShoppingCartForOneUser {
	var tempShoppingCartsForManyUser []ShoppingCartForOneUser

	for _, element := range FilterShoppingCartsById(shopping_carts) {
		tempShoppingCartsForManyUser = append(tempShoppingCartsForManyUser, FilterShoppingCartForOneUser(element))
	}

	return tempShoppingCartsForManyUser
}

func FilterShoppingCartsById(shopping_carts []ShoppingCart) [][]ShoppingCart {
	var tempShoppingCartForOneUser []ShoppingCart
	var tempShoppingCartsForManyUser [][]ShoppingCart
	for i := 0; i < len(shopping_carts); i++ {

		if i == len(shopping_carts)-1 {
			tempShoppingCartForOneUser = append(tempShoppingCartForOneUser, shopping_carts[i])
			tempShoppingCartsForManyUser = append(tempShoppingCartsForManyUser, tempShoppingCartForOneUser)
		} else {
			if shopping_carts[i].Id == shopping_carts[i+1].Id {
				tempShoppingCartForOneUser = append(tempShoppingCartForOneUser, shopping_carts[i])
			} else {
				tempShoppingCartForOneUser = append(tempShoppingCartForOneUser, shopping_carts[i])
				tempShoppingCartsForManyUser = append(tempShoppingCartsForManyUser, tempShoppingCartForOneUser)
				tempShoppingCartForOneUser = nil
			}
		}
	}
	return tempShoppingCartsForManyUser
}
