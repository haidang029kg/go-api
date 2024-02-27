package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"goapi/src/settings"
	"goapi/src/utils"

	"github.com/gin-gonic/gin"
)

func uploadHandler(c *gin.Context) {
	log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL)
	file, err := c.FormFile("myFile")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting file"})
		return
	}

	// Create the file
	dst, err := os.Create(filepath.Join(settings.SettingsIns.UploadPath, file.Filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating file"})
		return
	}
	defer dst.Close()

	// Write the content from POST to the file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error opening file"})
		return
	}
	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing file"})
		return
	}

	// split to get the file name
	filename := strings.Split(file.Filename, ".")[0]
	// generate a directory with the file name in the hls_converted folder
	os.Mkdir(filepath.Join(settings.SettingsIns.HlsPath, filename), 0o755)
	// generate the output file name with .m3u8 extension
	output := filename + ".m3u8"

	configQualities := []string{"720", "1080", "1440", "2160"}
	go utils.ConvertToHLS(filepath.Join(settings.SettingsIns.UploadPath, file.Filename), filepath.Join(settings.SettingsIns.HlsPath, filename, output), configQualities)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Uploaded File"})
}
