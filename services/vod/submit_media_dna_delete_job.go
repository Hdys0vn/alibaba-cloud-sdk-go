package vod

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

// SubmitMediaDNADeleteJob invokes the vod.SubmitMediaDNADeleteJob API synchronously
func (client *Client) SubmitMediaDNADeleteJob(request *SubmitMediaDNADeleteJobRequest) (response *SubmitMediaDNADeleteJobResponse, err error) {
	response = CreateSubmitMediaDNADeleteJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitMediaDNADeleteJobWithChan invokes the vod.SubmitMediaDNADeleteJob API asynchronously
func (client *Client) SubmitMediaDNADeleteJobWithChan(request *SubmitMediaDNADeleteJobRequest) (<-chan *SubmitMediaDNADeleteJobResponse, <-chan error) {
	responseChan := make(chan *SubmitMediaDNADeleteJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitMediaDNADeleteJob(request)
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

// SubmitMediaDNADeleteJobWithCallback invokes the vod.SubmitMediaDNADeleteJob API asynchronously
func (client *Client) SubmitMediaDNADeleteJobWithCallback(request *SubmitMediaDNADeleteJobRequest, callback func(response *SubmitMediaDNADeleteJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitMediaDNADeleteJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitMediaDNADeleteJob(request)
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

// SubmitMediaDNADeleteJobRequest is the request struct for api SubmitMediaDNADeleteJob
type SubmitMediaDNADeleteJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	FpDBId               string `position:"Query" name:"FpDBId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
}

// SubmitMediaDNADeleteJobResponse is the response struct for api SubmitMediaDNADeleteJob
type SubmitMediaDNADeleteJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MediaId   string `json:"MediaId" xml:"MediaId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitMediaDNADeleteJobRequest creates a request to invoke SubmitMediaDNADeleteJob API
func CreateSubmitMediaDNADeleteJobRequest() (request *SubmitMediaDNADeleteJobRequest) {
	request = &SubmitMediaDNADeleteJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitMediaDNADeleteJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitMediaDNADeleteJobResponse creates a response to parse from SubmitMediaDNADeleteJob response
func CreateSubmitMediaDNADeleteJobResponse() (response *SubmitMediaDNADeleteJobResponse) {
	response = &SubmitMediaDNADeleteJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
