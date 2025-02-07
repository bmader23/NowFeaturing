package handler

import (
	"encoding/json"
	"fmt"
	"io"

	mrequest "github.com/bmader23/nowfeaturing/model/request"
	"github.com/bmader23/nowfeaturing/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type FlagHandler struct {
	FlagService service.IService
	validate    *validator.Validate
}

func NewFlagHandler(gingine *gin.Engine, flagService service.IService) *FlagHandler {
	r := &FlagHandler{
		FlagService: flagService,
		validate:    validator.New(),
	}
	r.registerRoutes(gingine)
	return r
}

func (fh FlagHandler) registerRoutes(gingine *gin.Engine) {
	gingine.GET("flags", fh.GetApplicationFlags)
	gingine.GET("flag", fh.GetApplicationFlag)

	gingine.POST("flags", fh.PostFeatureFlags)
	gingine.POST("flag", fh.PostFeatureFlag)

}

func (fh FlagHandler) GetApplicationFlags(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to read request body: %w", err).Error())
		return
	}
	var req mrequest.GetFeatureFlagsRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to unmarshal byte stream to json object: %w", err).Error())
		return
	}

	err = fh.validate.Struct(req)
	if err != nil {
		ctx.String(500, fmt.Errorf("failed validation of request: %w", err).Error())
		return
	}

	resp, err := fh.FlagService.GetFeatureFlags(&req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}

func (fh FlagHandler) GetApplicationFlag(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to read request body: %w", err).Error())
		return
	}
	var req mrequest.GetFeatureFlagRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to unmarshal byte stream to json object: %w", err).Error())
		return
	}

	err = fh.validate.Struct(req)
	if err != nil {
		ctx.String(500, fmt.Errorf("failed validation of request: %w", err).Error())
		return
	}

	resp, err := fh.FlagService.GetFeatureFlag(&req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp.Flag)
}

func (fh FlagHandler) PostFeatureFlag(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to read request body: %w", err).Error())
		return
	}
	var req mrequest.UpdateFeatureFlagRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to unmarshal byte stream to json object: %w", err).Error())
		return
	}

	err = fh.validate.Struct(req)
	if err != nil {
		ctx.String(500, fmt.Errorf("failed validation of request: %w", err).Error())
		return
	}

	resp, err := fh.FlagService.UpdateFeatureFlag(&req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}

func (fh FlagHandler) PostFeatureFlags(ctx *gin.Context) {
	jsonData, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to read request body: %w", err).Error())
		return
	}
	var req mrequest.UpdateFeatureFlagRequest
	err = json.Unmarshal(jsonData, &req)
	if err != nil {
		ctx.JSON(500, fmt.Errorf("failed to unmarshal byte stream to json object: %w", err).Error())
		return
	}

	err = fh.validate.Struct(req)
	if err != nil {
		ctx.String(500, fmt.Errorf("failed validation of request: %w", err).Error())
		return
	}

	resp, err := fh.FlagService.UpdateFeatureFlag(&req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}
