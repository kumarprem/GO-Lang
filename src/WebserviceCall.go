package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://webapps-qa.homedepot.com/TMSCommon/rs/distanceTimeLookup/2/param;routeType=PRACTICAL;originFacilityAliasId=5147;destFacilityAliasId=0364"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "a6192b70-a7de-b3c9-e801-8952d21b6df0")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}