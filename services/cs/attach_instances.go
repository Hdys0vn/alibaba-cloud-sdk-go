package cs

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

// AttachInstances invokes the cs.AttachInstances API synchronously
func (client *Client) AttachInstances(request *AttachInstancesRequest) (response *AttachInstancesResponse, err error) {
	response = CreateAttachInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// AttachInstancesWithChan invokes the cs.AttachInstances API asynchronously
func (client *Client) AttachInstancesWithChan(request *AttachInstancesRequest) (<-chan *AttachInstancesResponse, <-chan error) {
	responseChan := make(chan *AttachInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachInstances(request)
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

// AttachInstancesWithCallback invokes the cs.AttachInstances API asynchronously
func (client *Client) AttachInstancesWithCallback(request *AttachInstancesRequest, callback func(response *AttachInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachInstancesResponse
		var err error
		defer close(result)
		response, err = client.AttachInstances(request)
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

// AttachInstancesRequest is the request struct for api AttachInstances
type AttachInstancesRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// AttachInstancesResponse is the response struct for api AttachInstances
type AttachInstancesResponse struct {
	*responses.BaseResponse
	TaskId string     `json:"task_id" xml:"task_id"`
	List   []ListItem `json:"list" xml:"list"`
}

// CreateAttachInstancesRequest creates a request to invoke AttachInstances API
func CreateAttachInstancesRequest() (request *AttachInstancesRequest) {
	request = &AttachInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "AttachInstances", "/clusters/[ClusterId]/attach", "", "")
	request.Method = requests.POST
	return
}

// CreateAttachInstancesResponse creates a response to parse from AttachInstances response
func CreateAttachInstancesResponse() (response *AttachInstancesResponse) {
	response = &AttachInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
