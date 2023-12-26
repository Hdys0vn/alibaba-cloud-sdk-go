package resourcemanager

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

// GetServiceLinkedRoleDeletionStatus invokes the resourcemanager.GetServiceLinkedRoleDeletionStatus API synchronously
func (client *Client) GetServiceLinkedRoleDeletionStatus(request *GetServiceLinkedRoleDeletionStatusRequest) (response *GetServiceLinkedRoleDeletionStatusResponse, err error) {
	response = CreateGetServiceLinkedRoleDeletionStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceLinkedRoleDeletionStatusWithChan invokes the resourcemanager.GetServiceLinkedRoleDeletionStatus API asynchronously
func (client *Client) GetServiceLinkedRoleDeletionStatusWithChan(request *GetServiceLinkedRoleDeletionStatusRequest) (<-chan *GetServiceLinkedRoleDeletionStatusResponse, <-chan error) {
	responseChan := make(chan *GetServiceLinkedRoleDeletionStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceLinkedRoleDeletionStatus(request)
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

// GetServiceLinkedRoleDeletionStatusWithCallback invokes the resourcemanager.GetServiceLinkedRoleDeletionStatus API asynchronously
func (client *Client) GetServiceLinkedRoleDeletionStatusWithCallback(request *GetServiceLinkedRoleDeletionStatusRequest, callback func(response *GetServiceLinkedRoleDeletionStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceLinkedRoleDeletionStatusResponse
		var err error
		defer close(result)
		response, err = client.GetServiceLinkedRoleDeletionStatus(request)
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

// GetServiceLinkedRoleDeletionStatusRequest is the request struct for api GetServiceLinkedRoleDeletionStatus
type GetServiceLinkedRoleDeletionStatusRequest struct {
	*requests.RpcRequest
	DeletionTaskId string `position:"Query" name:"DeletionTaskId"`
}

// GetServiceLinkedRoleDeletionStatusResponse is the response struct for api GetServiceLinkedRoleDeletionStatus
type GetServiceLinkedRoleDeletionStatusResponse struct {
	*responses.BaseResponse
	Status    string                                     `json:"Status" xml:"Status"`
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Reason    ReasonInGetServiceLinkedRoleDeletionStatus `json:"Reason" xml:"Reason"`
}

// CreateGetServiceLinkedRoleDeletionStatusRequest creates a request to invoke GetServiceLinkedRoleDeletionStatus API
func CreateGetServiceLinkedRoleDeletionStatusRequest() (request *GetServiceLinkedRoleDeletionStatusRequest) {
	request = &GetServiceLinkedRoleDeletionStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "GetServiceLinkedRoleDeletionStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetServiceLinkedRoleDeletionStatusResponse creates a response to parse from GetServiceLinkedRoleDeletionStatus response
func CreateGetServiceLinkedRoleDeletionStatusResponse() (response *GetServiceLinkedRoleDeletionStatusResponse) {
	response = &GetServiceLinkedRoleDeletionStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
