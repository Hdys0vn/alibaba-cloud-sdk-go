package opensearch

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

// CreateAppGroup invokes the opensearch.CreateAppGroup API synchronously
func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (response *CreateAppGroupResponse, err error) {
	response = CreateCreateAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAppGroupWithChan invokes the opensearch.CreateAppGroup API asynchronously
func (client *Client) CreateAppGroupWithChan(request *CreateAppGroupRequest) (<-chan *CreateAppGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAppGroup(request)
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

// CreateAppGroupWithCallback invokes the opensearch.CreateAppGroup API asynchronously
func (client *Client) CreateAppGroupWithCallback(request *CreateAppGroupRequest, callback func(response *CreateAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAppGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAppGroup(request)
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

// CreateAppGroupRequest is the request struct for api CreateAppGroup
type CreateAppGroupRequest struct {
	*requests.RoaRequest
}

// CreateAppGroupResponse is the response struct for api CreateAppGroup
type CreateAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    ResultInCreateAppGroup `json:"result" xml:"result"`
}

// CreateCreateAppGroupRequest creates a request to invoke CreateAppGroup API
func CreateCreateAppGroupRequest() (request *CreateAppGroupRequest) {
	request = &CreateAppGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateAppGroup", "/v4/openapi/app-groups", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAppGroupResponse creates a response to parse from CreateAppGroup response
func CreateCreateAppGroupResponse() (response *CreateAppGroupResponse) {
	response = &CreateAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
