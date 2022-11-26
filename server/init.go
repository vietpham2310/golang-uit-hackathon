package server

import (
	hdl "github.com/vietpham1023/golang-uit-hackathon/handler"
	"github.com/vietpham1023/golang-uit-hackathon/internal/repos"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/book"
	campaign2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/campaign"
	"github.com/vietpham1023/golang-uit-hackathon/internal/services/merchant_campaign"
	user2 "github.com/vietpham1023/golang-uit-hackathon/internal/services/user"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {
	book := book.NewBook(s.cfg, repo)
	user := user2.NewUser(s.cfg, repo)
	merchantCampaign := merchant_campaign.NewMerchantCampaign(s.cfg, repo)
	campaign := campaign2.NewCampaign(s.cfg, repo)
	return &ServiceList{
		book:             book,
		user:             user,
		merchantCampaign: merchantCampaign,
		campaign:         campaign,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {

	handler := hdl.NewHandler(
		s.cfg,
		s.router,
		serviceList.book,
		serviceList.user,
		serviceList.merchantCampaign,
		serviceList.campaign,
	)

	handler.ConfigureAPIRoute(s.router)

}
