package model

import "context"

type DbConnector interface {
	Insert(ctx *context.Context, model *ShortUrl) error
	FindOne(ctx *context.Context, hashID string) (*ShortUrl, error)
}

type ShortUrlReq struct {
	Url string `json:"url" validate:"required"`
}

type ShortUrl struct {
	Url  string `json:"url" bson:"Url" validate:"required"`
	Hash string `json:"hash" bson:"Hash" validate:"required"`
}
