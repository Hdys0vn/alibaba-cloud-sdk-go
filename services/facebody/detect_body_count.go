package facebody

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

// DetectBodyCount invokes the facebody.DetectBodyCount API synchronously
func (client *Client) DetectBodyCount(request *DetectBodyCountRequest) (response *DetectBodyCountResponse, err error) {
	response = CreateDetectBodyCountResponse()
	err = client.DoAction(request, response)
	return
}

// DetectBodyCountWithChan invokes the facebody.DetectBodyCount API asynchronously
func (client *Client) DetectBodyCountWithChan(request *DetectBodyCountRequest) (<-chan *DetectBodyCountResponse, <-chan error) {
	responseChan := make(chan *DetectBodyCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectBodyCount(request)
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

// DetectBodyCountWithCallback invokes the facebody.DetectBodyCount API asynchronously
func (client *Client) DetectBodyCountWithCallback(request *DetectBodyCountRequest, callback func(response *DetectBodyCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectBodyCountResponse
		var err error
		defer close(result)
		response, err = client.DetectBodyCount(request)
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

// DetectBodyCountRequest is the request struct for api DetectBodyCount
type DetectBodyCountRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	ImageURL           string           `position:"Body" name:"ImageURL"`
}

// DetectBodyCountResponse is the response struct for api DetectBodyCount
type DetectBodyCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectBodyCountRequest creates a request to invoke DetectBodyCount API
func CreateDetectBodyCountRequest() (request *DetectBodyCountRequest) {
	request = &DetectBodyCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DetectBodyCount", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectBodyCountResponse creates a response to parse from DetectBodyCount response
func CreateDetectBodyCountResponse() (response *DetectBodyCountResponse) {
	response = &DetectBodyCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
