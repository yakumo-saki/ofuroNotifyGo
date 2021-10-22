package db

import (
	"time"

	"github.com/yakumo-saki/ofuroNotifyGo/ylog"
)

var logger = ylog.GetLogger("dynamo")

func Test() {

	db := getConnection()

	lastTable := db.Table(LAST_TBL)
	histTable := db.Table(HIST_TBL)

	if err := db.CreateTable(LAST_TBL, LastOfuro{}).Run(); err != nil {
		logger.D(LAST_TBL+" Table already created.", err)
	}
	if err := db.CreateTable(HIST_TBL, OfuroHistory{}).Run(); err != nil {
		logger.D(HIST_TBL+" Table already created.", err)
	}
	logger.D("Table Created.")

	// put item
	l := LastOfuro{Key: LAST_TBL_KEY, UnixTime: int(time.Now().Unix()), InOut: "In", DateTime: ""}
	err := lastTable.Put(l).Run()
	if err != nil {
		logger.E("failed: put", err)
	}
	logger.D("LastOfuro PUT OK.")

	h := LastOfuroToHistory(l)
	err = histTable.Put(h).Run()
	if err != nil {
		logger.E("Histry failed: put", err)
	}
	logger.D("Histry PUT OK.")

	// get the same item
	var result LastOfuro
	err = lastTable.Get("key", LAST_TBL_KEY).One(&result)
	if err != nil {
		logger.D("failed GET", err)
	}
	logger.D(result)

	// get all items
	var results []LastOfuro
	err = lastTable.Scan().All(&results)
	if err != nil {
		logger.D("failed SCAN", err)
	}
	logger.D(results)

	// use placeholders in filter expressions (see Expressions section below)
	// var filtered []widget
	// err = table.Scan().Filter("'Count' > ?", 10).All(&filtered)
}
