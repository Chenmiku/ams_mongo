package transaction

import ()

func GetByID(id string) (*Transaction, error) {
	var w Transaction
	return &w, TableTransaction.ReadByID(id, &w)
}

func GetByHash(hash string) (*Transaction, error) {
	var w Transaction
	return &w, TableTransaction.ReadOne(map[string]interface{}{
		"hash": hash,
		"dtime": 0,
	}, &w)
}

// func GetAllByWalletID(pageSize int, pageNumber int, sortBy string, sortOrder string, walletid string, pubaddress *[]PublicAddress) (int, error) {
// 	var where map[string]interface{}
// 	if walletid == "" {
// 		where =  map[string]interface{}{
// 			"dtime": 0,
// 		}
// 	} else {
// 		where = map[string]interface{}{
// 			"dtime": 0,
// 			"wallet_id": walletid,
// 		}
// 	}

// 	exclude := []string{}
// 	return TablePublicAddress.ReadPagingSortWithExclude(where, pageSize, pageNumber, sortBy, sortOrder, pubaddress, exclude)
// }
