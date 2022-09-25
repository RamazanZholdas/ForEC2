package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/structs"
	"github.com/gin-gonic/gin"
)

/*
	Json format:
	{
		"name":"example",
		"duration":"example",
		"genre":"example",
		"artist":"example"
	}
*/
func CreateSong(c *gin.Context) {
	var song structs.Song

	err := json.NewDecoder(c.Request.Body).Decode(&song)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Json": errors.New("cant decode")})
		return
	}

	databaseConn.Db.Select("Name", "Duration", "Genre", "Artist").Create(&song)
	json.NewEncoder(c.Writer).Encode(song)
}
