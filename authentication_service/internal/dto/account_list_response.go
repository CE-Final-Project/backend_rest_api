package dto

import accountService "github.com/ce-final-project/backend_rest_api/account_service/proto/account"

type AccountsListResponse struct {
	TotalCount int64              `json:"totalCount" bson:"totalCount"`
	TotalPages int64              `json:"totalPages" bson:"totalPages"`
	Page       int64              `json:"page" bson:"page"`
	Size       int64              `json:"size" bson:"size"`
	HasMore    bool               `json:"hasMore" bson:"hasMore"`
	Accounts   []*AccountResponse `json:"accounts" bson:"accounts"`
}

func AccountsListResponseFromGrpc(listResponse *accountService.SearchRes) *AccountsListResponse {
	list := make([]*AccountResponse, 0, len(listResponse.GetAccounts()))
	for _, account := range listResponse.GetAccounts() {
		list = append(list, AccountResponseFromGrpc(account))
	}

	return &AccountsListResponse{
		TotalCount: listResponse.GetTotalCount(),
		TotalPages: listResponse.GetTotalPages(),
		Page:       listResponse.GetPage(),
		Size:       listResponse.GetSize(),
		HasMore:    listResponse.GetHasMore(),
		Accounts:   list,
	}
}
