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

// EvaluateVideoQuality invokes the videorecog.EvaluateVideoQuality API synchronously
func (client *Client) EvaluateVideoQuality(request *EvaluateVideoQualityRequest) (response *EvaluateVideoQualityResponse, err error) {
	response = CreateEvaluateVideoQualityResponse()
	err = client.DoAction(request, response)
	return
}

// EvaluateVideoQualityWithChan invokes the videorecog.EvaluateVideoQuality API asynchronously
func (client *Client) EvaluateVideoQualityWithChan(request *EvaluateVideoQualityRequest) (<-chan *EvaluateVideoQualityResponse, <-chan error) {
	responseChan := make(chan *EvaluateVideoQualityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EvaluateVideoQuality(request)
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

// EvaluateVideoQualityWithCallback invokes the videorecog.EvaluateVideoQuality API asynchronously
func (client *Client) EvaluateVideoQualityWithCallback(request *EvaluateVideoQualityRequest, callback func(response *EvaluateVideoQualityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EvaluateVideoQualityResponse
		var err error
		defer close(result)
		response, err = client.EvaluateVideoQuality(request)
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

// EvaluateVideoQualityRequest is the request struct for api EvaluateVideoQuality
type EvaluateVideoQualityRequest struct {
	*requests.RpcRequest
	Mode     string           `position:"Body" name:"Mode"`
	Async    requests.Boolean `position:"Body" name:"Async"`
	VideoUrl string           `position:"Body" name:"VideoUrl"`
}

// EvaluateVideoQualityResponse is the response struct for api EvaluateVideoQuality
type EvaluateVideoQualityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEvaluateVideoQualityRequest creates a request to invoke EvaluateVideoQuality API
func CreateEvaluateVideoQualityRequest() (request *EvaluateVideoQualityRequest) {
	request = &EvaluateVideoQualityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videorecog", "2020-03-20", "EvaluateVideoQuality", "videorecog", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEvaluateVideoQualityResponse creates a response to parse from EvaluateVideoQuality response
func CreateEvaluateVideoQualityResponse() (response *EvaluateVideoQualityResponse) {
	response = &EvaluateVideoQualityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
