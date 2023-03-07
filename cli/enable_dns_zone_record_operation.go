// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/dalet-oss/iris-api/client/dns"

	"github.com/spf13/cobra"
)

// makeOperationDNSEnableDNSZoneRecordCmd returns a cmd to handle operation enableDnsZoneRecord
func makeOperationDNSEnableDNSZoneRecordCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "EnableDNSZoneRecord",
		Short: `Enable a given DNS Zone record.`,
		RunE:  runOperationDNSEnableDNSZoneRecord,
	}

	if err := registerOperationDNSEnableDNSZoneRecordParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDNSEnableDNSZoneRecord uses cmd flags to call endpoint api
func runOperationDNSEnableDNSZoneRecord(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dns.NewEnableDNSZoneRecordParams()
	if err, _ := retrieveOperationDNSEnableDNSZoneRecordRecordIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDNSEnableDNSZoneRecordZoneIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDNSEnableDNSZoneRecordResult(appCli.DNS.EnableDNSZoneRecord(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDNSEnableDNSZoneRecordParamFlags registers all flags needed to fill params
func registerOperationDNSEnableDNSZoneRecordParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDNSEnableDNSZoneRecordRecordIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDNSEnableDNSZoneRecordZoneIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDNSEnableDNSZoneRecordRecordIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	recordIdDescription := `Required. The ID of the DNS record to enable.`

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
func registerOperationDNSEnableDNSZoneRecordZoneIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	zoneIdDescription := `Required. The ID of the zone's record to enable.`

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

func retrieveOperationDNSEnableDNSZoneRecordRecordIDFlag(m *dns.EnableDNSZoneRecordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDNSEnableDNSZoneRecordZoneIDFlag(m *dns.EnableDNSZoneRecordParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationDNSEnableDNSZoneRecordResult parses request result and return the string content
func parseOperationDNSEnableDNSZoneRecordResult(resp0 *dns.EnableDNSZoneRecordCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning enableDnsZoneRecordCreated is not supported

		// Non schema case: warning enableDnsZoneRecordConflict is not supported

		// Non schema case: warning enableDnsZoneRecordInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response enableDnsZoneRecordCreated is not supported by go-swagger cli yet.

	return "", nil
}