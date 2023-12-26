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

// DeleteJobs invokes the ehpc.DeleteJobs API synchronously
func (client *Client) DeleteJobs(request *DeleteJobsRequest) (response *DeleteJobsResponse, err error) {
	response = CreateDeleteJobsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteJobsWithChan invokes the ehpc.DeleteJobs API asynchronously
func (client *Client) DeleteJobsWithChan(request *DeleteJobsRequest) (<-chan *DeleteJobsResponse, <-chan error) {
	responseChan := make(chan *DeleteJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteJobs(request)
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

// DeleteJobsWithCallback invokes the ehpc.DeleteJobs API asynchronously
func (client *Client) DeleteJobsWithCallback(request *DeleteJobsRequest, callback func(response *DeleteJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteJobsResponse
		var err error
		defer close(result)
		response, err = client.DeleteJobs(request)
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

// DeleteJobsRequest is the request struct for api DeleteJobs
type DeleteJobsRequest struct {
	*requests.RpcRequest
	Jobs      string           `position:"Query" name:"Jobs"`
	ClusterId string           `position:"Query" name:"ClusterId"`
	Async     requests.Boolean `position:"Query" name:"Async"`
}

// DeleteJobsResponse is the response struct for api DeleteJobs
type DeleteJobsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteJobsRequest creates a request to invoke DeleteJobs API
func CreateDeleteJobsRequest() (request *DeleteJobsRequest) {
	request = &DeleteJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DeleteJobs", "", "")
	request.Method = requests.GET
	return
}

// CreateDeleteJobsResponse creates a response to parse from DeleteJobs response
func CreateDeleteJobsResponse() (response *DeleteJobsResponse) {
	response = &DeleteJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
