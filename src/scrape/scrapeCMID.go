package scrape

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CMID struct {
	CMIDUrl string
}

func GetCMID(url string) (string, error) {
	// Here all we are doing is making the inital request to the url given by the user
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request: ", err)
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
		return "", err
	}

	// Redundant check for error status code just incase the request was successful but the status code was not 200
	if resp.StatusCode >= 400 {
		log.Println("Error: ", resp.StatusCode)
		return "", errors.New("Error: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body: ", err)
		return "", err
	}

	// Now here we start to search for the CMID static js file
	// by just splitting the body of the response
	findUrl := strings.Split(string(body), "<script src=\"")[1]

	findUrl2 := strings.Split(findUrl, "\"")[0]
	c := CMID{
		CMIDUrl: url + findUrl2,
	}

	// Now we make a call to the funciton to actually find the CMID
	cmid, err := c.findCMID()
	if err != nil {
		log.Println("Error finding CMID: ", err)
		return "", err
	}
	log.Println("CMID found: ", cmid)

	return cmid, nil
}

func (c *CMID) findCMID() (string, error) {
	// Here we make the request to the CMID url
	// using the js static file we found earlier
	// and using the origin url the user gave us
	req, err := http.NewRequest("GET", c.CMIDUrl, nil)
	if err != nil {
		log.Println("Error creating request")
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error sending request")
		return "", err
	}

	// Another redundant check for error status code
	if resp.StatusCode >= 400 {
		log.Println("Error: ", resp.StatusCode)
		return "", errors.New("Error: " + resp.Status)
	}
	log.Println("Searching for CMID...")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body")
		return "", err
	}

	// Now we start to search for the CMID in the js file
	// by just splitting the body of the response
	findCMID := strings.Split(string(body), "REACT_APP_CANDY_MACHINE_ID:\"")[1]

	return strings.Split(findCMID, "\"")[0], nil
}
