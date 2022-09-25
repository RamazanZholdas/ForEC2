package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
	Send id in url parameter Example: localhost/updateSong?id=1
	Json format:
	{
		"name":"example",
		"duration":"example",
		"genre":"example",
		"artist":"example"
	}
*/
func UpdateSong(c *gin.Context) {
	id := c.Param("id")

	var song structs.Song
	var updatedSong structs.SimpleSong

	query := databaseConn.Db.First(&song, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	} else {
		err := json.NewDecoder(c.Request.Body).Decode(&updatedSong)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Json": errors.New("cant decode")})
			return
		}

		song.Name = updatedSong.Name
		song.Artist = updatedSong.Artist
		song.Genre = updatedSong.Genre
		song.Duration = updatedSong.Duration
		databaseConn.Db.Save(&song)

		json.NewEncoder(c.Writer).Encode(updatedSong)
	}
}
