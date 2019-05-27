package logic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateAgeHandler(name string, updateAge int) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "", "127.0.0.1", "3306", "DB_TEST"))
	if err != nil {
		return err
	}
	defer func() {
		_ = db.Close()
	}()

	_, err = db.Exec(fmt.Sprintf("UPDATE TB_USER SET age=%d WHERE name=\"%s\";", updateAge, name))
	if err != nil {
		return err
	}

	return nil
}
