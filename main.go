package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"opencoredata.org/csdcoMendeley/structs"
	// "time"
)

func main() {
	// url := "https://api.mendeley.com/documents?group_id=4cdf9c25-5a17-3644-8087-301999b5ba63&limit=50"
	url := "https://api.mendeley.com/documents?group_id=4cdf9c25-5a17-3644-8087-301999b5ba63&limit=500"

	req, _ := http.NewRequest("GET", url, nil)

	// When getting a new token (every hour for some dumb reason) both the authorization and postman-token will change
	req.Header.Add("authorization", "Bearer MSwxNDY0ODExNDU5MTM0LDE5NjQyMDEsMzE2MSxhbGwsLCxjOWQ0OTBjNWMzZDA1NTExY2I2ZDE2NmYyOTZjZjg1YjQwZTcxYWUtaCw5MzgyMDgxNi1lNTU3LTMzMmUtYjJkOS04ZDBhOGQ0MGQ1MGQsRXU3dFBSUlBCeEFnT3NBd0dKZElTbEIzSGlZ")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "2c9cded0-a9a5-4032-7028-32d94e1972bb")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	docsStruct := structs.Mdocs{}

	if err := json.Unmarshal(body, &docsStruct); err != nil {
		panic(err)
	}

	for _, elem := range docsStruct {
		doi := elem.Identifiers.Doi   // pass along to dataset func to print out along with the file.
		fmt.Println(doi)
		dataset(elem.ID)
	}

}

func dataset(id string) {

	url := fmt.Sprintf("https://api.mendeley.com/documents/%s", id)

	// fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	// When getting a new token (every hour for some dumb reason) both the authorization and postman-token will change
	req.Header.Add("authorization", "Bearer MSwxNDY0ODExNDU5MTM0LDE5NjQyMDEsMzE2MSxhbGwsLCxjOWQ0OTBjNWMzZDA1NTExY2I2ZDE2NmYyOTZjZjg1YjQwZTcxYWUtaCw5MzgyMDgxNi1lNTU3LTMzMmUtYjJkOS04ZDBhOGQ0MGQ1MGQsRXU3dFBSUlBCeEFnT3NBd0dKZElTbEIzSGlZ")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "2c9cded0-a9a5-4032-7028-32d94e1972bb")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	docStruct := structs.Mdoc{}

	if err := json.Unmarshal(body, &docStruct); err != nil {
		log.Println(err)
	}
	
	// Now that is unmarshalled we can add in the DOI and then 
	// re-marshall to JSON ..  prety it and send it out to a file.

	fmt.Printf("ID: %s\n", docStruct.ID)
	fmt.Printf("Tag: %s\n", docStruct.Tags)
	fmt.Printf("Abstract: %s\n\n", docStruct.Abstract)

}
