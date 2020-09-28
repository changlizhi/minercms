package service

import (
	"minercms/daosnews"
)

func QueryNewsService(canShu map[string]interface{}) map[string]interface{} {
	ret := daosnews.QueryNewsList(canShu)
	return ret
}
