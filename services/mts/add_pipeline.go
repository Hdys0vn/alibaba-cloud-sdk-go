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

// AddPipeline invokes the mts.AddPipeline API synchronously
func (client *Client) AddPipeline(request *AddPipelineRequest) (response *AddPipelineResponse, err error) {
	response = CreateAddPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// AddPipelineWithChan invokes the mts.AddPipeline API asynchronously
func (client *Client) AddPipelineWithChan(request *AddPipelineRequest) (<-chan *AddPipelineResponse, <-chan error) {
	responseChan := make(chan *AddPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddPipeline(request)
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

// AddPipelineWithCallback invokes the mts.AddPipeline API asynchronously
func (client *Client) AddPipelineWithCallback(request *AddPipelineRequest, callback func(response *AddPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddPipeline(request)
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

// AddPipelineRequest is the request struct for api AddPipeline
type AddPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Role                 string           `position:"Query" name:"Role"`
	Speed                string           `position:"Query" name:"Speed"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	SpeedLevel           requests.Integer `position:"Query" name:"SpeedLevel"`
}

// AddPipelineResponse is the response struct for api AddPipeline
type AddPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateAddPipelineRequest creates a request to invoke AddPipeline API
func CreateAddPipelineRequest() (request *AddPipelineRequest) {
	request = &AddPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddPipeline", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddPipelineResponse creates a response to parse from AddPipeline response
func CreateAddPipelineResponse() (response *AddPipelineResponse) {
	response = &AddPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
