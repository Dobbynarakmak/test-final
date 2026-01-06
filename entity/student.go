package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FullName	string	`valid:"required~FullName is required"`
	Age			int		`valid:"int,range(18|200)~Age must be at least 18"`
	Email		string	`valid:"email~Email is invalid,required~Email is required"`
	GPA			float32	`valid:"float,range(0|4)~GPA must be between 0.00 and 4.00"`
}