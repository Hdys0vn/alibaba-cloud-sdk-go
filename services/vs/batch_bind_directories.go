package vs

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

// BatchBindDirectories invokes the vs.BatchBindDirectories API synchronously
func (client *Client) BatchBindDirectories(request *BatchBindDirectoriesRequest) (response *BatchBindDirectoriesResponse, err error) {
	response = CreateBatchBindDirectoriesResponse()
	err = client.DoAction(request, response)
	return
}

// BatchBindDirectoriesWithChan invokes the vs.BatchBindDirectories API asynchronously
func (client *Client) BatchBindDirectoriesWithChan(request *BatchBindDirectoriesRequest) (<-chan *BatchBindDirectoriesResponse, <-chan error) {
	responseChan := make(chan *BatchBindDirectoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchBindDirectories(request)
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

// BatchBindDirectoriesWithCallback invokes the vs.BatchBindDirectories API asynchronously
func (client *Client) BatchBindDirectoriesWithCallback(request *BatchBindDirectoriesRequest, callback func(response *BatchBindDirectoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchBindDirectoriesResponse
		var err error
		defer close(result)
		response, err = client.BatchBindDirectories(request)
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

// BatchBindDirectoriesRequest is the request struct for api BatchBindDirectories
type BatchBindDirectoriesRequest struct {
	*requests.RpcRequest
	DirectoryId string           `position:"Query" name:"DirectoryId"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId    string           `position:"Query" name:"DeviceId"`
}

// BatchBindDirectoriesResponse is the response struct for api BatchBindDirectories
type BatchBindDirectoriesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateBatchBindDirectoriesRequest creates a request to invoke BatchBindDirectories API
func CreateBatchBindDirectoriesRequest() (request *BatchBindDirectoriesRequest) {
	request = &BatchBindDirectoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchBindDirectories", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchBindDirectoriesResponse creates a response to parse from BatchBindDirectories response
func CreateBatchBindDirectoriesResponse() (response *BatchBindDirectoriesResponse) {
	response = &BatchBindDirectoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
