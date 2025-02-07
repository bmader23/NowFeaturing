package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/bmader23/nowfeaturing/model"
)

func TestFlagRepository_FileFlagRepository_ReadFlags(t *testing.T) {
	repository := FileFlagRepository{
		PathRoot: "../testdata",
	}
	t.Run("Valid feature flag file", func(t *testing.T) {
		applicationId := "feature_flag_test"
		flags, err := repository.ReadFlags(applicationId)
		if err != nil {
			fmt.Printf("No error expected. Received: %v", err)
			t.FailNow()
		}
		if len(flags) < 2 {
			fmt.Printf("Fewer flags received (%d) than expected.", len(flags))
			t.FailNow()
		}
	})
	t.Run("Invalid application id", func(t *testing.T) {
		_, err := repository.ReadFlags("invalid_app_id")
		if err == nil {
			fmt.Print("Error expected but none received.")
			t.FailNow()
		}
	})
}
func TestFlagRepository_FileFlagRepository_UpdateFeatureFlags(t *testing.T) {
	repository := FileFlagRepository{
		PathRoot: "../testdata",
	}
	t.Run("Create new flags file", func(t *testing.T) {
		applicationId := "new_application_id"
		r, err := repository.UpdateFeatureFlags(applicationId, []model.FeatureFlag{
			{
				Key:   "TestFlag",
				Value: "Test1",
			}, {
				Key:   "TestFlag2",
				Value: "Test2",
			},
		})
		if err != nil {
			fmt.Printf("No error expected. Received: %v", err)
			t.FailNow()
		}
		if !r {
			fmt.Printf("CreateUpdateFlags returned false: %v", err)
			t.FailNow()
		}
	})
}

func TestFlagRepository_FileFlagRepository_UpdateFeatureFlag(t *testing.T) {
	repository := FileFlagRepository{
		PathRoot: "../testdata",
	}

	initialFile := fmt.Sprint(repository.PathRoot, "/_starting_state.json")
	newFile := fmt.Sprint(repository.PathRoot, "/update_feature_flag_test.json")
	b, err := os.ReadFile(initialFile)
	if err != nil {
		t.FailNow()
	}
	err = os.WriteFile(newFile, b, 0333)
	if err != nil {
		t.FailNow()
	}

	t.Run("Create new flags file", func(t *testing.T) {
		applicationId := "update_feature_flag_test"
		r, err := repository.UpdateFeatureFlag(applicationId, model.FeatureFlag{
			Key:   "FirstFeatureFlag",
			Value: "VAL1",
		})
		if err != nil {
			fmt.Printf("No error expected. Received: %v", err)
			t.FailNow()
		}
		if !r {
			fmt.Printf("Returned false response: %v", b)
			t.FailNow()
		}
		b, err := os.ReadFile(newFile)
		if err != nil {
			t.FailNow()
		}
		var result []model.FeatureFlag
		err = json.Unmarshal(b, &result)
		if err != nil {
			t.FailNow()
		}
		for _, ff := range result {
			if ff.Key == "FirstFeatureFlag" && ff.Value != "VAL1" {
				t.FailNow()
			}
		}
	})
}

func TestFlagRepository_FileFlagRepository_DeleteFeatureFlag(t *testing.T) {
	repository := FileFlagRepository{
		PathRoot: "../testdata",
	}

	initialFile := fmt.Sprint(repository.PathRoot, "/_starting_state.json")
	newFile := fmt.Sprint(repository.PathRoot, "/delete_feature_flag_test.json")

	b, err := os.ReadFile(initialFile)
	if err != nil {
		t.FailNow()
	}
	err = os.WriteFile(newFile, b, 0333)
	if err != nil {
		t.FailNow()
	}

	t.Run("Delete flag from file", func(t *testing.T) {
		applicationId := "delete_feature_flag_test"
		err := repository.DeleteFeatureFlag(applicationId, "FirstFeatureFlag")
		if err != nil {
			fmt.Printf("No error expected. Received: %v", err)
			t.FailNow()
		}
		b, err := os.ReadFile(newFile)
		if err != nil {
			t.FailNow()
		}
		var result []model.FeatureFlag
		err = json.Unmarshal(b, &result)
		if err != nil {
			t.FailNow()
		}
		for _, ff := range result {
			if ff.Key == "FirstFeatureFlag" {
				t.FailNow()
			}
		}
	})
}
