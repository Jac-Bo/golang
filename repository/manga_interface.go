package repository

import "github.com/Jac-Bo/golang/model"

type MangaRepositoryInterface interface {
	InsertManga(model.PostManga) bool              // Function that Inserts Data from Func, returns a bool
	GetAllManga() []model.Manga                    // Function that Retrieves Data From Func
	GetOneManga(uint) model.Manga                  // Gets One Manga with Parameter of an ID
	UpdateManga(uint, model.PostManga) model.Manga // Updates manga via reference of integer and returns the model of type Manga
	DeleteManga(uint) bool                         // Deletes manga and returns bool for the verification of deleted/not
}
