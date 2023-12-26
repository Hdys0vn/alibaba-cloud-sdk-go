package sas

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

// CheckUserHasEcs invokes the sas.CheckUserHasEcs API synchronously
func (client *Client) CheckUserHasEcs(request *CheckUserHasEcsRequest) (response *CheckUserHasEcsResponse, err error) {
	response = CreateCheckUserHasEcsResponse()
	err = client.DoAction(request, response)
	return
}

// CheckUserHasEcsWithChan invokes the sas.CheckUserHasEcs API asynchronously
func (client *Client) CheckUserHasEcsWithChan(request *CheckUserHasEcsRequest) (<-chan *CheckUserHasEcsResponse, <-chan error) {
	responseChan := make(chan *CheckUserHasEcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckUserHasEcs(request)
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

// CheckUserHasEcsWithCallback invokes the sas.CheckUserHasEcs API asynchronously
func (client *Client) CheckUserHasEcsWithCallback(request *CheckUserHasEcsRequest, callback func(response *CheckUserHasEcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckUserHasEcsResponse
		var err error
		defer close(result)
		response, err = client.CheckUserHasEcs(request)
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

// CheckUserHasEcsRequest is the request struct for api CheckUserHasEcs
type CheckUserHasEcsRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Lang        string           `position:"Query" name:"Lang"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
}

// CheckUserHasEcsResponse is the response struct for api CheckUserHasEcs
type CheckUserHasEcsResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckUserHasEcsRequest creates a request to invoke CheckUserHasEcs API
func CreateCheckUserHasEcsRequest() (request *CheckUserHasEcsRequest) {
	request = &CheckUserHasEcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "CheckUserHasEcs", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckUserHasEcsResponse creates a response to parse from CheckUserHasEcs response
func CreateCheckUserHasEcsResponse() (response *CheckUserHasEcsResponse) {
	response = &CheckUserHasEcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
