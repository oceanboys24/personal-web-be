package main

import (
	"personal-web-be/config"

)


func main () {
	//Load Env
	config.LoadEnv()
	//Connect Supabase 
	config.ConnectSupabase()	
	

}	