package usecase

import (
	"fmt"
	"gis-project-backend/pkg/module/quest5/model"
)

func (service *quest5UseCase) QuestF(year string) (*model.Quest5FResponse, error) {
	// query := fmt.Sprintf("SELECT [ID],[Country],[City],[Year],[Pm25],[Latitude],[Longtitude],[Population],[Wbinc16_text],[Region],[Conc_pm25],[Color_pm25] FROM dbo.AirPollutionPM25 WHERE Year=%s", year)
	query := fmt.Sprintf("SELECT [ID],[Country],[City],[Latitude],[Longtitude] FROM dbo.AirPollutionPM25 WHERE Year=%s AND Wbinc16_text='low income'", year)
	results, err := service.CoreRegistry.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var resultQuest model.Quest5FResponse
	for results.Next() {
		var tempData model.Quest5FData
		// err = results.Scan(&tempData.ID, &tempData.Country, &tempData.City, &tempData.Year, &tempData.Pm25, &tempData.Latitude, &tempData.Longtitude,
		// 	&tempData.Population, &tempData.Wbinc16, &tempData.Region, &tempData.ConcPm25, &tempData.ColorPm25)
		err = results.Scan(&tempData.ID, &tempData.Country, &tempData.City, &tempData.Latitude, &tempData.Longtitude)
		if err != nil {
			return nil, err
		}
		resultQuest.Data = append(resultQuest.Data, tempData)
	}
	fmt.Println(query)
	fmt.Println(len(resultQuest.Data))
	return &resultQuest, nil
}
