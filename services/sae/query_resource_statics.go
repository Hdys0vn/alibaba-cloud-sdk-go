package sae

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

// QueryResourceStatics invokes the sae.QueryResourceStatics API synchronously
func (client *Client) QueryResourceStatics(request *QueryResourceStaticsRequest) (response *QueryResourceStaticsResponse, err error) {
	response = CreateQueryResourceStaticsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryResourceStaticsWithChan invokes the sae.QueryResourceStatics API asynchronously
func (client *Client) QueryResourceStaticsWithChan(request *QueryResourceStaticsRequest) (<-chan *QueryResourceStaticsResponse, <-chan error) {
	responseChan := make(chan *QueryResourceStaticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryResourceStatics(request)
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

// QueryResourceStaticsWithCallback invokes the sae.QueryResourceStatics API asynchronously
func (client *Client) QueryResourceStaticsWithCallback(request *QueryResourceStaticsRequest, callback func(response *QueryResourceStaticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryResourceStaticsResponse
		var err error
		defer close(result)
		response, err = client.QueryResourceStatics(request)
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

// QueryResourceStaticsRequest is the request struct for api QueryResourceStatics
type QueryResourceStaticsRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// QueryResourceStaticsResponse is the response struct for api QueryResourceStatics
type QueryResourceStaticsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryResourceStaticsRequest creates a request to invoke QueryResourceStatics API
func CreateQueryResourceStaticsRequest() (request *QueryResourceStaticsRequest) {
	request = &QueryResourceStaticsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "QueryResourceStatics", "/pop/v1/paas/quota/queryResourceStatics", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryResourceStaticsResponse creates a response to parse from QueryResourceStatics response
func CreateQueryResourceStaticsResponse() (response *QueryResourceStaticsResponse) {
	response = &QueryResourceStaticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
