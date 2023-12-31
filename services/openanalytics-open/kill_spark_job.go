package openanalytics_open

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

// KillSparkJob invokes the openanalytics_open.KillSparkJob API synchronously
func (client *Client) KillSparkJob(request *KillSparkJobRequest) (response *KillSparkJobResponse, err error) {
	response = CreateKillSparkJobResponse()
	err = client.DoAction(request, response)
	return
}

// KillSparkJobWithChan invokes the openanalytics_open.KillSparkJob API asynchronously
func (client *Client) KillSparkJobWithChan(request *KillSparkJobRequest) (<-chan *KillSparkJobResponse, <-chan error) {
	responseChan := make(chan *KillSparkJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.KillSparkJob(request)
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

// KillSparkJobWithCallback invokes the openanalytics_open.KillSparkJob API asynchronously
func (client *Client) KillSparkJobWithCallback(request *KillSparkJobRequest, callback func(response *KillSparkJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *KillSparkJobResponse
		var err error
		defer close(result)
		response, err = client.KillSparkJob(request)
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

// KillSparkJobRequest is the request struct for api KillSparkJob
type KillSparkJobRequest struct {
	*requests.RpcRequest
	JobId  string `position:"Body" name:"JobId"`
	VcName string `position:"Body" name:"VcName"`
}

// KillSparkJobResponse is the response struct for api KillSparkJob
type KillSparkJobResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateKillSparkJobRequest creates a request to invoke KillSparkJob API
func CreateKillSparkJobRequest() (request *KillSparkJobRequest) {
	request = &KillSparkJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "KillSparkJob", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateKillSparkJobResponse creates a response to parse from KillSparkJob response
func CreateKillSparkJobResponse() (response *KillSparkJobResponse) {
	response = &KillSparkJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
