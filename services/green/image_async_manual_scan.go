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

// ImageAsyncManualScan invokes the green.ImageAsyncManualScan API synchronously
func (client *Client) ImageAsyncManualScan(request *ImageAsyncManualScanRequest) (response *ImageAsyncManualScanResponse, err error) {
	response = CreateImageAsyncManualScanResponse()
	err = client.DoAction(request, response)
	return
}

// ImageAsyncManualScanWithChan invokes the green.ImageAsyncManualScan API asynchronously
func (client *Client) ImageAsyncManualScanWithChan(request *ImageAsyncManualScanRequest) (<-chan *ImageAsyncManualScanResponse, <-chan error) {
	responseChan := make(chan *ImageAsyncManualScanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImageAsyncManualScan(request)
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

// ImageAsyncManualScanWithCallback invokes the green.ImageAsyncManualScan API asynchronously
func (client *Client) ImageAsyncManualScanWithCallback(request *ImageAsyncManualScanRequest, callback func(response *ImageAsyncManualScanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImageAsyncManualScanResponse
		var err error
		defer close(result)
		response, err = client.ImageAsyncManualScan(request)
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

// ImageAsyncManualScanRequest is the request struct for api ImageAsyncManualScan
type ImageAsyncManualScanRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// ImageAsyncManualScanResponse is the response struct for api ImageAsyncManualScan
type ImageAsyncManualScanResponse struct {
	*responses.BaseResponse
}

// CreateImageAsyncManualScanRequest creates a request to invoke ImageAsyncManualScan API
func CreateImageAsyncManualScanRequest() (request *ImageAsyncManualScanRequest) {
	request = &ImageAsyncManualScanRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "ImageAsyncManualScan", "/green/image/manual/asyncScan", "", "")
	request.Method = requests.POST
	return
}

// CreateImageAsyncManualScanResponse creates a response to parse from ImageAsyncManualScan response
func CreateImageAsyncManualScanResponse() (response *ImageAsyncManualScanResponse) {
	response = &ImageAsyncManualScanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
