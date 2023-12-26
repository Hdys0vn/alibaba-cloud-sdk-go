package viapi

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

// QueryAsyncJobList invokes the viapi.QueryAsyncJobList API synchronously
func (client *Client) QueryAsyncJobList(request *QueryAsyncJobListRequest) (response *QueryAsyncJobListResponse, err error) {
	response = CreateQueryAsyncJobListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAsyncJobListWithChan invokes the viapi.QueryAsyncJobList API asynchronously
func (client *Client) QueryAsyncJobListWithChan(request *QueryAsyncJobListRequest) (<-chan *QueryAsyncJobListResponse, <-chan error) {
	responseChan := make(chan *QueryAsyncJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAsyncJobList(request)
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

// QueryAsyncJobListWithCallback invokes the viapi.QueryAsyncJobList API asynchronously
func (client *Client) QueryAsyncJobListWithCallback(request *QueryAsyncJobListRequest, callback func(response *QueryAsyncJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAsyncJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryAsyncJobList(request)
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

// QueryAsyncJobListRequest is the request struct for api QueryAsyncJobList
type QueryAsyncJobListRequest struct {
	*requests.RpcRequest
	StartTime  string `position:"Body" name:"StartTime"`
	PageNum    string `position:"Body" name:"PageNum"`
	JobId      string `position:"Body" name:"JobId"`
	PopApiName string `position:"Body" name:"PopApiName"`
	PageSize   string `position:"Body" name:"PageSize"`
	PopProduct string `position:"Body" name:"PopProduct"`
	EndTime    string `position:"Body" name:"EndTime"`
	Status     string `position:"Body" name:"Status"`
}

// QueryAsyncJobListResponse is the response struct for api QueryAsyncJobList
type QueryAsyncJobListResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Message   string                  `json:"Message" xml:"Message"`
	Success   bool                    `json:"Success" xml:"Success"`
	HttpCode  string                  `json:"HttpCode" xml:"HttpCode"`
	Data      DataInQueryAsyncJobList `json:"Data" xml:"Data"`
}

// CreateQueryAsyncJobListRequest creates a request to invoke QueryAsyncJobList API
func CreateQueryAsyncJobListRequest() (request *QueryAsyncJobListRequest) {
	request = &QueryAsyncJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi", "2023-01-17", "QueryAsyncJobList", "viapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAsyncJobListResponse creates a response to parse from QueryAsyncJobList response
func CreateQueryAsyncJobListResponse() (response *QueryAsyncJobListResponse) {
	response = &QueryAsyncJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
