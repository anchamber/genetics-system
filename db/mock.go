package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"

	apiModel "github.com/anchamber/genetics-api/model"
	"github.com/anchamber/genetics-system/db/model"
)

type SytemDBMock struct {
	DB *sqlx.DB
}

var systems []*model.System = []*model.System{
	{Name: "doctor", Location: "tardis", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "rick", Location: "c-137", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "morty", Location: "herry-herpson", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
	{Name: "obi", Location: "high_ground", Type: model.Techniplast, CleaningInterval: 90, LastCleaned: time.Now()},
}

func NewMockDB() SytemDBMock {
	mock := SytemDBMock{DB: initDB("test.sqlite")}

	for _, system := range systems {
		mock.Insert(system)
	}

	return mock
}

func (o *Options) createPaginationClause() string {
	if o.Pageination == nil {
		return ""
	}
	var limit int64 = int64(o.Pageination.Limit)
	if limit <= 0 {
		limit = -1
	}
	return fmt.Sprintf("LIMIT %d OFFSET %d", limit, o.Pageination.Offset)
}

func getOperatorAsString(operator apiModel.Operator) string {
	switch operator {
	case apiModel.EQ:
		return "="
	case apiModel.GREATER:
		return ">"
	case apiModel.GREATER_EQ:
		return ">="
	case apiModel.SMALLLER:
		return "<"
	case apiModel.SMALLER_EQ:
		return "<="
	case apiModel.CONTAINS:
		return "LIKE"
	default:
		return "="
	}
}

func (o *Options) createFilterClause() string {
	if len(o.Filters) == 0 {
		return ""
	}
	whereClause := "WHERE "

	for index, filter := range o.Filters {
		if index > 0 {
			whereClause += " AND "
		}
		whereClause += fmt.Sprintf("%s %v :%s", filter.Key, getOperatorAsString(filter.Operator), filter.Key)
	}
	fmt.Println(whereClause)
	return whereClause
}

func (o *Options) createFilterMap() map[string]interface{} {
	values := make(map[string]interface{})
	for _, filter := range o.Filters {
		values[filter.Key] = filter.Value
	}
	return values
}

func (systemDB SytemDBMock) Select(options Options) ([]*model.System, error) {
	selectStatement := fmt.Sprintf("SELECT name, location, type, cleaning_interval, last_cleaned FROM systems %s %s;", options.createFilterClause(), options.createPaginationClause())
	fmt.Println(selectStatement)
	filterValues := options.createFilterMap()
	rows, err := systemDB.DB.NamedQuery(selectStatement, filterValues)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	defer rows.Close()
	var data []*model.System
	for rows.Next() {
		var entry model.System
		err = rows.Scan(&entry.Name, &entry.Location, &entry.Type, &entry.CleaningInterval, &entry.LastCleaned)
		if err != nil {
			return nil, err
		}
		data = append(data, &entry)
	}

	return data, nil
}

func (systemDB SytemDBMock) SelectByName(name string) (*model.System, error) {
	selectStatement := `
		SELECT name, location, type, cleaning_interval, last_cleaned 
		FROM systems
		WHERE name = $1;
	`
	rows, err := systemDB.DB.Query(selectStatement, name)
	if err != nil {
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	var entry model.System
	if !rows.Next() {
		return nil, nil
	}
	defer rows.Close()
	err = rows.Scan(&entry.Name, &entry.Location, &entry.Type, &entry.CleaningInterval, &entry.LastCleaned)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (systemDB SytemDBMock) Insert(system *model.System) error {
	var errorString string = ""
	insertStatement := `
		INSERT INTO systems (name, location, type, cleaning_interval, last_cleaned)
			VALUES (?, ?, ?, ?, ?);
	`
	tx, err := systemDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(system.Name, system.Location, system.Type, system.CleaningInterval, system.LastCleaned)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			switch sqliteErr.Code {
			case sqlite3.ErrConstraint:
				errorString = string(SystemAlreadyExists)
			default:
				fmt.Printf("%v\n", sqliteErr)
				errorString = string(Unknown)
			}
		} else {
			fmt.Printf("%v\n", err.Error())
			errorString = string(Unknown)
		}
		tx.Rollback()
	} else {
		numberCreated, _ := result.RowsAffected()
		fmt.Printf("created %d entries\n", numberCreated)
		tx.Commit()
	}
	if errorString == "" {
		return nil
	}
	return errors.New(errorString)
}

func (systemDB SytemDBMock) Update(system *model.System) error {
	insertStatement := `
		UPDATE systems 
			SET name = $1, location = $2 type = $3, cleaning_interval = $4, last_cleaned = $5
			WHERE name = $1;
	`
	tx, err := systemDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(system.Name, system.Location, system.Type, system.CleaningInterval, system.LastCleaned)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		return err
	}
	tx.Commit()
	return nil
}

func initDB(filepath string) *sqlx.DB {
	os.Remove(filepath)
	db, err := sqlx.Connect("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	CreateTables(db)
	CreateIndexes(db)

	return db
}

func CreateTables(db *sqlx.DB) error {

	systemTable := `
		CREATE TABLE IF NOT EXISTS systems(
			name							TEXT		PRIMARY KEY NOT NULL,
			location					TEXT,
			type							TEXT,
			cleaning_interval INT,
			last_cleaned			DATE
		);
	`

	_, err := db.Exec(systemTable)
	if err != nil {
		log.Fatalf("failed to create table\n")
		return err
	}
	return nil
}

func CreateIndexes(db *sqlx.DB) error {

	systemIndex := `
		CREATE UNIQUE INDEX IF NOT EXISTS idx_system_name ON systems(name);
	`

	_, err := db.Exec(systemIndex)
	if err != nil {
		log.Fatalf("failed to create table\n")
		return err
	}
	return nil
}
