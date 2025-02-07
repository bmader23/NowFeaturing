package mrequest

type GetFeatureFlagRequest struct {
	ApplicationId string `validate:"required"`
	FlagName      string `validate:"required"`
}
