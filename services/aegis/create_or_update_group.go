package aegis

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

// CreateOrUpdateGroup invokes the aegis.CreateOrUpdateGroup API synchronously
// api document: https://help.aliyun.com/api/aegis/createorupdategroup.html
func (client *Client) CreateOrUpdateGroup(request *CreateOrUpdateGroupRequest) (response *CreateOrUpdateGroupResponse, err error) {
	response = CreateCreateOrUpdateGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateGroupWithChan invokes the aegis.CreateOrUpdateGroup API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateGroupWithChan(request *CreateOrUpdateGroupRequest) (<-chan *CreateOrUpdateGroupResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateGroup(request)
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

// CreateOrUpdateGroupWithCallback invokes the aegis.CreateOrUpdateGroup API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdategroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateGroupWithCallback(request *CreateOrUpdateGroupRequest, callback func(response *CreateOrUpdateGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateGroup(request)
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

// CreateOrUpdateGroupRequest is the request struct for api CreateOrUpdateGroup
type CreateOrUpdateGroupRequest struct {
	*requests.RpcRequest
	RuleIds         string           `position:"Query" name:"RuleIds"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	MachineGroupIds string           `position:"Query" name:"MachineGroupIds"`
	Description     string           `position:"Query" name:"Description"`
	Id              requests.Integer `position:"Query" name:"Id"`
	Lang            string           `position:"Query" name:"Lang"`
	GroupName       string           `position:"Query" name:"GroupName"`
}

// CreateOrUpdateGroupResponse is the response struct for api CreateOrUpdateGroup
type CreateOrUpdateGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrUpdateGroupRequest creates a request to invoke CreateOrUpdateGroup API
func CreateCreateOrUpdateGroupRequest() (request *CreateOrUpdateGroupRequest) {
	request = &CreateOrUpdateGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "CreateOrUpdateGroup", "vipaegis", "openAPI")
	return
}

// CreateCreateOrUpdateGroupResponse creates a response to parse from CreateOrUpdateGroup response
func CreateCreateOrUpdateGroupResponse() (response *CreateOrUpdateGroupResponse) {
	response = &CreateOrUpdateGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
