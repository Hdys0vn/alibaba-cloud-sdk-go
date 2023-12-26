package pairecservice

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

// CreateCalculationJobs invokes the pairecservice.CreateCalculationJobs API synchronously
func (client *Client) CreateCalculationJobs(request *CreateCalculationJobsRequest) (response *CreateCalculationJobsResponse, err error) {
	response = CreateCreateCalculationJobsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCalculationJobsWithChan invokes the pairecservice.CreateCalculationJobs API asynchronously
func (client *Client) CreateCalculationJobsWithChan(request *CreateCalculationJobsRequest) (<-chan *CreateCalculationJobsResponse, <-chan error) {
	responseChan := make(chan *CreateCalculationJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCalculationJobs(request)
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

// CreateCalculationJobsWithCallback invokes the pairecservice.CreateCalculationJobs API asynchronously
func (client *Client) CreateCalculationJobsWithCallback(request *CreateCalculationJobsRequest, callback func(response *CreateCalculationJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCalculationJobsResponse
		var err error
		defer close(result)
		response, err = client.CreateCalculationJobs(request)
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

// CreateCalculationJobsRequest is the request struct for api CreateCalculationJobs
type CreateCalculationJobsRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateCalculationJobsResponse is the response struct for api CreateCalculationJobs
type CreateCalculationJobsResponse struct {
	*responses.BaseResponse
	RequestId         string   `json:"RequestId" xml:"RequestId"`
	CalculationJobIds []string `json:"CalculationJobIds" xml:"CalculationJobIds"`
}

// CreateCreateCalculationJobsRequest creates a request to invoke CreateCalculationJobs API
func CreateCreateCalculationJobsRequest() (request *CreateCalculationJobsRequest) {
	request = &CreateCalculationJobsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "CreateCalculationJobs", "/api/v1/batch/calculationjobs/create", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCalculationJobsResponse creates a response to parse from CreateCalculationJobs response
func CreateCreateCalculationJobsResponse() (response *CreateCalculationJobsResponse) {
	response = &CreateCalculationJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
