package mts

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

// UpdatePornPipeline invokes the mts.UpdatePornPipeline API synchronously
func (client *Client) UpdatePornPipeline(request *UpdatePornPipelineRequest) (response *UpdatePornPipelineResponse, err error) {
	response = CreateUpdatePornPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePornPipelineWithChan invokes the mts.UpdatePornPipeline API asynchronously
func (client *Client) UpdatePornPipelineWithChan(request *UpdatePornPipelineRequest) (<-chan *UpdatePornPipelineResponse, <-chan error) {
	responseChan := make(chan *UpdatePornPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePornPipeline(request)
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

// UpdatePornPipelineWithCallback invokes the mts.UpdatePornPipeline API asynchronously
func (client *Client) UpdatePornPipelineWithCallback(request *UpdatePornPipelineRequest, callback func(response *UpdatePornPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePornPipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdatePornPipeline(request)
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

// UpdatePornPipelineRequest is the request struct for api UpdatePornPipeline
type UpdatePornPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Name                 string           `position:"Query" name:"Name"`
}

// UpdatePornPipelineResponse is the response struct for api UpdatePornPipeline
type UpdatePornPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateUpdatePornPipelineRequest creates a request to invoke UpdatePornPipeline API
func CreateUpdatePornPipelineRequest() (request *UpdatePornPipelineRequest) {
	request = &UpdatePornPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdatePornPipeline", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePornPipelineResponse creates a response to parse from UpdatePornPipeline response
func CreateUpdatePornPipelineResponse() (response *UpdatePornPipelineResponse) {
	response = &UpdatePornPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
