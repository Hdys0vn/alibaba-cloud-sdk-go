package live

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

// CreateEdgeTranscodeJob invokes the live.CreateEdgeTranscodeJob API synchronously
func (client *Client) CreateEdgeTranscodeJob(request *CreateEdgeTranscodeJobRequest) (response *CreateEdgeTranscodeJobResponse, err error) {
	response = CreateCreateEdgeTranscodeJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateEdgeTranscodeJobWithChan invokes the live.CreateEdgeTranscodeJob API asynchronously
func (client *Client) CreateEdgeTranscodeJobWithChan(request *CreateEdgeTranscodeJobRequest) (<-chan *CreateEdgeTranscodeJobResponse, <-chan error) {
	responseChan := make(chan *CreateEdgeTranscodeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateEdgeTranscodeJob(request)
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

// CreateEdgeTranscodeJobWithCallback invokes the live.CreateEdgeTranscodeJob API asynchronously
func (client *Client) CreateEdgeTranscodeJobWithCallback(request *CreateEdgeTranscodeJobRequest, callback func(response *CreateEdgeTranscodeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateEdgeTranscodeJobResponse
		var err error
		defer close(result)
		response, err = client.CreateEdgeTranscodeJob(request)
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

// CreateEdgeTranscodeJobRequest is the request struct for api CreateEdgeTranscodeJob
type CreateEdgeTranscodeJobRequest struct {
	*requests.RpcRequest
	StreamInput  string           `position:"Query" name:"StreamInput"`
	StreamOutput string           `position:"Query" name:"StreamOutput"`
	ClusterId    string           `position:"Query" name:"ClusterId"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId   string           `position:"Query" name:"TemplateId"`
	Name         string           `position:"Query" name:"Name"`
}

// CreateEdgeTranscodeJobResponse is the response struct for api CreateEdgeTranscodeJob
type CreateEdgeTranscodeJobResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateEdgeTranscodeJobRequest creates a request to invoke CreateEdgeTranscodeJob API
func CreateCreateEdgeTranscodeJobRequest() (request *CreateEdgeTranscodeJobRequest) {
	request = &CreateEdgeTranscodeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateEdgeTranscodeJob", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateEdgeTranscodeJobResponse creates a response to parse from CreateEdgeTranscodeJob response
func CreateCreateEdgeTranscodeJobResponse() (response *CreateEdgeTranscodeJobResponse) {
	response = &CreateEdgeTranscodeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}