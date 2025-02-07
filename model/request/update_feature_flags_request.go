package mrequest

import "github.com/bmader23/nowfeaturing/model"

type UpdateFeatureFlagsRequest struct {
	ApplicationId string
	Flags         []model.FeatureFlag
}
