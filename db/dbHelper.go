package db

import (
	"strconv"
	"time"
	"www/loghelper"
	// "datebase/sql"
	"database/sql"
	// github
	_ "github.com/lib/pq"
)

//Add data
func Add(title string, newName string) {
	db, err := sql.Open("postgres", "user=linweiyan password=linweiyan dbname=linweiyan sslmode=disable")
	checkErr(err, 0)

	stmt, err := db.Prepare("insert into imagemaptable(title, newname, uploaddate) values($1, $2, $3)")
	checkErr(err, 1)

	_, err = stmt.Exec("ajdfjadjfjadfja", "499274592952945235242541", time.Now())
	checkErr(err, 2)
}

func checkErr(err error, index int) {
	if err != nil {
		loghelper.Info.Println("error:" + err.Error() + ":" + strconv.Itoa(index))
		panic(err)
	}
}
