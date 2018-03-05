package utils

import (
	"log"

	"gopkg.in/robfig/cron.v2"
	"pershotravndk.com/models"
)

func TokensMonitor() {
	c := cron.New()
	c.AddFunc("@daily", DeleteToken)
	c.Start()
}

func DeleteToken() {
	db, err := models.NewDB("postgres://jsopcnfzumgznz:20807490dae09e58991e7a56179e659d80a9169bafbba8b01bb996464fed4347@ec2-107-22-175-33.compute-1.amazonaws.com:5432/db56m8m3hnlru2")
	if err != nil {
		log.Panic(err)
	}
	rows, err := db.Query("Select user_id from confirmation_token WHERE date_expires = CURRENT_DATE and date_used = null;")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		var userID string
		err = rows.Scan(&userID)
		if err != nil {
			log.Print(err)
		}
		err = models.DeleteToken(userID, db)
		if err != nil {
			log.Print(err)
		}
		err = models.DeleteProfile(userID, db)
		if err != nil {
			log.Print(err)
		}
		err = models.DeleteUser(userID, db)
		if err != nil {
			log.Print(err)
		}
	}
}
