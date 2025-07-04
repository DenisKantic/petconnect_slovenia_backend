package models

type AdoptPost struct {
	Category    string `form:"category" binding:"required"`
	PetName     string `form:"petName" binding:"required"`
	Description string `form:"description" binding:"required"`
	Sex         string `form:"sex" binding:"required"`
	Vaccinated  bool   `form:"vaccinated"`
	Chipped     bool   `form:"chipped"`
	Location    string `form:"location" binding:"required"`
	Images      []string
}
