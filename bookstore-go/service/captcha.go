package service

import (
	"bookstore/global"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mojocn/base64Captcha"
)

type CaptchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

type CaptchaResponse struct {
	CaptchaID     string `json:"captcha_id"`
	CaptchaBase64 string `json:"captcha_base64"`
}

func (c *CaptchaService) GenerateCaptcha() (*CaptchaResponse, error) {
	driver := base64Captcha.NewDriverDigit(
		80,  // 高
		240, // 宽
		4,   // 验证码长度
		0.7, // 干扰强度
		80,  // 干扰数量
	)
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	// 用redis存储有效期的图片验证码
	log.Println("图片验证码真实 answer:", answer)
	redisKey := fmt.Sprintf("captcha:%s", id)
	err = global.RedisClient.Set(context.TODO(), redisKey, answer, time.Minute*1).Err()
	if err != nil {
		return nil, err
	}
	return &CaptchaResponse{
		CaptchaID:     id,
		CaptchaBase64: b64s,
	}, nil
}

func (c *CaptchaService) VerifyCaptcha(id, answer string) bool {
	if id == "" || answer == "" {
		return false
	}
	ctx := context.Background()
	captchaID := fmt.Sprintf("captcha:%s", id)
	captchaAnswer, err := global.RedisClient.Get(ctx, captchaID).Result()
	if err != nil {
		return false
	}

	isValid := captchaAnswer == answer
	if isValid {
		global.RedisClient.Del(ctx, captchaID)
	}
	return isValid
}
