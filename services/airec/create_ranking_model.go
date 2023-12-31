package airec

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

// CreateRankingModel invokes the airec.CreateRankingModel API synchronously
func (client *Client) CreateRankingModel(request *CreateRankingModelRequest) (response *CreateRankingModelResponse, err error) {
	response = CreateCreateRankingModelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRankingModelWithChan invokes the airec.CreateRankingModel API asynchronously
func (client *Client) CreateRankingModelWithChan(request *CreateRankingModelRequest) (<-chan *CreateRankingModelResponse, <-chan error) {
	responseChan := make(chan *CreateRankingModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRankingModel(request)
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

// CreateRankingModelWithCallback invokes the airec.CreateRankingModel API asynchronously
func (client *Client) CreateRankingModelWithCallback(request *CreateRankingModelRequest, callback func(response *CreateRankingModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRankingModelResponse
		var err error
		defer close(result)
		response, err = client.CreateRankingModel(request)
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

// CreateRankingModelRequest is the request struct for api CreateRankingModel
type CreateRankingModelRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"instanceId"`
	DryRun     requests.Boolean `position:"Query" name:"dryRun"`
}

// CreateRankingModelResponse is the response struct for api CreateRankingModel
type CreateRankingModelResponse struct {
	*responses.BaseResponse
	Code      string                     `json:"code" xml:"code"`
	RequestId string                     `json:"requestId" xml:"requestId"`
	Message   string                     `json:"message" xml:"message"`
	Result    ResultInCreateRankingModel `json:"result" xml:"result"`
}

// CreateCreateRankingModelRequest creates a request to invoke CreateRankingModel API
func CreateCreateRankingModelRequest() (request *CreateRankingModelRequest) {
	request = &CreateRankingModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "CreateRankingModel", "/v2/openapi/instances/[instanceId]/ranking-models", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRankingModelResponse creates a response to parse from CreateRankingModel response
func CreateCreateRankingModelResponse() (response *CreateRankingModelResponse) {
	response = &CreateRankingModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
