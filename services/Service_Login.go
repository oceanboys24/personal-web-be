package services

import (
	"context"

	"personal-web-be/config"
	model "personal-web-be/models"

	"github.com/nedpals/supabase-go"
)



func LoginServices(login model.LoginModel) (string, error) {
	ctx := context.Background()

	user, err := config.SupaClient.Auth.SignIn(ctx, supabase.UserCredentials{
		Email: login.Email,
		Password: login.Password,
	})

	if err != nil {
		return "" , err
	}


	return user.AccessToken, nil
}