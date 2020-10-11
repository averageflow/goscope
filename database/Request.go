// License: MIT
// Authors:
// 		- Josep Jesus Bigorra Algaba (@averageflow)
package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/averageflow/goscope/utils"
)

func GetDetailedRequest(db *sql.DB, connection, requestUID string) *sql.Row {
	var query string
	if connection == MySQL || connection == SQLite {
		query = "SELECT `uid`, `client_ip`, `method`, `path`, `url`, " +
			"`host`, `time`, `headers`, `body`, `referrer`, `user_agent` FROM `requests` WHERE `uid` = ? LIMIT 1;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "client_ip", "method", "path", "url", 
			"host", "time", "headers", "body", "referrer", "user_agent" FROM "requests" WHERE "uid" = ? LIMIT 1;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func GetRequests(db *sql.DB, connection string, offset int) *sql.Rows {
	var query string
	if connection == MySQL || connection == SQLite {
		query = "SELECT `requests`.`uid`, `requests`.`method`, `requests`.`path`, `requests`.`time`, " +
			"`responses`.`status` FROM `requests` " +
			"INNER JOIN `responses` ON `requests`.`uid` = `responses`.`request_uid` " +
			"WHERE `requests`.`application` = ? ORDER BY `requests`.`time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "requests"."uid", "requests"."method", "requests"."path", "requests"."time", 
			"responses"."status" FROM "requests" 
			INNER JOIN "responses" ON "requests"."uid" = "responses"."request_uid" 
			WHERE "requests"."application" = ? ORDER BY "requests"."time" DESC LIMIT ? OFFSET ?;`
	}

	rows, err := db.Query(
		query,
		utils.Config.ApplicationID,
		utils.Config.GoScopeEntriesPerPage,
		offset,
	)

	if err != nil {
		log.Println(err.Error())

		return nil
	}

	return rows
}

func SearchRequests(db *sql.DB, connection, searchWildcard string, offset int) *sql.Rows {
	var query string

	if connection == MySQL || connection == SQLite {
		query = "SELECT `requests`.`uid`, `requests`.`method`, `requests`.`path`, `requests`.`time`, " +
			"`responses`.`status` FROM `requests` " +
			"INNER JOIN `responses` ON `requests`.`uid` = `responses`.`request_uid` " +
			"WHERE `requests`.`application` = ? AND " +
			"(`requests`.`uid` LIKE ? OR `requests`.`application` LIKE ? " +
			"OR `requests`.`client_ip` LIKE ? OR `requests`.`method` LIKE ? " +
			"OR `requests`.`path` LIKE ? " +
			"OR `requests`.`url` LIKE ? OR `requests`.`host` LIKE ? " +
			"OR `requests`.`body` LIKE ? OR `requests`.`referrer` LIKE ? " +
			"OR `requests`.`user_agent` LIKE ? OR `requests`.`time` LIKE ? " +
			"OR `responses`.`uid` LIKE ? OR `responses`.`request_uid` LIKE ? " +
			"OR `responses`.`application` LIKE ? OR `responses`.`client_ip` LIKE ? " +
			"OR `responses`.`status` LIKE ? " +
			"OR `responses`.`body` LIKE ? OR `responses`.`path` LIKE ? " +
			"OR `responses`.`headers` LIKE ? OR `responses`.`size` LIKE ? " +
			"OR `responses`.`time` LIKE ?) ORDER BY `requests`.`time` DESC LIMIT ? OFFSET ?;"
	} else if connection == PostgreSQL {
		query = `SELECT "requests"."uid", "requests"."method", "requests"."path", 
			"requests"."time", "responses"."status" FROM "requests" 
			INNER JOIN "responses" ON "requests"."uid" = "responses"."request_uid" 
			WHERE "requests"."application" = ? AND 
			("requests"."uid" LIKE ? OR "requests"."application" LIKE ? 
			OR "requests"."client_ip" LIKE ? OR "requests"."method" LIKE ? 
			OR "requests"."path" LIKE ? 
			OR "requests"."url" LIKE ? OR "requests"."host" LIKE ? 
			OR "requests"."body" LIKE ? OR "requests"."referrer" LIKE ?
			OR "requests"."user_agent" LIKE ? OR "requests"."time" LIKE ? 
			OR "responses"."uid" LIKE ? OR "responses"."request_uid" LIKE ? 
			OR "responses"."application" LIKE ? OR "responses"."client_ip" LIKE ?
			OR "responses"."status" LIKE ? 
			OR "responses"."body" LIKE ? OR "responses"."path" LIKE ? 
			OR "responses"."headers" LIKE ? OR "responses"."size" LIKE ?
			OR "responses"."time" LIKE ?) ORDER BY "requests"."time" DESC LIMIT ? OFFSET ?;`
	}

	rows, err := db.Query(
		query,
		utils.Config.ApplicationID,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		searchWildcard, searchWildcard, searchWildcard, searchWildcard,
		searchWildcard,
		utils.Config.GoScopeEntriesPerPage,
		offset,
	)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return rows
}

func GetDetailedResponse(db *sql.DB, connection, requestUID string) *sql.Row {
	var query string

	if connection == MySQL || connection == SQLite {
		query = "SELECT `uid`, `client_ip`, `status`, `time`, " +
			"`body`, `path`, `headers`, `size` FROM `responses` " +
			"WHERE `request_uid` = ? LIMIT 1;"
	} else if connection == PostgreSQL {
		query = `SELECT "uid", "client_ip", "status", "time",
			"body", "path", "headers", "size" FROM "responses"
			WHERE "request_uid" = ? LIMIT 1;`
	}

	row := db.QueryRow(query, requestUID)

	return row
}

func DumpResponse(c *gin.Context, responsePayload utils.DumpResponsePayload, body string) {
	now := time.Now().Unix()
	requestUID, _ := uuid.NewV4()
	headers, _ := json.Marshal(c.Request.Header)
	query := "INSERT INTO requests (uid, application, client_ip, method, path, host, time, " +
		"headers, body, referrer, url, user_agent) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	requestPath := c.FullPath()
	if requestPath == "" {
		requestPath = c.Request.URL.String()
	}

	_, err := utils.DB.Exec(
		query,
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		c.Request.Method,
		requestPath,
		c.Request.Host,
		now,
		string(headers),
		body,
		c.Request.Referer(),
		c.Request.RequestURI,
		c.Request.UserAgent(),
	)

	if err != nil {
		log.Println(err.Error())
	}

	responseUID, _ := uuid.NewV4()
	headers, _ = json.Marshal(responsePayload.Headers)
	query = "INSERT INTO responses (uid, request_uid, application, client_ip, status, time, " +
		"body, path, headers, size) VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err = utils.DB.Exec(
		query,
		responseUID.String(),
		requestUID.String(),
		utils.Config.ApplicationID,
		c.ClientIP(),
		responsePayload.Status,
		now,
		responsePayload.Body.String(),
		c.FullPath(),
		string(headers),
		responsePayload.Body.Len(),
	)

	if err != nil {
		log.Println(err.Error())
	}
}