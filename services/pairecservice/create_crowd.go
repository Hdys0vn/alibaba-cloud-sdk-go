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

// CreateCrowd invokes the pairecservice.CreateCrowd API synchronously
func (client *Client) CreateCrowd(request *CreateCrowdRequest) (response *CreateCrowdResponse, err error) {
	response = CreateCreateCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCrowdWithChan invokes the pairecservice.CreateCrowd API asynchronously
func (client *Client) CreateCrowdWithChan(request *CreateCrowdRequest) (<-chan *CreateCrowdResponse, <-chan error) {
	responseChan := make(chan *CreateCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCrowd(request)
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

// CreateCrowdWithCallback invokes the pairecservice.CreateCrowd API asynchronously
func (client *Client) CreateCrowdWithCallback(request *CreateCrowdRequest, callback func(response *CreateCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCrowdResponse
		var err error
		defer close(result)
		response, err = client.CreateCrowd(request)
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

// CreateCrowdRequest is the request struct for api CreateCrowd
type CreateCrowdRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateCrowdResponse is the response struct for api CreateCrowd
type CreateCrowdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CrowdId   string `json:"CrowdId" xml:"CrowdId"`
}

// CreateCreateCrowdRequest creates a request to invoke CreateCrowd API
func CreateCreateCrowdRequest() (request *CreateCrowdRequest) {
	request = &CreateCrowdRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateCrowd", "/api/v1/crowds", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCrowdResponse creates a response to parse from CreateCrowd response
func CreateCreateCrowdResponse() (response *CreateCrowdResponse) {
	response = &CreateCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
