package config

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)
var SupaClient *supa.Client


func ConnectSupabase() {
	key := os.Getenv("KEY")
	url := os.Getenv("URL")
	
	client := supa.CreateClient(url, key)

	SupaClient = client

}