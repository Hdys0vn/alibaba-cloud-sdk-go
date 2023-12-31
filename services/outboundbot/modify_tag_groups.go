package outboundbot

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

// ModifyTagGroups invokes the outboundbot.ModifyTagGroups API synchronously
func (client *Client) ModifyTagGroups(request *ModifyTagGroupsRequest) (response *ModifyTagGroupsResponse, err error) {
	response = CreateModifyTagGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTagGroupsWithChan invokes the outboundbot.ModifyTagGroups API asynchronously
func (client *Client) ModifyTagGroupsWithChan(request *ModifyTagGroupsRequest) (<-chan *ModifyTagGroupsResponse, <-chan error) {
	responseChan := make(chan *ModifyTagGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTagGroups(request)
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

// ModifyTagGroupsWithCallback invokes the outboundbot.ModifyTagGroups API asynchronously
func (client *Client) ModifyTagGroupsWithCallback(request *ModifyTagGroupsRequest, callback func(response *ModifyTagGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTagGroupsResponse
		var err error
		defer close(result)
		response, err = client.ModifyTagGroups(request)
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

// ModifyTagGroupsRequest is the request struct for api ModifyTagGroups
type ModifyTagGroupsRequest struct {
	*requests.RpcRequest
	Tags       string `position:"Query" name:"Tags"`
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	TagGroups  string `position:"Query" name:"TagGroups"`
}

// ModifyTagGroupsResponse is the response struct for api ModifyTagGroups
type ModifyTagGroupsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	TagGroups      []TagGroup `json:"TagGroups" xml:"TagGroups"`
	Tags           []Tag      `json:"Tags" xml:"Tags"`
}

// CreateModifyTagGroupsRequest creates a request to invoke ModifyTagGroups API
func CreateModifyTagGroupsRequest() (request *ModifyTagGroupsRequest) {
	request = &ModifyTagGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyTagGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyTagGroupsResponse creates a response to parse from ModifyTagGroups response
func CreateModifyTagGroupsResponse() (response *ModifyTagGroupsResponse) {
	response = &ModifyTagGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
