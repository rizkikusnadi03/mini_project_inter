package upload

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveImage(c *fiber.Ctx, fieldName string, destFolder string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		if err.Error() == "there is no uploaded file associated with the given key" {
			return "", nil // Optional file
		}
		return "", err
	}

	// Create unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dir := fmt.Sprintf("./uploads/%s", destFolder)
	
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}
	
	path := fmt.Sprintf("%s/%s", dir, filename)

	if err := c.SaveFile(file, path); err != nil {
		return "", err
	}

	return path, nil
}
