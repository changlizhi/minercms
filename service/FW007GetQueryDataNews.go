package service

import "minercms/daosnews"


func GetPageDataQueryNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.PageDataQueryNewsList(canShu)
	return ret
}
