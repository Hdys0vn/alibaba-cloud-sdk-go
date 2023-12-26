package linkwan

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// CountNotifications invokes the linkwan.CountNotifications API synchronously
func (client *Client) CountNotifications(request *CountNotificationsRequest) (response *CountNotificationsResponse, err error) {
	response = CreateCountNotificationsResponse()
	err = client.DoAction(request, response)
	return
}

// CountNotificationsWithChan invokes the linkwan.CountNotifications API asynchronously
func (client *Client) CountNotificationsWithChan(request *CountNotificationsRequest) (<-chan *CountNotificationsResponse, <-chan error) {
	responseChan := make(chan *CountNotificationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountNotifications(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CountNotificationsWithCallback invokes the linkwan.CountNotifications API asynchronously
func (client *Client) CountNotificationsWithCallback(request *CountNotificationsRequest, callback func(response *CountNotificationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountNotificationsResponse
		var err error
		defer close(result)
		response, err = client.CountNotifications(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CountNotificationsRequest is the request struct for api CountNotifications
type CountNotificationsRequest struct {
	*requests.RpcRequest
	EndMillis   requests.Integer `position:"Query" name:"EndMillis"`
	HandleState string           `position:"Query" name:"HandleState"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
	Category    *[]string        `position:"Query" name:"Category"  type:"Repeated"`
	BeginMillis requests.Integer `position:"Query" name:"BeginMillis"`
}

// CountNotificationsResponse is the response struct for api CountNotifications
type CountNotificationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateCountNotificationsRequest creates a request to invoke CountNotifications API
func CreateCountNotificationsRequest() (request *CountNotificationsRequest) {
	request = &CountNotificationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CountNotifications", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCountNotificationsResponse creates a response to parse from CountNotifications response
func CreateCountNotificationsResponse() (response *CountNotificationsResponse) {
	response = &CountNotificationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}