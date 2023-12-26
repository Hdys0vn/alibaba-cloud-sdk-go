package videoenhan

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

// EnhancePortraitVideo invokes the videoenhan.EnhancePortraitVideo API synchronously
func (client *Client) EnhancePortraitVideo(request *EnhancePortraitVideoRequest) (response *EnhancePortraitVideoResponse, err error) {
	response = CreateEnhancePortraitVideoResponse()
	err = client.DoAction(request, response)
	return
}

// EnhancePortraitVideoWithChan invokes the videoenhan.EnhancePortraitVideo API asynchronously
func (client *Client) EnhancePortraitVideoWithChan(request *EnhancePortraitVideoRequest) (<-chan *EnhancePortraitVideoResponse, <-chan error) {
	responseChan := make(chan *EnhancePortraitVideoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnhancePortraitVideo(request)
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

// EnhancePortraitVideoWithCallback invokes the videoenhan.EnhancePortraitVideo API asynchronously
func (client *Client) EnhancePortraitVideoWithCallback(request *EnhancePortraitVideoRequest, callback func(response *EnhancePortraitVideoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnhancePortraitVideoResponse
		var err error
		defer close(result)
		response, err = client.EnhancePortraitVideo(request)
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

// EnhancePortraitVideoRequest is the request struct for api EnhancePortraitVideo
type EnhancePortraitVideoRequest struct {
	*requests.RpcRequest
	Async    requests.Boolean `position:"Body" name:"Async"`
	VideoUrl string           `position:"Body" name:"VideoUrl"`
}

// EnhancePortraitVideoResponse is the response struct for api EnhancePortraitVideo
type EnhancePortraitVideoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnhancePortraitVideoRequest creates a request to invoke EnhancePortraitVideo API
func CreateEnhancePortraitVideoRequest() (request *EnhancePortraitVideoRequest) {
	request = &EnhancePortraitVideoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "EnhancePortraitVideo", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnhancePortraitVideoResponse creates a response to parse from EnhancePortraitVideo response
func CreateEnhancePortraitVideoResponse() (response *EnhancePortraitVideoResponse) {
	response = &EnhancePortraitVideoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}