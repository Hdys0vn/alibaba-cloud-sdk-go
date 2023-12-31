package ots

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

// BindInstance2Vpc invokes the ots.BindInstance2Vpc API synchronously
// api document: https://help.aliyun.com/api/ots/bindinstance2vpc.html
func (client *Client) BindInstance2Vpc(request *BindInstance2VpcRequest) (response *BindInstance2VpcResponse, err error) {
	response = CreateBindInstance2VpcResponse()
	err = client.DoAction(request, response)
	return
}

// BindInstance2VpcWithChan invokes the ots.BindInstance2Vpc API asynchronously
// api document: https://help.aliyun.com/api/ots/bindinstance2vpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindInstance2VpcWithChan(request *BindInstance2VpcRequest) (<-chan *BindInstance2VpcResponse, <-chan error) {
	responseChan := make(chan *BindInstance2VpcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindInstance2Vpc(request)
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

// BindInstance2VpcWithCallback invokes the ots.BindInstance2Vpc API asynchronously
// api document: https://help.aliyun.com/api/ots/bindinstance2vpc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindInstance2VpcWithCallback(request *BindInstance2VpcRequest, callback func(response *BindInstance2VpcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindInstance2VpcResponse
		var err error
		defer close(result)
		response, err = client.BindInstance2Vpc(request)
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

// BindInstance2VpcRequest is the request struct for api BindInstance2Vpc
type BindInstance2VpcRequest struct {
	*requests.RpcRequest
	AccessKeyId     string           `position:"Query" name:"access_key_id"`
	InstanceVpcName string           `position:"Query" name:"InstanceVpcName"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceName    string           `position:"Query" name:"InstanceName"`
	VpcId           string           `position:"Query" name:"VpcId"`
	VirtualSwitchId string           `position:"Query" name:"VirtualSwitchId"`
	RegionNo        string           `position:"Query" name:"RegionNo"`
	Network         string           `position:"Query" name:"Network"`
}

// BindInstance2VpcResponse is the response struct for api BindInstance2Vpc
type BindInstance2VpcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Endpoint  string `json:"Endpoint" xml:"Endpoint"`
	Domain    string `json:"Domain" xml:"Domain"`
}

// CreateBindInstance2VpcRequest creates a request to invoke BindInstance2Vpc API
func CreateBindInstance2VpcRequest() (request *BindInstance2VpcRequest) {
	request = &BindInstance2VpcRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ots", "2016-06-20", "BindInstance2Vpc", "ots", "openAPI")
	return
}

// CreateBindInstance2VpcResponse creates a response to parse from BindInstance2Vpc response
func CreateBindInstance2VpcResponse() (response *BindInstance2VpcResponse) {
	response = &BindInstance2VpcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
