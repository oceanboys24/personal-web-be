package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"
)

func CreateLoggerActivity(activity model.Activity) error {
	err := config.SupaClient.
		DB.
		From("recent_activity").
		Insert(activity).
		Execute(nil) 

	if err != nil {
		return err
	}

	return nil
}

func GetLoggerActivity() ([]model.ActivityResponse,error) {
	var result []model.ActivityResponse

	err := config.SupaClient.DB.From("recent_activity").Select("*").Execute(&result)
	if err != nil {
		return nil, err
	}


	return result ,nil
}
