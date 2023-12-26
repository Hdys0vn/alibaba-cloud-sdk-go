package dms_enterprise

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

// SuspendTaskFlowInstance invokes the dms_enterprise.SuspendTaskFlowInstance API synchronously
func (client *Client) SuspendTaskFlowInstance(request *SuspendTaskFlowInstanceRequest) (response *SuspendTaskFlowInstanceResponse, err error) {
	response = CreateSuspendTaskFlowInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendTaskFlowInstanceWithChan invokes the dms_enterprise.SuspendTaskFlowInstance API asynchronously
func (client *Client) SuspendTaskFlowInstanceWithChan(request *SuspendTaskFlowInstanceRequest) (<-chan *SuspendTaskFlowInstanceResponse, <-chan error) {
	responseChan := make(chan *SuspendTaskFlowInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendTaskFlowInstance(request)
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

// SuspendTaskFlowInstanceWithCallback invokes the dms_enterprise.SuspendTaskFlowInstance API asynchronously
func (client *Client) SuspendTaskFlowInstanceWithCallback(request *SuspendTaskFlowInstanceRequest, callback func(response *SuspendTaskFlowInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendTaskFlowInstanceResponse
		var err error
		defer close(result)
		response, err = client.SuspendTaskFlowInstance(request)
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

// SuspendTaskFlowInstanceRequest is the request struct for api SuspendTaskFlowInstance
type SuspendTaskFlowInstanceRequest struct {
	*requests.RpcRequest
	DagId         requests.Integer `position:"Query" name:"DagId"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
	DagInstanceId requests.Integer `position:"Query" name:"DagInstanceId"`
}

// SuspendTaskFlowInstanceResponse is the response struct for api SuspendTaskFlowInstance
type SuspendTaskFlowInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateSuspendTaskFlowInstanceRequest creates a request to invoke SuspendTaskFlowInstance API
func CreateSuspendTaskFlowInstanceRequest() (request *SuspendTaskFlowInstanceRequest) {
	request = &SuspendTaskFlowInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SuspendTaskFlowInstance", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendTaskFlowInstanceResponse creates a response to parse from SuspendTaskFlowInstance response
func CreateSuspendTaskFlowInstanceResponse() (response *SuspendTaskFlowInstanceResponse) {
	response = &SuspendTaskFlowInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}