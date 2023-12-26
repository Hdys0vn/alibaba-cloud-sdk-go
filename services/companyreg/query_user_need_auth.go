package companyreg

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

// QueryUserNeedAuth invokes the companyreg.QueryUserNeedAuth API synchronously
func (client *Client) QueryUserNeedAuth(request *QueryUserNeedAuthRequest) (response *QueryUserNeedAuthResponse, err error) {
	response = CreateQueryUserNeedAuthResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUserNeedAuthWithChan invokes the companyreg.QueryUserNeedAuth API asynchronously
func (client *Client) QueryUserNeedAuthWithChan(request *QueryUserNeedAuthRequest) (<-chan *QueryUserNeedAuthResponse, <-chan error) {
	responseChan := make(chan *QueryUserNeedAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUserNeedAuth(request)
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

// QueryUserNeedAuthWithCallback invokes the companyreg.QueryUserNeedAuth API asynchronously
func (client *Client) QueryUserNeedAuthWithCallback(request *QueryUserNeedAuthRequest, callback func(response *QueryUserNeedAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUserNeedAuthResponse
		var err error
		defer close(result)
		response, err = client.QueryUserNeedAuth(request)
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

// QueryUserNeedAuthRequest is the request struct for api QueryUserNeedAuth
type QueryUserNeedAuthRequest struct {
	*requests.RpcRequest
}

// QueryUserNeedAuthResponse is the response struct for api QueryUserNeedAuth
type QueryUserNeedAuthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	NeedAuth  bool   `json:"NeedAuth" xml:"NeedAuth"`
}

// CreateQueryUserNeedAuthRequest creates a request to invoke QueryUserNeedAuth API
func CreateQueryUserNeedAuthRequest() (request *QueryUserNeedAuthRequest) {
	request = &QueryUserNeedAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "QueryUserNeedAuth", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryUserNeedAuthResponse creates a response to parse from QueryUserNeedAuth response
func CreateQueryUserNeedAuthResponse() (response *QueryUserNeedAuthResponse) {
	response = &QueryUserNeedAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
