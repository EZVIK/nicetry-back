package controller

import (
	"github.com/gofiber/fiber/v2"
	"nicetry/global"
	"nicetry/pkg/app"
	"nicetry/pkg/e"
	"nicetry/pkg/upload"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func (s *Controller) UploadImage(c *fiber.Ctx) error {

	code := e.SUCCESS
	data := make(map[string]string)

	image, err := c.FormFile("image")
	if err != nil {
		global.Logger.Infof("c.Request.FormFile: ", err)
		code = e.ERROR
		c.JSON(app.NewErrRes(code, e.GetMsg(code), ""))
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := global.AppSetting.ImageFilePath

		src := fullPath + imageName
		if ! upload.CheckImageExt(imageName) || image.Size > global.AppSetting.ImageMaxSize {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT.Code()
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				global.Logger.Info(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL.Code()
			} else if err := c.SaveFile(image, src); err != nil {
				global.Logger.Info(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL.Code()
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				//data["image_save_url"] = fullPath + imageName
			}
		}
	}

	c.JSON(app.NewRes(data))
	return nil
}
