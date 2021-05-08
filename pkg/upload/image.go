package upload

import (
	"fmt"
	"mime/multipart"
	"nicetry/global"
	"nicetry/pkg/files"
	"nicetry/pkg/utils"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return global.AppSetting.ImagePrefixUrl + name
}

func GetImageName(name string) string {

	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.Md5Crypt(fileName, utils.Upper(4))

	return fileName + ext
	//return fileName
}

func CheckImageExt(fileName string) bool {
	ext := files.GetExt(fileName)
	for _, allowExt := range global.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

func CheckImageSize(f *multipart.FileHeader) bool {
	return f.Size < global.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = files.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := files.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
