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

// QueryStudioAppList invokes the iot.QueryStudioAppList API synchronously
func (client *Client) QueryStudioAppList(request *QueryStudioAppListRequest) (response *QueryStudioAppListResponse, err error) {
	response = CreateQueryStudioAppListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryStudioAppListWithChan invokes the iot.QueryStudioAppList API asynchronously
func (client *Client) QueryStudioAppListWithChan(request *QueryStudioAppListRequest) (<-chan *QueryStudioAppListResponse, <-chan error) {
	responseChan := make(chan *QueryStudioAppListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryStudioAppList(request)
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

// QueryStudioAppListWithCallback invokes the iot.QueryStudioAppList API asynchronously
func (client *Client) QueryStudioAppListWithCallback(request *QueryStudioAppListRequest, callback func(response *QueryStudioAppListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryStudioAppListResponse
		var err error
		defer close(result)
		response, err = client.QueryStudioAppList(request)
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

// QueryStudioAppListRequest is the request struct for api QueryStudioAppList
type QueryStudioAppListRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ProjectId     string           `position:"Body" name:"ProjectId"`
	Types         *[]string        `position:"Body" name:"Types"  type:"Repeated"`
	FuzzyName     string           `position:"Body" name:"FuzzyName"`
	PageNo        requests.Integer `position:"Body" name:"PageNo"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryStudioAppListResponse is the response struct for api QueryStudioAppList
type QueryStudioAppListResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	Success      bool                     `json:"Success" xml:"Success"`
	Code         string                   `json:"Code" xml:"Code"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryStudioAppList `json:"Data" xml:"Data"`
}

// CreateQueryStudioAppListRequest creates a request to invoke QueryStudioAppList API
func CreateQueryStudioAppListRequest() (request *QueryStudioAppListRequest) {
	request = &QueryStudioAppListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryStudioAppList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryStudioAppListResponse creates a response to parse from QueryStudioAppList response
func CreateQueryStudioAppListResponse() (response *QueryStudioAppListResponse) {
	response = &QueryStudioAppListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
