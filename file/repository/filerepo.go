package filerepo

import (
	"context"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func UploadFile() UploadFileFunc {
	return func(ctx context.Context, pathfile string, file multipart.File) error {
		dst, err := os.Create(pathfile)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = io.Copy(dst, file)
		if err != nil {
			log.Println(err)
		}
		return err
	}
}
