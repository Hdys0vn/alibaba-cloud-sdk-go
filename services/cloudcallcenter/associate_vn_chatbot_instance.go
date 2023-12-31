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

// AssociateVnChatbotInstance invokes the cloudcallcenter.AssociateVnChatbotInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/associatevnchatbotinstance.html
func (client *Client) AssociateVnChatbotInstance(request *AssociateVnChatbotInstanceRequest) (response *AssociateVnChatbotInstanceResponse, err error) {
	response = CreateAssociateVnChatbotInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateVnChatbotInstanceWithChan invokes the cloudcallcenter.AssociateVnChatbotInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/associatevnchatbotinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateVnChatbotInstanceWithChan(request *AssociateVnChatbotInstanceRequest) (<-chan *AssociateVnChatbotInstanceResponse, <-chan error) {
	responseChan := make(chan *AssociateVnChatbotInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateVnChatbotInstance(request)
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

// AssociateVnChatbotInstanceWithCallback invokes the cloudcallcenter.AssociateVnChatbotInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/associatevnchatbotinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateVnChatbotInstanceWithCallback(request *AssociateVnChatbotInstanceRequest, callback func(response *AssociateVnChatbotInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateVnChatbotInstanceResponse
		var err error
		defer close(result)
		response, err = client.AssociateVnChatbotInstance(request)
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

// AssociateVnChatbotInstanceRequest is the request struct for api AssociateVnChatbotInstance
type AssociateVnChatbotInstanceRequest struct {
	*requests.RpcRequest
	InstanceId        string `position:"Query" name:"InstanceId"`
	ChatbotInstanceId string `position:"Query" name:"ChatbotInstanceId"`
	ChatbotName       string `position:"Query" name:"ChatbotName"`
}

// AssociateVnChatbotInstanceResponse is the response struct for api AssociateVnChatbotInstance
type AssociateVnChatbotInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateVnChatbotInstanceRequest creates a request to invoke AssociateVnChatbotInstance API
func CreateAssociateVnChatbotInstanceRequest() (request *AssociateVnChatbotInstanceRequest) {
	request = &AssociateVnChatbotInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AssociateVnChatbotInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateAssociateVnChatbotInstanceResponse creates a response to parse from AssociateVnChatbotInstance response
func CreateAssociateVnChatbotInstanceResponse() (response *AssociateVnChatbotInstanceResponse) {
	response = &AssociateVnChatbotInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
