package vkoauth

import (
	"context"
)

// Создает URL, на который нужно направить пользователя для проведения авторизации методом Authorization Code Flow
// После проведения авторизации, пользователь перейдет на redirect_uri, куда будет отправлен параметр code
// Вам нужно взять значение этого параметра и использовать его в методе config.ExchangeCode(context.Background(), code)
// Смотрите документацию: https://dev.vk.com/api/access-token/authcode-flow-user
func (v *Config) CodeFlowAuthUrl(params AuthParams, opts ...AuthOption) string {
	return v.buildAuthUrl(params, "code", opts...)
}

// Получает токен доступа на основе результата авторизации Authorization Code Flow
// code - параметр, полученный сервером при редиректе пользователя
func (v *Config) ExchangeCode(ctx context.Context, code string, opts ...AuthOption) (*Token, error) {
	exchangeOptions := []AuthOption{
		setParam{"code", code},
		setParam{"redirect_uri", v.RedirectUri},
	}

	exchangeOptions = append(exchangeOptions, opts...)
	return v.doTokenRequest(ctx, v.buildTokenUrl(v.endpoint().TokenUrl,
		exchangeOptions...,
	))
}
