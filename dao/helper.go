package dao

import "strconv"

func userID2CasbinKey(id uint) string {
	return "user_" + strconv.Itoa(int(id))
}

func roleID2CasbinKey(id uint) string {
	return "role_" + strconv.Itoa(int(id))
}
