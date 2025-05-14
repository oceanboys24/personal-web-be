package main

import (
	"personal-web-be/config"
	"personal-web-be/middleware"

	"personal-web-be/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main () {
	//Connect Redis
	config.ConnectRedis()

	//Connect Supabase 
	config.ConnectSupabase()	
	//Connnect Supabase Storage
	config.ConnectStorage()
	//init Fiber 
	app := fiber.New()

	//init CORS
	app.Use(cors.New())

	//Endpoint V1 
	v1Route := app.Group("/v1")

	v1Route.Use(middleware.ActiviyLogger())

	//Login Endpoint
	routes.LoginRoute(v1Route)	
	//Upload Endpoint
	routes.UploadRoute(v1Route)
	//Hero Endpoint 
	routes.HeroRoute(v1Route)
	//Stack Endpoint
	routes.StackRoute(v1Route)
	//Work Experience Endpoint
	routes.WorkRoute(v1Route)
	//Project Endpoint
	routes.ProjectRoute(v1Route)
	//Visitor Endpoint
	routes.VisitorRoute(v1Route)
	//Activity Endpoint
	routes.ActivityRoute(v1Route)

	//Running Server
	app.Listen("0.0.0.0:8080")
}	