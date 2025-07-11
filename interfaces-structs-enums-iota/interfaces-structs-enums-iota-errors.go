package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	idfyService := ekycIdfy{
		idfyClientId:     "1234567890",
		idfyClientSecret: "secret",
		idfyBaseUrl:      "idfy.com/api/v1/ekyc",
	}
	perfiosService := ekycPerfios{
		perfiosClientId:     "1234567890",
		perfiosClientSecret: "secret",
		perfiosBaseUrl:      "idfy.com/api/v1/ekyc",
	}

	var ekycService ekycVerifier

	var aadhaar uint
	fmt.Print("Enter aadhaar number: ")
	fmt.Scan(&aadhaar)

	var status ekycStatus
	ekycService = &idfyService
	status, err := ekycService.verifyAadhaar(aadhaar)
	if err != nil {
		if errors.Is(err, ErrorIdfyServiceDown) {
			fmt.Println("Idfy service failed. retrying with perfios...")
			ekycService = &perfiosService
			status, err = ekycService.verifyAadhaar(aadhaar)
			if err == nil {
				fmt.Printf("Aadhaar %v %v\n", aadhaar, status)
			} else {
				fmt.Print("Both services failed\n")
			}
		} else if errors.Is(err, ErrorAadhaarLength) {
			fmt.Println("Enter a valid aadhaar")
		}
	} else {
		fmt.Printf("Aadhaar %v %v", aadhaar, status)
		fmt.Println()
	}
}

// created a new type and constants for maintaining the ekycStatus
type ekycStatus int

const (
	Pending ekycStatus = iota
	Verified
	Rejected
)

// error message
/*
NOTE:
	errors.New creates a new unique error instance everytime.
	Comparing two different error instance created witht e same error message are NOT equal
	eg.
	const a = "error message"
	b := errors.New(a)
	c := errors.New(a)
	if b==c{
		fmt.Println("This won;t print")
	}else{
		fmt.Println("Hello")
	}
*/
const (
	AadhaarLengthErrorMessage      = "aadhaar number must be exactly 12 digits"
	IdfyServiceDownErrorMessage    = "idfy service is down"
	PerfiosServiceDownErrorMessage = "perfios service is Idown"
)
var (
	ErrorAadhaarLength      = errors.New(AadhaarLengthErrorMessage)
	ErrorIdfyServiceDown    = errors.New(IdfyServiceDownErrorMessage)
	ErrorPerfiosServiceDown = errors.New(PerfiosServiceDownErrorMessage)
)

// E-KYC verifier
type ekycVerifier interface {
	verifyAadhaar(aadhaarNo uint) (ekycStatus, error)
}

// IDFY implementation
type ekycIdfy struct {
	idfyClientId     string
	idfyClientSecret string
	idfyBaseUrl      string
}

// simulating idfy failure
func (idfy *ekycIdfy) verifyAadhaar(aadhaarNo uint) (ekycStatus, error) {
	if !validAadhaar(aadhaarNo) {
		return Rejected, ErrorAadhaarLength
	}
	return Rejected, ErrorIdfyServiceDown
}

// Perfios implementation
type ekycPerfios struct {
	perfiosClientId     string
	perfiosClientSecret string
	perfiosBaseUrl      string
}

// (perfios ekycPerfios) -> receiver; to signify that this function is linked to the structure
func (perfios *ekycPerfios) verifyAadhaar(aadhaarNo uint) (ekycStatus, error) {
	if !validAadhaar(aadhaarNo) {
		return Rejected, ErrorAadhaarLength
	}
	return Verified, nil
}

// utility
func validAadhaar(aadhaarNo uint) bool {
	return len(strconv.Itoa(int(aadhaarNo))) == 12
}
