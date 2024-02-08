package mysql

import (
	"AITUNews/pkg/models"
	"database/sql"
	"errors"
)

type NewsModel struct {
	DB *sql.DB
}

// This will return a specific snippet based on its id.
func (m *NewsModel) Get(id int) (*models.News, error) {

	stmt := `SELECT id, title, content, image_url FROM news WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.News{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Image_url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 50 most recently created news.
func (m *NewsModel) First() ([]*models.News, error) {
	stmt := `SELECT id, title, content, image_url FROM news
     ORDER BY id ASC LIMIT 50`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []*models.News{}

	for rows.Next() {
		s := &models.News{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Image_url)
		if err != nil {
			return nil, err
		}
		news = append(news, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

func (m *NewsModel) Insert(title, content, imageUrl string) (int, error) {
	stmt := `INSERT INTO news (title, content, image_url) VALUES(?, ?, ? )`

	result, err := m.DB.Exec(stmt, title, content, imageUrl)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
