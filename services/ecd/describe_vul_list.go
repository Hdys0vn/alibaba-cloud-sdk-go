package ecd

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

// DescribeVulList invokes the ecd.DescribeVulList API synchronously
func (client *Client) DescribeVulList(request *DescribeVulListRequest) (response *DescribeVulListResponse, err error) {
	response = CreateDescribeVulListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulListWithChan invokes the ecd.DescribeVulList API asynchronously
func (client *Client) DescribeVulListWithChan(request *DescribeVulListRequest) (<-chan *DescribeVulListResponse, <-chan error) {
	responseChan := make(chan *DescribeVulListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulList(request)
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

// DescribeVulListWithCallback invokes the ecd.DescribeVulList API asynchronously
func (client *Client) DescribeVulListWithCallback(request *DescribeVulListRequest, callback func(response *DescribeVulListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulListResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulList(request)
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

// DescribeVulListRequest is the request struct for api DescribeVulList
type DescribeVulListRequest struct {
	*requests.RpcRequest
	OfficeSiteId string           `position:"Query" name:"OfficeSiteId"`
	Dealed       string           `position:"Query" name:"Dealed"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	Type         string           `position:"Query" name:"Type"`
	AliasName    string           `position:"Query" name:"AliasName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
	Necessity    string           `position:"Query" name:"Necessity"`
}

// DescribeVulListResponse is the response struct for api DescribeVulList
type DescribeVulListResponse struct {
	*responses.BaseResponse
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	VulRecords  []VulRecord `json:"VulRecords" xml:"VulRecords"`
}

// CreateDescribeVulListRequest creates a request to invoke DescribeVulList API
func CreateDescribeVulListRequest() (request *DescribeVulListRequest) {
	request = &DescribeVulListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeVulList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVulListResponse creates a response to parse from DescribeVulList response
func CreateDescribeVulListResponse() (response *DescribeVulListResponse) {
	response = &DescribeVulListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
