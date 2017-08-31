package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)

// CandidateInfoStore is a high level smart contract that CandidateInfoStores together business artifact based smart contracts
type CandidateInfoStore struct {

}

type CandidateDetails struct{
    CandidateId string `json:"candidateId"`
	Title string `json:"title"`
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	DOB string `json:"dob"`
	EmailID string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
	UniqueIdType string `json:"uniqueIdType"`
	UniqueIdNumber string `json:"uniqueIdNumber"`
	Nationality string `json:"nationality"`
	Address string `json:"address"`
	Country string `json:"country"`
	City string `json:"city"`
	Zip string `json:"zip"`
	State string `json:"state"`
	VerifyStatus string `json:"verifyStatus"`
	}

type CertificateDetails struct{
	CertificateId string `json:"certificateId"`
	CandidateId string `json:"candidateId"`
	Degree string `json:"degree"`
	Marks string `json:"marks"`
	Grade string `json:"grade"`
	Year string `json:"year"`
	UniversityName string `json:"universityName"`
	}
	
type ExperienceDetails struct{
	ExperienceId string `json:"experienceId"`
	CandidateId string `json:"candidateId"`
	Organization string `json:"organization"`
	DOJ string `json:"doj"`
	Designation string `json:"designation"`
	Skillset string `json:"skillset"`
	Certification string `json:"certification"`
	Salary string `json:"salary"`
	DOL string `json:"dol"`
	}
type CertificatesDetails struct{	
	CertificateDetails []CertificateDetails `json:"certificatedetails"`
}

type ExperiencesDetails struct{	
	ExperienceDetails []ExperienceDetails `json:"experiencedetails"`
}

func (t *CandidateInfoStore) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("CandidateDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create application Table
	err = stub.CreateTable("CandidateDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "candidateId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "emailId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "phoneNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "uniqueIdType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "uniqueIdNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "state", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "verifyStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating CandidateDetails.")
	}
	
	_, err = stub.GetTable("CertificateDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create application Table
	err = stub.CreateTable("CertificateDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "certificateId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "candidateId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "degree", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "marks", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "grade", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "year", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "universityName", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating CertificateDetails.")
	}
	
	_, err = stub.GetTable("ExperienceDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create application Table
	err = stub.CreateTable("ExperienceDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "experienceId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "candidateId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "organization", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "doj", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "designation", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "skillset", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "certification", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "salary", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dol", Type: shim.ColumnDefinition_STRING, Key: false},
		})
	if err != nil {
		return nil, errors.New("Failed creating ExperienceDetails.")
	}
	
	//Initializing the counter
	//stub.PutState("CANDIDATEINCREAMENTER",[]byte("1"))
	stub.PutState("CERTIFICATEINCREAMENTER",[]byte("1"))
	stub.PutState("EXPERIENCEINCREAMENTER",[]byte("1"))
	
	
	// setting up the users role
	stub.PutState("user_type1_1", []byte("Govt"))
	stub.PutState("user_type1_2", []byte("IBM"))
	stub.PutState("user_type1_3", []byte("CTS"))
	stub.PutState("user_type1_4", []byte("University"))	
	
	return nil, nil
}

