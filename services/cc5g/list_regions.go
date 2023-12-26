package cc5g

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

// ListRegions invokes the cc5g.ListRegions API synchronously
func (client *Client) ListRegions(request *ListRegionsRequest) (response *ListRegionsResponse, err error) {
	response = CreateListRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRegionsWithChan invokes the cc5g.ListRegions API asynchronously
func (client *Client) ListRegionsWithChan(request *ListRegionsRequest) (<-chan *ListRegionsResponse, <-chan error) {
	responseChan := make(chan *ListRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRegions(request)
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

// ListRegionsWithCallback invokes the cc5g.ListRegions API asynchronously
func (client *Client) ListRegionsWithCallback(request *ListRegionsRequest, callback func(response *ListRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRegionsResponse
		var err error
		defer close(result)
		response, err = client.ListRegions(request)
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

// ListRegionsRequest is the request struct for api ListRegions
type ListRegionsRequest struct {
	*requests.RpcRequest
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ListRegionsResponse is the response struct for api ListRegions
type ListRegionsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Regions   []Region `json:"Regions" xml:"Regions"`
}

// CreateListRegionsRequest creates a request to invoke ListRegions API
func CreateListRegionsRequest() (request *ListRegionsRequest) {
	request = &ListRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "ListRegions", "fivegcc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListRegionsResponse creates a response to parse from ListRegions response
func CreateListRegionsResponse() (response *ListRegionsResponse) {
	response = &ListRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
