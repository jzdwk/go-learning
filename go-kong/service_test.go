/*
@Time : 2021/3/9
@Author : jzd
@Project: go-learning
*/
package go_kong

import "testing"

func TestRemoveAllServices(t *testing.T) {
	client := NewServiceClient(kongUrl)
	if err := client.RemoveAll(); err != nil {
		t.Error(err)
	}
	t.Log("success")
}
