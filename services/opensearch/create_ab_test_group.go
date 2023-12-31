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

// CreateABTestGroup invokes the opensearch.CreateABTestGroup API synchronously
func (client *Client) CreateABTestGroup(request *CreateABTestGroupRequest) (response *CreateABTestGroupResponse, err error) {
	response = CreateCreateABTestGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateABTestGroupWithChan invokes the opensearch.CreateABTestGroup API asynchronously
func (client *Client) CreateABTestGroupWithChan(request *CreateABTestGroupRequest) (<-chan *CreateABTestGroupResponse, <-chan error) {
	responseChan := make(chan *CreateABTestGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateABTestGroup(request)
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

// CreateABTestGroupWithCallback invokes the opensearch.CreateABTestGroup API asynchronously
func (client *Client) CreateABTestGroupWithCallback(request *CreateABTestGroupRequest, callback func(response *CreateABTestGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateABTestGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateABTestGroup(request)
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

// CreateABTestGroupRequest is the request struct for api CreateABTestGroup
type CreateABTestGroupRequest struct {
	*requests.RoaRequest
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// CreateABTestGroupResponse is the response struct for api CreateABTestGroup
type CreateABTestGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateCreateABTestGroupRequest creates a request to invoke CreateABTestGroup API
func CreateCreateABTestGroupRequest() (request *CreateABTestGroupRequest) {
	request = &CreateABTestGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateABTestGroup", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]/groups", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateABTestGroupResponse creates a response to parse from CreateABTestGroup response
func CreateCreateABTestGroupResponse() (response *CreateABTestGroupResponse) {
	response = &CreateABTestGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
