package library

import (
	"context"
	// "fmt"
	"mime/multipart"
	"os"
	// "strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UploadImage(file multipart.File, handler *multipart.FileHeader) (string, string, error) {
	
	cld_name := os.Getenv("CLOUD_NAME")
	api_key := os.Getenv("CLOUD_APIKEY")
	api_secret := os.Getenv("CLOUD_SECRET")

	ctx := context.Background()
	imageId := uuid.New()

	cld, err := cloudinary.NewFromParams(cld_name, api_key, api_secret)
	if err != nil {
		return "", "", err
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: imageId.String(), Folder: "vehicle_rental"})
	if err != nil {
		return "","", err
	}

	uidImage := resp.PublicID
	// imageName := strings.Replace(uidImage, "vehicle_rental/", "", -1)
	// fmt.Print(imageName)
	return resp.SecureURL, uidImage, nil

}

func DeleteImage(idImage string) (string, error) {

	cld_name := os.Getenv("CLOUD_NAME")
	api_key := os.Getenv("CLOUD_APIKEY")
	api_secret := os.Getenv("CLOUD_SECRET")

	ctx := context.Background()

	cld, err := cloudinary.NewFromParams(cld_name, api_key, api_secret)
	if err != nil {
		return "",  err
	}
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: idImage})
	if err != nil {
		return "", err
	}

	return resp.Result, nil

}
