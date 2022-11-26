package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"is-theatre/internal/datastruct"
)

type AuthorQuery interface {
	CreateAuthor(author datastruct.Author) (*int64, error)
	GetAuthor(id int64) (*datastruct.Author, error)
	UpdateAuthor(author *datastruct.Author) (*datastruct.Author, error)
	DeleteAuthor(id int64) error
	GetAuthors(limit, offset uint64) ([]datastruct.Author, error)
	GetAuthorsByCentury(beg, end int) ([]datastruct.Author, error)
	GetAuthorsByCountry(country string) ([]datastruct.Author, error)
	GetAuthorsByConditions(country *string, century *int) ([]datastruct.Author, error)
}

type authorQuery struct{}

func (a authorQuery) CreateAuthor(author datastruct.Author) (*int64, error) {
	db := dbQueryBuilder().
		Insert(datastruct.AuthorTableName).
		Columns("name", "secondName", "country", "century").
		Values(author.Name, author.SecondName, author.Country, author.Century).
		Suffix("RETURNING id")

	var id int64
	err := db.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("create author error: %w", err)
	}

	return &id, nil
}

func (a authorQuery) GetAuthor(id int64) (*datastruct.Author, error) {
	db := dbQueryBuilder().
		Select("name",
			"secondName",
			"country",
			"century",
			"id").
		From(datastruct.AuthorTableName).
		Where(squirrel.Eq{"id": id})

	au := datastruct.Author{}
	err := db.QueryRow().
		Scan(&au.Name,
			&au.SecondName,
			&au.Country,
			&au.Century,
			&au.Id)
	if err != nil {
		return nil, fmt.Errorf("get author: %w", err)
	}

	return &au, nil
}

func (a authorQuery) UpdateAuthor(author *datastruct.Author) (*datastruct.Author, error) {

	fromDB, err := a.GetAuthor(author.Id)
	if err != nil {
		return nil, fmt.Errorf("updated composition error: %w", err)
	}

	updated := updateAuthor(fromDB, author)

	db := dbQueryBuilder().
		Update(datastruct.UserTableName).
		SetMap(map[string]interface{}{
			"name":       updated.Name,
			"secondName": updated.SecondName,
			"country":    updated.Country,
			"century":    updated.Century,
		}).Where(squirrel.Eq{"id": author.Id}).
		Suffix("RETURNING name, secondName, country, century, id")

	uAuthor := datastruct.Author{}
	err = db.QueryRow().Scan(
		&uAuthor.Name,
		&uAuthor.SecondName,
		&uAuthor.Country,
		&uAuthor.Century,
		&uAuthor.Id,
	)
	if err != nil {
		return nil, fmt.Errorf("update author error: %w", err)
	}

	return &uAuthor, nil
}

func (a authorQuery) DeleteAuthor(id int64) error {
	db := dbQueryBuilder().
		Delete(datastruct.AuthorTableName).
		Where(squirrel.Eq{"id": id})

	_, err := db.Exec()
	if err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (a authorQuery) GetAuthors(limit, offset uint64) ([]datastruct.Author, error) {
	db := dbQueryBuilder().
		Select("name", "secondName", "country", "century", "id").
		From(datastruct.AuthorTableName).
		Limit(limit).Offset(offset)

	var authors []datastruct.Author
	var au datastruct.Author
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by centry error: %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&au.Name, &au.SecondName, &au.Country, &au.Century, &au.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by centry error: %w", err)
		}
		authors = append(authors, au)
	}

	return authors, nil
}

func (a authorQuery) GetAuthorsByCentury(beg, end int) ([]datastruct.Author, error) {
	db := dbQueryBuilder().
		Select("name", "secondName", "country", "century", "id").
		From(datastruct.AuthorTableName).
		Where(squirrel.And{squirrel.LtOrEq{"century": end}, squirrel.Gt{"century": beg}})

	var authors []datastruct.Author
	var au datastruct.Author
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by centry error: %w", err)
	}
	for rows.Next() {
		err = rows.Scan(&au.Name, &au.SecondName, &au.Country, &au.Century, &au.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by centry error: %w", err)
		}
		authors = append(authors, au)
	}

	return authors, nil
}

func (a authorQuery) GetAuthorsByCountry(country string) ([]datastruct.Author, error) {
	db := dbQueryBuilder().
		Select("name", "secondName", "country", "century", "id").
		From(datastruct.AuthorTableName).
		Where(squirrel.Eq{"country": country})

	var authors []datastruct.Author
	var au datastruct.Author
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by country error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&au.Name, &au.SecondName, &au.Country, &au.Century, &au.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by country error: %w", err)
		}
		authors = append(authors, au)
	}

	return authors, nil
}

func (a authorQuery) GetAuthorsByConditions(country *string, century *int) ([]datastruct.Author, error) {
	var conditions squirrel.And
	if country != nil {
		conditions = append(conditions, squirrel.Eq{"country": *country})
	}

	if century != nil {
		conditions = append(conditions, squirrel.Eq{"century": *century})
	}

	db := dbQueryBuilder().
		Select("name", "secondName", "country", "century", "id").
		From(datastruct.AuthorTableName).
		Where(squirrel.Eq{"country": country})

	var authors []datastruct.Author
	var au datastruct.Author
	rows, err := db.Query()
	if err != nil {
		return nil, fmt.Errorf("get authors by country error: %w", err)
	}

	for rows.Next() {
		err = rows.Scan(&au.Name, &au.SecondName, &au.Country, &au.Century, &au.Id)
		if err != nil {
			return nil, fmt.Errorf("get authors by country error: %w", err)
		}
		authors = append(authors, au)
	}

	return authors, nil
}

func updateAuthor(fromDB, new *datastruct.Author) (updated datastruct.Author) {
	updated = *fromDB
	if len(new.Name) > 0 {
		updated.Name = new.Name
	}
	if len(new.SecondName) > 0 {
		updated.SecondName = new.SecondName
	}
	if len(new.Country) > 0 {
		updated.Country = new.Country
	}
	if new.Century != 0 {
		updated.Century = new.Century
	}

	return
}
