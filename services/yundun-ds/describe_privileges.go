package yundun_ds

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

// DescribePrivileges invokes the yundun_ds.DescribePrivileges API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeprivileges.html
func (client *Client) DescribePrivileges(request *DescribePrivilegesRequest) (response *DescribePrivilegesResponse, err error) {
	response = CreateDescribePrivilegesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePrivilegesWithChan invokes the yundun_ds.DescribePrivileges API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeprivileges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePrivilegesWithChan(request *DescribePrivilegesRequest) (<-chan *DescribePrivilegesResponse, <-chan error) {
	responseChan := make(chan *DescribePrivilegesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePrivileges(request)
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

// DescribePrivilegesWithCallback invokes the yundun_ds.DescribePrivileges API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeprivileges.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePrivilegesWithCallback(request *DescribePrivilegesRequest, callback func(response *DescribePrivilegesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePrivilegesResponse
		var err error
		defer close(result)
		response, err = client.DescribePrivileges(request)
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

// DescribePrivilegesRequest is the request struct for api DescribePrivileges
type DescribePrivilegesRequest struct {
	*requests.RpcRequest
	AccountId    requests.Integer `position:"Query" name:"AccountId"`
	UseAccountId requests.Integer `position:"Query" name:"UseAccountId"`
	DataTypeIds  string           `position:"Query" name:"DataTypeIds"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	FeatureType  requests.Integer `position:"Query" name:"FeatureType"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	Lang         string           `position:"Query" name:"Lang"`
	Key          string           `position:"Query" name:"Key"`
}

// DescribePrivilegesResponse is the response struct for api DescribePrivileges
type DescribePrivilegesResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	Items       []Privilege `json:"Items" xml:"Items"`
}

// CreateDescribePrivilegesRequest creates a request to invoke DescribePrivileges API
func CreateDescribePrivilegesRequest() (request *DescribePrivilegesRequest) {
	request = &DescribePrivilegesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribePrivileges", "sddp", "openAPI")
	return
}

// CreateDescribePrivilegesResponse creates a response to parse from DescribePrivileges response
func CreateDescribePrivilegesResponse() (response *DescribePrivilegesResponse) {
	response = &DescribePrivilegesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
