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

// ListDefaultSLARules invokes the dms_enterprise.ListDefaultSLARules API synchronously
func (client *Client) ListDefaultSLARules(request *ListDefaultSLARulesRequest) (response *ListDefaultSLARulesResponse, err error) {
	response = CreateListDefaultSLARulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDefaultSLARulesWithChan invokes the dms_enterprise.ListDefaultSLARules API asynchronously
func (client *Client) ListDefaultSLARulesWithChan(request *ListDefaultSLARulesRequest) (<-chan *ListDefaultSLARulesResponse, <-chan error) {
	responseChan := make(chan *ListDefaultSLARulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDefaultSLARules(request)
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

// ListDefaultSLARulesWithCallback invokes the dms_enterprise.ListDefaultSLARules API asynchronously
func (client *Client) ListDefaultSLARulesWithCallback(request *ListDefaultSLARulesRequest, callback func(response *ListDefaultSLARulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDefaultSLARulesResponse
		var err error
		defer close(result)
		response, err = client.ListDefaultSLARules(request)
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

// ListDefaultSLARulesRequest is the request struct for api ListDefaultSLARules
type ListDefaultSLARulesRequest struct {
	*requests.RpcRequest
	DagId requests.Integer `position:"Query" name:"DagId"`
	Tid   requests.Integer `position:"Query" name:"Tid"`
}

// ListDefaultSLARulesResponse is the response struct for api ListDefaultSLARules
type ListDefaultSLARulesResponse struct {
	*responses.BaseResponse
	RequestId    string                           `json:"RequestId" xml:"RequestId"`
	ErrorCode    string                           `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool                             `json:"Success" xml:"Success"`
	SLARuleList  SLARuleListInListDefaultSLARules `json:"SLARuleList" xml:"SLARuleList"`
}

// CreateListDefaultSLARulesRequest creates a request to invoke ListDefaultSLARules API
func CreateListDefaultSLARulesRequest() (request *ListDefaultSLARulesRequest) {
	request = &ListDefaultSLARulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDefaultSLARules", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDefaultSLARulesResponse creates a response to parse from ListDefaultSLARules response
func CreateListDefaultSLARulesResponse() (response *ListDefaultSLARulesResponse) {
	response = &ListDefaultSLARulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
