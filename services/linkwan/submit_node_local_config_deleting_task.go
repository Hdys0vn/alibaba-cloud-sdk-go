package linkwan

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

// SubmitNodeLocalConfigDeletingTask invokes the linkwan.SubmitNodeLocalConfigDeletingTask API synchronously
func (client *Client) SubmitNodeLocalConfigDeletingTask(request *SubmitNodeLocalConfigDeletingTaskRequest) (response *SubmitNodeLocalConfigDeletingTaskResponse, err error) {
	response = CreateSubmitNodeLocalConfigDeletingTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitNodeLocalConfigDeletingTaskWithChan invokes the linkwan.SubmitNodeLocalConfigDeletingTask API asynchronously
func (client *Client) SubmitNodeLocalConfigDeletingTaskWithChan(request *SubmitNodeLocalConfigDeletingTaskRequest) (<-chan *SubmitNodeLocalConfigDeletingTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitNodeLocalConfigDeletingTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitNodeLocalConfigDeletingTask(request)
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

// SubmitNodeLocalConfigDeletingTaskWithCallback invokes the linkwan.SubmitNodeLocalConfigDeletingTask API asynchronously
func (client *Client) SubmitNodeLocalConfigDeletingTaskWithCallback(request *SubmitNodeLocalConfigDeletingTaskRequest, callback func(response *SubmitNodeLocalConfigDeletingTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitNodeLocalConfigDeletingTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitNodeLocalConfigDeletingTask(request)
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

// SubmitNodeLocalConfigDeletingTaskRequest is the request struct for api SubmitNodeLocalConfigDeletingTask
type SubmitNodeLocalConfigDeletingTaskRequest struct {
	*requests.RpcRequest
	DevEui      string `position:"Query" name:"DevEui"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// SubmitNodeLocalConfigDeletingTaskResponse is the response struct for api SubmitNodeLocalConfigDeletingTask
type SubmitNodeLocalConfigDeletingTaskResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         int64  `json:"Data" xml:"Data"`
}

// CreateSubmitNodeLocalConfigDeletingTaskRequest creates a request to invoke SubmitNodeLocalConfigDeletingTask API
func CreateSubmitNodeLocalConfigDeletingTaskRequest() (request *SubmitNodeLocalConfigDeletingTaskRequest) {
	request = &SubmitNodeLocalConfigDeletingTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "SubmitNodeLocalConfigDeletingTask", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitNodeLocalConfigDeletingTaskResponse creates a response to parse from SubmitNodeLocalConfigDeletingTask response
func CreateSubmitNodeLocalConfigDeletingTaskResponse() (response *SubmitNodeLocalConfigDeletingTaskResponse) {
	response = &SubmitNodeLocalConfigDeletingTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
