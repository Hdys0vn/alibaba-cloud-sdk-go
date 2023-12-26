package jarvis

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

// CreateAccessWhiteListGroup invokes the jarvis.CreateAccessWhiteListGroup API synchronously
// api document: https://help.aliyun.com/api/jarvis/createaccesswhitelistgroup.html
func (client *Client) CreateAccessWhiteListGroup(request *CreateAccessWhiteListGroupRequest) (response *CreateAccessWhiteListGroupResponse, err error) {
	response = CreateCreateAccessWhiteListGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAccessWhiteListGroupWithChan invokes the jarvis.CreateAccessWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createaccesswhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccessWhiteListGroupWithChan(request *CreateAccessWhiteListGroupRequest) (<-chan *CreateAccessWhiteListGroupResponse, <-chan error) {
	responseChan := make(chan *CreateAccessWhiteListGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAccessWhiteListGroup(request)
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

// CreateAccessWhiteListGroupWithCallback invokes the jarvis.CreateAccessWhiteListGroup API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createaccesswhitelistgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccessWhiteListGroupWithCallback(request *CreateAccessWhiteListGroupRequest, callback func(response *CreateAccessWhiteListGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAccessWhiteListGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateAccessWhiteListGroup(request)
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

// CreateAccessWhiteListGroupRequest is the request struct for api CreateAccessWhiteListGroup
type CreateAccessWhiteListGroupRequest struct {
	*requests.RpcRequest
	Note             string           `position:"Query" name:"Note"`
	ResourceOwnerId  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SrcIP            string           `position:"Query" name:"SrcIP"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	DstPort          requests.Integer `position:"Query" name:"DstPort"`
	InstanceIdList   string           `position:"Query" name:"InstanceIdList"`
	LiveTime         requests.Integer `position:"Query" name:"LiveTime"`
	ProductName      string           `position:"Query" name:"ProductName"`
	WhiteListType    requests.Integer `position:"Query" name:"WhiteListType"`
	InstanceInfoList string           `position:"Query" name:"InstanceInfoList"`
	Lang             string           `position:"Query" name:"Lang"`
	SourceCode       string           `position:"Query" name:"SourceCode"`
}

// CreateAccessWhiteListGroupResponse is the response struct for api CreateAccessWhiteListGroup
type CreateAccessWhiteListGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateCreateAccessWhiteListGroupRequest creates a request to invoke CreateAccessWhiteListGroup API
func CreateCreateAccessWhiteListGroupRequest() (request *CreateAccessWhiteListGroupRequest) {
	request = &CreateAccessWhiteListGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "CreateAccessWhiteListGroup", "jarvis", "openAPI")
	return
}

// CreateCreateAccessWhiteListGroupResponse creates a response to parse from CreateAccessWhiteListGroup response
func CreateCreateAccessWhiteListGroupResponse() (response *CreateAccessWhiteListGroupResponse) {
	response = &CreateAccessWhiteListGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
