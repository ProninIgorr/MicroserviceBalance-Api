package service

import "github.com/ProninIgorr/MicroserviceBalance-Api/pkg/repository"

type PurchaseService struct {
	repo repository.Purchase
}

func NewPurchaseService(repo repository.Purchase) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (s *PurchaseService) UpdatePurchase(id int, srv string) error {
	return s.repo.UpdatePurchase(id, srv)
}
