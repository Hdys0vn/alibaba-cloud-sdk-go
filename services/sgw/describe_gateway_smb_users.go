package sgw

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

// DescribeGatewaySMBUsers invokes the sgw.DescribeGatewaySMBUsers API synchronously
func (client *Client) DescribeGatewaySMBUsers(request *DescribeGatewaySMBUsersRequest) (response *DescribeGatewaySMBUsersResponse, err error) {
	response = CreateDescribeGatewaySMBUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewaySMBUsersWithChan invokes the sgw.DescribeGatewaySMBUsers API asynchronously
func (client *Client) DescribeGatewaySMBUsersWithChan(request *DescribeGatewaySMBUsersRequest) (<-chan *DescribeGatewaySMBUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewaySMBUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewaySMBUsers(request)
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

// DescribeGatewaySMBUsersWithCallback invokes the sgw.DescribeGatewaySMBUsers API asynchronously
func (client *Client) DescribeGatewaySMBUsersWithCallback(request *DescribeGatewaySMBUsersRequest, callback func(response *DescribeGatewaySMBUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewaySMBUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewaySMBUsers(request)
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

// DescribeGatewaySMBUsersRequest is the request struct for api DescribeGatewaySMBUsers
type DescribeGatewaySMBUsersRequest struct {
	*requests.RpcRequest
	Pattern       string           `position:"Query" name:"Pattern"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	GatewayId     string           `position:"Query" name:"GatewayId"`
}

// DescribeGatewaySMBUsersResponse is the response struct for api DescribeGatewaySMBUsers
type DescribeGatewaySMBUsersResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Users      Users  `json:"Users" xml:"Users"`
}

// CreateDescribeGatewaySMBUsersRequest creates a request to invoke DescribeGatewaySMBUsers API
func CreateDescribeGatewaySMBUsersRequest() (request *DescribeGatewaySMBUsersRequest) {
	request = &DescribeGatewaySMBUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewaySMBUsers", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewaySMBUsersResponse creates a response to parse from DescribeGatewaySMBUsers response
func CreateDescribeGatewaySMBUsersResponse() (response *DescribeGatewaySMBUsersResponse) {
	response = &DescribeGatewaySMBUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}