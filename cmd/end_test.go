package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
	"timetracker_cli/internal"
)

func TestCannotEndSessionThatHasNotStarted(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	cmd := NewEndCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got none")
	}
}

func TestCanAddDescriptionWhenEnding(t *testing.T) {
	tracker := internal.NewTestTracker(internal.NewStubClock(5), internal.TrackerConfig{})

	tracker.Start()

	expected := "test description"

	cmd := NewEndCmd(tracker)
	cmd.SetArgs([]string{expected})
	cmd.Execute()

	if len(tracker.Sessions) != 1 {
		t.Fatalf("Expected 1 saved session got %d", len(tracker.Sessions))
	}

	savedSession := tracker.Sessions[0]

	if savedSession.Description == "" {
		t.Errorf("Expected description '%s' got ''", expected)
	}
}

func TestEndAcceptsOnlyOneArgument(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Start()

	cmd := NewEndCmd(tracker)
	cmd.SetArgs([]string{"a", "b"})
	b := bytes.NewBufferString("")
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got ''")
	}

	if tracker.Current.Description == "a" {
		t.Error("Description shouldn't be set on error")
	}
}

func TestCannotEndSessionWithoutDescription(t *testing.T) {
	tracker := internal.NewTracker(internal.TrackerConfig{})

	tracker.Start()

	cmd := NewEndCmd(tracker)
	b := bytes.NewBufferString("")
	cmd.SetErr(b)
	cmd.Execute()

	out, err := ioutil.ReadAll(b)

	if err != nil {
		t.Fatal(err)
	}

	if string(out) == "" {
		t.Error("Expected error, got none")
	}
}

func TestCanEndASession(t *testing.T) {
	tracker := internal.NewTestTracker(internal.NewStubClock(5), internal.TrackerConfig{})

	tracker.Start()

	expected := "test description"

	cmd := NewEndCmd(tracker)
	cmd.SetArgs([]string{expected})
	cmd.Execute()

	if len(tracker.Sessions) != 1 {
		t.Errorf("Expected 1 saved session got %d", len(tracker.Sessions))
	}
}
