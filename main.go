package main

import "net/http"

url = "http://localhost:8080/v1/tasks/ga_month_report"

func main() {
	req, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
	}
}
