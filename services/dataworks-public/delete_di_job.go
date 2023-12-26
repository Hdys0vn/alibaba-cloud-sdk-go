package dataworks_public

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

// DeleteDIJob invokes the dataworks_public.DeleteDIJob API synchronously
func (client *Client) DeleteDIJob(request *DeleteDIJobRequest) (response *DeleteDIJobResponse, err error) {
	response = CreateDeleteDIJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDIJobWithChan invokes the dataworks_public.DeleteDIJob API asynchronously
func (client *Client) DeleteDIJobWithChan(request *DeleteDIJobRequest) (<-chan *DeleteDIJobResponse, <-chan error) {
	responseChan := make(chan *DeleteDIJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDIJob(request)
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

// DeleteDIJobWithCallback invokes the dataworks_public.DeleteDIJob API asynchronously
func (client *Client) DeleteDIJobWithCallback(request *DeleteDIJobRequest, callback func(response *DeleteDIJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDIJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteDIJob(request)
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

// DeleteDIJobRequest is the request struct for api DeleteDIJob
type DeleteDIJobRequest struct {
	*requests.RpcRequest
	DIJobId requests.Integer `position:"Body" name:"DIJobId"`
}

// DeleteDIJobResponse is the response struct for api DeleteDIJob
type DeleteDIJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDIJobRequest creates a request to invoke DeleteDIJob API
func CreateDeleteDIJobRequest() (request *DeleteDIJobRequest) {
	request = &DeleteDIJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteDIJob", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDIJobResponse creates a response to parse from DeleteDIJob response
func CreateDeleteDIJobResponse() (response *DeleteDIJobResponse) {
	response = &DeleteDIJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
