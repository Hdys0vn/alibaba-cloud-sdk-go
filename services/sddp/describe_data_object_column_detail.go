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

// DescribeDataObjectColumnDetail invokes the sddp.DescribeDataObjectColumnDetail API synchronously
func (client *Client) DescribeDataObjectColumnDetail(request *DescribeDataObjectColumnDetailRequest) (response *DescribeDataObjectColumnDetailResponse, err error) {
	response = CreateDescribeDataObjectColumnDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataObjectColumnDetailWithChan invokes the sddp.DescribeDataObjectColumnDetail API asynchronously
func (client *Client) DescribeDataObjectColumnDetailWithChan(request *DescribeDataObjectColumnDetailRequest) (<-chan *DescribeDataObjectColumnDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeDataObjectColumnDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataObjectColumnDetail(request)
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

// DescribeDataObjectColumnDetailWithCallback invokes the sddp.DescribeDataObjectColumnDetail API asynchronously
func (client *Client) DescribeDataObjectColumnDetailWithCallback(request *DescribeDataObjectColumnDetailRequest, callback func(response *DescribeDataObjectColumnDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataObjectColumnDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataObjectColumnDetail(request)
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

// DescribeDataObjectColumnDetailRequest is the request struct for api DescribeDataObjectColumnDetail
type DescribeDataObjectColumnDetailRequest struct {
	*requests.RpcRequest
	ProductId   requests.Integer `position:"Query" name:"ProductId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Id          requests.Integer `position:"Query" name:"Id"`
	Lang        string           `position:"Query" name:"Lang"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	TemplateId  requests.Integer `position:"Query" name:"TemplateId"`
}

// DescribeDataObjectColumnDetailResponse is the response struct for api DescribeDataObjectColumnDetail
type DescribeDataObjectColumnDetailResponse struct {
	*responses.BaseResponse
	CurrentPage int    `json:"CurrentPage" xml:"CurrentPage"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	Items       []Rule `json:"Items" xml:"Items"`
}

// CreateDescribeDataObjectColumnDetailRequest creates a request to invoke DescribeDataObjectColumnDetail API
func CreateDescribeDataObjectColumnDetailRequest() (request *DescribeDataObjectColumnDetailRequest) {
	request = &DescribeDataObjectColumnDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "DescribeDataObjectColumnDetail", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDataObjectColumnDetailResponse creates a response to parse from DescribeDataObjectColumnDetail response
func CreateDescribeDataObjectColumnDetailResponse() (response *DescribeDataObjectColumnDetailResponse) {
	response = &DescribeDataObjectColumnDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
