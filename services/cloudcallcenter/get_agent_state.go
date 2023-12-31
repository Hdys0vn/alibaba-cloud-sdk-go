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

// GetAgentState invokes the cloudcallcenter.GetAgentState API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getagentstate.html
func (client *Client) GetAgentState(request *GetAgentStateRequest) (response *GetAgentStateResponse, err error) {
	response = CreateGetAgentStateResponse()
	err = client.DoAction(request, response)
	return
}

// GetAgentStateWithChan invokes the cloudcallcenter.GetAgentState API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getagentstate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAgentStateWithChan(request *GetAgentStateRequest) (<-chan *GetAgentStateResponse, <-chan error) {
	responseChan := make(chan *GetAgentStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAgentState(request)
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

// GetAgentStateWithCallback invokes the cloudcallcenter.GetAgentState API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getagentstate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAgentStateWithCallback(request *GetAgentStateRequest, callback func(response *GetAgentStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAgentStateResponse
		var err error
		defer close(result)
		response, err = client.GetAgentState(request)
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

// GetAgentStateRequest is the request struct for api GetAgentState
type GetAgentStateRequest struct {
	*requests.RpcRequest
	AgentId    string `position:"Query" name:"AgentId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Dn         string `position:"Query" name:"Dn"`
}

// GetAgentStateResponse is the response struct for api GetAgentState
type GetAgentStateResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetAgentStateRequest creates a request to invoke GetAgentState API
func CreateGetAgentStateRequest() (request *GetAgentStateRequest) {
	request = &GetAgentStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetAgentState", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAgentStateResponse creates a response to parse from GetAgentState response
func CreateGetAgentStateResponse() (response *GetAgentStateResponse) {
	response = &GetAgentStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
