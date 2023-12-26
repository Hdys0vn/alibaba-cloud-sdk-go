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

// QueryMcdpZone invokes the mpaas.QueryMcdpZone API synchronously
func (client *Client) QueryMcdpZone(request *QueryMcdpZoneRequest) (response *QueryMcdpZoneResponse, err error) {
	response = CreateQueryMcdpZoneResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMcdpZoneWithChan invokes the mpaas.QueryMcdpZone API asynchronously
func (client *Client) QueryMcdpZoneWithChan(request *QueryMcdpZoneRequest) (<-chan *QueryMcdpZoneResponse, <-chan error) {
	responseChan := make(chan *QueryMcdpZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMcdpZone(request)
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

// QueryMcdpZoneWithCallback invokes the mpaas.QueryMcdpZone API asynchronously
func (client *Client) QueryMcdpZoneWithCallback(request *QueryMcdpZoneRequest, callback func(response *QueryMcdpZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMcdpZoneResponse
		var err error
		defer close(result)
		response, err = client.QueryMcdpZone(request)
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

// QueryMcdpZoneRequest is the request struct for api QueryMcdpZone
type QueryMcdpZoneRequest struct {
	*requests.RpcRequest
	TenantId    string           `position:"Body" name:"TenantId"`
	Id          requests.Integer `position:"Body" name:"Id"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// QueryMcdpZoneResponse is the response struct for api QueryMcdpZone
type QueryMcdpZoneResponse struct {
	*responses.BaseResponse
	ResultMessage string                       `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                       `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                       `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInQueryMcdpZone `json:"ResultContent" xml:"ResultContent"`
}

// CreateQueryMcdpZoneRequest creates a request to invoke QueryMcdpZone API
func CreateQueryMcdpZoneRequest() (request *QueryMcdpZoneRequest) {
	request = &QueryMcdpZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "QueryMcdpZone", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMcdpZoneResponse creates a response to parse from QueryMcdpZone response
func CreateQueryMcdpZoneResponse() (response *QueryMcdpZoneResponse) {
	response = &QueryMcdpZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
