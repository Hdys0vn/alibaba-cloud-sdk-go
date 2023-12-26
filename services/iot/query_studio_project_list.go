package iot

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

// QueryStudioProjectList invokes the iot.QueryStudioProjectList API synchronously
func (client *Client) QueryStudioProjectList(request *QueryStudioProjectListRequest) (response *QueryStudioProjectListResponse, err error) {
	response = CreateQueryStudioProjectListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryStudioProjectListWithChan invokes the iot.QueryStudioProjectList API asynchronously
func (client *Client) QueryStudioProjectListWithChan(request *QueryStudioProjectListRequest) (<-chan *QueryStudioProjectListResponse, <-chan error) {
	responseChan := make(chan *QueryStudioProjectListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryStudioProjectList(request)
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

// QueryStudioProjectListWithCallback invokes the iot.QueryStudioProjectList API asynchronously
func (client *Client) QueryStudioProjectListWithCallback(request *QueryStudioProjectListRequest, callback func(response *QueryStudioProjectListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryStudioProjectListResponse
		var err error
		defer close(result)
		response, err = client.QueryStudioProjectList(request)
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

// QueryStudioProjectListRequest is the request struct for api QueryStudioProjectList
type QueryStudioProjectListRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	PageNo        requests.Integer `position:"Body" name:"PageNo"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	Name          string           `position:"Body" name:"Name"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryStudioProjectListResponse is the response struct for api QueryStudioProjectList
type QueryStudioProjectListResponse struct {
	*responses.BaseResponse
	RequestId    string                       `json:"RequestId" xml:"RequestId"`
	Success      bool                         `json:"Success" xml:"Success"`
	Code         string                       `json:"Code" xml:"Code"`
	ErrorMessage string                       `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryStudioProjectList `json:"Data" xml:"Data"`
}

// CreateQueryStudioProjectListRequest creates a request to invoke QueryStudioProjectList API
func CreateQueryStudioProjectListRequest() (request *QueryStudioProjectListRequest) {
	request = &QueryStudioProjectListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryStudioProjectList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryStudioProjectListResponse creates a response to parse from QueryStudioProjectList response
func CreateQueryStudioProjectListResponse() (response *QueryStudioProjectListResponse) {
	response = &QueryStudioProjectListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}