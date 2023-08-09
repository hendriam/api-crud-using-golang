package model

type Brand struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

func (model ModelApp) SelectBrand() ([]Brand, error) {
	var brands []Brand

	ctx, cancel := defaultContext()
	defer cancel()

	rows, err := model.db.QueryContext(
		ctx,
		`SELECT * FROM brands`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var brand Brand
		if err := rows.Scan(
			&brand.Id,
			&brand.Name,
		); nil != err {
			return nil, err
		}

		brands = append(brands, brand)
	}

	return brands, nil
}

func (model ModelApp) SelectBrandById(brandId int) (Brand, error) {
	var brand Brand

	ctx, cancel := defaultContext()
	defer cancel()

	err := model.db.QueryRowContext(
		ctx,
		"SELECT * FROM brands WHERE id=?",
		brandId,
	).Scan(&brand.Id, &brand.Name)
	if err != nil {
		return Brand{}, err
	}

	return brand, err
}

func (model ModelApp) InsertBrand(body Brand) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	brand, err := model.db.ExecContext(
		ctx,
		"INSERT INTO brands (name) VALUE (?)",
		body.Name,
	)
	if err != nil {
		return 0, err
	}

	row, err := brand.LastInsertId()
	if err != nil {
		return 0, err
	}

	return row, err
}

func (model ModelApp) UpdateBrand(id int, body Brand) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	brand, err := model.db.ExecContext(
		ctx,
		"UPDATE brands SET name=? WHERE id =?",
		body.Name,
		id,
	)
	if err != nil {
		return 0, err
	}

	row, err := brand.RowsAffected()
	if err != nil {
		return 0, err
	}

	return row, err
}

func (model ModelApp) DeleteBrand(id int) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	brand, err := model.db.ExecContext(
		ctx,
		"DELETE FROM brands WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}

	row, err := brand.RowsAffected()
	if err != nil {
		return 0, err
	}

	return row, err
}
