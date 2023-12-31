package api

type EnrichedFio struct {
	ID         int     `json:"id" validate:"gt=0"`
	Name       *string `json:"name,omitempty" validate:"omitempty,name"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
	Age        *int    `json:"age,omitempty" validate:"omitempty,gte=0"`
	Gender     *string `json:"gender,omitempty"`
	Country    *string `json:"country,omitempty" validate:"omitempty,uppercase,len=2"`
}

type FailedFio struct {
	Name       string  `json:"name" validate:"name"`
	Surname    string  `json:"surname" validate:"required"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type Filter struct {
	Limit  int     `json:"limit" validate:"gt=0"`
	Offset int     `json:"offset" validate:"gte=0,ltcsfield=Limit"`
	Order  *string `json:"order,omitempty" validate:"omitempty,order"`
	MinAge *int    `json:"minAge,omitempty" validate:"omitempty,gte=0"`
	MaxAge *int    `json:"maxAge,omitempty" validate:"omitempty,gte=0"`
}

type Fio struct {
	Name       string  `json:"name" validate:"name"`
	Surname    string  `json:"surname" validate:"required"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type User struct {
	ID         int     `json:"id" validate:"gt=0"`
	Name       string  `json:"name" validate:"name"`
	Surname    string  `json:"surname" validate:"required"`
	Patronymic *string `json:"patronymic,omitempty"`
	Age        int     `json:"age" validate:"gte=0"`
	Gender     string  `json:"gender" validate:"required"`
	Country    string  `json:"country" validate:"required,uppercase,len=2"`
}
