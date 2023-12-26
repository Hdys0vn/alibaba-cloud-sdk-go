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

// ListRealTimeAgent invokes the cloudcallcenter.ListRealTimeAgent API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrealtimeagent.html
func (client *Client) ListRealTimeAgent(request *ListRealTimeAgentRequest) (response *ListRealTimeAgentResponse, err error) {
	response = CreateListRealTimeAgentResponse()
	err = client.DoAction(request, response)
	return
}

// ListRealTimeAgentWithChan invokes the cloudcallcenter.ListRealTimeAgent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrealtimeagent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRealTimeAgentWithChan(request *ListRealTimeAgentRequest) (<-chan *ListRealTimeAgentResponse, <-chan error) {
	responseChan := make(chan *ListRealTimeAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRealTimeAgent(request)
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

// ListRealTimeAgentWithCallback invokes the cloudcallcenter.ListRealTimeAgent API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listrealtimeagent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRealTimeAgentWithCallback(request *ListRealTimeAgentRequest, callback func(response *ListRealTimeAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRealTimeAgentResponse
		var err error
		defer close(result)
		response, err = client.ListRealTimeAgent(request)
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

// ListRealTimeAgentRequest is the request struct for api ListRealTimeAgent
type ListRealTimeAgentRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListRealTimeAgentResponse is the response struct for api ListRealTimeAgent
type ListRealTimeAgentResponse struct {
	*responses.BaseResponse
	RequestId      string                  `json:"RequestId" xml:"RequestId"`
	Success        bool                    `json:"Success" xml:"Success"`
	Code           string                  `json:"Code" xml:"Code"`
	Message        string                  `json:"Message" xml:"Message"`
	HttpStatusCode int                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListRealTimeAgent `json:"Data" xml:"Data"`
}

// CreateListRealTimeAgentRequest creates a request to invoke ListRealTimeAgent API
func CreateListRealTimeAgentRequest() (request *ListRealTimeAgentRequest) {
	request = &ListRealTimeAgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListRealTimeAgent", "", "")
	request.Method = requests.POST
	return
}

// CreateListRealTimeAgentResponse creates a response to parse from ListRealTimeAgent response
func CreateListRealTimeAgentResponse() (response *ListRealTimeAgentResponse) {
	response = &ListRealTimeAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
