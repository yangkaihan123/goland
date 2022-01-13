package func_02_test

import (
	"fmt"
	"testing"
)

func TestDefer4(t *testing.T) {
	doDBOperations()
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning form function here!")
	return
	//terminate the program
	//deferred function executed here just before actually returning, even if
	//there is a return or abnormal termination before
}

func disconnectFromDB() {
	fmt.Println("ok, the database is disconnected")
}

func connectToDB() {
	fmt.Println("ok, the database is connected")
}
