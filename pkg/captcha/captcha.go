package captcha

import (
	"github.com/mojocn/base64Captcha"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"sync"
)

type Capthcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Capthcha

func NewCaptcha() *Capthcha {
	once.Do(func() {
		internalCaptcha = &Capthcha{}

		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha",
		}

		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),      // 宽
			config.GetInt("captcha.width"),       // 高
			config.GetInt("captcha.length"),      // 长度
			config.GetFloat64("captcha.maxskew"), // 数字的最大倾斜角度
			config.GetInt("captcha.dotcount"),    // 图片背景里的混淆点数量
		)

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

func (c *Capthcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Capthcha) VerifyCaptcha(id string, answer string) (match bool) {

	return c.Base64Captcha.Verify(id, answer, false)
}
