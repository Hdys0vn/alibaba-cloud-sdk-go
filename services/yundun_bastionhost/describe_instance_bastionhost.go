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

// DescribeInstanceBastionhost invokes the yundun_bastionhost.DescribeInstanceBastionhost API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstancebastionhost.html
func (client *Client) DescribeInstanceBastionhost(request *DescribeInstanceBastionhostRequest) (response *DescribeInstanceBastionhostResponse, err error) {
	response = CreateDescribeInstanceBastionhostResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceBastionhostWithChan invokes the yundun_bastionhost.DescribeInstanceBastionhost API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstancebastionhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceBastionhostWithChan(request *DescribeInstanceBastionhostRequest) (<-chan *DescribeInstanceBastionhostResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceBastionhostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceBastionhost(request)
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

// DescribeInstanceBastionhostWithCallback invokes the yundun_bastionhost.DescribeInstanceBastionhost API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describeinstancebastionhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceBastionhostWithCallback(request *DescribeInstanceBastionhostRequest, callback func(response *DescribeInstanceBastionhostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceBastionhostResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceBastionhost(request)
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

// DescribeInstanceBastionhostRequest is the request struct for api DescribeInstanceBastionhost
type DescribeInstanceBastionhostRequest struct {
	*requests.RpcRequest
	ResourceGroupId string                            `position:"Query" name:"ResourceGroupId"`
	SourceIp        string                            `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer                  `position:"Query" name:"PageSize"`
	Tag             *[]DescribeInstanceBastionhostTag `position:"Query" name:"Tag"  type:"Repeated"`
	Lang            string                            `position:"Query" name:"Lang"`
	InstanceStatus  string                            `position:"Query" name:"InstanceStatus"`
	CurrentPage     requests.Integer                  `position:"Query" name:"CurrentPage"`
	InstanceId      *[]string                         `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// DescribeInstanceBastionhostTag is a repeated param struct in DescribeInstanceBastionhostRequest
type DescribeInstanceBastionhostTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeInstanceBastionhostResponse is the response struct for api DescribeInstanceBastionhost
type DescribeInstanceBastionhostResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TotalCount int64      `json:"TotalCount" xml:"TotalCount"`
	Instances  []Instance `json:"Instances" xml:"Instances"`
}

// CreateDescribeInstanceBastionhostRequest creates a request to invoke DescribeInstanceBastionhost API
func CreateDescribeInstanceBastionhostRequest() (request *DescribeInstanceBastionhostRequest) {
	request = &DescribeInstanceBastionhostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "DescribeInstanceBastionhost", "bastionhost", "openAPI")
	return
}

// CreateDescribeInstanceBastionhostResponse creates a response to parse from DescribeInstanceBastionhost response
func CreateDescribeInstanceBastionhostResponse() (response *DescribeInstanceBastionhostResponse) {
	response = &DescribeInstanceBastionhostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}