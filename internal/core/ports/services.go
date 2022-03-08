package ports

import (
	"github.com/ashish041/html-parser/internal/core/domain"
)

type DomainService interface {
	New(id string) (*domain.Url, error)
	Get(response *domain.Url) (*domain.Informations, error)
}
