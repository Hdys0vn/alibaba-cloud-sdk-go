package schedulerx2

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

// EnableJob invokes the schedulerx2.EnableJob API synchronously
func (client *Client) EnableJob(request *EnableJobRequest) (response *EnableJobResponse, err error) {
	response = CreateEnableJobResponse()
	err = client.DoAction(request, response)
	return
}

// EnableJobWithChan invokes the schedulerx2.EnableJob API asynchronously
func (client *Client) EnableJobWithChan(request *EnableJobRequest) (<-chan *EnableJobResponse, <-chan error) {
	responseChan := make(chan *EnableJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableJob(request)
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

// EnableJobWithCallback invokes the schedulerx2.EnableJob API asynchronously
func (client *Client) EnableJobWithCallback(request *EnableJobRequest, callback func(response *EnableJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableJobResponse
		var err error
		defer close(result)
		response, err = client.EnableJob(request)
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

// EnableJobRequest is the request struct for api EnableJob
type EnableJobRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Query" name:"NamespaceSource"`
	GroupId         string           `position:"Query" name:"GroupId"`
	JobId           requests.Integer `position:"Query" name:"JobId"`
	Namespace       string           `position:"Query" name:"Namespace"`
}

// EnableJobResponse is the response struct for api EnableJob
type EnableJobResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEnableJobRequest creates a request to invoke EnableJob API
func CreateEnableJobRequest() (request *EnableJobRequest) {
	request = &EnableJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "EnableJob", "", "")
	request.Method = requests.GET
	return
}

// CreateEnableJobResponse creates a response to parse from EnableJob response
func CreateEnableJobResponse() (response *EnableJobResponse) {
	response = &EnableJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
