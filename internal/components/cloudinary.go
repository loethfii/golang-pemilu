package components

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"luthfi/pemilu/internal/config"
)

type CloudinaryUploadReq struct {
	CTX       context.Context
	ImageFile any
	FileName  string
}

func CloudinaryUploadConfig() *cloudinary.Cloudinary {
	cnf := config.Get()
	cld, _ := cloudinary.NewFromParams(cnf.Cloudinary.CloudName, cnf.Cloudinary.ApiKey, cnf.Cloudinary.ApiSecret)
	return cld
}

func CloudinaryUploadImage(cld *cloudinary.Cloudinary, cldReq CloudinaryUploadReq) (string, error) {
	resp, err := cld.Upload.Upload(cldReq.CTX, cldReq.ImageFile, uploader.UploadParams{
		PublicID: cldReq.FileName,
	})
	if err != nil {
		return "", err
	}
	
	return resp.SecureURL, nil
	
}
