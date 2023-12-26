package ecs

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

// ReleaseCapacityReservation invokes the ecs.ReleaseCapacityReservation API synchronously
func (client *Client) ReleaseCapacityReservation(request *ReleaseCapacityReservationRequest) (response *ReleaseCapacityReservationResponse, err error) {
	response = CreateReleaseCapacityReservationResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseCapacityReservationWithChan invokes the ecs.ReleaseCapacityReservation API asynchronously
func (client *Client) ReleaseCapacityReservationWithChan(request *ReleaseCapacityReservationRequest) (<-chan *ReleaseCapacityReservationResponse, <-chan error) {
	responseChan := make(chan *ReleaseCapacityReservationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseCapacityReservation(request)
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

// ReleaseCapacityReservationWithCallback invokes the ecs.ReleaseCapacityReservation API asynchronously
func (client *Client) ReleaseCapacityReservationWithCallback(request *ReleaseCapacityReservationRequest, callback func(response *ReleaseCapacityReservationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseCapacityReservationResponse
		var err error
		defer close(result)
		response, err = client.ReleaseCapacityReservation(request)
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

// ReleaseCapacityReservationRequest is the request struct for api ReleaseCapacityReservation
type ReleaseCapacityReservationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PrivatePoolOptionsId string           `position:"Query" name:"PrivatePoolOptions.Id"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReleaseCapacityReservationResponse is the response struct for api ReleaseCapacityReservation
type ReleaseCapacityReservationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseCapacityReservationRequest creates a request to invoke ReleaseCapacityReservation API
func CreateReleaseCapacityReservationRequest() (request *ReleaseCapacityReservationRequest) {
	request = &ReleaseCapacityReservationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ReleaseCapacityReservation", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseCapacityReservationResponse creates a response to parse from ReleaseCapacityReservation response
func CreateReleaseCapacityReservationResponse() (response *ReleaseCapacityReservationResponse) {
	response = &ReleaseCapacityReservationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
