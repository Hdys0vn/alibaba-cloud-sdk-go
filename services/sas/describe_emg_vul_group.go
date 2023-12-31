package sas

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

// DescribeEmgVulGroup invokes the sas.DescribeEmgVulGroup API synchronously
func (client *Client) DescribeEmgVulGroup(request *DescribeEmgVulGroupRequest) (response *DescribeEmgVulGroupResponse, err error) {
	response = CreateDescribeEmgVulGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEmgVulGroupWithChan invokes the sas.DescribeEmgVulGroup API asynchronously
func (client *Client) DescribeEmgVulGroupWithChan(request *DescribeEmgVulGroupRequest) (<-chan *DescribeEmgVulGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeEmgVulGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEmgVulGroup(request)
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

// DescribeEmgVulGroupWithCallback invokes the sas.DescribeEmgVulGroup API asynchronously
func (client *Client) DescribeEmgVulGroupWithCallback(request *DescribeEmgVulGroupRequest, callback func(response *DescribeEmgVulGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEmgVulGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeEmgVulGroup(request)
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

// DescribeEmgVulGroupRequest is the request struct for api DescribeEmgVulGroup
type DescribeEmgVulGroupRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeEmgVulGroupResponse is the response struct for api DescribeEmgVulGroup
type DescribeEmgVulGroupResponse struct {
	*responses.BaseResponse
	RequestId       string        `json:"RequestId" xml:"RequestId"`
	TotalCount      int           `json:"TotalCount" xml:"TotalCount"`
	EmgVulGroupList []EmgVulGroup `json:"EmgVulGroupList" xml:"EmgVulGroupList"`
}

// CreateDescribeEmgVulGroupRequest creates a request to invoke DescribeEmgVulGroup API
func CreateDescribeEmgVulGroupRequest() (request *DescribeEmgVulGroupRequest) {
	request = &DescribeEmgVulGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeEmgVulGroup", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEmgVulGroupResponse creates a response to parse from DescribeEmgVulGroup response
func CreateDescribeEmgVulGroupResponse() (response *DescribeEmgVulGroupResponse) {
	response = &DescribeEmgVulGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
