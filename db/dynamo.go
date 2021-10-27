package db

import (
	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

func MakeSureTableExist() {
	logger := ylog.GetLogger()

	db := getConnection()

	if err := db.CreateTable(LAST_TBL, LastOfuro{}).Run(); err != nil {
		logger.D(LAST_TBL+" Table already created.", err)
	}
	if err := db.CreateTable(HIST_TBL, OfuroHistory{}).Run(); err != nil {
		logger.D(HIST_TBL+" Table already created.", err)
	}
}

func GetLastOfuro() *LastOfuro {
	logger := ylog.GetLogger()

	db := getConnection()

	lastTable := db.Table(LAST_TBL)

	// get the same item
	var result LastOfuro
	err := lastTable.Get("key", LAST_TBL_KEY).One(&result)
	if err != nil {
		logger.D("failed GET:", err)
		return nil
	}
	logger.D(result)

	return &result
}

func PutLastOfuro(lastOfuro *LastOfuro) error {

	db := getConnection()
	lastTable := db.Table(LAST_TBL)
	err := lastTable.Put(lastOfuro).Run()

	return err
}

func PutHistory(history *OfuroHistory) error {
	db := getConnection()
	lastTable := db.Table(HIST_TBL)
	err := lastTable.Put(history).Run()

	return err
}
