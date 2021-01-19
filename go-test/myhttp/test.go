/*
@Time : 21-1-6
@Author : jzd
@Project: go-learning
*/

/*func main() {
	//time out = 8s
	client := &http.Client{Timeout: time.Second * 8}

	//process body
	//var reader *bytes.Reader
	//body request param
	req, err := http.NewRequest("GET", "http://myharbor.com:8000/headers", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	repo, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(repo.Status)
}*/


