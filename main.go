package main

import (
  "os/exec"
  "fmt"
  "strings"
)

func whichRPM() (string, error) {
  rpmPath, err := exec.LookPath("rpm")
  if err != nil {
    // This means, rpm utility is not installed in the system
    return "", err
  }
  return rpmPath, nil
}

func rpmVa(rpmPath string) ([]byte, error) {
  cmd := exec.Command(rpmPath, "-V", "docker", "httpd")
  out, err := cmd.CombinedOutput()
  return out, err
}

func processRPMVa(output string) []string {
  lines := strings.Split(output, "\n")
  return lines
}


func main() {
  // Find
  rpmPath, err := whichRPM()
  if err != nil {
    fmt.Println("Error: rpm is not installed in given system.")
    // TODO: Exit here
  }
  // fmt.Printf("rpm command line is at %s\n", rpmPath)
  outString, err := rpmVa(rpmPath)
  if err != nil{
    fmt.Printf("Error: %v\n", err)
  }
  out := processRPMVa(string(outString))
  fmt.Printf("Output: %d\n", len(out))
}







