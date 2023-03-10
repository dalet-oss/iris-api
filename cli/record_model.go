// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/dalet-oss/iris-api/models"
	"github.com/spf13/cobra"
)

// Schema cli for Record

// register flags to command
func registerModelRecordFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRecordID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRecordTTL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRecordType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRecordValues(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRecordID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The DNS record ID.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerRecordTTL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ttlDescription := `The DNS record time-to-live.`

	var ttlFlagName string
	if cmdPrefix == "" {
		ttlFlagName = "ttl"
	} else {
		ttlFlagName = fmt.Sprintf("%v.ttl", cmdPrefix)
	}

	var ttlFlagDefault int32

	_ = cmd.PersistentFlags().Int32(ttlFlagName, ttlFlagDefault, ttlDescription)

	return nil
}

func registerRecordType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Required. The DNS record type.`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerRecordValues(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: values []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRecordFlags(depth int, m *models.Record, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveRecordIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ttlAdded := retrieveRecordTTLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ttlAdded

	err, typeAdded := retrieveRecordTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, valuesAdded := retrieveRecordValuesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valuesAdded

	return nil, retAdded
}

func retrieveRecordIDFlags(depth int, m *models.Record, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRecordTTLFlags(depth int, m *models.Record, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ttlFlagName := fmt.Sprintf("%v.ttl", cmdPrefix)
	if cmd.Flags().Changed(ttlFlagName) {

		var ttlFlagName string
		if cmdPrefix == "" {
			ttlFlagName = "ttl"
		} else {
			ttlFlagName = fmt.Sprintf("%v.ttl", cmdPrefix)
		}

		ttlFlagValue, err := cmd.Flags().GetInt32(ttlFlagName)
		if err != nil {
			return err, false
		}
		m.TTL = ttlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRecordTypeFlags(depth int, m *models.Record, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = &typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRecordValuesFlags(depth int, m *models.Record, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valuesFlagName := fmt.Sprintf("%v.values", cmdPrefix)
	if cmd.Flags().Changed(valuesFlagName) {
		// warning: values array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
