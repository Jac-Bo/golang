package repository

import (
	"database/sql"
	"log"

	"github.com/Jac-Bo/golang/model"
)

type MangaRepository struct {
	DB *sql.DB
}

func NewMangaRepository(db *sql.DB) MangaRepositoryInterface {
	return &MangaRepository{DB: db}
}

// Implements delete Manga
func (m *MangaRepository) DeleteManga(id uint) bool {
	_, err := m.DB.Exec("DELETE FROM manga WHERE ID = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetOneManga implements MangaRepositoryInterface
func (m *MangaRepository) GetAllManga() []model.Manga {
	query, err := m.DB.Query("SELECT * FROM manga")
	if err != nil {
		log.Println(err)
		return nil
	}
	var mangas []model.Manga
	if query != nil {
		for query.Next() {
			var (
				id       uint
				title    string
				genre    string
				volumes  uint8
				chapters uint16
				author   string
			)
			err := query.Scan(&id, &title, &genre, &volumes, &chapters, &author)
			if err != nil {
				log.Println(err)
			}
			//Creates a new manga of type model.Manga and adds the new manga var into the manga db
			manga := model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
			mangas = append(mangas, manga)

		}
	}
	return mangas
}

// GetOneManga implements MangaRepositoryInterface
func (m *MangaRepository) GetOneManga(id uint) model.Manga {
	query, err := m.DB.Query("SELECT * FROM manga WHERE id = $1", id) // query variable refers to a dataset that is queried

	if err != nil { //Error handler
		log.Println(err)
		return model.Manga{}
	}

	var manga model.Manga // Initializes the model "manga"

	if query != nil {
		for query.Next() {
			var ( // creates a variable that are a reference and will be used to apply data via the "&" operand
				id       uint
				title    string
				genre    string
				volumes  uint8
				chapters uint16
				author   string
			)
			err := query.Scan(&id, &title, &genre, &volumes, &chapters, &chapters, &author) //
			if err != nil {
				log.Println(err)
			} /*   Creates a model of the retrieved values to be returned to the user.  */
			manga = model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
		}
	}
	return manga // Return The Manga that has been retrieved via the Query
}

// InsertManga implements MangaRepositoryInterface
func (m *MangaRepository) InsertManga(post model.PostManga) bool {
	stmt, err := m.DB.Prepare("INSERT INTO manga(title, genre, volumes, chapters, author) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Title, post.Genre, post.Volumes, post.Chapters, post.Author)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// UpdateManga implements MangaRepositoryInterface
func (m *MangaRepository) UpdateManga(id uint, post model.PostManga) model.Manga {
	_, err := m.DB.Exec("UPDATE manga SET title = $1, genre = $2, volumes = $3, chapters = $4, author = $5 WHERE id = $6", post.Title, post.Genre, post.Volumes, post.Chapters, post.Author, id)
	if err != nil {
		log.Println(err)
		return model.Manga{}
	}
	return m.GetOneManga(id)
}
