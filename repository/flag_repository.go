package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bmader23/nowfeaturing/model"
)

type FlagRepository interface {
	ReadFlags(string) ([]model.FeatureFlag, error)
	ReadFlag(string, string) (*model.FeatureFlag, error)
	UpdateFeatureFlags(string, []model.FeatureFlag) (bool, error)
	UpdateFeatureFlag(string, model.FeatureFlag) (bool, error)
	DeleteFeatureFlag(applicationId string, key string) error
}

type FileFlagRepository struct {
	PathRoot string
}

func (ffr FileFlagRepository) ReadFlags(applicationId string) ([]model.FeatureFlag, error) {
	filepath := fmt.Sprintf("%s/%s.json", ffr.PathRoot, applicationId)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", filepath)
	}
	var r []model.FeatureFlag
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (ffr FileFlagRepository) ReadFlag(applicationId string, flagKey string) (*model.FeatureFlag, error) {
	filepath := fmt.Sprintf("%s/%s.json", ffr.PathRoot, applicationId)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", filepath)
	}
	var r []model.FeatureFlag
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	for _, ele := range r {
		if ele.Key == flagKey {
			return &ele, nil
		}
	}

	return nil, fmt.Errorf("failed to find flag with name %s", flagKey)
}

func (ffr FileFlagRepository) UpdateFeatureFlags(applicationId string, flags []model.FeatureFlag) (bool, error) {
	filepath := fmt.Sprintf("%s/%s.json", ffr.PathRoot, applicationId)
	b, err := json.Marshal(flags)
	if err != nil {
		return false, err
	}

	err = os.WriteFile(filepath, b, 0333)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ffr FileFlagRepository) UpdateFeatureFlag(applicationId string, flag model.FeatureFlag) (bool, error) {
	filepath := fmt.Sprintf("%s/%s.json", ffr.PathRoot, applicationId)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return false, fmt.Errorf("failed to open file: %s", filepath)
	}
	var r []model.FeatureFlag
	err = json.Unmarshal(data, &r)
	if err != nil {
		return false, err
	}

	for i, ele := range r {
		if ele.Key == flag.Key {
			r[i].Value = flag.Value
		}
	}

	b, err := json.Marshal(r)
	if err != nil {
		return false, err
	}

	err = os.WriteFile(filepath, b, 0333)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ffr FileFlagRepository) DeleteFeatureFlag(applicationId string, key string) error {
	filepath := fmt.Sprintf("%s/%s.json", ffr.PathRoot, applicationId)

	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	var r []model.FeatureFlag
	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	for i, ele := range r {
		if ele.Key == key {
			r = append(r[:i], r[i+1:]...)
		}
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, b, 0333)
	if err != nil {
		return err
	}
	return nil
}

type DbFlagRepository struct {
	ConnectionString string
}

func (dfr DbFlagRepository) GetFlags(string) ([]model.FeatureFlag, error) {
	return []model.FeatureFlag{}, nil
}

func (dfr DbFlagRepository) CreateUpdateFlags(applicationId string, flags []model.FeatureFlag) error {
	return nil
}
