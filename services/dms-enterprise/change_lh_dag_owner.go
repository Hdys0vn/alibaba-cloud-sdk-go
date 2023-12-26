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

// ChangeLhDagOwner invokes the dms_enterprise.ChangeLhDagOwner API synchronously
func (client *Client) ChangeLhDagOwner(request *ChangeLhDagOwnerRequest) (response *ChangeLhDagOwnerResponse, err error) {
	response = CreateChangeLhDagOwnerResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeLhDagOwnerWithChan invokes the dms_enterprise.ChangeLhDagOwner API asynchronously
func (client *Client) ChangeLhDagOwnerWithChan(request *ChangeLhDagOwnerRequest) (<-chan *ChangeLhDagOwnerResponse, <-chan error) {
	responseChan := make(chan *ChangeLhDagOwnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeLhDagOwner(request)
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

// ChangeLhDagOwnerWithCallback invokes the dms_enterprise.ChangeLhDagOwner API asynchronously
func (client *Client) ChangeLhDagOwnerWithCallback(request *ChangeLhDagOwnerRequest, callback func(response *ChangeLhDagOwnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeLhDagOwnerResponse
		var err error
		defer close(result)
		response, err = client.ChangeLhDagOwner(request)
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

// ChangeLhDagOwnerRequest is the request struct for api ChangeLhDagOwner
type ChangeLhDagOwnerRequest struct {
	*requests.RpcRequest
	DagId       requests.Integer `position:"Query" name:"DagId"`
	Tid         requests.Integer `position:"Query" name:"Tid"`
	OwnerUserId requests.Integer `position:"Query" name:"OwnerUserId"`
}

// ChangeLhDagOwnerResponse is the response struct for api ChangeLhDagOwner
type ChangeLhDagOwnerResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateChangeLhDagOwnerRequest creates a request to invoke ChangeLhDagOwner API
func CreateChangeLhDagOwnerRequest() (request *ChangeLhDagOwnerRequest) {
	request = &ChangeLhDagOwnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ChangeLhDagOwner", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeLhDagOwnerResponse creates a response to parse from ChangeLhDagOwner response
func CreateChangeLhDagOwnerResponse() (response *ChangeLhDagOwnerResponse) {
	response = &ChangeLhDagOwnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
