package service

import (
	"e-commerce/config"
	"e-commerce/internal/app/storage"
)

type Service struct {
	Storage storage.Storage
	Config  config.Config
}
