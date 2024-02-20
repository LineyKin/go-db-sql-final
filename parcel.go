package main

import (
	"database/sql"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

// метод, добавляющий новую посылку
func (s ParcelStore) Add(p Parcel) (int, error) {
	sqlPattern := "INSERT INTO parcel (client, status, address, created_at) VALUES(:client, :status, :address, :created_at)"

	res, err := s.db.Exec(
		sqlPattern,
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// метод, получающий информацию по конкретной посылке (по трекеру)
func (s ParcelStore) Get(number int) (Parcel, error) {
	sqlPattern := "SELECT number, client, status, address, created_at FROM parcel WHERE number = :number"
	row := s.db.QueryRow(sqlPattern, sql.Named("number", number))

	// пустой экземпляр структуры Parcel
	p := Parcel{}

	// заполняем экземпляр p данными из БД
	err := row.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)

	return p, err
}

// метод, возвращающий все посылки конкретного клиента
func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {

	sqlPattern := "SELECT number, client, status, address, created_at FROM parcel WHERE client = :client"
	rows, err := s.db.Query(sqlPattern, sql.Named("client", client))

	if err != nil {
		return nil, err
	}

	// заполните срез Parcel данными из таблицы
	var res []Parcel

	for rows.Next() {
		p := Parcel{}
		err := rows.Scan(&p.Number, &p.Client, &p.Status, &p.Address, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}

// метод, меняющий стату посылки
func (s ParcelStore) SetStatus(number int, status string) error {
	sqlPattern := "UPDATE parcel SET status = :status WHERE number = :number"

	_, err := s.db.Exec(
		sqlPattern,
		sql.Named("status", status),
		sql.Named("number", number))

	return err
}

func (s ParcelStore) SetAddress(number int, address string) error {
	// реализуйте обновление адреса в таблице parcel
	// менять адрес можно только если значение статуса registered

	return nil
}

func (s ParcelStore) Delete(number int) error {
	// реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

	return nil
}
