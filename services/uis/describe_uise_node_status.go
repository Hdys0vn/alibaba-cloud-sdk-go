package uis

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

// DescribeUiseNodeStatus invokes the uis.DescribeUiseNodeStatus API synchronously
// api document: https://help.aliyun.com/api/uis/describeuisenodestatus.html
func (client *Client) DescribeUiseNodeStatus(request *DescribeUiseNodeStatusRequest) (response *DescribeUiseNodeStatusResponse, err error) {
	response = CreateDescribeUiseNodeStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUiseNodeStatusWithChan invokes the uis.DescribeUiseNodeStatus API asynchronously
// api document: https://help.aliyun.com/api/uis/describeuisenodestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUiseNodeStatusWithChan(request *DescribeUiseNodeStatusRequest) (<-chan *DescribeUiseNodeStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUiseNodeStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUiseNodeStatus(request)
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

// DescribeUiseNodeStatusWithCallback invokes the uis.DescribeUiseNodeStatus API asynchronously
// api document: https://help.aliyun.com/api/uis/describeuisenodestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUiseNodeStatusWithCallback(request *DescribeUiseNodeStatusRequest, callback func(response *DescribeUiseNodeStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUiseNodeStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUiseNodeStatus(request)
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

// DescribeUiseNodeStatusRequest is the request struct for api DescribeUiseNodeStatus
type DescribeUiseNodeStatusRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UisNodeId            string           `position:"Query" name:"UisNodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Ip                   string           `position:"Query" name:"Ip"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeUiseNodeStatusResponse is the response struct for api DescribeUiseNodeStatus
type DescribeUiseNodeStatusResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	IpStatusList IpStatusList `json:"IpStatusList" xml:"IpStatusList"`
}

// CreateDescribeUiseNodeStatusRequest creates a request to invoke DescribeUiseNodeStatus API
func CreateDescribeUiseNodeStatusRequest() (request *DescribeUiseNodeStatusRequest) {
	request = &DescribeUiseNodeStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "DescribeUiseNodeStatus", "uis", "openAPI")
	return
}

// CreateDescribeUiseNodeStatusResponse creates a response to parse from DescribeUiseNodeStatus response
func CreateDescribeUiseNodeStatusResponse() (response *DescribeUiseNodeStatusResponse) {
	response = &DescribeUiseNodeStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
