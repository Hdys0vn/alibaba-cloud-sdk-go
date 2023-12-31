package eas

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

// CreateResourceInstances invokes the eas.CreateResourceInstances API synchronously
func (client *Client) CreateResourceInstances(request *CreateResourceInstancesRequest) (response *CreateResourceInstancesResponse, err error) {
	response = CreateCreateResourceInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourceInstancesWithChan invokes the eas.CreateResourceInstances API asynchronously
func (client *Client) CreateResourceInstancesWithChan(request *CreateResourceInstancesRequest) (<-chan *CreateResourceInstancesResponse, <-chan error) {
	responseChan := make(chan *CreateResourceInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResourceInstances(request)
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

// CreateResourceInstancesWithCallback invokes the eas.CreateResourceInstances API asynchronously
func (client *Client) CreateResourceInstancesWithCallback(request *CreateResourceInstancesRequest, callback func(response *CreateResourceInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourceInstancesResponse
		var err error
		defer close(result)
		response, err = client.CreateResourceInstances(request)
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

// CreateResourceInstancesRequest is the request struct for api CreateResourceInstances
type CreateResourceInstancesRequest struct {
	*requests.RoaRequest
	ResourceId string `position:"Path" name:"ResourceId"`
	ClusterId  string `position:"Path" name:"ClusterId"`
	Body       string `position:"Body" name:"body"`
}

// CreateResourceInstancesResponse is the response struct for api CreateResourceInstances
type CreateResourceInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateResourceInstancesRequest creates a request to invoke CreateResourceInstances API
func CreateCreateResourceInstancesRequest() (request *CreateResourceInstancesRequest) {
	request = &CreateResourceInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "CreateResourceInstances", "/api/v2/resources/[ClusterId]/[ResourceId]/instances", "eas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateResourceInstancesResponse creates a response to parse from CreateResourceInstances response
func CreateCreateResourceInstancesResponse() (response *CreateResourceInstancesResponse) {
	response = &CreateResourceInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
