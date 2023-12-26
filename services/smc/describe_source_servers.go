package smc

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

// DescribeSourceServers invokes the smc.DescribeSourceServers API synchronously
func (client *Client) DescribeSourceServers(request *DescribeSourceServersRequest) (response *DescribeSourceServersResponse, err error) {
	response = CreateDescribeSourceServersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSourceServersWithChan invokes the smc.DescribeSourceServers API asynchronously
func (client *Client) DescribeSourceServersWithChan(request *DescribeSourceServersRequest) (<-chan *DescribeSourceServersResponse, <-chan error) {
	responseChan := make(chan *DescribeSourceServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSourceServers(request)
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

// DescribeSourceServersWithCallback invokes the smc.DescribeSourceServers API asynchronously
func (client *Client) DescribeSourceServersWithCallback(request *DescribeSourceServersRequest, callback func(response *DescribeSourceServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSourceServersResponse
		var err error
		defer close(result)
		response, err = client.DescribeSourceServers(request)
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

// DescribeSourceServersRequest is the request struct for api DescribeSourceServers
type DescribeSourceServersRequest struct {
	*requests.RpcRequest
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	JobId                string           `position:"Query" name:"JobId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	State                string           `position:"Query" name:"State"`
	SourceId             *[]string        `position:"Query" name:"SourceId"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// DescribeSourceServersResponse is the response struct for api DescribeSourceServers
type DescribeSourceServersResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	TotalCount    int           `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	SourceServers SourceServers `json:"SourceServers" xml:"SourceServers"`
}

// CreateDescribeSourceServersRequest creates a request to invoke DescribeSourceServers API
func CreateDescribeSourceServersRequest() (request *DescribeSourceServersRequest) {
	request = &DescribeSourceServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("smc", "2019-06-01", "DescribeSourceServers", "smc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSourceServersResponse creates a response to parse from DescribeSourceServers response
func CreateDescribeSourceServersResponse() (response *DescribeSourceServersResponse) {
	response = &DescribeSourceServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
