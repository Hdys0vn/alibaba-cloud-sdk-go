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

// ModifyAppGroup invokes the opensearch.ModifyAppGroup API synchronously
func (client *Client) ModifyAppGroup(request *ModifyAppGroupRequest) (response *ModifyAppGroupResponse, err error) {
	response = CreateModifyAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAppGroupWithChan invokes the opensearch.ModifyAppGroup API asynchronously
func (client *Client) ModifyAppGroupWithChan(request *ModifyAppGroupRequest) (<-chan *ModifyAppGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAppGroup(request)
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

// ModifyAppGroupWithCallback invokes the opensearch.ModifyAppGroup API asynchronously
func (client *Client) ModifyAppGroupWithCallback(request *ModifyAppGroupRequest, callback func(response *ModifyAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAppGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyAppGroup(request)
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

// ModifyAppGroupRequest is the request struct for api ModifyAppGroup
type ModifyAppGroupRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// ModifyAppGroupResponse is the response struct for api ModifyAppGroup
type ModifyAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    ResultInModifyAppGroup `json:"result" xml:"result"`
}

// CreateModifyAppGroupRequest creates a request to invoke ModifyAppGroup API
func CreateModifyAppGroupRequest() (request *ModifyAppGroupRequest) {
	request = &ModifyAppGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ModifyAppGroup", "/v4/openapi/app-groups/[appGroupIdentity]", "", "")
	request.Method = requests.PUT
	return
}

// CreateModifyAppGroupResponse creates a response to parse from ModifyAppGroup response
func CreateModifyAppGroupResponse() (response *ModifyAppGroupResponse) {
	response = &ModifyAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