//registerUser to register a user
func (t *CandidateInfoStore) RegisterCandidate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 16 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 16. Got: %d.", len(args))
		}
		candidateId:=args[0]
		title:=args[1]
		gender:=args[2]
		firstName:=args[3]
		lastName:=args[4]
		dob:=args[5]
		emailId:=args[6]
		phoneNumber:=args[7]
		uniqueIdType:=args[8]
		uniqueIdNumber:=args[9]
		nationality:=args[10]
		address :=args[11]
		country:=args[12]
		city:=args[13]
		zip:=args[14]
		state:=args[15]
		verifyStatus:="false"
			
			
		//stub.PutState(uniqueIdNumber,[]byte(candidateId))
		/*	
		assignerOrg, err := stub.GetState(candidateId)
		if assignerOrg !=nil{
			return nil, fmt.Errorf("Candidate already registered %s",candidateId)
		} else if err !=nil{
			return nil, fmt.Errorf("System error")
		}
		*/
		
		// Insert a row
		ok, err := stub.InsertRow("CandidateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: candidateId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: emailId}},
				&shim.Column{Value: &shim.Column_String_{String_: phoneNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdType}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: state}},
				&shim.Column{Value: &shim.Column_String_{String_: verifyStatus}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//UpdateCandidateDetails to verified a user
func (t *CandidateInfoStore) ApproveCandidateDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) < 1 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 1. Got: %d.", len(args))
		}
		candidateId:=args[0]
	
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: candidateId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("CandidateDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving application with candidateId %s. Error %s", candidateId, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}

		//End- Check that the currentStatus to newStatus transition is accurate
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"CandidateDetails",
			columns,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
		
		status:="true"
		
		//candidateId:=row.Columns[0].GetString_()
		title:=row.Columns[1].GetString_()
		gender:=row.Columns[2].GetString_()
		firstName:=row.Columns[3].GetString_()
		lastName:=row.Columns[4].GetString_()
		dob:=row.Columns[5].GetString_()
		emailId:=row.Columns[6].GetString_()
		phoneNumber:=row.Columns[7].GetString_()
		uniqueIdType:=row.Columns[8].GetString_()
		uniqueIdNumber:=row.Columns[9].GetString_()
		nationality:=row.Columns[10].GetString_()
		address :=row.Columns[11].GetString_()
		country:=row.Columns[12].GetString_()
		city:=row.Columns[13].GetString_()
		zip:=row.Columns[14].GetString_()
		state:=row.Columns[15].GetString_()
		verifyStatus:=status
		
		
		err = stub.DeleteRow(
					"CandidateDetails",
					columns,
				)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}

		// Insert a row
		ok, err := stub.InsertRow("CandidateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: candidateId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: emailId}},
				&shim.Column{Value: &shim.Column_String_{String_: phoneNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdType}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: state}},
				&shim.Column{Value: &shim.Column_String_{String_: verifyStatus}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

func (t *CandidateInfoStore) getCandidate(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting candidateId to query")
	}

	candidateId := args[0]


	// Get the row pertaining to this candidateId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: candidateId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("CandidateDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + candidateId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + candidateId + "\"}"
		return nil, errors.New(jsonResp)
	}


	res2E := CandidateDetails{}
	res2E.CandidateId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.DOB = row.Columns[5].GetString_()
	res2E.EmailID = row.Columns[6].GetString_()
	res2E.PhoneNumber = row.Columns[7].GetString_()
	res2E.UniqueIdType = row.Columns[8].GetString_()
	res2E.UniqueIdNumber = row.Columns[9].GetString_()
	res2E.Nationality = row.Columns[10].GetString_()
	res2E.Address = row.Columns[11].GetString_()
	res2E.Country = row.Columns[12].GetString_()
	res2E.City = row.Columns[13].GetString_()
	res2E.Zip = row.Columns[14].GetString_()
	res2E.State = row.Columns[15].GetString_()
	res2E.VerifyStatus = row.Columns[16].GetString_()

    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))

	return mapB, nil
	
}

func (t *CandidateInfoStore) getAllCandidate(stub shim.ChaincodeStubInterface , args []string) ([]byte, error) {

	// Get the row pertaining to this candidateId
	var columns []shim.Column
	rows, err := stub.GetRows("CandidateDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application \"}"
		return nil, errors.New(jsonResp)
	}

	res2E := []*CandidateDetails{}
	
	for row := range rows{
	newCan := new (CandidateDetails)
	newCan.CandidateId = row.Columns[0].GetString_()
	newCan.Title = row.Columns[1].GetString_()
	newCan.Gender = row.Columns[2].GetString_()
	newCan.FirstName = row.Columns[3].GetString_()
	newCan.LastName = row.Columns[4].GetString_()
	newCan.DOB = row.Columns[5].GetString_()
	newCan.EmailID = row.Columns[6].GetString_()
	newCan.PhoneNumber = row.Columns[7].GetString_()
	newCan.UniqueIdType = row.Columns[8].GetString_()
	newCan.UniqueIdNumber = row.Columns[9].GetString_()
	newCan.Nationality = row.Columns[10].GetString_()
	newCan.Address = row.Columns[11].GetString_()
	newCan.Country = row.Columns[12].GetString_()
	newCan.City = row.Columns[13].GetString_()
	newCan.Zip = row.Columns[14].GetString_()
	newCan.State = row.Columns[15].GetString_()
	newCan.VerifyStatus = row.Columns[16].GetString_()
	res2E=append(res2E,newCan)
}

    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))

	return mapB, nil

}


