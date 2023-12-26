package elasticsearch

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

// ResumeLogstashTask invokes the elasticsearch.ResumeLogstashTask API synchronously
func (client *Client) ResumeLogstashTask(request *ResumeLogstashTaskRequest) (response *ResumeLogstashTaskResponse, err error) {
	response = CreateResumeLogstashTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeLogstashTaskWithChan invokes the elasticsearch.ResumeLogstashTask API asynchronously
func (client *Client) ResumeLogstashTaskWithChan(request *ResumeLogstashTaskRequest) (<-chan *ResumeLogstashTaskResponse, <-chan error) {
	responseChan := make(chan *ResumeLogstashTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeLogstashTask(request)
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

// ResumeLogstashTaskWithCallback invokes the elasticsearch.ResumeLogstashTask API asynchronously
func (client *Client) ResumeLogstashTaskWithCallback(request *ResumeLogstashTaskRequest, callback func(response *ResumeLogstashTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeLogstashTaskResponse
		var err error
		defer close(result)
		response, err = client.ResumeLogstashTask(request)
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

// ResumeLogstashTaskRequest is the request struct for api ResumeLogstashTask
type ResumeLogstashTaskRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// ResumeLogstashTaskResponse is the response struct for api ResumeLogstashTask
type ResumeLogstashTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateResumeLogstashTaskRequest creates a request to invoke ResumeLogstashTask API
func CreateResumeLogstashTaskRequest() (request *ResumeLogstashTaskRequest) {
	request = &ResumeLogstashTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ResumeLogstashTask", "/openapi/logstashes/[InstanceId]/actions/resume", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResumeLogstashTaskResponse creates a response to parse from ResumeLogstashTask response
func CreateResumeLogstashTaskResponse() (response *ResumeLogstashTaskResponse) {
	response = &ResumeLogstashTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}