package mts

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

// QueryInferenceServer invokes the mts.QueryInferenceServer API synchronously
func (client *Client) QueryInferenceServer(request *QueryInferenceServerRequest) (response *QueryInferenceServerResponse, err error) {
	response = CreateQueryInferenceServerResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInferenceServerWithChan invokes the mts.QueryInferenceServer API asynchronously
func (client *Client) QueryInferenceServerWithChan(request *QueryInferenceServerRequest) (<-chan *QueryInferenceServerResponse, <-chan error) {
	responseChan := make(chan *QueryInferenceServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInferenceServer(request)
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

// QueryInferenceServerWithCallback invokes the mts.QueryInferenceServer API asynchronously
func (client *Client) QueryInferenceServerWithCallback(request *QueryInferenceServerRequest, callback func(response *QueryInferenceServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInferenceServerResponse
		var err error
		defer close(result)
		response, err = client.QueryInferenceServer(request)
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

// QueryInferenceServerRequest is the request struct for api QueryInferenceServer
type QueryInferenceServerRequest struct {
	*requests.RpcRequest
	CreateTime  requests.Integer `position:"Query" name:"CreateTime"`
	MaxPageSize requests.Integer `position:"Query" name:"MaxPageSize"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	ModelType   string           `position:"Query" name:"ModelType"`
}

// QueryInferenceServerResponse is the response struct for api QueryInferenceServer
type QueryInferenceServerResponse struct {
	*responses.BaseResponse
	Message   string          `json:"Message" xml:"Message"`
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Code      string          `json:"Code" xml:"Code"`
	TotalSize int64           `json:"TotalSize" xml:"TotalSize"`
	Functions []FunctionsItem `json:"Functions" xml:"Functions"`
}

// CreateQueryInferenceServerRequest creates a request to invoke QueryInferenceServer API
func CreateQueryInferenceServerRequest() (request *QueryInferenceServerRequest) {
	request = &QueryInferenceServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryInferenceServer", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryInferenceServerResponse creates a response to parse from QueryInferenceServer response
func CreateQueryInferenceServerResponse() (response *QueryInferenceServerResponse) {
	response = &QueryInferenceServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