//Issue Certificate to register a user
func (t *CandidateInfoStore) CertificateIssue(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 6 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 6. Got: %d.", len(args))
		}
		
		//getting certifiactionId
		
		Avalbytes, err := stub.GetState("CERTIFICATEINCREAMENTER")
		Aval, _ := strconv.ParseInt(string(Avalbytes), 10, 0)
		newAval:=int(Aval) + 1

		newASNincrement:= strconv.Itoa(newAval)
		stub.PutState("ASNincrement", []byte(newASNincrement))

		

		certUniqueid:=string(Avalbytes)
		
		certificateId:=certUniqueid
		candidateId:=args[0]
		degree:=args[1]
		marks:=args[2]
		grade:=args[3]
		year:=args[4]
		universityName:=args[5]
		
		
		assignerOrg, err := stub.GetState(certificateId)
		if assignerOrg !=nil{
			return nil, fmt.Errorf("Candidate already registered %s",certificateId)
		} else if err !=nil{
			return nil, fmt.Errorf("System error")
		}
		
		
		// Insert a row
		ok, err := stub.InsertRow("CertificateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: certificateId}},
				&shim.Column{Value: &shim.Column_String_{String_: candidateId}},
				&shim.Column{Value: &shim.Column_String_{String_: degree}},
				&shim.Column{Value: &shim.Column_String_{String_: marks}},
				&shim.Column{Value: &shim.Column_String_{String_: grade}},
				&shim.Column{Value: &shim.Column_String_{String_: year}},
				&shim.Column{Value: &shim.Column_String_{String_: universityName}},
				}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		//append the certificate against candidateId
		certificate, err := stub.GetState("CERTIFICATE:"+candidateId)
		if certificate !=nil{
			var certificateString []string
			err = json.Unmarshal([]byte(certificate), &certificateString)
			if err != nil {
				return nil, errors.New("Row already exists.")
			}
			certificateString = append(certificateString, certificateId)
			outputMapBytes, _ := json.Marshal(certificateString)			
			stub.PutState("CERTIFICATE:"+candidateId, []byte(outputMapBytes))				
		} else{
			var certificate []string
			
			certificate = append(certificate, certificateId)
			outputMapBytes, _ := json.Marshal(certificate)			
			stub.PutState("CERTIFICATE:"+candidateId, []byte(outputMapBytes))	
		}
		
			
			
		return nil, nil

}

