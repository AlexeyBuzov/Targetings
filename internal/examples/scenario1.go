package examples

import (
	"TargetingsLogic/internal/enums"
	"TargetingsLogic/internal/models"
	"github.com/AlekSi/pointer"
)

var Scenario1 = models.Scenario{
	Requests: []models.ForecastRequest{
		{
			Targetings: []models.Targeting{
				{
					Type:      enums.LOCATION,
					Value:     []int64{1, 2},             // Мос. обл., Лен. обл.
					Exception: pointer.To([]int64{1, 2}), // Москва, СПБ
				},
				{
					Type:      enums.MICROCAT,
					Value:     []int64{1},             // Авто
					Exception: pointer.To([]int64{2}), // С пробегом
				},
				{
					Type:  enums.INTEREST,
					Value: []int64{1}, // Покупка недвижимости
				},
			},
		},
		{
			Targetings: []models.Targeting{
				{
					Type:      enums.LOCATION,
					Value:     []int64{1, 2},             // Мос. обл., Лен. обл.
					Exception: pointer.To([]int64{1, 2}), // Москва, СПБ
				},
				{
					Type:      enums.MICROCAT,
					Value:     []int64{1},             // Авто
					Exception: pointer.To([]int64{2}), // С пробегом
				},
				{
					Type:  enums.INTEREST,
					Value: []int64{1}, // Покупка недвижимости
				},
			},
		},
	},
}
