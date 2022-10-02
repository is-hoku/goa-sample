package sample

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func newFirebase(ctx context.Context) (*auth.Client, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type contextKey string

const userInfoContextKey contextKey = "user_info"

func SetUserInfo(ctx context.Context, token *auth.Token) (context.Context, error) {
	userInfo, err := GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, userInfoContextKey, userInfo.UserInfo), nil
}

func GetUserInfo(ctx context.Context) (*auth.UserInfo, error) {
	v := ctx.Value(userInfoContextKey)
	userInfo, ok := v.(*auth.UserInfo)
	if !ok {
		return nil, errors.New("User Info not found")
	}
	return userInfo, nil
}

func GetUserByToken(ctx context.Context, token *auth.Token) (*auth.UserRecord, error) {
	client, err := newFirebase(ctx)
	if err != nil {
		return nil, err
	}
	userRecord, err := client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}
	return userRecord, nil
}
