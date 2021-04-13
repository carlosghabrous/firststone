package main

import "testing"

func TestValidCommandsAreAccepted(t *testing.T) {

}

func TestInvalidCommandsThrowError(t *testing.T) {
	command := "fake command"
	if err := dispatchCommand(command); err == nil {
		t.Errorf("dispatchCommand(%s) should return error, got %v\n", command, err)
	}
}
