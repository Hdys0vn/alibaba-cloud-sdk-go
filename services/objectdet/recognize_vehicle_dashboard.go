package objectdet

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

// RecognizeVehicleDashboard invokes the objectdet.RecognizeVehicleDashboard API synchronously
func (client *Client) RecognizeVehicleDashboard(request *RecognizeVehicleDashboardRequest) (response *RecognizeVehicleDashboardResponse, err error) {
	response = CreateRecognizeVehicleDashboardResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeVehicleDashboardWithChan invokes the objectdet.RecognizeVehicleDashboard API asynchronously
func (client *Client) RecognizeVehicleDashboardWithChan(request *RecognizeVehicleDashboardRequest) (<-chan *RecognizeVehicleDashboardResponse, <-chan error) {
	responseChan := make(chan *RecognizeVehicleDashboardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeVehicleDashboard(request)
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

// RecognizeVehicleDashboardWithCallback invokes the objectdet.RecognizeVehicleDashboard API asynchronously
func (client *Client) RecognizeVehicleDashboardWithCallback(request *RecognizeVehicleDashboardRequest, callback func(response *RecognizeVehicleDashboardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeVehicleDashboardResponse
		var err error
		defer close(result)
		response, err = client.RecognizeVehicleDashboard(request)
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

// RecognizeVehicleDashboardRequest is the request struct for api RecognizeVehicleDashboard
type RecognizeVehicleDashboardRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// RecognizeVehicleDashboardResponse is the response struct for api RecognizeVehicleDashboard
type RecognizeVehicleDashboardResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeVehicleDashboardRequest creates a request to invoke RecognizeVehicleDashboard API
func CreateRecognizeVehicleDashboardRequest() (request *RecognizeVehicleDashboardRequest) {
	request = &RecognizeVehicleDashboardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("objectdet", "2019-12-30", "RecognizeVehicleDashboard", "objectdet", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeVehicleDashboardResponse creates a response to parse from RecognizeVehicleDashboard response
func CreateRecognizeVehicleDashboardResponse() (response *RecognizeVehicleDashboardResponse) {
	response = &RecognizeVehicleDashboardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