//UpdateCertificateDetails to verified a user
func (t *CandidateInfoStore) UpdateCertificateDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) < 1 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 1. Got: %d.", len(args))
		}
		certificateId:=args[0]
		degree:=args[1]
		marks:=args[2]
		grade:=args[3]
		year:=args[4]
		universityName:=args[5]
		
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: certificateId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("CertificateDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving application with certificateId %s. Error %s", certificateId, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}

		//End- Check that the currentStatus to newStatus transition is accurate
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"CertificateDetails",
			columns,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
				
		//certificateId =row.Columns[0].GetString_()
		degree=row.Columns[1].GetString_()
		marks=row.Columns[2].GetString_()
		grade=row.Columns[3].GetString_()
		year=row.Columns[4].GetString_()
		universityName=row.Columns[5].GetString_()
		
		// Insert a row
		ok, err := stub.InsertRow("CandidateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: certificateId}},
				&shim.Column{Value: &shim.Column_String_{String_: degree}},
				&shim.Column{Value: &shim.Column_String_{String_: marks}},
				&shim.Column{Value: &shim.Column_String_{String_: grade}},
				&shim.Column{Value: &shim.Column_String_{String_: year}},
				&shim.Column{Value: &shim.Column_String_{String_: universityName}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

func (t *CandidateInfoStore) getAllCertificateByCandidateId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting certificateId to query")
	}
	
	
	
	

	candidateId := args[0]
	//append the certificate against candidateId
	certificate, err := stub.GetState("CERTIFICATE:"+candidateId)
	var certificateString []string
		err = json.Unmarshal([]byte(certificate), &certificateString)
		if err != nil {
			return nil, errors.New("Row already exists.")
		}
	
	arrayCertificate := CertificatesDetails{}
	arrayCertificate.CertificateDetails=make([]CertificateDetails,0)
	
	for _, certificateId := range certificateString {
		// Get the row pertaining to this certificateId
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: certificateId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("CertificateDetails", columns)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get the data for the application " + certificateId + "\"}"
			return nil, errors.New(jsonResp)
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			jsonResp := "{\"Error\":\"Failed to get the data for the application " + certificateId + "\"}"
			return nil, errors.New(jsonResp)
		}


		newCan := CertificateDetails{}
		newCan.CertificateId = row.Columns[0].GetString_()
		newCan.Degree = row.Columns[1].GetString_()
		newCan.Marks = row.Columns[2].GetString_()
		newCan.Grade = row.Columns[3].GetString_()
		newCan.Year = row.Columns[4].GetString_()
		newCan.UniversityName = row.Columns[5].GetString_()
		arrayCertificate.CertificateDetails=append(arrayCertificate.CertificateDetails,newCan)		
	}

	
		
    mapB, _ := json.Marshal(arrayCertificate)
    fmt.Println(string(mapB))

	return mapB, nil

}


//Issue experience to register a user
func (t *CandidateInfoStore) addExperienceDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 6 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 6. Got: %d.", len(args))
		}
		
		//getting certifiactionId
		
		Avalbytes, err := stub.GetState("EXPERIENCEINCREAMENTER")
		Aval, _ := strconv.ParseInt(string(Avalbytes), 10, 0)
		newAval:=int(Aval) + 1

		newASNincrement:= strconv.Itoa(newAval)
		stub.PutState("ASNincrement", []byte(newASNincrement))

		

		experienceUniqueid:=string(Avalbytes)
		
		experienceId:=experienceUniqueid
		candidateId:=args[0]
		organization:=args[1]
		doj:=args[2]
		designation:=args[3]
		skillset:=args[4]
		certification:=args[5]
		salary:=args[6]
		dol:=args[7]
		
		
		assignerOrg, err := stub.GetState(experienceId)
		if assignerOrg !=nil{
			return nil, fmt.Errorf("Candidate already registered %s",experienceId)
		} else if err !=nil{
			return nil, fmt.Errorf("System error")
		}
		
		
		// Insert a row
		ok, err := stub.InsertRow("ExperienceDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: experienceId}},
				&shim.Column{Value: &shim.Column_String_{String_: organization}},
				&shim.Column{Value: &shim.Column_String_{String_: doj}},
				&shim.Column{Value: &shim.Column_String_{String_: designation}},
				&shim.Column{Value: &shim.Column_String_{String_: skillset}},
				&shim.Column{Value: &shim.Column_String_{String_: certification}},
				&shim.Column{Value: &shim.Column_String_{String_: salary}},
				&shim.Column{Value: &shim.Column_String_{String_: dol}},
				}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		//append the experience against candidateId
		experience, err := stub.GetState("EXPERIENCE:"+candidateId)
		if experience !=nil{
			var experienceString []string
			err = json.Unmarshal([]byte(experience), &experienceString)
			if err != nil {
				return nil, errors.New("Row already exists.")
			}
			experienceString =append(experienceString, experienceId)
			outputMapBytes, _ := json.Marshal(experienceString)			
			stub.PutState("EXPERIENCE:"+candidateId, []byte(outputMapBytes))				
		} else{
			var experience []string
			
			experience=append(experience, experienceId)
			outputMapBytes, _ := json.Marshal(experience)			
			stub.PutState("EXPERIENCE:"+candidateId, []byte(outputMapBytes))	
		}
					
		return nil, nil

}



