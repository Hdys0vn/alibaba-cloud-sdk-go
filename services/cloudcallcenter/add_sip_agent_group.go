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

// AddSipAgentGroup invokes the cloudcallcenter.AddSipAgentGroup API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addsipagentgroup.html
func (client *Client) AddSipAgentGroup(request *AddSipAgentGroupRequest) (response *AddSipAgentGroupResponse, err error) {
	response = CreateAddSipAgentGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddSipAgentGroupWithChan invokes the cloudcallcenter.AddSipAgentGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addsipagentgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSipAgentGroupWithChan(request *AddSipAgentGroupRequest) (<-chan *AddSipAgentGroupResponse, <-chan error) {
	responseChan := make(chan *AddSipAgentGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSipAgentGroup(request)
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

// AddSipAgentGroupWithCallback invokes the cloudcallcenter.AddSipAgentGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addsipagentgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSipAgentGroupWithCallback(request *AddSipAgentGroupRequest, callback func(response *AddSipAgentGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSipAgentGroupResponse
		var err error
		defer close(result)
		response, err = client.AddSipAgentGroup(request)
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

// AddSipAgentGroupRequest is the request struct for api AddSipAgentGroup
type AddSipAgentGroupRequest struct {
	*requests.RpcRequest
	Ip            string `position:"Query" name:"Ip"`
	SendInterface string `position:"Query" name:"SendInterface"`
	Provider      string `position:"Query" name:"Provider"`
	Port          string `position:"Query" name:"Port"`
	Name          string `position:"Query" name:"Name"`
}

// AddSipAgentGroupResponse is the response struct for api AddSipAgentGroup
type AddSipAgentGroupResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Provider       string     `json:"Provider" xml:"Provider"`
	SipAgents      []SipAgent `json:"SipAgents" xml:"SipAgents"`
}

// CreateAddSipAgentGroupRequest creates a request to invoke AddSipAgentGroup API
func CreateAddSipAgentGroupRequest() (request *AddSipAgentGroupRequest) {
	request = &AddSipAgentGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AddSipAgentGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateAddSipAgentGroupResponse creates a response to parse from AddSipAgentGroup response
func CreateAddSipAgentGroupResponse() (response *AddSipAgentGroupResponse) {
	response = &AddSipAgentGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
