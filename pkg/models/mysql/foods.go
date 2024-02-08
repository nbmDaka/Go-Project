package mysql

//import (
//	"AITUNews/pkg/models"
//	"database/sql"
//)
//
//type FoodModel struct {
//	DB *sql.DB
//}
//
//// This will return the 50 most recently created news.
//func (m *FoodModel) Latest() ([]*models.Foods, error) {
//	stmt := `SELECT id, meal_name, weekday, quantity FROM canteen_menu
//     ORDER BY id DESC LIMIT 50`
//
//	rows, err := m.DB.Query(stmt)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	foods := []*models.Foods{}
//
//	for rows.Next() {
//		s := &models.Foods{}
//		err = rows.Scan(&s.ID, &s.Meal_name, &s.Weekday, &s.Quantity)
//		if err != nil {
//			return nil, err
//		}
//		foods = append(foods, s)
//	}
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return foods, nil
//}
//
//func (m *FoodModel) InsertFood(meal_name, weekday, quantity string) (int, error) {
//	stmt := `INSERT INTO canteen_menu (meal_name, weekday, quantity) VALUES(?, ?, ? )`
//
//	result, err := m.DB.Exec(stmt, meal_name, weekday, quantity)
//	if err != nil {
//		return 0, err
//	}
//
//	id, err := result.LastInsertId()
//	if err != nil {
//		return 0, err
//	}
//
//	return int(id), nil
//
//}
