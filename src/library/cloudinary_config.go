package library

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func CloudUpload(src string) (string, error) {
	cld_name := os.Getenv("CLOUD_NAME")
	api_key := os.Getenv("CLOUD_APIKEY")
	api_secret := os.Getenv("CLOUD_SECRET")

	ctx := context.Background()
	imageId := uuid.New()

	cld, err := cloudinary.NewFromParams(cld_name, api_key, api_secret)
	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{PublicID: imageId.String(), Folder: "vehicle_rental"})
	if err != nil {
		return "", err
	}

	if err := os.Remove(src); err != nil {
		log.Println(err.Error())
	}

	return resp.SecureURL, nil
}