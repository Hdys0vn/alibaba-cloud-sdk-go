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

// SubmitFpShotJob invokes the mts.SubmitFpShotJob API synchronously
func (client *Client) SubmitFpShotJob(request *SubmitFpShotJobRequest) (response *SubmitFpShotJobResponse, err error) {
	response = CreateSubmitFpShotJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitFpShotJobWithChan invokes the mts.SubmitFpShotJob API asynchronously
func (client *Client) SubmitFpShotJobWithChan(request *SubmitFpShotJobRequest) (<-chan *SubmitFpShotJobResponse, <-chan error) {
	responseChan := make(chan *SubmitFpShotJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitFpShotJob(request)
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

// SubmitFpShotJobWithCallback invokes the mts.SubmitFpShotJob API asynchronously
func (client *Client) SubmitFpShotJobWithCallback(request *SubmitFpShotJobRequest, callback func(response *SubmitFpShotJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitFpShotJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitFpShotJob(request)
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

// SubmitFpShotJobRequest is the request struct for api SubmitFpShotJob
type SubmitFpShotJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UserData             string           `position:"Query" name:"UserData"`
	FpShotConfig         string           `position:"Query" name:"FpShotConfig"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitFpShotJobResponse is the response struct for api SubmitFpShotJob
type SubmitFpShotJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitFpShotJobRequest creates a request to invoke SubmitFpShotJob API
func CreateSubmitFpShotJobRequest() (request *SubmitFpShotJobRequest) {
	request = &SubmitFpShotJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitFpShotJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitFpShotJobResponse creates a response to parse from SubmitFpShotJob response
func CreateSubmitFpShotJobResponse() (response *SubmitFpShotJobResponse) {
	response = &SubmitFpShotJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
