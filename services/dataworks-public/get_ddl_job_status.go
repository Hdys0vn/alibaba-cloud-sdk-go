package dataworks_public

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

// GetDDLJobStatus invokes the dataworks_public.GetDDLJobStatus API synchronously
func (client *Client) GetDDLJobStatus(request *GetDDLJobStatusRequest) (response *GetDDLJobStatusResponse, err error) {
	response = CreateGetDDLJobStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetDDLJobStatusWithChan invokes the dataworks_public.GetDDLJobStatus API asynchronously
func (client *Client) GetDDLJobStatusWithChan(request *GetDDLJobStatusRequest) (<-chan *GetDDLJobStatusResponse, <-chan error) {
	responseChan := make(chan *GetDDLJobStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDDLJobStatus(request)
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

// GetDDLJobStatusWithCallback invokes the dataworks_public.GetDDLJobStatus API asynchronously
func (client *Client) GetDDLJobStatusWithCallback(request *GetDDLJobStatusRequest, callback func(response *GetDDLJobStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDDLJobStatusResponse
		var err error
		defer close(result)
		response, err = client.GetDDLJobStatus(request)
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

// GetDDLJobStatusRequest is the request struct for api GetDDLJobStatus
type GetDDLJobStatusRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

// GetDDLJobStatusResponse is the response struct for api GetDDLJobStatus
type GetDDLJobStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDDLJobStatusRequest creates a request to invoke GetDDLJobStatus API
func CreateGetDDLJobStatusRequest() (request *GetDDLJobStatusRequest) {
	request = &GetDDLJobStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetDDLJobStatus", "", "")
	request.Method = requests.GET
	return
}

// CreateGetDDLJobStatusResponse creates a response to parse from GetDDLJobStatus response
func CreateGetDDLJobStatusResponse() (response *GetDDLJobStatusResponse) {
	response = &GetDDLJobStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
