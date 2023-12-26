package oos

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

// DeletePatchBaseline invokes the oos.DeletePatchBaseline API synchronously
func (client *Client) DeletePatchBaseline(request *DeletePatchBaselineRequest) (response *DeletePatchBaselineResponse, err error) {
	response = CreateDeletePatchBaselineResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePatchBaselineWithChan invokes the oos.DeletePatchBaseline API asynchronously
func (client *Client) DeletePatchBaselineWithChan(request *DeletePatchBaselineRequest) (<-chan *DeletePatchBaselineResponse, <-chan error) {
	responseChan := make(chan *DeletePatchBaselineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePatchBaseline(request)
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

// DeletePatchBaselineWithCallback invokes the oos.DeletePatchBaseline API asynchronously
func (client *Client) DeletePatchBaselineWithCallback(request *DeletePatchBaselineRequest, callback func(response *DeletePatchBaselineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePatchBaselineResponse
		var err error
		defer close(result)
		response, err = client.DeletePatchBaseline(request)
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

// DeletePatchBaselineRequest is the request struct for api DeletePatchBaseline
type DeletePatchBaselineRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// DeletePatchBaselineResponse is the response struct for api DeletePatchBaseline
type DeletePatchBaselineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePatchBaselineRequest creates a request to invoke DeletePatchBaseline API
func CreateDeletePatchBaselineRequest() (request *DeletePatchBaselineRequest) {
	request = &DeletePatchBaselineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "DeletePatchBaseline", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePatchBaselineResponse creates a response to parse from DeletePatchBaseline response
func CreateDeletePatchBaselineResponse() (response *DeletePatchBaselineResponse) {
	response = &DeletePatchBaselineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
