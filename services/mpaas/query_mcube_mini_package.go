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

// QueryMcubeMiniPackage invokes the mpaas.QueryMcubeMiniPackage API synchronously
func (client *Client) QueryMcubeMiniPackage(request *QueryMcubeMiniPackageRequest) (response *QueryMcubeMiniPackageResponse, err error) {
	response = CreateQueryMcubeMiniPackageResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMcubeMiniPackageWithChan invokes the mpaas.QueryMcubeMiniPackage API asynchronously
func (client *Client) QueryMcubeMiniPackageWithChan(request *QueryMcubeMiniPackageRequest) (<-chan *QueryMcubeMiniPackageResponse, <-chan error) {
	responseChan := make(chan *QueryMcubeMiniPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMcubeMiniPackage(request)
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

// QueryMcubeMiniPackageWithCallback invokes the mpaas.QueryMcubeMiniPackage API asynchronously
func (client *Client) QueryMcubeMiniPackageWithCallback(request *QueryMcubeMiniPackageRequest, callback func(response *QueryMcubeMiniPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMcubeMiniPackageResponse
		var err error
		defer close(result)
		response, err = client.QueryMcubeMiniPackage(request)
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

// QueryMcubeMiniPackageRequest is the request struct for api QueryMcubeMiniPackage
type QueryMcubeMiniPackageRequest struct {
	*requests.RpcRequest
	H5Id        string `position:"Body" name:"H5Id"`
	TenantId    string `position:"Body" name:"TenantId"`
	Id          string `position:"Body" name:"Id"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// QueryMcubeMiniPackageResponse is the response struct for api QueryMcubeMiniPackage
type QueryMcubeMiniPackageResponse struct {
	*responses.BaseResponse
	ResultMessage          string                 `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode             string                 `json:"ResultCode" xml:"ResultCode"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	QueryMiniPackageResult QueryMiniPackageResult `json:"QueryMiniPackageResult" xml:"QueryMiniPackageResult"`
}

// CreateQueryMcubeMiniPackageRequest creates a request to invoke QueryMcubeMiniPackage API
func CreateQueryMcubeMiniPackageRequest() (request *QueryMcubeMiniPackageRequest) {
	request = &QueryMcubeMiniPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "QueryMcubeMiniPackage", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMcubeMiniPackageResponse creates a response to parse from QueryMcubeMiniPackage response
func CreateQueryMcubeMiniPackageResponse() (response *QueryMcubeMiniPackageResponse) {
	response = &QueryMcubeMiniPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
