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

// SubmitTerrorismJob invokes the mts.SubmitTerrorismJob API synchronously
func (client *Client) SubmitTerrorismJob(request *SubmitTerrorismJobRequest) (response *SubmitTerrorismJobResponse, err error) {
	response = CreateSubmitTerrorismJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitTerrorismJobWithChan invokes the mts.SubmitTerrorismJob API asynchronously
func (client *Client) SubmitTerrorismJobWithChan(request *SubmitTerrorismJobRequest) (<-chan *SubmitTerrorismJobResponse, <-chan error) {
	responseChan := make(chan *SubmitTerrorismJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitTerrorismJob(request)
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

// SubmitTerrorismJobWithCallback invokes the mts.SubmitTerrorismJob API asynchronously
func (client *Client) SubmitTerrorismJobWithCallback(request *SubmitTerrorismJobRequest, callback func(response *SubmitTerrorismJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitTerrorismJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitTerrorismJob(request)
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

// SubmitTerrorismJobRequest is the request struct for api SubmitTerrorismJob
type SubmitTerrorismJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TerrorismConfig      string           `position:"Query" name:"TerrorismConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitTerrorismJobResponse is the response struct for api SubmitTerrorismJob
type SubmitTerrorismJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitTerrorismJobRequest creates a request to invoke SubmitTerrorismJob API
func CreateSubmitTerrorismJobRequest() (request *SubmitTerrorismJobRequest) {
	request = &SubmitTerrorismJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitTerrorismJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitTerrorismJobResponse creates a response to parse from SubmitTerrorismJob response
func CreateSubmitTerrorismJobResponse() (response *SubmitTerrorismJobResponse) {
	response = &SubmitTerrorismJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
