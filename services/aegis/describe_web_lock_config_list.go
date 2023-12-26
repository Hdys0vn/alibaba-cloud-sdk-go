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

// DescribeWebLockConfigList invokes the aegis.DescribeWebLockConfigList API synchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockconfiglist.html
func (client *Client) DescribeWebLockConfigList(request *DescribeWebLockConfigListRequest) (response *DescribeWebLockConfigListResponse, err error) {
	response = CreateDescribeWebLockConfigListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebLockConfigListWithChan invokes the aegis.DescribeWebLockConfigList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockconfiglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLockConfigListWithChan(request *DescribeWebLockConfigListRequest) (<-chan *DescribeWebLockConfigListResponse, <-chan error) {
	responseChan := make(chan *DescribeWebLockConfigListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebLockConfigList(request)
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

// DescribeWebLockConfigListWithCallback invokes the aegis.DescribeWebLockConfigList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeweblockconfiglist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLockConfigListWithCallback(request *DescribeWebLockConfigListRequest, callback func(response *DescribeWebLockConfigListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebLockConfigListResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebLockConfigList(request)
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

// DescribeWebLockConfigListRequest is the request struct for api DescribeWebLockConfigList
type DescribeWebLockConfigListRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
	Uuid     string `position:"Query" name:"Uuid"`
}

// DescribeWebLockConfigListResponse is the response struct for api DescribeWebLockConfigList
type DescribeWebLockConfigListResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	TotalCount int          `json:"TotalCount" xml:"TotalCount"`
	ConfigList []ConfigInfo `json:"ConfigList" xml:"ConfigList"`
}

// CreateDescribeWebLockConfigListRequest creates a request to invoke DescribeWebLockConfigList API
func CreateDescribeWebLockConfigListRequest() (request *DescribeWebLockConfigListRequest) {
	request = &DescribeWebLockConfigListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeWebLockConfigList", "vipaegis", "openAPI")
	return
}

// CreateDescribeWebLockConfigListResponse creates a response to parse from DescribeWebLockConfigList response
func CreateDescribeWebLockConfigListResponse() (response *DescribeWebLockConfigListResponse) {
	response = &DescribeWebLockConfigListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}