/* This is is a consumer for the calculator web service      *
** Part of the assignment for Essential Skills module Lab 10.*
** University of Amsterdam                                   *
** MSc In System and Network Engineering 					 *
** Nick Triantafyllidis										 *
** This will be a RESTful service							 *
 */

package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main(){
	
	operation := ""
	host := "host"
	first := 0
	second := 0
	
	fmt.Printf("Tell me the hostname or ip of the web service: ")
	fmt.Scanf("%s", &host)
	fmt.Printf("What would you like to do (sum,difference,product,quotient): ")
	fmt.Scanf("%s", &operation)
	fmt.Printf("Give first number: ")
	fmt.Scanf("%d", &first)
	fmt.Printf("Give second number: ")
	fmt.Scanf("%d", &second)
	
	url := fmt.Sprintf("http://%s/api/v1/%s?a=%d&b=%d", host , operation, first, second)
	
	res, _  := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(body))
	
	
}
