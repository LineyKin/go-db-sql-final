package main

import (
	"database/sql"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

var (
	// randSource источник псевдо случайных чисел.
	// Для повышения уникальности в качестве seed
	// используется текущее время в unix формате (в виде числа)
	randSource = rand.NewSource(time.Now().UnixNano())
	// randRange использует randSource для генерации случайных чисел
	randRange = rand.New(randSource)
)

// getTestParcel возвращает тестовую посылку
func getTestParcel() Parcel {
	return Parcel{
		Client:    1000,
		Status:    ParcelStatusRegistered,
		Address:   "test",
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
}

// TestAddGetDelete проверяет добавление, получение и удаление посылки
func TestAddGetDelete(t *testing.T) {
	// подключаемся к БД
	db, err := sql.Open("sqlite", "demo.db")
	// тест подключения к БД. Останавливаем, если ошибка
	require.NoError(t, err)
	defer db.Close()

	store := NewParcelStore(db)
	parcel := getTestParcel()

	// add
	// добавим новую запись
	newParcelNumber, err := store.Add(parcel)
	// тест, что нет ошибки
	require.NoError(t, err)

	// тест на возврат идентификатора новой записи
	require.NotEmpty(t, newParcelNumber)

	// добавим полученный идентификатор нашёй тестовой посылке
	parcel.Number := newParcelNumber

	// get
	// получим нашу новую запись
	newParcel, err := store.Get(newParcelNumber)

	// тест на отсутствие ошибок. Если есть - тормозим
	require.NoError(t, err)

	// тест на одиннаковое заполнение полей
	assert.Equal(t, newParcel, parcel)

	// delete
	// удаляем нашу новую посылку
	err = store.Delete(newParcelNumber)

	// тест на отсутствие ошибок при удалении
	require.NoError(t, err)

	// попробуем найти удалённую запись в базе
	_, err = store.Get(newParcelNumber)

	// тест на наличие ошибки sql.ErrNoRows
	require.Equal(t, err, sql.ErrNoRows)
}

// TestSetAddress проверяет обновление адреса
func TestSetAddress(t *testing.T) {
	// подключаемся к БД
	db, err := sql.Open("sqlite", "demo.db")
	// тест подключения к БД. Останавливаем, если ошибка
	require.NoError(t, err)
	defer db.Close()

	store := NewParcelStore(db)
	parcel := getTestParcel()

	// add
	// добавим новую запись
	newParcelNumber, err := store.Add(parcel)
	// тест, что нет ошибки
	require.NoError(t, err)

	// тест на возврат идентификатора новой записи
	require.NotEmpty(t, newParcelNumber)

	// set address
	newAddress := "new test address"
	// обновим адрес
	err = store.SetAddress(newParcelNumber, newAddress)

	// тест, что нет ошибок
	require.NoError(t, err)

	// check
	// получите добавленную посылку и убедитесь, что адрес обновился
	// получим нашу новую запись
	newParcel, err := store.Get(newParcelNumber)

	// тест на отсутствие ошибок. Если есть - тормозим
	require.NoError(t, err)

	// тест на обновление адреса
	require.Equal(t, newParcel.Address, newAddress)
}

// TestSetStatus проверяет обновление статуса
func TestSetStatus(t *testing.T) {
	// prepare
	db, err := // настройте подключение к БД

	// add
	// добавьте новую посылку в БД, убедитесь в отсутствии ошибки и наличии идентификатора

	// set status
	// обновите статус, убедитесь в отсутствии ошибки

	// check
	// получите добавленную посылку и убедитесь, что статус обновился
}

// TestGetByClient проверяет получение посылок по идентификатору клиента
func TestGetByClient(t *testing.T) {
	// prepare
	db, err := // настройте подключение к БД

	parcels := []Parcel{
		getTestParcel(),
		getTestParcel(),
		getTestParcel(),
	}
	parcelMap := map[int]Parcel{}

	// задаём всем посылкам один и тот же идентификатор клиента
	client := randRange.Intn(10_000_000)
	parcels[0].Client = client
	parcels[1].Client = client
	parcels[2].Client = client

	// add
	for i := 0; i < len(parcels); i++ {
		id, err := // добавьте новую посылку в БД, убедитесь в отсутствии ошибки и наличии идентификатора

		// обновляем идентификатор добавленной у посылки
		parcels[i].Number = id

		// сохраняем добавленную посылку в структуру map, чтобы её можно было легко достать по идентификатору посылки
		parcelMap[id] = parcels[i]
	}

	// get by client
	storedParcels, err := // получите список посылок по идентификатору клиента, сохранённого в переменной client
	// убедитесь в отсутствии ошибки
	// убедитесь, что количество полученных посылок совпадает с количеством добавленных

	// check
	for _, parcel := range storedParcels {
		// в parcelMap лежат добавленные посылки, ключ - идентификатор посылки, значение - сама посылка
		// убедитесь, что все посылки из storedParcels есть в parcelMap
		// убедитесь, что значения полей полученных посылок заполнены верно
	}
}
