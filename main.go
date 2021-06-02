package main
import (
	"net/http"
	"fmt"
	"log"

)



func main(){

	http.HandleFunc("/",func(rw  http.ResponseWriter, r *http.Request ){
		fmt.Println("hi there")
	})
	http.HandleFunc("/register",func(rw http.ResponseWriter, r *http.Request){
		log.Println("register Section....!")
		registeredList :=[]string{"sakib","Sajid","Ashraf"}
		for key,val := range registeredList{
			log.Println(key, " ",val)
		}
	})

	http.ListenAndServe(":8000",nil)


}