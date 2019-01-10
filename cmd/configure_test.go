package cmd

import (
	"testing"
	"bytes"
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestGetConfigValuesFromPrompts(t *testing.T) {

  var in bytes.Buffer
  var gotOut, wantOut bytes.Buffer

  expectedtfeURL      := "example-url.com"
  expectedtfeAPIToken := "exampletoken"

  // The reader should read to the \n each of two times.
  in.Write([]byte(fmt.Sprintf("%s\n%s\n", expectedtfeURL, expectedtfeAPIToken)))

  // wantOut could just be []byte, but for symmetry's sake I've used another buffer
  wantOut.Write([]byte("TFE URL:\nEnter a value (Default is https://app.terraform.io): \nTFE API Token (Create one at example-url.com/app/settings/tokens)\nEnter a value: \n"))

  tfeUrl, tfeAPIToken, err := GetConfigValuesFromPrompts(&in, &gotOut)

  // verify that correct prompts were sent to the writer
  if err != nil {
     t.Error("Unexpected error return", err)
  }

  // verify that correct prompts were sent to the writer
  if gotOut.String() != wantOut.String() {
    dmp   := diffmatchpatch.New()
    diffs := dmp.DiffMain(gotOut.String(), wantOut.String(), true)
    t.Errorf("\nPrompt comparison failed!\nDiff is below:\n" + dmp.DiffPrettyText(diffs))
  }

  // verify that correct prompts were sent to the writer
  if tfeUrl != expectedtfeURL {
    t.Errorf("expected %s to be equal to %s", tfeUrl, expectedtfeURL)
  }

  // verify that correct prompts were sent to the writer
  if tfeAPIToken != expectedtfeAPIToken {
    t.Errorf("expected %s to be equal to %s", tfeAPIToken, expectedtfeAPIToken)
  }

}
