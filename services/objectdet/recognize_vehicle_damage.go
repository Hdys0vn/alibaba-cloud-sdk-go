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

// RecognizeVehicleDamage invokes the objectdet.RecognizeVehicleDamage API synchronously
func (client *Client) RecognizeVehicleDamage(request *RecognizeVehicleDamageRequest) (response *RecognizeVehicleDamageResponse, err error) {
	response = CreateRecognizeVehicleDamageResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeVehicleDamageWithChan invokes the objectdet.RecognizeVehicleDamage API asynchronously
func (client *Client) RecognizeVehicleDamageWithChan(request *RecognizeVehicleDamageRequest) (<-chan *RecognizeVehicleDamageResponse, <-chan error) {
	responseChan := make(chan *RecognizeVehicleDamageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeVehicleDamage(request)
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

// RecognizeVehicleDamageWithCallback invokes the objectdet.RecognizeVehicleDamage API asynchronously
func (client *Client) RecognizeVehicleDamageWithCallback(request *RecognizeVehicleDamageRequest, callback func(response *RecognizeVehicleDamageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeVehicleDamageResponse
		var err error
		defer close(result)
		response, err = client.RecognizeVehicleDamage(request)
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

// RecognizeVehicleDamageRequest is the request struct for api RecognizeVehicleDamage
type RecognizeVehicleDamageRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// RecognizeVehicleDamageResponse is the response struct for api RecognizeVehicleDamage
type RecognizeVehicleDamageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeVehicleDamageRequest creates a request to invoke RecognizeVehicleDamage API
func CreateRecognizeVehicleDamageRequest() (request *RecognizeVehicleDamageRequest) {
	request = &RecognizeVehicleDamageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("objectdet", "2019-12-30", "RecognizeVehicleDamage", "objectdet", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeVehicleDamageResponse creates a response to parse from RecognizeVehicleDamage response
func CreateRecognizeVehicleDamageResponse() (response *RecognizeVehicleDamageResponse) {
	response = &RecognizeVehicleDamageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
