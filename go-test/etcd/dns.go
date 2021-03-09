/*
@Time : 21-1-21
@Author : jzd
@Project: go-learning
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dnsRule1 := &DnsRule{Domain: "domain1.com", IpAddress: "localhost;127.0.0.1", IpType: "IPV4", Ttl: 12345}
	dnsRules := []*DnsRule{dnsRule1}
	maps := formatToEtctd(dnsRules)
	fmt.Println(maps)
}

func formatToEtctd(dnsrules []*DnsRule) map[string]string {
	etct := make(map[string]string)
	for _, v := range dnsrules {
		domain1 := strings.Split(v.Domain, ".")
		domain2 := ""
		for _, v := range domain1 {
			domain2 = "/" + v + domain2
		}
		ipList := strings.Split(v.IpAddress, ";")
		for i, ip := range ipList {
			domainkey := domain2 + "/" + strconv.Itoa(i)
			domainVal := "{\"host\":\"" + ip + "\",\"ttl\":" +
				strconv.FormatInt(v.Ttl, 10) + "}"
			etct[domainkey] = domainVal
		}
	}
	return etct
}

type DnsRule struct {
	Domain    string
	IpAddress string
	IpType    string
	Ttl       int64
}
