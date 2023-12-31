package openanalytics_open

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

// GrantPrivileges invokes the openanalytics_open.GrantPrivileges API synchronously
func (client *Client) GrantPrivileges(request *GrantPrivilegesRequest) (response *GrantPrivilegesResponse, err error) {
	response = CreateGrantPrivilegesResponse()
	err = client.DoAction(request, response)
	return
}

// GrantPrivilegesWithChan invokes the openanalytics_open.GrantPrivileges API asynchronously
func (client *Client) GrantPrivilegesWithChan(request *GrantPrivilegesRequest) (<-chan *GrantPrivilegesResponse, <-chan error) {
	responseChan := make(chan *GrantPrivilegesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantPrivileges(request)
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

// GrantPrivilegesWithCallback invokes the openanalytics_open.GrantPrivileges API asynchronously
func (client *Client) GrantPrivilegesWithCallback(request *GrantPrivilegesRequest, callback func(response *GrantPrivilegesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantPrivilegesResponse
		var err error
		defer close(result)
		response, err = client.GrantPrivileges(request)
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

// GrantPrivilegesRequest is the request struct for api GrantPrivileges
type GrantPrivilegesRequest struct {
	*requests.RpcRequest
	PrivilegeBag string `position:"Query" name:"PrivilegeBag"`
}

// GrantPrivilegesResponse is the response struct for api GrantPrivileges
type GrantPrivilegesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Data      bool   `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateGrantPrivilegesRequest creates a request to invoke GrantPrivileges API
func CreateGrantPrivilegesRequest() (request *GrantPrivilegesRequest) {
	request = &GrantPrivilegesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2020-09-28", "GrantPrivileges", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGrantPrivilegesResponse creates a response to parse from GrantPrivileges response
func CreateGrantPrivilegesResponse() (response *GrantPrivilegesResponse) {
	response = &GrantPrivilegesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
