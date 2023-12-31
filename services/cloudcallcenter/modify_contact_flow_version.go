package cloudcallcenter

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

// ModifyContactFlowVersion invokes the cloudcallcenter.ModifyContactFlowVersion API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflowversion.html
func (client *Client) ModifyContactFlowVersion(request *ModifyContactFlowVersionRequest) (response *ModifyContactFlowVersionResponse, err error) {
	response = CreateModifyContactFlowVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyContactFlowVersionWithChan invokes the cloudcallcenter.ModifyContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyContactFlowVersionWithChan(request *ModifyContactFlowVersionRequest) (<-chan *ModifyContactFlowVersionResponse, <-chan error) {
	responseChan := make(chan *ModifyContactFlowVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyContactFlowVersion(request)
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

// ModifyContactFlowVersionWithCallback invokes the cloudcallcenter.ModifyContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyContactFlowVersionWithCallback(request *ModifyContactFlowVersionRequest, callback func(response *ModifyContactFlowVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyContactFlowVersionResponse
		var err error
		defer close(result)
		response, err = client.ModifyContactFlowVersion(request)
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

// ModifyContactFlowVersionRequest is the request struct for api ModifyContactFlowVersion
type ModifyContactFlowVersionRequest struct {
	*requests.RpcRequest
	Canvas               string `position:"Query" name:"Canvas"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ContactFlowVersionId string `position:"Query" name:"ContactFlowVersionId"`
	Content              string `position:"Query" name:"Content"`
}

// ModifyContactFlowVersionResponse is the response struct for api ModifyContactFlowVersion
type ModifyContactFlowVersionResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactFlow    ContactFlow `json:"ContactFlow" xml:"ContactFlow"`
}

// CreateModifyContactFlowVersionRequest creates a request to invoke ModifyContactFlowVersion API
func CreateModifyContactFlowVersionRequest() (request *ModifyContactFlowVersionRequest) {
	request = &ModifyContactFlowVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyContactFlowVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyContactFlowVersionResponse creates a response to parse from ModifyContactFlowVersion response
func CreateModifyContactFlowVersionResponse() (response *ModifyContactFlowVersionResponse) {
	response = &ModifyContactFlowVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
