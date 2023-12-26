package afs

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

// DescribePersonMachineList invokes the afs.DescribePersonMachineList API synchronously
// api document: https://help.aliyun.com/api/afs/describepersonmachinelist.html
func (client *Client) DescribePersonMachineList(request *DescribePersonMachineListRequest) (response *DescribePersonMachineListResponse, err error) {
	response = CreateDescribePersonMachineListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePersonMachineListWithChan invokes the afs.DescribePersonMachineList API asynchronously
// api document: https://help.aliyun.com/api/afs/describepersonmachinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePersonMachineListWithChan(request *DescribePersonMachineListRequest) (<-chan *DescribePersonMachineListResponse, <-chan error) {
	responseChan := make(chan *DescribePersonMachineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePersonMachineList(request)
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

// DescribePersonMachineListWithCallback invokes the afs.DescribePersonMachineList API asynchronously
// api document: https://help.aliyun.com/api/afs/describepersonmachinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePersonMachineListWithCallback(request *DescribePersonMachineListRequest, callback func(response *DescribePersonMachineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePersonMachineListResponse
		var err error
		defer close(result)
		response, err = client.DescribePersonMachineList(request)
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

// DescribePersonMachineListRequest is the request struct for api DescribePersonMachineList
type DescribePersonMachineListRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribePersonMachineListResponse is the response struct for api DescribePersonMachineList
type DescribePersonMachineListResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	BizCode          string           `json:"BizCode" xml:"BizCode"`
	PersonMachineRes PersonMachineRes `json:"PersonMachineRes" xml:"PersonMachineRes"`
}

// CreateDescribePersonMachineListRequest creates a request to invoke DescribePersonMachineList API
func CreateDescribePersonMachineListRequest() (request *DescribePersonMachineListRequest) {
	request = &DescribePersonMachineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "DescribePersonMachineList", "afs", "openAPI")
	return
}

// CreateDescribePersonMachineListResponse creates a response to parse from DescribePersonMachineList response
func CreateDescribePersonMachineListResponse() (response *DescribePersonMachineListResponse) {
	response = &DescribePersonMachineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}