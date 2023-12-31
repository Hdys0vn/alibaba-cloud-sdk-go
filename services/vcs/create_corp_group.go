package vcs

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

// CreateCorpGroup invokes the vcs.CreateCorpGroup API synchronously
func (client *Client) CreateCorpGroup(request *CreateCorpGroupRequest) (response *CreateCorpGroupResponse, err error) {
	response = CreateCreateCorpGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCorpGroupWithChan invokes the vcs.CreateCorpGroup API asynchronously
func (client *Client) CreateCorpGroupWithChan(request *CreateCorpGroupRequest) (<-chan *CreateCorpGroupResponse, <-chan error) {
	responseChan := make(chan *CreateCorpGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCorpGroup(request)
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

// CreateCorpGroupWithCallback invokes the vcs.CreateCorpGroup API asynchronously
func (client *Client) CreateCorpGroupWithCallback(request *CreateCorpGroupRequest, callback func(response *CreateCorpGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCorpGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateCorpGroup(request)
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

// CreateCorpGroupRequest is the request struct for api CreateCorpGroup
type CreateCorpGroupRequest struct {
	*requests.RpcRequest
	CorpId      string `position:"Body" name:"CorpId"`
	ClientToken string `position:"Body" name:"ClientToken"`
	GroupId     string `position:"Body" name:"GroupId"`
}

// CreateCorpGroupResponse is the response struct for api CreateCorpGroup
type CreateCorpGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateCorpGroupRequest creates a request to invoke CreateCorpGroup API
func CreateCreateCorpGroupRequest() (request *CreateCorpGroupRequest) {
	request = &CreateCorpGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "CreateCorpGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCorpGroupResponse creates a response to parse from CreateCorpGroup response
func CreateCreateCorpGroupResponse() (response *CreateCorpGroupResponse) {
	response = &CreateCorpGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
