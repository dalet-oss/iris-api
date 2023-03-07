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

// makeOperationDNSGetAllDNSServersCmd returns a cmd to handle operation getAllDnsServers
func makeOperationDNSGetAllDNSServersCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetAllDNSServers",
		Short: `Returns the IDs of DNS servers.`,
		RunE:  runOperationDNSGetAllDNSServers,
	}

	if err := registerOperationDNSGetAllDNSServersParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDNSGetAllDNSServers uses cmd flags to call endpoint api
func runOperationDNSGetAllDNSServers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dns.NewGetAllDNSServersParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDNSGetAllDNSServersResult(appCli.DNS.GetAllDNSServers(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDNSGetAllDNSServersParamFlags registers all flags needed to fill params
func registerOperationDNSGetAllDNSServersParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationDNSGetAllDNSServersResult parses request result and return the string content
func parseOperationDNSGetAllDNSServersResult(resp0 *dns.GetAllDNSServersOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dns.GetAllDNSServersOK)
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
