package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"

	"github.com/google/uuid"
)

func GetStackService() ( []model.StackModel,error) {
	var result []model.StackModel

	err := config.SupaClient.DB.From("stack").Select("*").Execute(&result)
	if err != nil {
		return nil, err
	}

	return result ,nil
}





func CreateStackService(stack model.StackModel)  error {
	var result []model.StackModel

	err := config.SupaClient.DB.From("stack").Insert(stack).Execute(&result)
	if err != nil {
		return err
	}

	return nil
}

func UpdateStackService(id uuid.UUID,stack map[string]interface{}) error {
	finalId := id.String()
	var result map[string]interface{}
	err := config.SupaClient.DB.From("stack").Update(stack).Eq("id", finalId).Execute(&result) 
	if err != nil {
		return nil
	}

	return nil
}



func DeleteStackService(id uuid.UUID) error {
	finalId := id.String()
	err := config.SupaClient.DB.From("stack").Delete().Eq("id",finalId).Execute("")
	if err != nil {
		return err
	}

	return nil
}