//Update Experience Details to verified a user
func (t *CandidateInfoStore) UpdateExperienceDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) < 1 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 1. Got: %d.", len(args))
		}
				
		experienceId:=args[0]
		uniqueIdNumber:=args[1]
		organization:=args[2]
		doj:=args[3]
		designation:=args[4]
		skillset:=args[5]
		certification:=args[6]
		salary:=args[7]
		dol:=args[8]
		
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: experienceId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("ExperienceDetails", columns)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving application with experienceId %s. Error %s", experienceId, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}

		//End- Check that the currentStatus to newStatus transition is accurate
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"ExperienceDetails",
			columns,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
				
		experienceId =row.Columns[0].GetString_()
		organization=row.Columns[1].GetString_()
		doj=row.Columns[2].GetString_()
		designation=row.Columns[3].GetString_()
		skillset=row.Columns[4].GetString_()
		certification=row.Columns[5].GetString_()
		salary=row.Columns[6].GetString_()
		dol=row.Columns[7].GetString_()
		
		// Insert a row
		ok, err := stub.InsertRow("CandidateDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: experienceId}},
				&shim.Column{Value: &shim.Column_String_{String_: uniqueIdNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: organization}},
				&shim.Column{Value: &shim.Column_String_{String_: doj}},
				&shim.Column{Value: &shim.Column_String_{String_: designation}},
				&shim.Column{Value: &shim.Column_String_{String_: skillset}},
				&shim.Column{Value: &shim.Column_String_{String_: certification}},
				&shim.Column{Value: &shim.Column_String_{String_: salary}},
				&shim.Column{Value: &shim.Column_String_{String_: dol}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

func (t *CandidateInfoStore) getAllCertificateByCandidateId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting certificateId to query")
	}
	candidateId := args[0]
	//append the certificate against candidateId
	experience, err := stub.GetState("EXPERIENCE:"+candidateId)
	var experienceString []string
		err = json.Unmarshal([]byte(experience), &experienceString)
		if err != nil {
			return nil, errors.New("Row already exists.")
		}
	
	arrayExperience := ExperiencesDetails{}
	arrayExperience.ExperienceDetails=make([]ExperienceDetails,0)
	
	for _, experienceId := range experienceString {
		// Get the row pertaining to this certificateId
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: experienceId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("ExperienceDetails", columns)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get the data for the application " + experienceId + "\"}"
			return nil, errors.New(jsonResp)
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			jsonResp := "{\"Error\":\"Failed to get the data for the application " + experienceId + "\"}"
			return nil, errors.New(jsonResp)
		}


		newCan := ExperienceDetails{}
		newCan.ExperienceId = row.Columns[0].GetString_()
		newCan.Organization = row.Columns[1].GetString_()
		newCan.DOJ = row.Columns[2].GetString_()
		newCan.Designation = row.Columns[3].GetString_()
		newCan.Skillset = row.Columns[4].GetString_()
		newCan.Certification = row.Columns[5].GetString_()
		newCan.Salary = row.Columns[6].GetString_()
		newCan.DOJ = row.Columns[7].GetString_()
		arrayExperience.ExperienceDetails=append(arrayExperience.ExperienceDetails,newCan)		
	}

	
		
    mapB, _ := json.Marshal(arrayExperience)
    fmt.Println(string(mapB))

	return mapB, nil

}

