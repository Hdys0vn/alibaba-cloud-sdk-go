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

// DetectVehicleIllegalParking invokes the objectdet.DetectVehicleIllegalParking API synchronously
func (client *Client) DetectVehicleIllegalParking(request *DetectVehicleIllegalParkingRequest) (response *DetectVehicleIllegalParkingResponse, err error) {
	response = CreateDetectVehicleIllegalParkingResponse()
	err = client.DoAction(request, response)
	return
}

// DetectVehicleIllegalParkingWithChan invokes the objectdet.DetectVehicleIllegalParking API asynchronously
func (client *Client) DetectVehicleIllegalParkingWithChan(request *DetectVehicleIllegalParkingRequest) (<-chan *DetectVehicleIllegalParkingResponse, <-chan error) {
	responseChan := make(chan *DetectVehicleIllegalParkingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVehicleIllegalParking(request)
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

// DetectVehicleIllegalParkingWithCallback invokes the objectdet.DetectVehicleIllegalParking API asynchronously
func (client *Client) DetectVehicleIllegalParkingWithCallback(request *DetectVehicleIllegalParkingRequest, callback func(response *DetectVehicleIllegalParkingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVehicleIllegalParkingResponse
		var err error
		defer close(result)
		response, err = client.DetectVehicleIllegalParking(request)
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

// DetectVehicleIllegalParkingRequest is the request struct for api DetectVehicleIllegalParking
type DetectVehicleIllegalParkingRequest struct {
	*requests.RpcRequest
	RoadRegions string `position:"Body" name:"RoadRegions"`
	ImageURL    string `position:"Body" name:"ImageURL"`
}

// DetectVehicleIllegalParkingResponse is the response struct for api DetectVehicleIllegalParking
type DetectVehicleIllegalParkingResponse struct {
	*responses.BaseResponse
	Message   string                            `json:"Message" xml:"Message"`
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Code      string                            `json:"Code" xml:"Code"`
	Data      DataInDetectVehicleIllegalParking `json:"Data" xml:"Data"`
}

// CreateDetectVehicleIllegalParkingRequest creates a request to invoke DetectVehicleIllegalParking API
func CreateDetectVehicleIllegalParkingRequest() (request *DetectVehicleIllegalParkingRequest) {
	request = &DetectVehicleIllegalParkingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("objectdet", "2019-12-30", "DetectVehicleIllegalParking", "objectdet", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectVehicleIllegalParkingResponse creates a response to parse from DetectVehicleIllegalParking response
func CreateDetectVehicleIllegalParkingResponse() (response *DetectVehicleIllegalParkingResponse) {
	response = &DetectVehicleIllegalParkingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
