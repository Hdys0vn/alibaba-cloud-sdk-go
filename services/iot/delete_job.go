package iot

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

// DeleteJob invokes the iot.DeleteJob API synchronously
func (client *Client) DeleteJob(request *DeleteJobRequest) (response *DeleteJobResponse, err error) {
	response = CreateDeleteJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteJobWithChan invokes the iot.DeleteJob API asynchronously
func (client *Client) DeleteJobWithChan(request *DeleteJobRequest) (<-chan *DeleteJobResponse, <-chan error) {
	responseChan := make(chan *DeleteJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteJob(request)
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

// DeleteJobWithCallback invokes the iot.DeleteJob API asynchronously
func (client *Client) DeleteJobWithCallback(request *DeleteJobRequest, callback func(response *DeleteJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteJob(request)
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

// DeleteJobRequest is the request struct for api DeleteJob
type DeleteJobRequest struct {
	*requests.RpcRequest
	JobId         string `position:"Query" name:"JobId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteJobResponse is the response struct for api DeleteJob
type DeleteJobResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteJobRequest creates a request to invoke DeleteJob API
func CreateDeleteJobRequest() (request *DeleteJobRequest) {
	request = &DeleteJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteJob", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteJobResponse creates a response to parse from DeleteJob response
func CreateDeleteJobResponse() (response *DeleteJobResponse) {
	response = &DeleteJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
