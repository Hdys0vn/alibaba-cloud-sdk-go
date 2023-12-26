package apds

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

// StopSyncMigrationJob invokes the apds.StopSyncMigrationJob API synchronously
func (client *Client) StopSyncMigrationJob(request *StopSyncMigrationJobRequest) (response *StopSyncMigrationJobResponse, err error) {
	response = CreateStopSyncMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// StopSyncMigrationJobWithChan invokes the apds.StopSyncMigrationJob API asynchronously
func (client *Client) StopSyncMigrationJobWithChan(request *StopSyncMigrationJobRequest) (<-chan *StopSyncMigrationJobResponse, <-chan error) {
	responseChan := make(chan *StopSyncMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopSyncMigrationJob(request)
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

// StopSyncMigrationJobWithCallback invokes the apds.StopSyncMigrationJob API asynchronously
func (client *Client) StopSyncMigrationJobWithCallback(request *StopSyncMigrationJobRequest, callback func(response *StopSyncMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopSyncMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.StopSyncMigrationJob(request)
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

// StopSyncMigrationJobRequest is the request struct for api StopSyncMigrationJob
type StopSyncMigrationJobRequest struct {
	*requests.RoaRequest
	JobType string `position:"Query" name:"jobType"`
}

// StopSyncMigrationJobResponse is the response struct for api StopSyncMigrationJob
type StopSyncMigrationJobResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateStopSyncMigrationJobRequest creates a request to invoke StopSyncMigrationJob API
func CreateStopSyncMigrationJobRequest() (request *StopSyncMigrationJobRequest) {
	request = &StopSyncMigrationJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "StopSyncMigrationJob", "/okss-services/migration-job/unsync-migration-job", "", "")
	request.Method = requests.POST
	return
}

// CreateStopSyncMigrationJobResponse creates a response to parse from StopSyncMigrationJob response
func CreateStopSyncMigrationJobResponse() (response *StopSyncMigrationJobResponse) {
	response = &StopSyncMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}