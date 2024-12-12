package router

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
)

func NewRouter() error {

	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	return err
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 10)

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	img, format, err := image.Decode(file)

	img2 := imaging.Blur(img, 4)
	img2 = imaging.Invert(img2)

	outputFile, err := os.Create("/" + format)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if format == "jpeg" {
		jpeg.Encode(outputFile, img, nil)
	} else if format == "png" {
		png.Encode(outputFile, img2)
	} else {
		http.Error(w, "Unsupported format", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Image uploaded and resized successfully!"))

}
