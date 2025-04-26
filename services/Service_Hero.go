package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"
)


func GetServiceHero() (*model.HeroModel,error) {
	var result model.HeroModel
	err := config.SupaClient.DB.From("hero").Select("*").Single().Execute(&result)
	if err != nil {
		return nil , err
	}
	return &result ,nil
}


func UpdateServiceHero(updatedFiels map[string]interface{}) (model.HeroModel, error) {
	var result []model.HeroModel
	err := config.SupaClient.DB.From("hero").Update(updatedFiels).Eq("id","dc81c56d-420f-4eea-aff3-d0503f2732a0").Execute(&result)

	if err != nil {
		return result[0] , err
	}

	return  result[0] , nil
}