package api_github

import "sync"

var (
	once sync.Once
	api  *Api
)

type Api struct {
	token string
}

func New(token string) *Api {
	once.Do(func() {
		api = &Api{
			token: token,
		}
	})

	return api
}
