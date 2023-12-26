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

// ImageScanFeedback invokes the green.ImageScanFeedback API synchronously
func (client *Client) ImageScanFeedback(request *ImageScanFeedbackRequest) (response *ImageScanFeedbackResponse, err error) {
	response = CreateImageScanFeedbackResponse()
	err = client.DoAction(request, response)
	return
}

// ImageScanFeedbackWithChan invokes the green.ImageScanFeedback API asynchronously
func (client *Client) ImageScanFeedbackWithChan(request *ImageScanFeedbackRequest) (<-chan *ImageScanFeedbackResponse, <-chan error) {
	responseChan := make(chan *ImageScanFeedbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImageScanFeedback(request)
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

// ImageScanFeedbackWithCallback invokes the green.ImageScanFeedback API asynchronously
func (client *Client) ImageScanFeedbackWithCallback(request *ImageScanFeedbackRequest, callback func(response *ImageScanFeedbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImageScanFeedbackResponse
		var err error
		defer close(result)
		response, err = client.ImageScanFeedback(request)
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

// ImageScanFeedbackRequest is the request struct for api ImageScanFeedback
type ImageScanFeedbackRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// ImageScanFeedbackResponse is the response struct for api ImageScanFeedback
type ImageScanFeedbackResponse struct {
	*responses.BaseResponse
}

// CreateImageScanFeedbackRequest creates a request to invoke ImageScanFeedback API
func CreateImageScanFeedbackRequest() (request *ImageScanFeedbackRequest) {
	request = &ImageScanFeedbackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "ImageScanFeedback", "/green/image/feedback", "", "")
	request.Method = requests.POST
	return
}

// CreateImageScanFeedbackResponse creates a response to parse from ImageScanFeedback response
func CreateImageScanFeedbackResponse() (response *ImageScanFeedbackResponse) {
	response = &ImageScanFeedbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
