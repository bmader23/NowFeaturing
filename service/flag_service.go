package service

import (
	"fmt"

	mrequest "github.com/bmader23/nowfeaturing/model/request"
	mresponse "github.com/bmader23/nowfeaturing/model/response"
	"github.com/bmader23/nowfeaturing/repository"
)

type IService interface {
	GetFeatureFlags(*mrequest.GetFeatureFlagsRequest) (*mresponse.GetFeatureFlagsResponse, error)
	GetFeatureFlag(*mrequest.GetFeatureFlagRequest) (*mresponse.GetFeatureFlagResponse, error)
	UpdateFeatureFlags(*mrequest.UpdateFeatureFlagsRequest) (*mresponse.UpdateFeatureFlagsResponse, error)
	UpdateFeatureFlag(*mrequest.UpdateFeatureFlagRequest) (*mresponse.UpdateFeatureFlagResponse, error)
}

type FlagService struct {
	Rep repository.FlagRepository
}

func (fs FlagService) GetFeatureFlags(req *mrequest.GetFeatureFlagsRequest) (*mresponse.GetFeatureFlagsResponse, error) {
	flags, err := fs.Rep.ReadFlags(req.ApplicationId)
	if err != nil {
		return nil, err
	}
	r := &mresponse.GetFeatureFlagsResponse{
		Flags: flags,
	}
	return r, nil
}

func (fs FlagService) GetFeatureFlag(req *mrequest.GetFeatureFlagRequest) (*mresponse.GetFeatureFlagResponse, error) {
	flag, err := fs.Rep.ReadFlag(req.ApplicationId, req.FlagName)
	if err != nil {
		return nil, err
	}
	r := &mresponse.GetFeatureFlagResponse{
		Flag: *flag,
	}
	return r, nil
}

func (fs FlagService) UpdateFeatureFlags(req *mrequest.UpdateFeatureFlagsRequest) (*mresponse.UpdateFeatureFlagsResponse, error) {

	r, err := fs.Rep.UpdateFeatureFlags(req.ApplicationId, req.Flags)
	if err != nil {
		return nil, fmt.Errorf("failed to create or update flags: %w", err)
	}

	return &mresponse.UpdateFeatureFlagsResponse{
		Success: r,
	}, nil
}

func (fs FlagService) UpdateFeatureFlag(req *mrequest.UpdateFeatureFlagRequest) (*mresponse.UpdateFeatureFlagResponse, error) {

	r, err := fs.Rep.UpdateFeatureFlag(req.ApplicationId, req.Flags)
	if err != nil {
		return nil, err
	}
	return &mresponse.UpdateFeatureFlagResponse{
		Success: r,
	}, nil
}
