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

// ModifyContactFlow invokes the cloudcallcenter.ModifyContactFlow API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflow.html
func (client *Client) ModifyContactFlow(request *ModifyContactFlowRequest) (response *ModifyContactFlowResponse, err error) {
	response = CreateModifyContactFlowResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyContactFlowWithChan invokes the cloudcallcenter.ModifyContactFlow API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyContactFlowWithChan(request *ModifyContactFlowRequest) (<-chan *ModifyContactFlowResponse, <-chan error) {
	responseChan := make(chan *ModifyContactFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyContactFlow(request)
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

// ModifyContactFlowWithCallback invokes the cloudcallcenter.ModifyContactFlow API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycontactflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyContactFlowWithCallback(request *ModifyContactFlowRequest, callback func(response *ModifyContactFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyContactFlowResponse
		var err error
		defer close(result)
		response, err = client.ModifyContactFlow(request)
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

// ModifyContactFlowRequest is the request struct for api ModifyContactFlow
type ModifyContactFlowRequest struct {
	*requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	Name          string `position:"Query" name:"Name"`
	Description   string `position:"Query" name:"Description"`
	Type          string `position:"Query" name:"Type"`
}

// ModifyContactFlowResponse is the response struct for api ModifyContactFlow
type ModifyContactFlowResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactFlow    ContactFlow `json:"ContactFlow" xml:"ContactFlow"`
}

// CreateModifyContactFlowRequest creates a request to invoke ModifyContactFlow API
func CreateModifyContactFlowRequest() (request *ModifyContactFlowRequest) {
	request = &ModifyContactFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyContactFlow", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyContactFlowResponse creates a response to parse from ModifyContactFlow response
func CreateModifyContactFlowResponse() (response *ModifyContactFlowResponse) {
	response = &ModifyContactFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
