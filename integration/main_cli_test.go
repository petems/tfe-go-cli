// +build integration

package main

import (
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "path"
  "path/filepath"
  "runtime"
  "testing"
  "reflect"
  "github.com/sergi/go-diff/diffmatchpatch"
)

var update = flag.Bool("update", false, "update golden files")

func tmpFixturePath(t *testing.T) string {
  _, filename, _, ok := runtime.Caller(0)
  if !ok {
    t.Fatalf("problems recovering caller information")
  }

  return filepath.Join(filepath.Dir(filename), "tmp/")
}

func fixturePath(t *testing.T, fixture string) string {
  _, filename, _, ok := runtime.Caller(0)
  if !ok {
    t.Fatalf("problems recovering caller information")
  }

  return filepath.Join(filepath.Dir(filename), fixture)
}

func writeFixture(t *testing.T, fixture string, content []byte) {
  err := ioutil.WriteFile(fixturePath(t, fixture), content, 0644)
  if err != nil {
    t.Fatal(err)
  }
}

func loadFixture(t *testing.T, fixture string) string {
  content, err := ioutil.ReadFile(fixturePath(t, fixture))
  if err != nil {
    t.Fatal(err)
  }

  return string(content)
}

func loadTempFixtureFile(t *testing.T, tempfixturefile string) string {
  content, err := ioutil.ReadFile(tmpFixturePath(t) + "/" + tempfixturefile)
  if err != nil {
    t.Fatal(err)
  }

  return string(content)
}


func TestCliArgs(t *testing.T) {
  tests := []struct {
    cmdPath string
    name    string
    args    []string
    fixture string
  }{
    {"tfe-go-cli-int-testing", "no arguments", []string{}, "no-args.golden"},
    {"tfe-go-cli-int-testing", "version argument", []string{"version"}, "version.golden"},
    {"integration/authorize.exp", "authorize command", []string{""}, "authorize.golden"},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      dir, err := os.Getwd()
      if err != nil {
        t.Fatal(err)
      }

      cmd := exec.Command(path.Join(dir, tt.cmdPath), tt.args...)
      output, err := cmd.CombinedOutput()
      if err != nil {
        t.Fatal(err, fmt.Sprintf("%s", output))
      }

      if *update {
        writeFixture(t, tt.fixture, output)
      }

      actual := string(output)

      expected := loadFixture(t, tt.fixture)

      if !reflect.DeepEqual(actual, expected) {
        dmp := diffmatchpatch.New()
        diffs := dmp.DiffMain(actual, expected, false)
        t.Errorf("\nGolden file comparison failed!\nDiff is below:\n" + dmp.DiffPrettyText(diffs))
      }

      if tt.name == "authorize command" {

        actual := loadTempFixtureFile(t, "int.yaml")

        expected := loadFixture(t, "config.yaml.golden")

        if !reflect.DeepEqual(actual, expected) {
          dmp := diffmatchpatch.New()
          diffs := dmp.DiffMain(actual, expected, false)
          t.Errorf("\nGolden file comparison failed!\nDiff is below:\n" + dmp.DiffPrettyText(diffs))
        }

      }
    })
  }
}

func TestMain(m *testing.M) {
  err := os.Chdir("..")
  if err != nil {
    fmt.Printf("could not change dir: %v", err)
    os.Exit(1)
  }
  make := exec.Command("make", "int_build")
  err = make.Run()
  if err != nil {
    fmt.Printf("could not make binary for %s: %v", "tfe-go-cli", err)
    os.Exit(1)
  }

  expectCmd := exec.Command("expect")
  if err := expectCmd.Run(); err != nil {
    fmt.Printf("expect needs to be installed for acceptence tests: %s", err)
    os.Exit(1)
  }

  os.Exit(m.Run())
}
