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

// WebpageAsyncScan invokes the green.WebpageAsyncScan API synchronously
func (client *Client) WebpageAsyncScan(request *WebpageAsyncScanRequest) (response *WebpageAsyncScanResponse, err error) {
	response = CreateWebpageAsyncScanResponse()
	err = client.DoAction(request, response)
	return
}

// WebpageAsyncScanWithChan invokes the green.WebpageAsyncScan API asynchronously
func (client *Client) WebpageAsyncScanWithChan(request *WebpageAsyncScanRequest) (<-chan *WebpageAsyncScanResponse, <-chan error) {
	responseChan := make(chan *WebpageAsyncScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WebpageAsyncScan(request)
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

// WebpageAsyncScanWithCallback invokes the green.WebpageAsyncScan API asynchronously
func (client *Client) WebpageAsyncScanWithCallback(request *WebpageAsyncScanRequest, callback func(response *WebpageAsyncScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WebpageAsyncScanResponse
		var err error
		defer close(result)
		response, err = client.WebpageAsyncScan(request)
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

// WebpageAsyncScanRequest is the request struct for api WebpageAsyncScan
type WebpageAsyncScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// WebpageAsyncScanResponse is the response struct for api WebpageAsyncScan
type WebpageAsyncScanResponse struct {
	*responses.BaseResponse
}

// CreateWebpageAsyncScanRequest creates a request to invoke WebpageAsyncScan API
func CreateWebpageAsyncScanRequest() (request *WebpageAsyncScanRequest) {
	request = &WebpageAsyncScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "WebpageAsyncScan", "/green/webpage/asyncscan", "", "")
	request.Method = requests.POST
	return
}

// CreateWebpageAsyncScanResponse creates a response to parse from WebpageAsyncScan response
func CreateWebpageAsyncScanResponse() (response *WebpageAsyncScanResponse) {
	response = &WebpageAsyncScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
