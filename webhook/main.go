package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*
curl -s -H "Authorization: Bearer $ACCESS_TOKEN" -H "Content-Type: application/json" \
  -d '{"summary":"Important Meeting","starts_at":"2015-06-04T00:00:00Z","ends_at":"2015-06-04T00:00:00Z"}' \
  https://3.basecampapi.com/$ACCOUNT_ID/buckets/1/schedules/3/entries.json
*/
const (
	accountID   = "xxx"
	accessToken = "xx--42af0416989acce685133c8f55cf40855dfbf4d3"
)

func main() {
	/*
		fmt.Println("webhook")
		url := "https://3.basecamp.com/xxx/integrations/VA5bYy3rjSySL9LNRSRBFwAo/buckets/34862946/chats/6688955697/lines"
	*/

	//requestProjects()
	//getScheduleEntries()
	addScheduleEntry()
}

func getScheduleEntries() {
	//url := "https://3.basecampapi.com/xxx/buckets/34862946/schedules/6679430062.json"
	url := "https://3.basecampapi.com/xxx/buckets/34862946/schedules/6679430062/entries.json"
	result, err := doRequest(url, http.MethodGet, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// POST /buckets/1/schedules/3/entries.json
func addScheduleEntry() {
	projectID := "34862946"
	scheduleID := "6679430062"
	url := "https://3.basecampapi.com/xxx/buckets/" + projectID + "/schedules/" + scheduleID + "/entries.json"

	result, err := doRequest(url, http.MethodPost,
		strings.NewReader(`{"summary":"Important Meeting","starts_at":"2023-11-04T00:00:00Z","ends_at":"2023-11-04T00:00:00Z"}`))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func requestProjects() {
	url := "https://3.basecampapi.com/" + accountID + "/projects.json"
	result, err := doRequest(url, http.MethodGet, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func doRequest(url, method string, body io.Reader) (string, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")

	resp, err := newClient().Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 查看一下
	// io.Copy(os.Stdout, resp.Body)

	// log.Println(resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("cannot read body: %w", err)
		}
		return "", fmt.Errorf("%w: %s, %s",
			err, http.StatusText(resp.StatusCode), msg)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func newClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
