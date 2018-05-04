package appdrest

import (
	"fmt"
)

// allApplicationTypes is a wrapper on the json response of GetApplicationAllTypes
type applicationAllTypes struct {
	Applications []*Application `json:"apmApplications"`
}

// Application represents a single Business Application within AppDynamics
// Note that the REST version only has ID, Name and Description
type Application struct {
	ID                    int           `json:"id"`
	Version               int           `json:"version"`
	Name                  string        `json:"name"`
	NameUnique            bool          `json:"nameUnique"`
	BuiltIn               bool          `json:"builtIn"`
	CreatedBy             string        `json:"createdBy"`
	CreatedOn             int64         `json:"createdOn"`
	ModifiedBy            string        `json:"modifiedBy"`
	ModifiedOn            int64         `json:"modifiedOn"`
	Description           string        `json:"description"`
	Template              bool          `json:"template"`
	Active                bool          `json:"active"`
	Running               bool          `json:"running"`
	RunningSince          interface{}   `json:"runningSince"`
	DeployWorkflowID      int           `json:"deployWorkflowId"`
	UndeployWorkflowID    int           `json:"undeployWorkflowId"`
	Visualization         interface{}   `json:"visualization"`
	EnvironmentProperties []interface{} `json:"environmentProperties"`
	EumAppName            string        `json:"eumAppName"`
	ApplicationTypeInfo   struct {
		ApplicationTypes   []string `json:"applicationTypes"`
		EumEnabled         bool     `json:"eumEnabled"`
		EumWebEnabled      bool     `json:"eumWebEnabled"`
		EumMobileEnabled   bool     `json:"eumMobileEnabled"`
		EumIotEnabled      bool     `json:"eumIotEnabled"`
		HasEumWebEntities  bool     `json:"hasEumWebEntities"`
		HasMobileApps      bool     `json:"hasMobileApps"`
		HasTiers           bool     `json:"hasTiers"`
		NumberOfMobileApps int      `json:"numberOfMobileApps"`
	} `json:"applicationTypeInfo"`
}

// ApplicationService intermediates Application requests
type ApplicationService service

// GetApplications obtains all applications from a controller
func (s *ApplicationService) GetApplications() ([]*Application, error) {

	url := "controller/rest/applications?output=json"

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	err = s.client.Do(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetApplication gets an Application by Name or ID
func (s *ApplicationService) GetApplication(appNameOrID string) (*Application, error) {

	url := fmt.Sprintf("controller/rest/applications/%s?output=json", appNameOrID)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	err = s.client.Do(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps[0], nil
}

// GetApplicationsAllTypes is a RESTUI call.
// It might break in future versions of AppDynamics
func (s *ApplicationService) GetApplicationsAllTypes() ([]*Application, error) {

	url := fmt.Sprintf("controller/restui/applicationManagerUiBean/getApplicationsAllTypes")

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var apps applicationAllTypes
	err = s.client.DoRestUI(req, &apps)
	if err != nil {
		return nil, err
	}

	return apps.Applications, nil

}
