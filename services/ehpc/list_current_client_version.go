package ehpc

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

// ListCurrentClientVersion invokes the ehpc.ListCurrentClientVersion API synchronously
func (client *Client) ListCurrentClientVersion(request *ListCurrentClientVersionRequest) (response *ListCurrentClientVersionResponse, err error) {
	response = CreateListCurrentClientVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ListCurrentClientVersionWithChan invokes the ehpc.ListCurrentClientVersion API asynchronously
func (client *Client) ListCurrentClientVersionWithChan(request *ListCurrentClientVersionRequest) (<-chan *ListCurrentClientVersionResponse, <-chan error) {
	responseChan := make(chan *ListCurrentClientVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCurrentClientVersion(request)
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

// ListCurrentClientVersionWithCallback invokes the ehpc.ListCurrentClientVersion API asynchronously
func (client *Client) ListCurrentClientVersionWithCallback(request *ListCurrentClientVersionRequest, callback func(response *ListCurrentClientVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCurrentClientVersionResponse
		var err error
		defer close(result)
		response, err = client.ListCurrentClientVersion(request)
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

// ListCurrentClientVersionRequest is the request struct for api ListCurrentClientVersion
type ListCurrentClientVersionRequest struct {
	*requests.RpcRequest
}

// ListCurrentClientVersionResponse is the response struct for api ListCurrentClientVersion
type ListCurrentClientVersionResponse struct {
	*responses.BaseResponse
	ClientVersion string `json:"ClientVersion" xml:"ClientVersion"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateListCurrentClientVersionRequest creates a request to invoke ListCurrentClientVersion API
func CreateListCurrentClientVersionRequest() (request *ListCurrentClientVersionRequest) {
	request = &ListCurrentClientVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListCurrentClientVersion", "", "")
	request.Method = requests.GET
	return
}

// CreateListCurrentClientVersionResponse creates a response to parse from ListCurrentClientVersion response
func CreateListCurrentClientVersionResponse() (response *ListCurrentClientVersionResponse) {
	response = &ListCurrentClientVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
