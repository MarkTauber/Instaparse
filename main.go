package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "net"
	"net/http"
	"net/url"
	"strings"
)

type Subscriber struct {
	Pk          int    `json:"pk"`
	Uname       string `json:"username"`
	Full_name   string `json:"full_name"`
	Is_private  bool   `json:"is_private"`
	Is_verified bool   `json:"is_verified"`
	AnonPic     bool   `json:"has_anonymous_profile_picture"`
}

type subResponse struct {
	Users  []Subscriber `json:"users"`
	NextId string       `json:"next_max_id"`
}

var idlog []string
var out string
var reed_headers, _ = ioutil.ReadFile("headers.txt")
var recon1 = strings.Replace(string(reed_headers), "gzip, deflate, br", "plain", -1)

const id = 23325871562 //!!!ID аккаунта!!!
const count = 1000000  //!!!запросы не менять!!!

func main() {
	//recon1 := ``
	cli := new(http.Client)
	req := new(http.Request)
	req.Method = "GET"
	req.Header = make(http.Header)
	req.Host = "i.instagram.com"
	req.URL, _ = url.Parse("https://i.instagram.com/api/v1/friendships/" + fmt.Sprint(id) + "/followers/?count=" + fmt.Sprint(count) + "&search_surface=follow_list_page")
	//req.
	//Set Cookies
	//Set Headers
	//req.
	for _, splits_range := range strings.Split(recon1, "\n") {
		part := strings.Split(splits_range, ": ")
		//req.Header.
		//req.
		req.Header.Set(part[0], part[1])
	}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	//var bdy []byte
	var responseDecoded = new(subResponse)
	suckme, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(suckme) + "\n")
	json.Unmarshal(suckme, responseDecoded)
	//temp := string(suckme)
	//fmt.Println(temp)

	//fmt.Println(len(responseDecoded.Users))
	for _, i := range responseDecoded.Users {
		//var filter bool = !i.AnonPic && !i.Is_private
		//var filter bool = !i.Is_private
		var filter bool = true
		if filter {
			fmt.Println(i.Uname, ";", i.Pk)
			out += i.Uname + ":" + fmt.Sprint(i.Pk) + "\n"
		}
	}

	fmt.Println("\n\n" + responseDecoded.NextId)
	for responseDecoded.NextId != "" {
		responseDecoded = continueAss(*responseDecoded)
		idlog = append(idlog, responseDecoded.NextId)
		//fmt.Println("Таймаут")
		//time.Sleep(100 * time.Millisecond)
	}
	ioutil.WriteFile(fmt.Sprint(id)+"_NEXTID.txt", []byte(strings.Join(idlog, "\n")), 0644)
	ioutil.WriteFile(fmt.Sprint(id)+".txt", []byte(out), 0644)
}

func continueAss(resppenis subResponse) *subResponse {
	cli := new(http.Client)
	req := new(http.Request)
	req.Method = "GET"
	req.Header = make(http.Header)
	req.Host = "i.instagram.com"
	req.URL, _ = url.Parse("https://i.instagram.com/api/v1/friendships/" + fmt.Sprint(id) + "/followers/?count=" + fmt.Sprint(count) + "&max_id=" + resppenis.NextId + "&search_surface=follow_list_page")

	for _, splits_range := range strings.Split(recon1, "\n") {
		part := strings.Split(splits_range, ": ")
		req.Header.Set(part[0], part[1])
	}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}

	//var bdy []byte
	var responseDecoded = new(subResponse)
	suckme, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(suckme) + "\n")
	json.Unmarshal(suckme, responseDecoded)

	//const anonclosed = !i.AnonPic
	//fmt.Println(len(responseDecoded.Users))
	for _, i := range responseDecoded.Users {
		//var filter bool = !i.AnonPic && !i.Is_private
		//var filter bool = !i.Is_private
		var filter bool = true
		if filter {
			fmt.Println(i.Uname)
			out += i.Uname + "\n"
		}
	}
	//fmt.Println("\n\n" + responseDecoded.NextId)
	return responseDecoded
}
