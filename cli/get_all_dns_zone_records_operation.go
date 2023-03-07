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

// makeOperationDNSGetAllDNSZoneRecordsCmd returns a cmd to handle operation getAllDnsZoneRecords
func makeOperationDNSGetAllDNSZoneRecordsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetAllDNSZoneRecords",
		Short: `Returns the list of zone's record IDs.`,
		RunE:  runOperationDNSGetAllDNSZoneRecords,
	}

	if err := registerOperationDNSGetAllDNSZoneRecordsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDNSGetAllDNSZoneRecords uses cmd flags to call endpoint api
func runOperationDNSGetAllDNSZoneRecords(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dns.NewGetAllDNSZoneRecordsParams()
	if err, _ := retrieveOperationDNSGetAllDNSZoneRecordsZoneIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDNSGetAllDNSZoneRecordsResult(appCli.DNS.GetAllDNSZoneRecords(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDNSGetAllDNSZoneRecordsParamFlags registers all flags needed to fill params
func registerOperationDNSGetAllDNSZoneRecordsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDNSGetAllDNSZoneRecordsZoneIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDNSGetAllDNSZoneRecordsZoneIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	zoneIdDescription := `Required. The ID of the DNS zone to get records from.`

	var zoneIdFlagName string
	if cmdPrefix == "" {
		zoneIdFlagName = "zoneId"
	} else {
		zoneIdFlagName = fmt.Sprintf("%v.zoneId", cmdPrefix)
	}

	var zoneIdFlagDefault string

	_ = cmd.PersistentFlags().String(zoneIdFlagName, zoneIdFlagDefault, zoneIdDescription)

	return nil
}

func retrieveOperationDNSGetAllDNSZoneRecordsZoneIDFlag(m *dns.GetAllDNSZoneRecordsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("zoneId") {

		var zoneIdFlagName string
		if cmdPrefix == "" {
			zoneIdFlagName = "zoneId"
		} else {
			zoneIdFlagName = fmt.Sprintf("%v.zoneId", cmdPrefix)
		}

		zoneIdFlagValue, err := cmd.Flags().GetString(zoneIdFlagName)
		if err != nil {
			return err, false
		}
		m.ZoneID = zoneIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDNSGetAllDNSZoneRecordsResult parses request result and return the string content
func parseOperationDNSGetAllDNSZoneRecordsResult(resp0 *dns.GetAllDNSZoneRecordsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*dns.GetAllDNSZoneRecordsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getAllDnsZoneRecordsNotFound is not supported

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