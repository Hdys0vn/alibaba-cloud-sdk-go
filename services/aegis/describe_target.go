package aegis

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

// DescribeTarget invokes the aegis.DescribeTarget API synchronously
// api document: https://help.aliyun.com/api/aegis/describetarget.html
func (client *Client) DescribeTarget(request *DescribeTargetRequest) (response *DescribeTargetResponse, err error) {
	response = CreateDescribeTargetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTargetWithChan invokes the aegis.DescribeTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/describetarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTargetWithChan(request *DescribeTargetRequest) (<-chan *DescribeTargetResponse, <-chan error) {
	responseChan := make(chan *DescribeTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTarget(request)
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

// DescribeTargetWithCallback invokes the aegis.DescribeTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/describetarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTargetWithCallback(request *DescribeTargetRequest, callback func(response *DescribeTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTargetResponse
		var err error
		defer close(result)
		response, err = client.DescribeTarget(request)
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

// DescribeTargetRequest is the request struct for api DescribeTarget
type DescribeTargetRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Type     string `position:"Query" name:"Type"`
	Config   string `position:"Query" name:"Config"`
}

// DescribeTargetResponse is the response struct for api DescribeTarget
type DescribeTargetResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	Targets    []Target `json:"Targets" xml:"Targets"`
}

// CreateDescribeTargetRequest creates a request to invoke DescribeTarget API
func CreateDescribeTargetRequest() (request *DescribeTargetRequest) {
	request = &DescribeTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeTarget", "vipaegis", "openAPI")
	return
}

// CreateDescribeTargetResponse creates a response to parse from DescribeTarget response
func CreateDescribeTargetResponse() (response *DescribeTargetResponse) {
	response = &DescribeTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
