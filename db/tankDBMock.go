package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"

	"github.com/anchamber/genetics-system/db/model"
)

type TankDBMock struct {
	DB *sqlx.DB
}

var MockDataTanks = []*model.Tank{}

func NewTankDBMock(initialData []*model.Tank) TankDBMock {
	if initialData == nil {
		initialData = MockDataTanks
	}
	mock := TankDBMock{
		DB: initTankDB(),
	}
	mock.DB.SetMaxOpenConns(1)
	for _, tank := range initialData {
		err := mock.Insert(tank)
		if err != nil {
			panic(err)
		}
	}

	return mock
}

func (tankDB TankDBMock) Select(options Options) ([]*model.Tank, error) {
	selectStatement := fmt.Sprintf("SELECT id, number, system, active, size, fish_count FROM tanks %s %s;", options.createFilterClause(), options.createPaginationClause())
	// fmt.Println(selectStatement)
	filterValues := options.createFilterMap()
	rows, err := tankDB.DB.NamedQuery(selectStatement, filterValues)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf(fmt.Sprintf("failed closing rows %v\n", err))
		}
	}(rows)
	var data []*model.Tank
	for rows.Next() {
		var entry model.Tank
		err = rows.Scan(&entry.ID, &entry.Number, &entry.System, &entry.Active, &entry.Size, &entry.FishCount)
		if err != nil {
			return nil, err
		}
		data = append(data, &entry)
	}

	return data, nil
}

func (tankDB TankDBMock) SelectByNumber(number uint32) (*model.Tank, error) {
	//goland:noinspection ALL
	selectStatement := `
		SELECT id, number, system, active, size, fish_count
		FROM tanks
		WHERE number = $1;
	`
	rows, err := tankDB.DB.Query(selectStatement, number)
	if err != nil {
		log.Fatalf(`failed to select all`)
		return nil, err
	}

	var entry model.Tank
	if !rows.Next() {
		return nil, nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Printf(fmt.Sprintf("failed closing rows %v\n", err))
		}
	}(rows)
	err = rows.Scan(&entry.ID, &entry.Number, &entry.System, &entry.Active, &entry.Size, &entry.FishCount)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (tankDB TankDBMock) Insert(tank *model.Tank) error {
	//goland:noinspection ALL
	insertStatement := `
		INSERT INTO tanks (number, system, active, size, fish_count)
			VALUES (?, ?, ?, ?, ?);
	`
	tx, err := tankDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Printf(fmt.Sprintf("failed closing statement %v\n", err))
		}
	}(statement)

	_, err = statement.Exec(tank.Number, tank.System, tank.Active, tank.Size, tank.FishCount)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			switch sqliteErr.Code {
			case sqlite3.ErrConstraint:
				return &EntityAlreadyExists{entity: "tank"}
			default:
				fmt.Printf("%v\n", sqliteErr)
				return &UnknownDBError{message: "Unknown error occurred"}
			}
		} else {
			fmt.Printf("%v\n", err.Error())
			return &UnknownDBError{message: "Unknown error occurred"}
		}
	} else {
		// numberCreated, _ := result.RowsAffected()
		// fmt.Printf("created %d entries\n", numberCreated)
		err := tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func (tankDB TankDBMock) Update(tank *model.Tank) error {
	//goland:noinspection ALL
	insertStatement := `
		UPDATE tanks 
			SET number = $1, system = $2, active = $3, size = $4, fish_count = $5
			WHERE number = $1;
	`
	tx, err := tankDB.DB.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction\n")
		return err
	}

	statement, err := tx.Prepare(insertStatement)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Printf(fmt.Sprintf("failed closing statement %v\n", err))
		}
	}(statement)

	_, err = statement.Exec(tank.Number, tank.System, tank.Active, tank.Size, tank.FishCount)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (tankDB TankDBMock) Delete(number uint32) error {
	//goland:noinspection ALL
	statementString := `
		DELETE FROM tanks WHERE number = ?;
	`
	statement, err := tankDB.DB.Prepare(statementString)
	if err != nil {
		fmt.Printf("failed to prepare statement\n")
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Printf(fmt.Sprintf("failed closing statement %v\n", err))
		}
	}(statement)

	_, err = statement.Exec(number)
	if err != nil {
		fmt.Printf("failed to execute statement\n")
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			switch sqliteErr.Code {
			case sqlite3.ErrConstraint:
			default:
			}
		} else {
			fmt.Printf("%v\n", err.Error())
			return &UnknownDBError{message: "Unknown error occurred"}
		}
	}
	return nil
}

func initTankDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	err = CreateTankTables(db)
	if err != nil {
		return nil
	}

	return db
}

func CreateTankTables(db *sqlx.DB) error {
	//goland:noinspection ALL
	tankTable := `
		CREATE TABLE IF NOT EXISTS tanks(
			id					INTEGER	PRIMARY KEY AUTOINCREMENT,
			number				INT UNIQUE,
			system				string,
			active				bit ,
			size				INT,
			fish_count 			INT
		);
	`

	_, err := db.Exec(tankTable)
	if err != nil {
		log.Fatalf("failed to create table\n")
		return err
	}
	return nil
}
