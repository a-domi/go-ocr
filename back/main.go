package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	e.POST("/", analysisMedia)
	e.POST("/movie", analysisMovie)
	e.Logger.Fatal(e.Start(":8080"))
}

func analysisMedia(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return nil
	}
	src, err := file.Open()
	if err != nil {
		return nil
	}
	defer src.Close()

	fileModel := strings.Split(file.Filename, ".")
	fileName := fileModel[0]
	extension := fileModel[1]

	f, err := os.Create(fmt.Sprintf("%s_out.%s", fileName, extension))
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(f, src); err != nil {
		return err
	}
	defer os.Remove(f.Name())

	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("eng", "jpn")
	client.SetImage(f.Name())
	text, _ := client.Text()

	return c.String(http.StatusOK, text)
}

func analysisMovie(c echo.Context) error {
	fmt.Println("movie")
	// videoFile := "99f109b8-a210-45eb-b58f-fe17619104fc.mp4"
	// // outputDir := "./frames"

	// // ビデオキャプチャの作成
	// video, err := gocv.VideoCaptureFile(videoFile)
	// if err != nil {
	// 	fmt.Printf("Error opening video file: %v\n", videoFile)
	// 	return nil
	// }
	// defer video.Close()

	return c.String(http.StatusOK, "movie")

}
