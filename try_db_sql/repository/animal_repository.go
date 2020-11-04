package repository

import (
	"context"

	"go_playground/try_db_sql/entity"
)

func (r *Repository) FindAllAnimals(
	ctx context.Context,
) ([]*entity.Animal, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := r.db.Query("select * from animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []*entity.Animal
	for rows.Next() {
		animal := &entity.Animal{}
		err := rows.Scan(&animal.ID, &animal.Name)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}

	return animals, nil
}
