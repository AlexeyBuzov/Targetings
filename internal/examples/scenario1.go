package examples

import (
	"TargetingsLogic/internal/models"
	"github.com/AlekSi/pointer"
)

var Scenario1 = models.Scenario{
	Requests: []models.ForecastRequest{
		{
			Targetings: []models.Targeting{
				{
					Location: pointer.To(models.LocationTargeting{
						Include: []int64{1, 2},             // Мос. обл., Лен. обл.
						Exclude: pointer.To([]int64{1, 2}), // Москва, СПБ
					}),
				},
				{
					Microcategory: pointer.To(models.MicrocategoryTargeting{
						Include: []int64{1},             // Авто
						Exclude: pointer.To([]int64{2}), // С пробегом
					}),
				},
				{
					Interest: pointer.To(models.InterestTargeting{
						Include: []models.Interest{
							{
								Id:    1, // Покупка недвижимости
								Power: "h",
							},
						},
					}),
				},
			},
		},
		{
			Targetings: []models.Targeting{
				{
					Location: pointer.To(models.LocationTargeting{
						Include: []int64{1, 2},             // Мос. обл., Лен. обл.
						Exclude: pointer.To([]int64{1, 2}), // Москва, СПБ
					}),
				},
				{
					Microcategory: pointer.To(models.MicrocategoryTargeting{
						Include: []int64{1},             // Авто
						Exclude: pointer.To([]int64{2}), // С пробегом
					}),
				},
				{
					Interest: pointer.To(models.InterestTargeting{
						Include: []models.Interest{
							{
								Id:    1, // Покупка недвижимости
								Power: "h",
							},
						},
					}),
				},
			},
		},
	},
}
