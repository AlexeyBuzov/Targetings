package examples

import (
	"TargetingsLogic/internal/enums"
	"TargetingsLogic/internal/models"
	"github.com/AlekSi/pointer"
)

var Scenario1 = models.Node{
	Type:  enums.LOGICS,
	Logic: pointer.ToString(enums.AND),
	Children: &[]models.Node{
		{
			Type: enums.VALUE,
			Value: &models.Targeting{
				Type:  enums.SEGMENT,
				Value: 1234, // Кой-то сегмент
			},
		},
		{
			Type:  enums.LOGICS,
			Logic: pointer.ToString(enums.OR),
			Children: &[]models.Node{
				{
					Type: enums.VALUE,
					Value: &models.Targeting{
						Type:  enums.LOCATION,
						Value: 1, // Москва
						Children: &[]models.Node{
							{
								Type:  enums.LOGICS,
								Logic: pointer.ToString(enums.AND),
								Children: &[]models.Node{
									{
										Type:  enums.LOGICS,
										Logic: pointer.ToString(enums.OR),
										Children: &[]models.Node{
											{
												Type: enums.VALUE,
												Value: &models.Targeting{
													Type:  enums.MICROCAT,
													Value: 1, // Транспорт
												},
											},
											{
												Type: enums.VALUE,
												Value: &models.Targeting{
													Type:  enums.MICROCAT,
													Value: 2, // Недвижимость
												},
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Type: enums.VALUE,
					Value: &models.Targeting{
						Type:  enums.LOCATION,
						Value: 2, // СПб
						Children: &[]models.Node{
							{
								Type:  enums.LOGICS,
								Logic: pointer.ToString(enums.AND),
								Children: &[]models.Node{
									{
										Type:  enums.LOGICS,
										Logic: pointer.ToString(enums.OR),
										Children: &[]models.Node{
											{
												Type: enums.VALUE,
												Value: &models.Targeting{
													Type:  enums.MICROCAT,
													Value: 1, // Транспорт
												},
											},
											{
												Type:  enums.LOGICS,
												Logic: pointer.ToString(enums.NOT),
												Children: &[]models.Node{
													{
														Type: enums.VALUE,
														Value: &models.Targeting{
															Type:  enums.MICROCAT,
															Value: 3, // Мотоциклы
														},
													},
												},
											},
											{
												Type:  enums.LOGICS,
												Logic: pointer.ToString(enums.AND),
												Children: &[]models.Node{
													{
														Type: enums.VALUE,
														Value: &models.Targeting{
															Type:  enums.MICROCAT,
															Value: 2, // Недвижимость
														},
													},
													{
														Type: enums.VALUE,
														Value: &models.Targeting{
															Type:  enums.SEGMENT,
															Value: 4321, // Кой-то сегмент
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
