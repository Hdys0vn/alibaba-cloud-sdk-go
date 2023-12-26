package videorecog

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

// DetectVideoShot invokes the videorecog.DetectVideoShot API synchronously
func (client *Client) DetectVideoShot(request *DetectVideoShotRequest) (response *DetectVideoShotResponse, err error) {
	response = CreateDetectVideoShotResponse()
	err = client.DoAction(request, response)
	return
}

// DetectVideoShotWithChan invokes the videorecog.DetectVideoShot API asynchronously
func (client *Client) DetectVideoShotWithChan(request *DetectVideoShotRequest) (<-chan *DetectVideoShotResponse, <-chan error) {
	responseChan := make(chan *DetectVideoShotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVideoShot(request)
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

// DetectVideoShotWithCallback invokes the videorecog.DetectVideoShot API asynchronously
func (client *Client) DetectVideoShotWithCallback(request *DetectVideoShotRequest, callback func(response *DetectVideoShotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVideoShotResponse
		var err error
		defer close(result)
		response, err = client.DetectVideoShot(request)
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

// DetectVideoShotRequest is the request struct for api DetectVideoShot
type DetectVideoShotRequest struct {
	*requests.RpcRequest
	Async    requests.Boolean `position:"Body" name:"Async"`
	VideoUrl string           `position:"Body" name:"VideoUrl"`
}

// DetectVideoShotResponse is the response struct for api DetectVideoShot
type DetectVideoShotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectVideoShotRequest creates a request to invoke DetectVideoShot API
func CreateDetectVideoShotRequest() (request *DetectVideoShotRequest) {
	request = &DetectVideoShotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videorecog", "2020-03-20", "DetectVideoShot", "videorecog", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectVideoShotResponse creates a response to parse from DetectVideoShot response
func CreateDetectVideoShotResponse() (response *DetectVideoShotResponse) {
	response = &DetectVideoShotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}