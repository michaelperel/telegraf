package azure_iothub

import (
	"fmt"
	"testing"

	copier "github.com/jinzhu/copier"
)

// struct for generic test validation
type TestValidation struct {
	Name           string
	Input          interface{}
	ExpectedOutput interface{}
	ActualOutput   interface{}
}

// Check - performs check of TestValidation object, comparing ExpectedOutput to ActualOutput
func Check(tv TestValidation, t *testing.T) {
	if tv.ExpectedOutput != tv.ActualOutput {
		t.Errorf("\nERROR on %s\n\ninput: %s\nexpected: %s\nactual: %s\n", tv.Name, tv.Input, tv.ExpectedOutput, tv.ActualOutput)
	}
}

// reusable vars
var device_connection_string = "HostName=myiothub.azure-devices.net;DeviceId=mydevice;SharedAccessKey=My+L0ng+K3y="
var module_connection_string = "HostName=myiothub.azure-devices.net;DeviceId=mydevice;ModuleId=mymodule;SharedAccessKey=My+L0ng+K3y="
var invalid_connection_string = "He=myiothub.azure-devices.net;Dd=mydevice;Md=mymodule;SAK=My+L0ng+K3y="
var empty_string = ""
var whitespace_string = "     "

var iothub_all_connection_properties = &Iothub{
	HubName:         "myiothub",
	SharedAccessKey: "My+L0ng+K3y=",
	DeviceID:        "mydevice",
	ModuleID:        "mymodule",
}

var iothub_all_connection_properties_plus_policy = &Iothub{
	HubName:             "myiothub",
	SharedAccessKey:     "My+L0ng+K3y=",
	SharedAccessKeyName: "mykeyname",
	DeviceID:            "mydevice",
	ModuleID:            "mymodule",
}

var iothub_connection_string_only = &Iothub{
	ConnectionString: module_connection_string,
}

var iothub_missing_hubname = &Iothub{
	SharedAccessKey: "My+L0ng+K3y=",
	DeviceID:        "mydevice",
	ModuleID:        "mymodule",
}

var iothub_missing_sharedaccesskey = &Iothub{
	HubName:  "myiothub",
	DeviceID: "mydevice",
	ModuleID: "mymodule",
}

var iothub_missing_deviceid = &Iothub{
	HubName:         "myiothub",
	SharedAccessKey: "My+L0ng+K3y=",
	ModuleID:        "mymodule",
}

var iothub_missing_moduleid = &Iothub{
	HubName:         "myiothub",
	SharedAccessKey: "My+L0ng+K3y=",
	DeviceID:        "mydevice",
}

// Start of tests

// ####################################
// ###### TestHasConnectionString #####
// ####################################

// Tests to run, format: input, expected output, actual output
var TableTestHasConnectionString = []TestValidation{
	{"Device connection string", device_connection_string, true, nil},
	{"Module connection string", module_connection_string, true, nil},
	{"Invalid connection string", invalid_connection_string, true, nil},
	{"Empty connection string", empty_string, false, nil},
	{"Whitespace connection string", whitespace_string, false, nil},
}

// TestHasConnectionString - iterate over tests to run and check results.
func TestHasConnectionString(t *testing.T) {

	t.Log("##### Running tests for TestHasConnectionString #####")

	// for each TestValidation item in TableTestConnectionString
	for _, row := range TableTestHasConnectionString {

		t.Logf("Testing against: %s\n", row.Name)

		// create new Iothub
		temp := &Iothub{
			ConnectionString: fmt.Sprint(row.Input),
		}

		// test hasConnectionString() against Iothub
		result := temp.hasConnectionString()

		// update TestValidation item with result
		row.ActualOutput = result

		// check result against expected result
		Check(row, t)
	}
}

// ###########################
// ###### TestHasHubName #####
// ###########################

// Tests to run, format: input, expected output, actual output
var TableTestHasHubName = []TestValidation{
	{"Iothub with all connection properties", iothub_all_connection_properties, true, nil},
	{"Iothub missing HubName", iothub_missing_hubname, false, nil},
	{"Iothub missing SharedAccessKey", iothub_missing_sharedaccesskey, true, nil},
	{"Iothub missing DeviceID", iothub_missing_deviceid, true, nil},
	{"Iothub missing ModuleID", iothub_missing_moduleid, true, nil},
}

func TestHasHubName(t *testing.T) {

	t.Log("##### Running tests for TestHasHubName #####")

	// for each TestValidation item in TableTestHubName
	for _, row := range TableTestHasHubName {

		t.Logf("Testing against: %s\n", row.Name)

		temp := &Iothub{}
		copier.Copy(temp, row.Input)

		// test hasHubName() against Iothub
		result := temp.hasHubName()

		// update TestValidation item with result
		row.ActualOutput = result

		// check result against expected result
		Check(row, t)
	}
}

// ###################################
// ###### TestHasSharedAccessKey #####
// ###################################

// Tests to run, format: input, expected output, actual output
var TableTestHasSharedAccessKey = []TestValidation{
	{"Iothub with all connection properties", iothub_all_connection_properties, true, nil},
	{"Iothub missing HubName", iothub_missing_hubname, true, nil},
	{"Iothub missing SharedAccessKey", iothub_missing_sharedaccesskey, false, nil},
	{"Iothub missing DeviceID", iothub_missing_deviceid, true, nil},
	{"Iothub missing ModuleID", iothub_missing_moduleid, true, nil},
}

func TestHasSharedAccessKey(t *testing.T) {

	t.Log("##### Running tests for TestHasSharedAcessKey #####")

	// for each TestValidation item in TableTestHubName
	for _, row := range TableTestHasSharedAccessKey {

		t.Logf("Testing against: %s\n", row.Name)

		temp := &Iothub{}
		copier.Copy(temp, row.Input)

		// test hasSharedAccessKey() against Iothub
		result := temp.hasSharedAccessKey()

		// update TestValidation item with result
		row.ActualOutput = result

		// check result against expected result
		Check(row, t)
	}
}

// #######################################
// ###### TestHasSharedAccessKeyName #####
// #######################################

// Tests to run, format: input, expected output, actual output
var TableTestHasSharedAccessKeyName = []TestValidation{
	{"Iothub with all connection properties", iothub_all_connection_properties, false, nil},
	{"Iothub missing HubName", iothub_missing_hubname, false, nil},
	{"Iothub missing SharedAccessKey", iothub_missing_sharedaccesskey, false, nil},
	{"Iothub missing DeviceID", iothub_missing_deviceid, false, nil},
	{"Iothub missing ModuleID", iothub_missing_moduleid, false, nil},
	{"Iothub with all connection properties + policy name", iothub_all_connection_properties_plus_policy, true, nil},
}

func TestHasSharedAccessKeyName(t *testing.T) {

	t.Log("##### Running tests for TestHasSharedAccessKeyName #####")

	// for each TestValidation item in TableTestHasSharedAccessKeyName
	for _, row := range TableTestHasSharedAccessKeyName {

		t.Logf("Testing against: %s\n", row.Name)

		temp := &Iothub{}
		copier.Copy(temp, row.Input)

		// test hasSharedAccessKeyName() against Iothub
		result := temp.hasSharedAccessKeyName()

		// update TestValidation item with result
		row.ActualOutput = result

		// check result against expected result
		Check(row, t)
	}
}
