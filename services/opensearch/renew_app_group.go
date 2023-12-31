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

// RenewAppGroup invokes the opensearch.RenewAppGroup API synchronously
func (client *Client) RenewAppGroup(request *RenewAppGroupRequest) (response *RenewAppGroupResponse, err error) {
	response = CreateRenewAppGroupResponse()
	err = client.DoAction(request, response)
	return
}

// RenewAppGroupWithChan invokes the opensearch.RenewAppGroup API asynchronously
func (client *Client) RenewAppGroupWithChan(request *RenewAppGroupRequest) (<-chan *RenewAppGroupResponse, <-chan error) {
	responseChan := make(chan *RenewAppGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewAppGroup(request)
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

// RenewAppGroupWithCallback invokes the opensearch.RenewAppGroup API asynchronously
func (client *Client) RenewAppGroupWithCallback(request *RenewAppGroupRequest, callback func(response *RenewAppGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewAppGroupResponse
		var err error
		defer close(result)
		response, err = client.RenewAppGroup(request)
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

// RenewAppGroupRequest is the request struct for api RenewAppGroup
type RenewAppGroupRequest struct {
	*requests.RoaRequest
	ClientToken      string `position:"Query" name:"clientToken"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// RenewAppGroupResponse is the response struct for api RenewAppGroup
type RenewAppGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    bool   `json:"result" xml:"result"`
}

// CreateRenewAppGroupRequest creates a request to invoke RenewAppGroup API
func CreateRenewAppGroupRequest() (request *RenewAppGroupRequest) {
	request = &RenewAppGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RenewAppGroup", "/v4/openapi/app-groups/[appGroupIdentity]/actions/renew", "", "")
	request.Method = requests.POST
	return
}

// CreateRenewAppGroupResponse creates a response to parse from RenewAppGroup response
func CreateRenewAppGroupResponse() (response *RenewAppGroupResponse) {
	response = &RenewAppGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
