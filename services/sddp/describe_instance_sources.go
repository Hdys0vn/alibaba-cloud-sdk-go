package sddp

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

// DescribeInstanceSources invokes the sddp.DescribeInstanceSources API synchronously
func (client *Client) DescribeInstanceSources(request *DescribeInstanceSourcesRequest) (response *DescribeInstanceSourcesResponse, err error) {
	response = CreateDescribeInstanceSourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSourcesWithChan invokes the sddp.DescribeInstanceSources API asynchronously
func (client *Client) DescribeInstanceSourcesWithChan(request *DescribeInstanceSourcesRequest) (<-chan *DescribeInstanceSourcesResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSources(request)
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

// DescribeInstanceSourcesWithCallback invokes the sddp.DescribeInstanceSources API asynchronously
func (client *Client) DescribeInstanceSourcesWithCallback(request *DescribeInstanceSourcesRequest, callback func(response *DescribeInstanceSourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSourcesResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSources(request)
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

// DescribeInstanceSourcesRequest is the request struct for api DescribeInstanceSources
type DescribeInstanceSourcesRequest struct {
	*requests.RpcRequest
	ProductCode     string           `position:"Query" name:"ProductCode"`
	ProductId       requests.Integer `position:"Query" name:"ProductId"`
	SearchKey       string           `position:"Query" name:"SearchKey"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	SearchType      string           `position:"Query" name:"SearchType"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	ServiceRegionId string           `position:"Query" name:"ServiceRegionId"`
	EngineType      string           `position:"Query" name:"EngineType"`
	AuditStatus     requests.Integer `position:"Query" name:"AuditStatus"`
	AuthStatus      requests.Integer `position:"Query" name:"AuthStatus"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	Authed          requests.Boolean `position:"Query" name:"Authed"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	DbName          string           `position:"Query" name:"DbName"`
}

// DescribeInstanceSourcesResponse is the response struct for api DescribeInstanceSources
type DescribeInstanceSourcesResponse struct {
	*responses.BaseResponse
	CurrentPage int              `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string           `json:"RequestId" xml:"RequestId"`
	PageSize    int              `json:"PageSize" xml:"PageSize"`
	TotalCount  int              `json:"TotalCount" xml:"TotalCount"`
	Items       []InstanceSource `json:"Items" xml:"Items"`
}

// CreateDescribeInstanceSourcesRequest creates a request to invoke DescribeInstanceSources API
func CreateDescribeInstanceSourcesRequest() (request *DescribeInstanceSourcesRequest) {
	request = &DescribeInstanceSourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeInstanceSources", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceSourcesResponse creates a response to parse from DescribeInstanceSources response
func CreateDescribeInstanceSourcesResponse() (response *DescribeInstanceSourcesResponse) {
	response = &DescribeInstanceSourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
