package filerepo

import (
	"context"
	"mime/multipart"
)

type UploadFileFunc func(ctx context.Context, path string, file multipart.File) error
