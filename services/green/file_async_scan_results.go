package green

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

// FileAsyncScanResults invokes the green.FileAsyncScanResults API synchronously
func (client *Client) FileAsyncScanResults(request *FileAsyncScanResultsRequest) (response *FileAsyncScanResultsResponse, err error) {
	response = CreateFileAsyncScanResultsResponse()
	err = client.DoAction(request, response)
	return
}

// FileAsyncScanResultsWithChan invokes the green.FileAsyncScanResults API asynchronously
func (client *Client) FileAsyncScanResultsWithChan(request *FileAsyncScanResultsRequest) (<-chan *FileAsyncScanResultsResponse, <-chan error) {
	responseChan := make(chan *FileAsyncScanResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FileAsyncScanResults(request)
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

// FileAsyncScanResultsWithCallback invokes the green.FileAsyncScanResults API asynchronously
func (client *Client) FileAsyncScanResultsWithCallback(request *FileAsyncScanResultsRequest, callback func(response *FileAsyncScanResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FileAsyncScanResultsResponse
		var err error
		defer close(result)
		response, err = client.FileAsyncScanResults(request)
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

// FileAsyncScanResultsRequest is the request struct for api FileAsyncScanResults
type FileAsyncScanResultsRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// FileAsyncScanResultsResponse is the response struct for api FileAsyncScanResults
type FileAsyncScanResultsResponse struct {
	*responses.BaseResponse
}

// CreateFileAsyncScanResultsRequest creates a request to invoke FileAsyncScanResults API
func CreateFileAsyncScanResultsRequest() (request *FileAsyncScanResultsRequest) {
	request = &FileAsyncScanResultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "FileAsyncScanResults", "/green/file/results", "", "")
	request.Method = requests.POST
	return
}

// CreateFileAsyncScanResultsResponse creates a response to parse from FileAsyncScanResults response
func CreateFileAsyncScanResultsResponse() (response *FileAsyncScanResultsResponse) {
	response = &FileAsyncScanResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
