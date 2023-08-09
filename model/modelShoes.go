package model

type Shoes struct {
	Id         int     `json:"id"`
	Name       string  `json:"name" binding:"required"`
	BrandId    int     `json:"brand_id" binding:"required"`
	BrandName  string  `json:"brand_name"`
	Price      float32 `json:"price" binding:"required"`
	PictureUrl string  `json:"picture_url"`
}

func (model ModelApp) SelectShoes() ([]Shoes, error) {
	var shoes []Shoes

	ctx, cancel := defaultContext()
	defer cancel()

	rows, err := model.db.QueryContext(
		ctx,
		`SELECT shoes.id, shoes.name, brands.id as brand_id, brands.name as brand_name, shoes.price, shoes.picture_url
			FROM shoes JOIN brands ON shoes.brand_id = brands.id ORDER BY shoes.id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var shoe Shoes
		if err := rows.Scan(
			&shoe.Id,
			&shoe.Name,
			&shoe.BrandId,
			&shoe.BrandName,
			&shoe.Price,
			&shoe.PictureUrl,
		); nil != err {
			return nil, err
		}

		shoes = append(shoes, shoe)
	}

	return shoes, err
}

func (model ModelApp) SelectShoesById(shoesId int) (Shoes, error) {
	var shoes Shoes

	ctx, cancel := defaultContext()
	defer cancel()

	err := model.db.QueryRowContext(
		ctx,
		`SELECT shoes.id, shoes.name, brands.id as brand_id, brands.name as brand_name, shoes.price, shoes.picture_url
			FROM shoes JOIN brands ON shoes.brand_id = brands.id WHERE shoes.id=?`,
		shoesId,
	).Scan(&shoes.Id, &shoes.Name, &shoes.BrandId, &shoes.BrandName, &shoes.Price, &shoes.PictureUrl)
	if err != nil {
		return Shoes{}, err
	}

	return shoes, err
}

func (model ModelApp) InsertShoes(body Shoes) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	shoes, err := model.db.ExecContext(
		ctx,
		`INSERT INTO shoes (name, brand_id, price, picture_url) VALUE (?, ?, ?, ?)`,
		body.Name,
		body.BrandId,
		body.Price,
		body.PictureUrl,
	)
	if err != nil {
		return 0, err
	}

	row, err := shoes.LastInsertId()
	if err != nil {
		return 0, err
	}

	return row, err
}

func (model ModelApp) UpdateShoes(id int, body Shoes) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	shoes, err := model.db.ExecContext(
		ctx,
		"UPDATE shoes SET name=?, brand_id=?, price=?, picture_url=? WHERE id =?",
		body.Name,
		body.BrandId,
		body.Price,
		body.PictureUrl,
		id,
	)
	if err != nil {
		return 0, err
	}

	row, err := shoes.RowsAffected()
	if err != nil {
		return 0, err
	}

	return row, err
}

func (model ModelApp) DeleteShoes(id int) (int64, error) {
	ctx, cancel := defaultContext()
	defer cancel()

	shoes, err := model.db.ExecContext(
		ctx,
		"DELETE FROM shoes WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}

	row, err := shoes.RowsAffected()
	if err != nil {
		return 0, err
	}

	return row, err
}
