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

// DescribeSyncInfo invokes the yundun_bastionhost.DescribeSyncInfo API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describesyncinfo.html
func (client *Client) DescribeSyncInfo(request *DescribeSyncInfoRequest) (response *DescribeSyncInfoResponse, err error) {
	response = CreateDescribeSyncInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSyncInfoWithChan invokes the yundun_bastionhost.DescribeSyncInfo API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describesyncinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSyncInfoWithChan(request *DescribeSyncInfoRequest) (<-chan *DescribeSyncInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeSyncInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSyncInfo(request)
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

// DescribeSyncInfoWithCallback invokes the yundun_bastionhost.DescribeSyncInfo API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/describesyncinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSyncInfoWithCallback(request *DescribeSyncInfoRequest, callback func(response *DescribeSyncInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSyncInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeSyncInfo(request)
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

// DescribeSyncInfoRequest is the request struct for api DescribeSyncInfo
type DescribeSyncInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeSyncInfoResponse is the response struct for api DescribeSyncInfo
type DescribeSyncInfoResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	InstanceInfo InstanceInfo `json:"InstanceInfo" xml:"InstanceInfo"`
}

// CreateDescribeSyncInfoRequest creates a request to invoke DescribeSyncInfo API
func CreateDescribeSyncInfoRequest() (request *DescribeSyncInfoRequest) {
	request = &DescribeSyncInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "DescribeSyncInfo", "bastionhost", "openAPI")
	return
}

// CreateDescribeSyncInfoResponse creates a response to parse from DescribeSyncInfo response
func CreateDescribeSyncInfoResponse() (response *DescribeSyncInfoResponse) {
	response = &DescribeSyncInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
