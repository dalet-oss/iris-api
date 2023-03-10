// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/dalet-oss/iris-api/client/dhcp"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDhcpGetDHCPSubnetCmd returns a cmd to handle operation getDHCPSubnet
func makeOperationDhcpGetDHCPSubnetCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetDHCPSubnet",
		Short: `Returns the requested DHCP Subnet object.`,
		RunE:  runOperationDhcpGetDHCPSubnet,
	}

	if err := registerOperationDhcpGetDHCPSubnetParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDhcpGetDHCPSubnet uses cmd flags to call endpoint api
func runOperationDhcpGetDHCPSubnet(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dhcp.NewGetDHCPSubnetParams()
	if err, _ := retrieveOperationDhcpGetDHCPSubnetSubnetIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDhcpGetDHCPSubnetResult(appCli.Dhcp.GetDHCPSubnet(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDhcpGetDHCPSubnetParamFlags registers all flags needed to fill params
func registerOperationDhcpGetDHCPSubnetParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDhcpGetDHCPSubnetSubnetIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDhcpGetDHCPSubnetSubnetIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	subnetIdDescription := `Required. The ID of the DHCP subnet to query.`

	var subnetIdFlagName string
	if cmdPrefix == "" {
		subnetIdFlagName = "subnetId"
	} else {
		subnetIdFlagName = fmt.Sprintf("%v.subnetId", cmdPrefix)
	}

	var subnetIdFlagDefault string

	_ = cmd.PersistentFlags().String(subnetIdFlagName, subnetIdFlagDefault, subnetIdDescription)

	return nil
}

func retrieveOperationDhcpGetDHCPSubnetSubnetIDFlag(m *dhcp.GetDHCPSubnetParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("subnetId") {

		var subnetIdFlagName string
		if cmdPrefix == "" {
			subnetIdFlagName = "subnetId"
		} else {
			subnetIdFlagName = fmt.Sprintf("%v.subnetId", cmdPrefix)
		}

		subnetIdFlagValue, err := cmd.Flags().GetString(subnetIdFlagName)
		if err != nil {
			return err, false
		}
		m.SubnetID = subnetIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDhcpGetDHCPSubnetResult parses request result and return the string content
func parseOperationDhcpGetDHCPSubnetResult(resp0 *dhcp.GetDHCPSubnetOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dhcp.GetDHCPSubnetOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getDHCPSubnetNotFound is not supported

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
