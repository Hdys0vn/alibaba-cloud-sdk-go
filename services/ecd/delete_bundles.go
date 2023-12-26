package ecd

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

// DeleteBundles invokes the ecd.DeleteBundles API synchronously
func (client *Client) DeleteBundles(request *DeleteBundlesRequest) (response *DeleteBundlesResponse, err error) {
	response = CreateDeleteBundlesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBundlesWithChan invokes the ecd.DeleteBundles API asynchronously
func (client *Client) DeleteBundlesWithChan(request *DeleteBundlesRequest) (<-chan *DeleteBundlesResponse, <-chan error) {
	responseChan := make(chan *DeleteBundlesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBundles(request)
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

// DeleteBundlesWithCallback invokes the ecd.DeleteBundles API asynchronously
func (client *Client) DeleteBundlesWithCallback(request *DeleteBundlesRequest, callback func(response *DeleteBundlesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBundlesResponse
		var err error
		defer close(result)
		response, err = client.DeleteBundles(request)
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

// DeleteBundlesRequest is the request struct for api DeleteBundles
type DeleteBundlesRequest struct {
	*requests.RpcRequest
	BundleId *[]string `position:"Query" name:"BundleId"  type:"Repeated"`
}

// DeleteBundlesResponse is the response struct for api DeleteBundles
type DeleteBundlesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteBundlesRequest creates a request to invoke DeleteBundles API
func CreateDeleteBundlesRequest() (request *DeleteBundlesRequest) {
	request = &DeleteBundlesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeleteBundles", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteBundlesResponse creates a response to parse from DeleteBundles response
func CreateDeleteBundlesResponse() (response *DeleteBundlesResponse) {
	response = &DeleteBundlesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
