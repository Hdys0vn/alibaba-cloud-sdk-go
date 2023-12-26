package das

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

// CreateCacheAnalysisJob invokes the das.CreateCacheAnalysisJob API synchronously
func (client *Client) CreateCacheAnalysisJob(request *CreateCacheAnalysisJobRequest) (response *CreateCacheAnalysisJobResponse, err error) {
	response = CreateCreateCacheAnalysisJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCacheAnalysisJobWithChan invokes the das.CreateCacheAnalysisJob API asynchronously
func (client *Client) CreateCacheAnalysisJobWithChan(request *CreateCacheAnalysisJobRequest) (<-chan *CreateCacheAnalysisJobResponse, <-chan error) {
	responseChan := make(chan *CreateCacheAnalysisJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCacheAnalysisJob(request)
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

// CreateCacheAnalysisJobWithCallback invokes the das.CreateCacheAnalysisJob API asynchronously
func (client *Client) CreateCacheAnalysisJobWithCallback(request *CreateCacheAnalysisJobRequest, callback func(response *CreateCacheAnalysisJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCacheAnalysisJobResponse
		var err error
		defer close(result)
		response, err = client.CreateCacheAnalysisJob(request)
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

// CreateCacheAnalysisJobRequest is the request struct for api CreateCacheAnalysisJob
type CreateCacheAnalysisJobRequest struct {
	*requests.RpcRequest
	BackupSetId string `position:"Query" name:"BackupSetId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	NodeId      string `position:"Query" name:"NodeId"`
	Separators  string `position:"Query" name:"Separators"`
}

// CreateCacheAnalysisJobResponse is the response struct for api CreateCacheAnalysisJob
type CreateCacheAnalysisJobResponse struct {
	*responses.BaseResponse
	Message   string                       `json:"Message" xml:"Message"`
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Code      string                       `json:"Code" xml:"Code"`
	Success   string                       `json:"Success" xml:"Success"`
	Data      DataInCreateCacheAnalysisJob `json:"Data" xml:"Data"`
}

// CreateCreateCacheAnalysisJobRequest creates a request to invoke CreateCacheAnalysisJob API
func CreateCreateCacheAnalysisJobRequest() (request *CreateCacheAnalysisJobRequest) {
	request = &CreateCacheAnalysisJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "CreateCacheAnalysisJob", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCacheAnalysisJobResponse creates a response to parse from CreateCacheAnalysisJob response
func CreateCreateCacheAnalysisJobResponse() (response *CreateCacheAnalysisJobResponse) {
	response = &CreateCacheAnalysisJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
