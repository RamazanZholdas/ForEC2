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

//Send id in url parameter Example: localhost/updateSong?id=1
func DeleteSong(c *gin.Context) {
	id := c.Param("id")

	var song structs.Song

	query := databaseConn.Db.First(&song, id)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"id does not exist": errors.New("song does not exist")})
	} else {
		databaseConn.Db.Delete(&song)

		json.NewEncoder(c.Writer).Encode("song deleted")
	}
}
