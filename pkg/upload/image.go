package upload

import (
	"fmt"
	"golang.org/x/xerrors"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"go-blog-step-by-step/pkg/file"
	"go-blog-step-by-step/pkg/logging"
	"go-blog-step-by-step/pkg/setting"
	"go-blog-step-by-step/pkg/util"
)

func Setup() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd err: %v", err)
	}

	filePath := GetImageFullPath()
	src := path.Join(dir, filePath)
	log.Println(src)

	perm := file.CheckPermission(src)
	if perm == true {
		log.Fatalf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)

	if err != nil {
		log.Fatal(err)
	}
}

func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) (string, error) {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName, err := util.EncodeMD5(fileName)

	if err != nil {
		return "", xerrors.Errorf("%v", err)
	}

	return fileName + ext, nil
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return xerrors.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return xerrors.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
