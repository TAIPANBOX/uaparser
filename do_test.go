package uaparser

import "testing"

func TestParse(t *testing.T) {
	var (
		expectedBrowserNames     map[string][]string = GetBrowserNames()
		expectedOperatingSystems map[string][]string = GetOSNames()
		expectedDeviceTypes      map[string][]string = GetDeviceTypes()
		totalTestCount                               = 0
	)

	totalTestCount += _checkExepectations(t, expectedOperatingSystems, "os")
	totalTestCount += _checkExepectations(t, expectedBrowserNames, "browser")
	totalTestCount += _checkExepectations(t, expectedDeviceTypes, "deviceType")
}

func _checkExepectations(t *testing.T, expectations map[string][]string, testType string) (testCount int) {
	var (
		uaParseResult *UAInfo
		testResult    bool
		comparedTo    string
	)

	testCount = 0

	for expectation := range expectations {
		for key := range expectations[expectation] {
			uaParseResult = Parse(expectations[expectation][key])

			if testType == "browser" {
				testResult, comparedTo = _checkBrowser(uaParseResult, expectation)
			} else if testType == "os" {
				testResult, comparedTo = _checkOs(uaParseResult, expectation)
			} else if testType == "deviceType" {
				testResult, comparedTo = _checkDeviceType(uaParseResult, expectation)
			}

			if !testResult {
				t.Fatalf("Expected: '%s' got '%s' on: '%s'", expectation, comparedTo, expectations[expectation][key])
			}
			testCount += 1
		}
	}

	return testCount
}

func _checkBrowser(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.Browser == nil {
		return false, ""
	}
	return (uainfo.Browser.Name == expectation), uainfo.Browser.Name
}

func _checkOs(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.OS == nil {
		return false, ""
	}
	return (uainfo.OS.Name == expectation), uainfo.OS.Name
}

func _checkDeviceType(uainfo *UAInfo, expectation string) (result bool, comparedTo string) {
	if uainfo.DeviceType == nil {
		return false, ""
	}
	return (uainfo.DeviceType.Name == expectation), uainfo.DeviceType.Name
}
