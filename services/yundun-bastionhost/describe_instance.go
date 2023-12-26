package yundun_bastionhost

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

// DescribeInstance invokes the yundun_bastionhost.DescribeInstance API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstance.html
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
	response = CreateDescribeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceWithChan invokes the yundun_bastionhost.DescribeInstance API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceWithChan(request *DescribeInstanceRequest) (<-chan *DescribeInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstance(request)
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

// DescribeInstanceWithCallback invokes the yundun_bastionhost.DescribeInstance API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceWithCallback(request *DescribeInstanceRequest, callback func(response *DescribeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstance(request)
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

// DescribeInstanceRequest is the request struct for api DescribeInstance
type DescribeInstanceRequest struct {
	*requests.RpcRequest
	RegionNo string           `position:"Query" name:"regionNo"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	PageNo   requests.Integer `position:"Query" name:"PageNo"`
	PageSize requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeInstanceResponse is the response struct for api DescribeInstance
type DescribeInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount int64      `json:"TotalCount" xml:"TotalCount"`
	Instances  []Instance `json:"Instances" xml:"Instances"`
}

// CreateDescribeInstanceRequest creates a request to invoke DescribeInstance API
func CreateDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "DescribeInstance", "bastionhost", "openAPI")
	return
}

// CreateDescribeInstanceResponse creates a response to parse from DescribeInstance response
func CreateDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
