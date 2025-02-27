package merchant_campaign

import (
	"fmt"
	"github.com/vietpham1023/golang-uit-hackathon/config"
	"github.com/vietpham1023/golang-uit-hackathon/internal/models"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
)

type IMerchantCampaign interface {
	ListByFilter(data map[string]string) ([]*models.MerchantCampaignResponse, error)
}

type MerchantCampaign struct {
	cfg  *config.AppConfig
	repo repos.IRepo
}

func NewMerchantCampaign(appConfig *config.AppConfig, repo repos.IRepo) IMerchantCampaign {
	return &MerchantCampaign{
		cfg:  appConfig,
		repo: repo,
	}
}

func (m *MerchantCampaign) ListByFilter(data map[string]string) ([]*models.MerchantCampaignResponse, error) {
	page, size, filter := getFilterList(data)
	records, err := m.repo.MerchantCampaign().ListMerchantCampaignRepo(page, size, filter)

	if err != nil {
		fmt.Printf("list book fail with err: %v\n", err)
		return nil, err
	}
	return records, nil
}

//func (m *MerchantCampaignSQLRepo) Create(ctx context.Context, record *handlerModels.MerchantCampaign) error {
//	err := m.db.Create(record).Error
//	if err != nil {
//		return err
//	}
//	return nil
//}
