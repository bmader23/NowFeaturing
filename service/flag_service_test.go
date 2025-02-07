package service

import (
	"testing"

	"github.com/bmader23/nowfeaturing/model"
	mrequest "github.com/bmader23/nowfeaturing/model/request"
)

type TestFlagRepository struct{}

func (tfr TestFlagRepository) ReadFlags(string) ([]model.FeatureFlag, error) {
	return []model.FeatureFlag{}, nil
}

func (tfr TestFlagRepository) ReadFlag(string, string) (*model.FeatureFlag, error) {
	return &model.FeatureFlag{}, nil
}

func (tfr TestFlagRepository) UpdateFeatureFlags(string, []model.FeatureFlag) (bool, error) {
	return true, nil
}

func (tfr TestFlagRepository) UpdateFeatureFlag(string, model.FeatureFlag) (bool, error) {
	return true, nil
}

func (tfr TestFlagRepository) DeleteFeatureFlag(string, string) error {
	return nil
}

func TestFlagService(t *testing.T) {
	service := FlagService{
		Rep: TestFlagRepository{},
	}
	t.Run("Valid feature flag file", func(t *testing.T) {
		service.GetFeatureFlags(&mrequest.GetFeatureFlagsRequest{
			ApplicationId: "TestApplicationId",
		})
	})
}
