package service

import "minercms/daosnews"

/*
func GetPageDataNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.QueryOneNews(canShu)
	return ret
}
*/

//

func GetPageDataNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.InsertNewsT(canShu)
	return ret
}
