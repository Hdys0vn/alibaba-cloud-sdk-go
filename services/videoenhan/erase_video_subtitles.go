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

// EraseVideoSubtitles invokes the videoenhan.EraseVideoSubtitles API synchronously
func (client *Client) EraseVideoSubtitles(request *EraseVideoSubtitlesRequest) (response *EraseVideoSubtitlesResponse, err error) {
	response = CreateEraseVideoSubtitlesResponse()
	err = client.DoAction(request, response)
	return
}

// EraseVideoSubtitlesWithChan invokes the videoenhan.EraseVideoSubtitles API asynchronously
func (client *Client) EraseVideoSubtitlesWithChan(request *EraseVideoSubtitlesRequest) (<-chan *EraseVideoSubtitlesResponse, <-chan error) {
	responseChan := make(chan *EraseVideoSubtitlesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EraseVideoSubtitles(request)
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

// EraseVideoSubtitlesWithCallback invokes the videoenhan.EraseVideoSubtitles API asynchronously
func (client *Client) EraseVideoSubtitlesWithCallback(request *EraseVideoSubtitlesRequest, callback func(response *EraseVideoSubtitlesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EraseVideoSubtitlesResponse
		var err error
		defer close(result)
		response, err = client.EraseVideoSubtitles(request)
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

// EraseVideoSubtitlesRequest is the request struct for api EraseVideoSubtitles
type EraseVideoSubtitlesRequest struct {
	*requests.RpcRequest
	BH       requests.Float   `position:"Body" name:"BH"`
	BW       requests.Float   `position:"Body" name:"BW"`
	BX       requests.Float   `position:"Body" name:"BX"`
	BY       requests.Float   `position:"Body" name:"BY"`
	Async    requests.Boolean `position:"Body" name:"Async"`
	VideoUrl string           `position:"Body" name:"VideoUrl"`
}

// EraseVideoSubtitlesResponse is the response struct for api EraseVideoSubtitles
type EraseVideoSubtitlesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEraseVideoSubtitlesRequest creates a request to invoke EraseVideoSubtitles API
func CreateEraseVideoSubtitlesRequest() (request *EraseVideoSubtitlesRequest) {
	request = &EraseVideoSubtitlesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videoenhan", "2020-03-20", "EraseVideoSubtitles", "videoenhan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEraseVideoSubtitlesResponse creates a response to parse from EraseVideoSubtitles response
func CreateEraseVideoSubtitlesResponse() (response *EraseVideoSubtitlesResponse) {
	response = &EraseVideoSubtitlesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
