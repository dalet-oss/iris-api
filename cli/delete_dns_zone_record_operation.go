// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/dalet-oss/iris-api/client/dns"

	"github.com/spf13/cobra"
)

// makeOperationDNSDeleteDNSZoneRecordCmd returns a cmd to handle operation deleteDnsZoneRecord
func makeOperationDNSDeleteDNSZoneRecordCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeleteDNSZoneRecord",
		Short: `Deletes an existing zone's record.`,
		RunE:  runOperationDNSDeleteDNSZoneRecord,
	}

	if err := registerOperationDNSDeleteDNSZoneRecordParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDNSDeleteDNSZoneRecord uses cmd flags to call endpoint api
func runOperationDNSDeleteDNSZoneRecord(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dns.NewDeleteDNSZoneRecordParams()
	if err, _ := retrieveOperationDNSDeleteDNSZoneRecordRecordIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDNSDeleteDNSZoneRecordZoneIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDNSDeleteDNSZoneRecordResult(appCli.DNS.DeleteDNSZoneRecord(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDNSDeleteDNSZoneRecordParamFlags registers all flags needed to fill params
func registerOperationDNSDeleteDNSZoneRecordParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDNSDeleteDNSZoneRecordRecordIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDNSDeleteDNSZoneRecordZoneIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDNSDeleteDNSZoneRecordRecordIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	recordIdDescription := `Required. The ID of the DNS record to delete.`

	var recordIdFlagName string
	if cmdPrefix == "" {
		recordIdFlagName = "recordId"
	} else {
		recordIdFlagName = fmt.Sprintf("%v.recordId", cmdPrefix)
	}

	var recordIdFlagDefault string

	_ = cmd.PersistentFlags().String(recordIdFlagName, recordIdFlagDefault, recordIdDescription)

	return nil
}
func registerOperationDNSDeleteDNSZoneRecordZoneIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	zoneIdDescription := `Required. The ID of the zone's record to delete.`

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

func retrieveOperationDNSDeleteDNSZoneRecordRecordIDFlag(m *dns.DeleteDNSZoneRecordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("recordId") {

		var recordIdFlagName string
		if cmdPrefix == "" {
			recordIdFlagName = "recordId"
		} else {
			recordIdFlagName = fmt.Sprintf("%v.recordId", cmdPrefix)
		}

		recordIdFlagValue, err := cmd.Flags().GetString(recordIdFlagName)
		if err != nil {
			return err, false
		}
		m.RecordID = recordIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationDNSDeleteDNSZoneRecordZoneIDFlag(m *dns.DeleteDNSZoneRecordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDNSDeleteDNSZoneRecordResult parses request result and return the string content
func parseOperationDNSDeleteDNSZoneRecordResult(resp0 *dns.DeleteDNSZoneRecordOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteDnsZoneRecordOK is not supported

		// Non schema case: warning deleteDnsZoneRecordNotFound is not supported

		// Non schema case: warning deleteDnsZoneRecordInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response deleteDnsZoneRecordOK is not supported by go-swagger cli yet.

	return "", nil
}
