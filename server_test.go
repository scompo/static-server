package main

import (
        "testing"
        "os"
)

const (
        defaultValue = "DEFAULT VALUE"
        envVar = "MY_TEST_VAR"
        expectedValue = "MY TEST VALUE"
)

func TestReadEnvVarReturnsDefaultValueWhenEnvVarNotSet(t *testing.T) {
        res := ReadEnvVar("NOT_PRESENT", defaultValue)
	if res != defaultValue {
                t.Errorf("Did not return the default value when the enviroinment variable wasn't set")
        }
}

func TestReadEnvVarReturnsTheExpectedValueWhenSet(t *testing.T) {
        os.Setenv(envVar, expectedValue)
        res := ReadEnvVar(envVar, defaultValue)
	if res != expectedValue {
                t.Errorf("Did not return the expected value set in the enviroinment variable")
        }
}