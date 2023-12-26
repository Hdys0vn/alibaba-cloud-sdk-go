package swas_open

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

// DescribeSecurityAgentStatus invokes the swas_open.DescribeSecurityAgentStatus API synchronously
func (client *Client) DescribeSecurityAgentStatus(request *DescribeSecurityAgentStatusRequest) (response *DescribeSecurityAgentStatusResponse, err error) {
	response = CreateDescribeSecurityAgentStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityAgentStatusWithChan invokes the swas_open.DescribeSecurityAgentStatus API asynchronously
func (client *Client) DescribeSecurityAgentStatusWithChan(request *DescribeSecurityAgentStatusRequest) (<-chan *DescribeSecurityAgentStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityAgentStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityAgentStatus(request)
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

// DescribeSecurityAgentStatusWithCallback invokes the swas_open.DescribeSecurityAgentStatus API asynchronously
func (client *Client) DescribeSecurityAgentStatusWithCallback(request *DescribeSecurityAgentStatusRequest, callback func(response *DescribeSecurityAgentStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityAgentStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityAgentStatus(request)
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

// DescribeSecurityAgentStatusRequest is the request struct for api DescribeSecurityAgentStatus
type DescribeSecurityAgentStatusRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// DescribeSecurityAgentStatusResponse is the response struct for api DescribeSecurityAgentStatus
type DescribeSecurityAgentStatusResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ClientStatus string `json:"ClientStatus" xml:"ClientStatus"`
}

// CreateDescribeSecurityAgentStatusRequest creates a request to invoke DescribeSecurityAgentStatus API
func CreateDescribeSecurityAgentStatusRequest() (request *DescribeSecurityAgentStatusRequest) {
	request = &DescribeSecurityAgentStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeSecurityAgentStatus", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecurityAgentStatusResponse creates a response to parse from DescribeSecurityAgentStatus response
func CreateDescribeSecurityAgentStatusResponse() (response *DescribeSecurityAgentStatusResponse) {
	response = &DescribeSecurityAgentStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
