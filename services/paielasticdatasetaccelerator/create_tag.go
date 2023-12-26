package paielasticdatasetaccelerator

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

// CreateTag invokes the paielasticdatasetaccelerator.CreateTag API synchronously
func (client *Client) CreateTag(request *CreateTagRequest) (response *CreateTagResponse, err error) {
	response = CreateCreateTagResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTagWithChan invokes the paielasticdatasetaccelerator.CreateTag API asynchronously
func (client *Client) CreateTagWithChan(request *CreateTagRequest) (<-chan *CreateTagResponse, <-chan error) {
	responseChan := make(chan *CreateTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTag(request)
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

// CreateTagWithCallback invokes the paielasticdatasetaccelerator.CreateTag API asynchronously
func (client *Client) CreateTagWithCallback(request *CreateTagRequest, callback func(response *CreateTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTagResponse
		var err error
		defer close(result)
		response, err = client.CreateTag(request)
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

// CreateTagRequest is the request struct for api CreateTag
type CreateTagRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateTagResponse is the response struct for api CreateTag
type CreateTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTagRequest creates a request to invoke CreateTag API
func CreateCreateTagRequest() (request *CreateTagRequest) {
	request = &CreateTagRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "CreateTag", "/api/v1/tags", "datasetacc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTagResponse creates a response to parse from CreateTag response
func CreateCreateTagResponse() (response *CreateTagResponse) {
	response = &CreateTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
