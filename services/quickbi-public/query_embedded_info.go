package quickbi_public

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

// QueryEmbeddedInfo invokes the quickbi_public.QueryEmbeddedInfo API synchronously
func (client *Client) QueryEmbeddedInfo(request *QueryEmbeddedInfoRequest) (response *QueryEmbeddedInfoResponse, err error) {
	response = CreateQueryEmbeddedInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEmbeddedInfoWithChan invokes the quickbi_public.QueryEmbeddedInfo API asynchronously
func (client *Client) QueryEmbeddedInfoWithChan(request *QueryEmbeddedInfoRequest) (<-chan *QueryEmbeddedInfoResponse, <-chan error) {
	responseChan := make(chan *QueryEmbeddedInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEmbeddedInfo(request)
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

// QueryEmbeddedInfoWithCallback invokes the quickbi_public.QueryEmbeddedInfo API asynchronously
func (client *Client) QueryEmbeddedInfoWithCallback(request *QueryEmbeddedInfoRequest, callback func(response *QueryEmbeddedInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEmbeddedInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryEmbeddedInfo(request)
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

// QueryEmbeddedInfoRequest is the request struct for api QueryEmbeddedInfo
type QueryEmbeddedInfoRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryEmbeddedInfoResponse is the response struct for api QueryEmbeddedInfo
type QueryEmbeddedInfoResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryEmbeddedInfoRequest creates a request to invoke QueryEmbeddedInfo API
func CreateQueryEmbeddedInfoRequest() (request *QueryEmbeddedInfoRequest) {
	request = &QueryEmbeddedInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryEmbeddedInfo", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEmbeddedInfoResponse creates a response to parse from QueryEmbeddedInfo response
func CreateQueryEmbeddedInfoResponse() (response *QueryEmbeddedInfoResponse) {
	response = &QueryEmbeddedInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}