package ecd

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

// DescribeSecurityEventOperationStatus invokes the ecd.DescribeSecurityEventOperationStatus API synchronously
func (client *Client) DescribeSecurityEventOperationStatus(request *DescribeSecurityEventOperationStatusRequest) (response *DescribeSecurityEventOperationStatusResponse, err error) {
	response = CreateDescribeSecurityEventOperationStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityEventOperationStatusWithChan invokes the ecd.DescribeSecurityEventOperationStatus API asynchronously
func (client *Client) DescribeSecurityEventOperationStatusWithChan(request *DescribeSecurityEventOperationStatusRequest) (<-chan *DescribeSecurityEventOperationStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityEventOperationStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityEventOperationStatus(request)
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

// DescribeSecurityEventOperationStatusWithCallback invokes the ecd.DescribeSecurityEventOperationStatus API asynchronously
func (client *Client) DescribeSecurityEventOperationStatusWithCallback(request *DescribeSecurityEventOperationStatusRequest, callback func(response *DescribeSecurityEventOperationStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityEventOperationStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityEventOperationStatus(request)
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

// DescribeSecurityEventOperationStatusRequest is the request struct for api DescribeSecurityEventOperationStatus
type DescribeSecurityEventOperationStatusRequest struct {
	*requests.RpcRequest
	SecurityEventId *[]string        `position:"Query" name:"SecurityEventId"  type:"Repeated"`
	TaskId          requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeSecurityEventOperationStatusResponse is the response struct for api DescribeSecurityEventOperationStatus
type DescribeSecurityEventOperationStatusResponse struct {
	*responses.BaseResponse
	TaskStatus                     string                         `json:"TaskStatus" xml:"TaskStatus"`
	RequestId                      string                         `json:"RequestId" xml:"RequestId"`
	SecurityEventOperationStatuses []SecurityEventOperationStatus `json:"SecurityEventOperationStatuses" xml:"SecurityEventOperationStatuses"`
}

// CreateDescribeSecurityEventOperationStatusRequest creates a request to invoke DescribeSecurityEventOperationStatus API
func CreateDescribeSecurityEventOperationStatusRequest() (request *DescribeSecurityEventOperationStatusRequest) {
	request = &DescribeSecurityEventOperationStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeSecurityEventOperationStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSecurityEventOperationStatusResponse creates a response to parse from DescribeSecurityEventOperationStatus response
func CreateDescribeSecurityEventOperationStatusResponse() (response *DescribeSecurityEventOperationStatusResponse) {
	response = &DescribeSecurityEventOperationStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
