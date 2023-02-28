// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/dalet-oss/iris-api/client/dhcp"
	"github.com/dalet-oss/iris-api/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationDhcpCreateDHCPSubnetCmd returns a cmd to handle operation createDHCPSubnet
func makeOperationDhcpCreateDHCPSubnetCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "CreateDHCPSubnet",
		Short: `Creates a new DHCPv4 subnet.`,
		RunE:  runOperationDhcpCreateDHCPSubnet,
	}

	if err := registerOperationDhcpCreateDHCPSubnetParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDhcpCreateDHCPSubnet uses cmd flags to call endpoint api
func runOperationDhcpCreateDHCPSubnet(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dhcp.NewCreateDHCPSubnetParams()
	if err, _ := retrieveOperationDhcpCreateDHCPSubnetBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDhcpCreateDHCPSubnetResult(appCli.Dhcp.CreateDHCPSubnet(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDhcpCreateDHCPSubnetParamFlags registers all flags needed to fill params
func registerOperationDhcpCreateDHCPSubnetParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDhcpCreateDHCPSubnetBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDhcpCreateDHCPSubnetBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelSubnetFlags(0, "subnet", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationDhcpCreateDHCPSubnetBodyFlag(m *dhcp.CreateDHCPSubnetParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Subnet{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Subnet: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Subnet{}
	}
	err, added := retrieveModelSubnetFlags(0, bodyValueModel, "subnet", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationDhcpCreateDHCPSubnetResult parses request result and return the string content
func parseOperationDhcpCreateDHCPSubnetResult(resp0 *dhcp.CreateDHCPSubnetCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dhcp.CreateDHCPSubnetCreated)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning createDHCPSubnetBadRequest is not supported

		// Non schema case: warning createDHCPSubnetConflict is not supported

		// Non schema case: warning createDHCPSubnetInternalServerError is not supported

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