func (t *CandidateInfoStore) getCandidateDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting uniqueIdNumber to query")
	}

	uniqueIdNumber := args[0]
	candidateid, err := stub.GetState(uniqueIdNumber)

	// Get the row pertaining to this uniqueIdNumber
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: string(candidateid)}}
	columns = append(columns, col1)
	
	var columns1 []shim.Column
	col2 := shim.Column{Value: &shim.Column_String_{String_: uniqueIdNumber}}
	columns1 = append(columns1, col2)

	row, err := stub.GetRow("CandidateDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uniqueIdNumber + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	row1, err1 := stub.GetRow("CertificateDetails", columns1)
	if err1 != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uniqueIdNumber + "\"}"
		return nil, errors.New(jsonResp)
	}
	row2, err2 := stub.GetRow("ExperienceDetails", columns1)
	if err2 != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uniqueIdNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uniqueIdNumber + "\"}"
		return nil, errors.New(jsonResp)
	}
	if len(row1.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uniqueIdNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	newCan1 := CertificateDetails{}
	newCan1.CertificateId = row1.Columns[0].GetString_()
	newCan1.CandidateId = row1.Columns[1].GetString_()
	newCan1.Degree = row1.Columns[2].GetString_()
	newCan1.Marks = row1.Columns[3].GetString_()
	newCan1.Grade = row1.Columns[4].GetString_()
	newCan1.Year = row1.Columns[5].GetString_()
	newCan1.UniversityName = row1.Columns[6].GetString_()
	
	newCan2 := ExperienceDetails{}
	newCan2.ExperienceId = row2.Columns[0].GetString_()
	newCan2.CandidateId = row2.Columns[1].GetString_()
	newCan2.DOJ = row2.Columns[2].GetString_()
	newCan2.Designation = row2.Columns[3].GetString_()
	newCan2.Skillset = row2.Columns[4].GetString_()
	newCan2.Certification = row2.Columns[5].GetString_()
	newCan2.Salary = row2.Columns[6].GetString_()
	

	newCan := CandidateDetails{}
	newCan.CandidateId = row.Columns[0].GetString_()
	newCan.Title = row.Columns[1].GetString_()
	newCan.Gender = row.Columns[2].GetString_()
	newCan.FirstName = row.Columns[3].GetString_()
	newCan.LastName = row.Columns[4].GetString_()
	newCan.DOB = row.Columns[5].GetString_()
	newCan.EmailID = row.Columns[6].GetString_()
	newCan.PhoneNumber = row.Columns[7].GetString_()
	newCan.UniqueIdType = row.Columns[8].GetString_()
	newCan.UniqueIdNumber = row.Columns[9].GetString_()
	newCan.Nationality = row.Columns[10].GetString_()
	newCan.Address = row.Columns[11].GetString_()
	newCan.Country = row.Columns[12].GetString_()
	newCan.City = row.Columns[13].GetString_()
	newCan.Zip = row.Columns[14].GetString_()
	newCan.State = row.Columns[15].GetString_()
	newCan.VerifyStatus = row.Columns[16].GetString_()
	
	type CandidateDetailsMerge struct{
    CertificateDetail CertificateDetails `json:"CertificateDetail"`
	ExperienceDetails ExperienceDetails `json:"ExperienceDetails"`
	CandidateDetails CandidateDetails `json:"CandidateDetails"`
	}
	
	newCan3 := CandidateDetailsMerge{}
	newCan3.CertificateDetail=newCan1
	newCan3.ExperienceDetails=newCan2
	newCan3.CandidateDetails=newCan
	
    mapB, _ := json.Marshal(newCan3)
    fmt.Println(string(mapB))

	return mapB, nil

}


// Invoke invokes the chaincode
func (t *CandidateInfoStore) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "RegisterCandidate" {
		t := CandidateInfoStore{}
		return t.RegisterCandidate(stub, args)	
	} else if function == "ApproveCandidateDetails" {
		t := CandidateInfoStore{}
		return t.ApproveCandidateDetails(stub, args)	
	} else if function == "CertificateIssue" {
		t := CandidateInfoStore{}
		return t.CertificateIssue(stub, args)	
	}else if function == "addExperienceDetails" {
		t := CandidateInfoStore{}
		return t.addExperienceDetails(stub, args)	
	}

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *CandidateInfoStore) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getCandidate" { 
		t := CandidateInfoStore{}
		return t.getCandidate(stub, args)
	}else if function == "getAllCandidate"{
        t := CandidateInfoStore{}
		return t.getAllCandidate(stub, args)  
    } else if function == "getAllCertificateByCandidateId"{
        t := CandidateInfoStore{}
		return t.getAllCertificateByCandidateId(stub, args)  
    } else if function == "getAllExperienceByCandidateId"{
        t := CandidateInfoStore{}
		return t.getAllExperienceByCandidateId(stub, args)  
    }
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(CandidateInfoStore))
	if err != nil {
		fmt.Printf("Error starting CandidateInfoStore: %s", err)
	}
} 