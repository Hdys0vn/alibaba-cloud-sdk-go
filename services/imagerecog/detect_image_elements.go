package imagerecog

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

// DetectImageElements invokes the imagerecog.DetectImageElements API synchronously
func (client *Client) DetectImageElements(request *DetectImageElementsRequest) (response *DetectImageElementsResponse, err error) {
	response = CreateDetectImageElementsResponse()
	err = client.DoAction(request, response)
	return
}

// DetectImageElementsWithChan invokes the imagerecog.DetectImageElements API asynchronously
func (client *Client) DetectImageElementsWithChan(request *DetectImageElementsRequest) (<-chan *DetectImageElementsResponse, <-chan error) {
	responseChan := make(chan *DetectImageElementsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectImageElements(request)
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

// DetectImageElementsWithCallback invokes the imagerecog.DetectImageElements API asynchronously
func (client *Client) DetectImageElementsWithCallback(request *DetectImageElementsRequest, callback func(response *DetectImageElementsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectImageElementsResponse
		var err error
		defer close(result)
		response, err = client.DetectImageElements(request)
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

// DetectImageElementsRequest is the request struct for api DetectImageElements
type DetectImageElementsRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	Url                string           `position:"Body" name:"Url"`
}

// DetectImageElementsResponse is the response struct for api DetectImageElements
type DetectImageElementsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectImageElementsRequest creates a request to invoke DetectImageElements API
func CreateDetectImageElementsRequest() (request *DetectImageElementsRequest) {
	request = &DetectImageElementsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imagerecog", "2019-09-30", "DetectImageElements", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectImageElementsResponse creates a response to parse from DetectImageElements response
func CreateDetectImageElementsResponse() (response *DetectImageElementsResponse) {
	response = &DetectImageElementsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
