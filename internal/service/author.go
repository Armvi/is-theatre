package service

import (
	"fmt"
	"is-theatre/internal/datastruct"
	"is-theatre/internal/entity"
	"is-theatre/internal/repository"
)

type AuthorService interface {
	CreateAuthor(userId int64, author entity.Author) (*int64, error)
	GetAuthor(id int64) (*entity.Author, error)
	GetAuthors(offset, limit int) ([]entity.Author, error)
	GetAuthorsByCountry(country string) ([]entity.Author, error)
	DeleteAuthor(userId int64, authorId int64) error
	UpdateAuthor(userId int64, author entity.Author) (*entity.Author, error)
}

type authorService struct {
	dao repository.DAO
}

func NewAuthorService(dao repository.DAO) AuthorService {
	return &authorService{dao: dao}
}

func (a *authorService) CreateAuthor(userId int64, author entity.Author) (*int64, error) {

	user, err := a.dao.NewUserQuery().GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("create  error: %w", err)
	}

	if user.Role != datastruct.ADMIN {
		return nil, fmt.Errorf("create author error: not admin")
	}

	newAuthor := datastruct.Author{
		Name:       author.Name,
		SecondName: author.SecondName,
		Country:    author.Country,
		Century:    author.Century,
	}

	i, err := a.dao.NewAuthorQuery().CreateAuthor(newAuthor)
	if err != nil {
		return nil, fmt.Errorf("create author error: %w", err)
	}

	return i, nil
}

func (a *authorService) GetAuthor(id int64) (*entity.Author, error) {
	author, err := a.dao.NewAuthorQuery().GetAuthor(id)
	if err != nil {
		return nil, fmt.Errorf("get author error: %w", err)
	}

	return &entity.Author{
		Id:         author.Id,
		Name:       author.Name,
		SecondName: author.SecondName,
		Country:    author.Country,
		Century:    author.Century,
	}, nil
}

func (a *authorService) GetAuthors(offset, limit int) ([]entity.Author, error) {
	var authors []entity.Author

	dtAuthors, err := a.dao.NewAuthorQuery().GetAuthorsByCentury(offset, limit)
	if err != nil {
		return nil, err
	}

	for _, author := range dtAuthors {
		authors = append(authors, entity.Author{
			Id:         author.Id,
			Name:       author.Name,
			SecondName: author.SecondName,
			Country:    author.Country,
			Century:    author.Century,
		})
	}

	return authors, nil
}

func (a *authorService) GetAuthorsByCountry(country string) ([]entity.Author, error) {
	var authors []entity.Author

	dtAuthors, err := a.dao.NewAuthorQuery().GetAuthorsByCountry(country)
	if err != nil {
		return nil, err
	}

	for _, author := range dtAuthors {
		authors = append(authors, entity.Author{
			Id:         author.Id,
			Name:       author.Name,
			SecondName: author.SecondName,
			Country:    author.Country,
			Century:    author.Century,
		})
	}

	return authors, nil
}

func (a *authorService) DeleteAuthor(userId int64, authorId int64) error {
	user, err := a.dao.NewUserQuery().GetUser(userId)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	if user.Role != datastruct.ADMIN {
		return fmt.Errorf("delete author error: not admin")
	}

	err = a.dao.NewAuthorQuery().DeleteAuthor(authorId)
	if err != nil {
		return fmt.Errorf("delete author error: %w", err)
	}
	return nil
}

func (a *authorService) UpdateAuthor(userId int64, author entity.Author) (*entity.Author, error) {
	user, err := a.dao.NewUserQuery().GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	if user.Role != datastruct.ADMIN {
		return nil, fmt.Errorf("update author error: not admin")
	}

	updated, err := a.dao.NewAuthorQuery().UpdateAuthor(&datastruct.Author{
		Id:         author.Id,
		Name:       author.Name,
		SecondName: author.SecondName,
		Country:    author.Country,
		Century:    author.Century,
	})

	if err != nil {
		return nil, fmt.Errorf("update author error: %w", err)
	}

	return &entity.Author{
		Id:         updated.Id,
		Name:       updated.Name,
		SecondName: updated.SecondName,
		Country:    updated.Country,
		Century:    updated.Century,
	}, nil
}
