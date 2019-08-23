package controllers

import (
	"api/models"
	"api/utils"
	"database/sql"
	"log"
	"net/http"
)

func (c Controller) GetGroups(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var group models.Search
		groups := make([]models.Search, 0)
		var errorObj models.Error
		rows, err := db.Query("select * from chats")
		if err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.Url, &group.Img, &group.Num)
			if err != nil {
				log.Println(err)
				errorObj.Message = "Server error"
				utils.Respond(w, http.StatusInternalServerError, errorObj)
				return
			}
			groups = append(groups, group)
		}
		if err := rows.Err(); err != nil {
			log.Println(err)
			errorObj.Message = "Server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		utils.Respond(w, http.StatusOK, groups)
	}
}
