package mpaas

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

// SetMsacAppInstance invokes the mpaas.SetMsacAppInstance API synchronously
func (client *Client) SetMsacAppInstance(request *SetMsacAppInstanceRequest) (response *SetMsacAppInstanceResponse, err error) {
	response = CreateSetMsacAppInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// SetMsacAppInstanceWithChan invokes the mpaas.SetMsacAppInstance API asynchronously
func (client *Client) SetMsacAppInstanceWithChan(request *SetMsacAppInstanceRequest) (<-chan *SetMsacAppInstanceResponse, <-chan error) {
	responseChan := make(chan *SetMsacAppInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetMsacAppInstance(request)
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

// SetMsacAppInstanceWithCallback invokes the mpaas.SetMsacAppInstance API asynchronously
func (client *Client) SetMsacAppInstanceWithCallback(request *SetMsacAppInstanceRequest, callback func(response *SetMsacAppInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetMsacAppInstanceResponse
		var err error
		defer close(result)
		response, err = client.SetMsacAppInstance(request)
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

// SetMsacAppInstanceRequest is the request struct for api SetMsacAppInstance
type SetMsacAppInstanceRequest struct {
	*requests.RpcRequest
	TenantId                                 string `position:"Body" name:"TenantId"`
	MpaasMappcenterMsacSetAppInstanceJsonStr string `position:"Body" name:"mpaasMappcenterMsacSetAppInstanceJsonStr"`
	AppId                                    string `position:"Body" name:"AppId"`
	WorkspaceId                              string `position:"Body" name:"WorkspaceId"`
}

// SetMsacAppInstanceResponse is the response struct for api SetMsacAppInstance
type SetMsacAppInstanceResponse struct {
	*responses.BaseResponse
	ResultMessage string                            `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                            `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                            `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInSetMsacAppInstance `json:"ResultContent" xml:"ResultContent"`
}

// CreateSetMsacAppInstanceRequest creates a request to invoke SetMsacAppInstance API
func CreateSetMsacAppInstanceRequest() (request *SetMsacAppInstanceRequest) {
	request = &SetMsacAppInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "SetMsacAppInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateSetMsacAppInstanceResponse creates a response to parse from SetMsacAppInstance response
func CreateSetMsacAppInstanceResponse() (response *SetMsacAppInstanceResponse) {
	response = &SetMsacAppInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}