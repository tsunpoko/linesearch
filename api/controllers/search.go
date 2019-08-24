package controllers

import (
	"api/models"
	"api/utils"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (c Controller) AddGroup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var group models.Search
		var errorObj models.Error

		ur := r.FormValue("url")

		u, err := url.Parse(ur)
		/*
			fmt.Println(u)
			fmt.Printf("URL: %s\n", u.String())
			fmt.Printf("Scheme: %s\n", u.Scheme)
			fmt.Printf("Opaque: %s\n", u.Opaque)
			fmt.Printf("User: %s\n", u.User)
			fmt.Printf("Host: %s\n", u.Host)
			fmt.Printf("Hostname(): %s\n", u.Hostname())
			fmt.Printf("Path: %s\n", u.Path)
			fmt.Printf("RawPath: %s\n", u.RawPath)
			fmt.Printf("RawQuery: %s\n", u.RawQuery)
			fmt.Printf("Fragment: %s\n", u.Fragment)
		*/
		if u.Host != "line.me" {
			errorObj.Message = "Please input http://line.me/~~~"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}

		resp, _ := http.Get(ur)
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			errorObj.Message = "Please input correct url."
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			errorObj.Message = "Please input correct url."
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}
		bodyString := string(bodyBytes)
		//log.Println(bodyString)

		slice := strings.Split(bodyString, "\n")
		base_num := 0

		for i, str := range slice {
			//fmt.Printf("%d: %s\n", i, str)
			if strings.Contains(strings.ToLower(str), "<!--contents-area-->") {
				base_num = i
			}
		}

		if !strings.HasPrefix(slice[base_num+3], `  <div class="`) {
			errorObj.Message = "server error"
			utils.Respond(w, http.StatusInternalServerError, errorObj)
			return
		}

		img_row := slice[base_num+5]
		fmt.Println("img ok")
		name_row := slice[base_num+17]
		fmt.Println("name ok")
		num_row := slice[base_num+18]
		fmt.Println("num ok")
		description_row := slice[base_num+19]
		fmt.Println("desc ok")
		url_row := slice[base_num+20]
		if strings.HasSuffix(description_row, "</p>") != true {
			for i := 0; i < 100; i++ {
				s := slice[base_num+19+i]
				if strings.HasSuffix(s, "</p>") {
					url_row = slice[base_num+20+i]
					fmt.Println("url ok")
					break
				}
			}
		}

		/*
			fmt.Println(name_row)
			fmt.Println(num_row)
			fmt.Println(description_row)
			fmt.Println(img_row)
			fmt.Println(url_row)
		*/
		name := name_row[strings.Index(name_row, `">`)+2 : strings.Index(name_row, `</p>`)]
		num, _ := strconv.Atoi(num_row[strings.Index(num_row, `">`)+10 : strings.Index(num_row, `</p>`)])
		description := strings.Replace(description_row[strings.Index(description_row, `">`)+2:], `</p>`, ``, -1)
		url := url_row[strings.Index(url_row, `href="`)+6 : strings.Index(url_row, `class=`)-2]
		img := img_row[strings.Index(img_row, `c="`)+3 : strings.Index(img_row, `" alt`)]

		description = strings.Split(description, "<")[0]

		fmt.Println(name)
		fmt.Println(num)
		fmt.Println(description)
		fmt.Println(url)
		fmt.Println("!!!")
		fmt.Println(img)

		//byteArray, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(byteArray)

		var tmp string
		err = db.QueryRow("select name from chats where url = $1", url).Scan(&tmp)
		fmt.Println(tmp)

		group.Name = name
		group.Description = description
		group.Url = url
		group.Img = img + `.jpg`
		group.Num = num
		if err == sql.ErrNoRows {
			err = db.QueryRow("($1, $2, $3, $4, $5) RETURNING id;", group.Name, group.Description, group.Url, group.Img, group.Num).Scan(&group.ID)
			if err != nil {
				log.Println(err)
				errorObj.Message = "Server error"
				utils.Respond(w, http.StatusInternalServerError, errorObj)
				return
			}

			utils.Respond(w, http.StatusOK, group)
			return
		}
		//err = db.QueryRow("insert into students (name, belong, age) values ($1, $2, $3) RETURNING id;", student.Name, student.Belong, student.Age).Scan(&student.ID)

		log.Println(err)
		errorObj.Message = "Server error"
		utils.Respond(w, http.StatusInternalServerError, errorObj)
		return
	}
}
