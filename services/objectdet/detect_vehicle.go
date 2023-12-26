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

// DetectVehicle invokes the objectdet.DetectVehicle API synchronously
func (client *Client) DetectVehicle(request *DetectVehicleRequest) (response *DetectVehicleResponse, err error) {
	response = CreateDetectVehicleResponse()
	err = client.DoAction(request, response)
	return
}

// DetectVehicleWithChan invokes the objectdet.DetectVehicle API asynchronously
func (client *Client) DetectVehicleWithChan(request *DetectVehicleRequest) (<-chan *DetectVehicleResponse, <-chan error) {
	responseChan := make(chan *DetectVehicleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVehicle(request)
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

// DetectVehicleWithCallback invokes the objectdet.DetectVehicle API asynchronously
func (client *Client) DetectVehicleWithCallback(request *DetectVehicleRequest, callback func(response *DetectVehicleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVehicleResponse
		var err error
		defer close(result)
		response, err = client.DetectVehicle(request)
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

// DetectVehicleRequest is the request struct for api DetectVehicle
type DetectVehicleRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// DetectVehicleResponse is the response struct for api DetectVehicle
type DetectVehicleResponse struct {
	*responses.BaseResponse
	RequestId string              `json:"RequestId" xml:"RequestId"`
	Data      DataInDetectVehicle `json:"Data" xml:"Data"`
}

// CreateDetectVehicleRequest creates a request to invoke DetectVehicle API
func CreateDetectVehicleRequest() (request *DetectVehicleRequest) {
	request = &DetectVehicleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("objectdet", "2019-12-30", "DetectVehicle", "objectdet", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectVehicleResponse creates a response to parse from DetectVehicle response
func CreateDetectVehicleResponse() (response *DetectVehicleResponse) {
	response = &DetectVehicleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
