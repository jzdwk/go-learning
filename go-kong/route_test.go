/*
@Time : 2021/3/9
@Author : jzd
@Project: go-learning
*/
package go_kong

import "testing"

const kongUrl = "http://127.0.0.1:8001"

func TestRemoveAllRoutes(t *testing.T) {
	client := NewRouteClient(kongUrl)
	if err := client.RemoveAll(); err != nil {
		t.Error(err)
	}
	t.Log("success")
}
