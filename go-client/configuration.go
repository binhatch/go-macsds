/*
Hella Gutmann - macsDS (Data Services) - API collection

This document contains all relevant APIs for diagnostics (incl. DTCs), repair & maintenance information (RMI) and vehicle identification.

API version: V20240702-130718
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gomacds

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
			"ComponentListForFeedbackAPIAPIService.Feedbackpredictioncomponents": {
				{
					URL: "https://dtc.hgs.cloud",
					Description: "DTC Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas31pinComponents": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas32pinData": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas33pinEcuImage": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas41pinComponents": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas42pinData": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ComponentMeasurementValuesAPIService.CompMeas43pinEcuImage": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"DTCAPIAPIService.Dtcs": {
				{
					URL: "https://dtc.hgs.cloud",
					Description: "DTC Production",
				},
			},
			"GuideAPIService.Guide1": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ImageAPIService.Image1": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ImageAPIService.Image2WiringDiagrams": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"KTypeAPIAPIService.Ktype1makes": {
				{
					URL: "https://ktype.hgs.cloud",
					Description: "KType Production",
				},
			},
			"KTypeAPIAPIService.Ktype2models": {
				{
					URL: "https://ktype.hgs.cloud",
					Description: "KType Production",
				},
			},
			"KTypeAPIAPIService.Ktype3types": {
				{
					URL: "https://ktype.hgs.cloud",
					Description: "KType Production",
				},
			},
			"KTypeComponentMappingAPIService.ComponentMapping": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"LocationAPIAPIService.Location1component": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"LocationAPIAPIService.Location2fusebox": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ManufacturerCampaignsAPIService.ManufacturerCampaigns1tree": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ManufacturerCampaignsAPIService.ManufacturerCampaigns2manual": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"RecallCampaignsAPIService.RecallCampaigns1tree": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"RepairInstructionsAPIService.RepairInstructions1tree": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"RepairInstructionsAPIService.RepairInstructions2manual": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"RepairTimesAPIService.RepairTimes1tree": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"RepairTimesAPIService.RepairTimes2details": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ServiceSchedulesAPIService.ServiceSchedules1main": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"ServiceSchedulesAPIService.ServiceSchedules2inspections": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"TechnicalDataAPIService.TechnicalData1main": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"VISAPIAPIService.Ident": {
				{
					URL: "https://vis.hgs.cloud",
					Description: "VIS Production",
				},
			},
			"VehicleAPIAPIService.GetVehicles": {
				{
					URL: "https://vin.hgs.cloud",
					Description: "Vehicle Production",
				},
			},
			"WiringDiagramsImageAPIAPIService.Image2WiringDiagrams": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"WiringDiagramsSystemsAPIAPIService.PostWiringDiagrams1Systems": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
			"WiringDiagramsSystemsAPIAPIService.PostWiringDiagrams2Diagram": {
				{
					URL: "https://rmi.hgs.cloud",
					Description: "RMI Production",
				},
			},
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
