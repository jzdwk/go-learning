/*
@Time : 20-10-12
@Author : jzd
@Project: go-learning
*/
package main

func main() {
	/*path := "/tet/{info}/{id}/"
	elm := strings.Split(path, "/")
	for _, value := range elm {
		//fmt.Println(value)
		if value != "" {
			/*if value[0:]=="{"&&value[len(value)-1:]=="}"{
				fmt.Println(value)
			}
			if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
				args := value[1 : len(value)-1]
				if args == "id" {
					path = strings.ReplaceAll(path, value, "\\\\d+")
				}
				if args == "info" {
					path = strings.ReplaceAll(path, value, ".*")
				}
			}
		}
	}
	fmt.Println(path)*/

	/*	akdata := GetRandomString(14)
		ak := base64.StdEncoding.EncodeToString([]byte(akdata))
		fmt.Println(ak)
		skdata := GetRandomString(28)
		sk := base64.StdEncoding.EncodeToString([]byte(skdata))
		fmt.Println(sk)*/
	//rep := strings.NewReplacer("\\", "_", "/", "_", ".", "_")

	/*	err := os.Chdir("/home/jzd/go/src/k8smgr/")
		if err != nil{
			fmt.Println(err.Error())
		}*/
	/*commentFilename, _ := filepath.Rel("/home/jzd/go/src/k8smgr/", "/home/jzd/go/src/k8smgr/controllers/auth")
	fmt.Println(commentFilename)
	rst := rep.Replace(commentFilename) + ".go"
	fmt.Println(rst)

	path := filepath.Dir(os.Args[0])
	path,_ = filepath.Abs(path)
	path,_ = os.Getwd()
	fmt.Println(path)
	for idx, args := range os.Args {
		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	}*/
	//addr := fmt.Sprintf("%v://%v%v","http","127.0.0.1:800","/test1")
	//fmt.Println(addr)

}

/*func GetRandomString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}*/
