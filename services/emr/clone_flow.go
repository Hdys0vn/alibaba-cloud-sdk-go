package emr

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

// CloneFlow invokes the emr.CloneFlow API synchronously
func (client *Client) CloneFlow(request *CloneFlowRequest) (response *CloneFlowResponse, err error) {
	response = CreateCloneFlowResponse()
	err = client.DoAction(request, response)
	return
}

// CloneFlowWithChan invokes the emr.CloneFlow API asynchronously
func (client *Client) CloneFlowWithChan(request *CloneFlowRequest) (<-chan *CloneFlowResponse, <-chan error) {
	responseChan := make(chan *CloneFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloneFlow(request)
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

// CloneFlowWithCallback invokes the emr.CloneFlow API asynchronously
func (client *Client) CloneFlowWithCallback(request *CloneFlowRequest, callback func(response *CloneFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloneFlowResponse
		var err error
		defer close(result)
		response, err = client.CloneFlow(request)
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

// CloneFlowRequest is the request struct for api CloneFlow
type CloneFlowRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// CloneFlowResponse is the response struct for api CloneFlow
type CloneFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateCloneFlowRequest creates a request to invoke CloneFlow API
func CreateCloneFlowRequest() (request *CloneFlowRequest) {
	request = &CloneFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "CloneFlow", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCloneFlowResponse creates a response to parse from CloneFlow response
func CreateCloneFlowResponse() (response *CloneFlowResponse) {
	response = &CloneFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
