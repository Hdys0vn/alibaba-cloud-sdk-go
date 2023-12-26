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

// SplitVideoParts invokes the videorecog.SplitVideoParts API synchronously
func (client *Client) SplitVideoParts(request *SplitVideoPartsRequest) (response *SplitVideoPartsResponse, err error) {
	response = CreateSplitVideoPartsResponse()
	err = client.DoAction(request, response)
	return
}

// SplitVideoPartsWithChan invokes the videorecog.SplitVideoParts API asynchronously
func (client *Client) SplitVideoPartsWithChan(request *SplitVideoPartsRequest) (<-chan *SplitVideoPartsResponse, <-chan error) {
	responseChan := make(chan *SplitVideoPartsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SplitVideoParts(request)
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

// SplitVideoPartsWithCallback invokes the videorecog.SplitVideoParts API asynchronously
func (client *Client) SplitVideoPartsWithCallback(request *SplitVideoPartsRequest, callback func(response *SplitVideoPartsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SplitVideoPartsResponse
		var err error
		defer close(result)
		response, err = client.SplitVideoParts(request)
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

// SplitVideoPartsRequest is the request struct for api SplitVideoParts
type SplitVideoPartsRequest struct {
	*requests.RpcRequest
	Template string           `position:"Body" name:"Template"`
	MinTime  requests.Integer `position:"Body" name:"MinTime"`
	MaxTime  requests.Integer `position:"Body" name:"MaxTime"`
	Async    requests.Boolean `position:"Body" name:"Async"`
	VideoUrl string           `position:"Body" name:"VideoUrl"`
}

// SplitVideoPartsResponse is the response struct for api SplitVideoParts
type SplitVideoPartsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSplitVideoPartsRequest creates a request to invoke SplitVideoParts API
func CreateSplitVideoPartsRequest() (request *SplitVideoPartsRequest) {
	request = &SplitVideoPartsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("videorecog", "2020-03-20", "SplitVideoParts", "videorecog", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSplitVideoPartsResponse creates a response to parse from SplitVideoParts response
func CreateSplitVideoPartsResponse() (response *SplitVideoPartsResponse) {
	response = &SplitVideoPartsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
