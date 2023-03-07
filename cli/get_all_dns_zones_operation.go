// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/dalet-oss/iris-api/client/dns"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDNSGetAllDNSZonesCmd returns a cmd to handle operation getAllDnsZones
func makeOperationDNSGetAllDNSZonesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetAllDNSZones",
		Short: `Returns the IDs of DNS zones.`,
		RunE:  runOperationDNSGetAllDNSZones,
	}

	if err := registerOperationDNSGetAllDNSZonesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDNSGetAllDNSZones uses cmd flags to call endpoint api
func runOperationDNSGetAllDNSZones(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dns.NewGetAllDNSZonesParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDNSGetAllDNSZonesResult(appCli.DNS.GetAllDNSZones(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDNSGetAllDNSZonesParamFlags registers all flags needed to fill params
func registerOperationDNSGetAllDNSZonesParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationDNSGetAllDNSZonesResult parses request result and return the string content
func parseOperationDNSGetAllDNSZonesResult(resp0 *dns.GetAllDNSZonesOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dns.GetAllDNSZonesOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
