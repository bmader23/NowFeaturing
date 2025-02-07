package mrequest

import "github.com/bmader23/nowfeaturing/model"

type UpdateFeatureFlagRequest struct {
	ApplicationId string
	Flags         model.FeatureFlag
}
