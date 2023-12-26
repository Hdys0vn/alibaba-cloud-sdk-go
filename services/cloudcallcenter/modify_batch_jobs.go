package cloudcallcenter

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

// ModifyBatchJobs invokes the cloudcallcenter.ModifyBatchJobs API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifybatchjobs.html
func (client *Client) ModifyBatchJobs(request *ModifyBatchJobsRequest) (response *ModifyBatchJobsResponse, err error) {
	response = CreateModifyBatchJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBatchJobsWithChan invokes the cloudcallcenter.ModifyBatchJobs API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifybatchjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBatchJobsWithChan(request *ModifyBatchJobsRequest) (<-chan *ModifyBatchJobsResponse, <-chan error) {
	responseChan := make(chan *ModifyBatchJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBatchJobs(request)
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

// ModifyBatchJobsWithCallback invokes the cloudcallcenter.ModifyBatchJobs API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifybatchjobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBatchJobsWithCallback(request *ModifyBatchJobsRequest, callback func(response *ModifyBatchJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBatchJobsResponse
		var err error
		defer close(result)
		response, err = client.ModifyBatchJobs(request)
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

// ModifyBatchJobsRequest is the request struct for api ModifyBatchJobs
type ModifyBatchJobsRequest struct {
	*requests.RpcRequest
	Description   string           `position:"Query" name:"Description"`
	JobFilePath   string           `position:"Query" name:"JobFilePath"`
	CallingNumber *[]string        `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Submitted     requests.Boolean `position:"Query" name:"Submitted"`
	StrategyJson  string           `position:"Query" name:"StrategyJson"`
	JobGroupId    string           `position:"Query" name:"JobGroupId"`
	Name          string           `position:"Query" name:"Name"`
	ScenarioId    string           `position:"Query" name:"ScenarioId"`
}

// ModifyBatchJobsResponse is the response struct for api ModifyBatchJobs
type ModifyBatchJobsResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateModifyBatchJobsRequest creates a request to invoke ModifyBatchJobs API
func CreateModifyBatchJobsRequest() (request *ModifyBatchJobsRequest) {
	request = &ModifyBatchJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyBatchJobs", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyBatchJobsResponse creates a response to parse from ModifyBatchJobs response
func CreateModifyBatchJobsResponse() (response *ModifyBatchJobsResponse) {
	response = &ModifyBatchJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
