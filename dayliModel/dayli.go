package dayliModel

import (
	"dayli/utils"
	"fmt"
	"log"
)

type DayliEntry struct {
	Id_entry int
	Text     string
	Id_user  int64
}

func (d DayliEntry) AddEntry() {
	tx, err := db.Begin()
	utils.CheckError(err)

	stmt, err := tx.Prepare(`insert into dayli_entry(text, id_user)
							 values(?, ?)`)
	utils.CheckError(err)
	defer stmt.Close()

	_, err = stmt.Exec(d.Text, d.Id_user)
	utils.CheckError(err)

	err = tx.Commit()
	utils.CheckError(err)
}

func GetEntry(sqlGetEntrys string) []DayliEntry {
	rows, err := db.Query(sqlGetEntrys)
	utils.CheckError(err)

	defer rows.Close()
	var dayliEntry []DayliEntry

	for rows.Next() {
		var de DayliEntry

		err := rows.Scan(&de.Id_entry, &de.Text, &de.Id_user)
		utils.CheckError(err)

		dayliEntry = append(dayliEntry, de)
	}
	utils.CheckError(err)

	return dayliEntry
}

func GetEntrysByUser(id_user int64) []DayliEntry {
	sqlGetEntrys := fmt.Sprintf("select * from dayli_entry where Id_user = %d", id_user)
	return GetEntry(sqlGetEntrys)
}

func GetEntryByIdEntryAndIdUser(id_entry int, id_user int64) []DayliEntry {
	sqlGetEntrys := fmt.Sprintf("select * from dayli_entry where Id_entry = %d and Id_user = %d", id_entry, id_user)
	return GetEntry(sqlGetEntrys)
}

func DeleteEntry(id_entry int, id_user int64) {
	sqlDelEntry := fmt.Sprintf("delete from dayli_entry where Id_entry = %d and Id_user = %d", id_entry, id_user)
	_, err := db.Exec(sqlDelEntry)
	utils.CheckError(err)
}

func CreateTableDayli() {
	prepareSql := `
		create table dayli_entry (id_entry integer primary key AUTOINCREMENT,
								  Text text,
								  Id_user integer)
	`
	_, err := db.Exec(prepareSql)
	if err != nil {
		log.Printf("%q: %s\n", err, prepareSql)
	}
}
