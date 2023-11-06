package auth

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"net/http"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, error := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(error)

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
