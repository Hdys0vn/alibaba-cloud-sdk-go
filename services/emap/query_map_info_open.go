package emap

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

// QueryMapInfoOpen invokes the emap.QueryMapInfoOpen API synchronously
func (client *Client) QueryMapInfoOpen(request *QueryMapInfoOpenRequest) (response *QueryMapInfoOpenResponse, err error) {
	response = CreateQueryMapInfoOpenResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMapInfoOpenWithChan invokes the emap.QueryMapInfoOpen API asynchronously
func (client *Client) QueryMapInfoOpenWithChan(request *QueryMapInfoOpenRequest) (<-chan *QueryMapInfoOpenResponse, <-chan error) {
	responseChan := make(chan *QueryMapInfoOpenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMapInfoOpen(request)
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

// QueryMapInfoOpenWithCallback invokes the emap.QueryMapInfoOpen API asynchronously
func (client *Client) QueryMapInfoOpenWithCallback(request *QueryMapInfoOpenRequest, callback func(response *QueryMapInfoOpenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMapInfoOpenResponse
		var err error
		defer close(result)
		response, err = client.QueryMapInfoOpen(request)
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

// QueryMapInfoOpenRequest is the request struct for api QueryMapInfoOpen
type QueryMapInfoOpenRequest struct {
	*requests.RpcRequest
	StoreId requests.Integer `position:"Body" name:"StoreId"`
}

// QueryMapInfoOpenResponse is the response struct for api QueryMapInfoOpen
type QueryMapInfoOpenResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RenderData     string `json:"RenderData" xml:"RenderData"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Code           string `json:"Code" xml:"Code"`
}

// CreateQueryMapInfoOpenRequest creates a request to invoke QueryMapInfoOpen API
func CreateQueryMapInfoOpenRequest() (request *QueryMapInfoOpenRequest) {
	request = &QueryMapInfoOpenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emap", "2020-10-10", "QueryMapInfoOpen", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMapInfoOpenResponse creates a response to parse from QueryMapInfoOpen response
func CreateQueryMapInfoOpenResponse() (response *QueryMapInfoOpenResponse) {
	response = &QueryMapInfoOpenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
