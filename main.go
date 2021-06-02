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
	http.HandleFunc("/login",func(rw http.ResponseWriter, r*http.Request){
		fmt.Println("Login Section...!")
	})
	http.HandleFunc("/contact",func(rw http.ResponseWriter, r *http.Request){
		myInformation := make(map[string]string)
		myInformation["name"]= "Md Ruhul Amin"
		myInformation["college"]="NotreDame College"
		myInformation["University"]="American International University-Bangladesh"
		for key,val := range  myInformation{
			fmt.Println(key,"  :=>   ", val)
		}
	})
	http.ListenAndServe(":8000",nil)


}