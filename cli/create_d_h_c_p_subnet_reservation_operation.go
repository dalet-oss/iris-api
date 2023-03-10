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

// makeOperationDhcpCreateDHCPSubnetReservationCmd returns a cmd to handle operation createDHCPSubnetReservation
func makeOperationDhcpCreateDHCPSubnetReservationCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "CreateDHCPSubnetReservation",
		Short: `Creates a new DHCPv4 subnet's reservation.`,
		RunE:  runOperationDhcpCreateDHCPSubnetReservation,
	}

	if err := registerOperationDhcpCreateDHCPSubnetReservationParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDhcpCreateDHCPSubnetReservation uses cmd flags to call endpoint api
func runOperationDhcpCreateDHCPSubnetReservation(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dhcp.NewCreateDHCPSubnetReservationParams()
	if err, _ := retrieveOperationDhcpCreateDHCPSubnetReservationBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDhcpCreateDHCPSubnetReservationSubnetIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDhcpCreateDHCPSubnetReservationResult(appCli.Dhcp.CreateDHCPSubnetReservation(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDhcpCreateDHCPSubnetReservationParamFlags registers all flags needed to fill params
func registerOperationDhcpCreateDHCPSubnetReservationParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDhcpCreateDHCPSubnetReservationBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDhcpCreateDHCPSubnetReservationSubnetIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDhcpCreateDHCPSubnetReservationBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelReservationFlags(0, "reservation", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationDhcpCreateDHCPSubnetReservationSubnetIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	subnetIdDescription := `Required. The ID of the DHCP subnet to create reservations on.`

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

func retrieveOperationDhcpCreateDHCPSubnetReservationBodyFlag(m *dhcp.CreateDHCPSubnetReservationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Reservation{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Reservation: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Reservation{}
	}
	err, added := retrieveModelReservationFlags(0, bodyValueModel, "reservation", cmd)
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
func retrieveOperationDhcpCreateDHCPSubnetReservationSubnetIDFlag(m *dhcp.CreateDHCPSubnetReservationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDhcpCreateDHCPSubnetReservationResult parses request result and return the string content
func parseOperationDhcpCreateDHCPSubnetReservationResult(resp0 *dhcp.CreateDHCPSubnetReservationCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dhcp.CreateDHCPSubnetReservationCreated)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning createDHCPSubnetReservationBadRequest is not supported

		// Non schema case: warning createDHCPSubnetReservationConflict is not supported

		// Non schema case: warning createDHCPSubnetReservationInternalServerError is not supported

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
