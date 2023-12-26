package cdn

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

// DeleteCdnDeliverTask invokes the cdn.DeleteCdnDeliverTask API synchronously
func (client *Client) DeleteCdnDeliverTask(request *DeleteCdnDeliverTaskRequest) (response *DeleteCdnDeliverTaskResponse, err error) {
	response = CreateDeleteCdnDeliverTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCdnDeliverTaskWithChan invokes the cdn.DeleteCdnDeliverTask API asynchronously
func (client *Client) DeleteCdnDeliverTaskWithChan(request *DeleteCdnDeliverTaskRequest) (<-chan *DeleteCdnDeliverTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteCdnDeliverTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCdnDeliverTask(request)
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

// DeleteCdnDeliverTaskWithCallback invokes the cdn.DeleteCdnDeliverTask API asynchronously
func (client *Client) DeleteCdnDeliverTaskWithCallback(request *DeleteCdnDeliverTaskRequest, callback func(response *DeleteCdnDeliverTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCdnDeliverTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteCdnDeliverTask(request)
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

// DeleteCdnDeliverTaskRequest is the request struct for api DeleteCdnDeliverTask
type DeleteCdnDeliverTaskRequest struct {
	*requests.RpcRequest
	DeliverId requests.Integer `position:"Query" name:"DeliverId"`
}

// DeleteCdnDeliverTaskResponse is the response struct for api DeleteCdnDeliverTask
type DeleteCdnDeliverTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCdnDeliverTaskRequest creates a request to invoke DeleteCdnDeliverTask API
func CreateDeleteCdnDeliverTaskRequest() (request *DeleteCdnDeliverTaskRequest) {
	request = &DeleteCdnDeliverTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DeleteCdnDeliverTask", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCdnDeliverTaskResponse creates a response to parse from DeleteCdnDeliverTask response
func CreateDeleteCdnDeliverTaskResponse() (response *DeleteCdnDeliverTaskResponse) {
	response = &DeleteCdnDeliverTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
