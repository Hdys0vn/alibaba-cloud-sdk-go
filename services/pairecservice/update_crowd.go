package pairecservice

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

// UpdateCrowd invokes the pairecservice.UpdateCrowd API synchronously
func (client *Client) UpdateCrowd(request *UpdateCrowdRequest) (response *UpdateCrowdResponse, err error) {
	response = CreateUpdateCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCrowdWithChan invokes the pairecservice.UpdateCrowd API asynchronously
func (client *Client) UpdateCrowdWithChan(request *UpdateCrowdRequest) (<-chan *UpdateCrowdResponse, <-chan error) {
	responseChan := make(chan *UpdateCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCrowd(request)
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

// UpdateCrowdWithCallback invokes the pairecservice.UpdateCrowd API asynchronously
func (client *Client) UpdateCrowdWithCallback(request *UpdateCrowdRequest, callback func(response *UpdateCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCrowdResponse
		var err error
		defer close(result)
		response, err = client.UpdateCrowd(request)
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

// UpdateCrowdRequest is the request struct for api UpdateCrowd
type UpdateCrowdRequest struct {
	*requests.RoaRequest
	Body    string `position:"Body" name:"body"`
	CrowdId string `position:"Path" name:"CrowdId"`
}

// UpdateCrowdResponse is the response struct for api UpdateCrowd
type UpdateCrowdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCrowdRequest creates a request to invoke UpdateCrowd API
func CreateUpdateCrowdRequest() (request *UpdateCrowdRequest) {
	request = &UpdateCrowdRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "UpdateCrowd", "/api/v1/crowds/[CrowdId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateUpdateCrowdResponse creates a response to parse from UpdateCrowd response
func CreateUpdateCrowdResponse() (response *UpdateCrowdResponse) {
	response = &UpdateCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
