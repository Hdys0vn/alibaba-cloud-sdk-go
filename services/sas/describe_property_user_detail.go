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

// DescribePropertyUserDetail invokes the sas.DescribePropertyUserDetail API synchronously
func (client *Client) DescribePropertyUserDetail(request *DescribePropertyUserDetailRequest) (response *DescribePropertyUserDetailResponse, err error) {
	response = CreateDescribePropertyUserDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertyUserDetailWithChan invokes the sas.DescribePropertyUserDetail API asynchronously
func (client *Client) DescribePropertyUserDetailWithChan(request *DescribePropertyUserDetailRequest) (<-chan *DescribePropertyUserDetailResponse, <-chan error) {
	responseChan := make(chan *DescribePropertyUserDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertyUserDetail(request)
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

// DescribePropertyUserDetailWithCallback invokes the sas.DescribePropertyUserDetail API asynchronously
func (client *Client) DescribePropertyUserDetailWithCallback(request *DescribePropertyUserDetailRequest, callback func(response *DescribePropertyUserDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertyUserDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertyUserDetail(request)
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

// DescribePropertyUserDetailRequest is the request struct for api DescribePropertyUserDetail
type DescribePropertyUserDetailRequest struct {
	*requests.RpcRequest
	Remark             string           `position:"Query" name:"Remark"`
	Uuid               string           `position:"Query" name:"Uuid"`
	SourceIp           string           `position:"Query" name:"SourceIp"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	LastLoginTimeStart requests.Integer `position:"Query" name:"LastLoginTimeStart"`
	CurrentPage        requests.Integer `position:"Query" name:"CurrentPage"`
	LastLoginTimeEnd   requests.Integer `position:"Query" name:"LastLoginTimeEnd"`
	Extend             string           `position:"Query" name:"Extend"`
	IsRoot             string           `position:"Query" name:"IsRoot"`
	User               string           `position:"Query" name:"User"`
}

// DescribePropertyUserDetailResponse is the response struct for api DescribePropertyUserDetail
type DescribePropertyUserDetailResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo       `json:"PageInfo" xml:"PageInfo"`
	Propertys []PropertyUser `json:"Propertys" xml:"Propertys"`
}

// CreateDescribePropertyUserDetailRequest creates a request to invoke DescribePropertyUserDetail API
func CreateDescribePropertyUserDetailRequest() (request *DescribePropertyUserDetailRequest) {
	request = &DescribePropertyUserDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertyUserDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePropertyUserDetailResponse creates a response to parse from DescribePropertyUserDetail response
func CreateDescribePropertyUserDetailResponse() (response *DescribePropertyUserDetailResponse) {
	response = &DescribePropertyUserDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
