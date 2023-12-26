package yundun_dbaudit

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

// DescribeInstanceStorage invokes the yundun_dbaudit.DescribeInstanceStorage API synchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstancestorage.html
func (client *Client) DescribeInstanceStorage(request *DescribeInstanceStorageRequest) (response *DescribeInstanceStorageResponse, err error) {
	response = CreateDescribeInstanceStorageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceStorageWithChan invokes the yundun_dbaudit.DescribeInstanceStorage API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstancestorage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceStorageWithChan(request *DescribeInstanceStorageRequest) (<-chan *DescribeInstanceStorageResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceStorage(request)
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

// DescribeInstanceStorageWithCallback invokes the yundun_dbaudit.DescribeInstanceStorage API asynchronously
// api document: https://help.aliyun.com/api/yundun-dbaudit/describeinstancestorage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceStorageWithCallback(request *DescribeInstanceStorageRequest, callback func(response *DescribeInstanceStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceStorageResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceStorage(request)
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

// DescribeInstanceStorageRequest is the request struct for api DescribeInstanceStorage
type DescribeInstanceStorageRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeInstanceStorageResponse is the response struct for api DescribeInstanceStorage
type DescribeInstanceStorageResponse struct {
	*responses.BaseResponse
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	InstanceStorages []InstanceStorage `json:"InstanceStorages" xml:"InstanceStorages"`
}

// CreateDescribeInstanceStorageRequest creates a request to invoke DescribeInstanceStorage API
func CreateDescribeInstanceStorageRequest() (request *DescribeInstanceStorageRequest) {
	request = &DescribeInstanceStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-dbaudit", "2018-10-29", "DescribeInstanceStorage", "dbaudit", "openAPI")
	return
}

// CreateDescribeInstanceStorageResponse creates a response to parse from DescribeInstanceStorage response
func CreateDescribeInstanceStorageResponse() (response *DescribeInstanceStorageResponse) {
	response = &DescribeInstanceStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
