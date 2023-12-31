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

// VideoAsyncManualScanResults invokes the green.VideoAsyncManualScanResults API synchronously
func (client *Client) VideoAsyncManualScanResults(request *VideoAsyncManualScanResultsRequest) (response *VideoAsyncManualScanResultsResponse, err error) {
	response = CreateVideoAsyncManualScanResultsResponse()
	err = client.DoAction(request, response)
	return
}

// VideoAsyncManualScanResultsWithChan invokes the green.VideoAsyncManualScanResults API asynchronously
func (client *Client) VideoAsyncManualScanResultsWithChan(request *VideoAsyncManualScanResultsRequest) (<-chan *VideoAsyncManualScanResultsResponse, <-chan error) {
	responseChan := make(chan *VideoAsyncManualScanResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VideoAsyncManualScanResults(request)
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

// VideoAsyncManualScanResultsWithCallback invokes the green.VideoAsyncManualScanResults API asynchronously
func (client *Client) VideoAsyncManualScanResultsWithCallback(request *VideoAsyncManualScanResultsRequest, callback func(response *VideoAsyncManualScanResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VideoAsyncManualScanResultsResponse
		var err error
		defer close(result)
		response, err = client.VideoAsyncManualScanResults(request)
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

// VideoAsyncManualScanResultsRequest is the request struct for api VideoAsyncManualScanResults
type VideoAsyncManualScanResultsRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VideoAsyncManualScanResultsResponse is the response struct for api VideoAsyncManualScanResults
type VideoAsyncManualScanResultsResponse struct {
	*responses.BaseResponse
}

// CreateVideoAsyncManualScanResultsRequest creates a request to invoke VideoAsyncManualScanResults API
func CreateVideoAsyncManualScanResultsRequest() (request *VideoAsyncManualScanResultsRequest) {
	request = &VideoAsyncManualScanResultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VideoAsyncManualScanResults", "/green/video/manual/scan/results", "", "")
	request.Method = requests.POST
	return
}

// CreateVideoAsyncManualScanResultsResponse creates a response to parse from VideoAsyncManualScanResults response
func CreateVideoAsyncManualScanResultsResponse() (response *VideoAsyncManualScanResultsResponse) {
	response = &VideoAsyncManualScanResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
