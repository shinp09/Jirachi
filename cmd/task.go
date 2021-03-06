package cmd

import (
	"log"
	"net/http"
)

func GetTask(user User) {
	req, err := http.NewRequest(http.MethodGet, user.requestUrl, nil)
	if err != nil {
		log.Println(err)
	}

	// 認証情報を付与し、リクエストと一緒に送る
	req.SetBasicAuth(user.name, user.password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	ConsistencyTasks(resp, err)
}
