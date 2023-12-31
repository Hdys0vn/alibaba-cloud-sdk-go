package imageprocess

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

// DetectSpineMRI invokes the imageprocess.DetectSpineMRI API synchronously
func (client *Client) DetectSpineMRI(request *DetectSpineMRIRequest) (response *DetectSpineMRIResponse, err error) {
	response = CreateDetectSpineMRIResponse()
	err = client.DoAction(request, response)
	return
}

// DetectSpineMRIWithChan invokes the imageprocess.DetectSpineMRI API asynchronously
func (client *Client) DetectSpineMRIWithChan(request *DetectSpineMRIRequest) (<-chan *DetectSpineMRIResponse, <-chan error) {
	responseChan := make(chan *DetectSpineMRIResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectSpineMRI(request)
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

// DetectSpineMRIWithCallback invokes the imageprocess.DetectSpineMRI API asynchronously
func (client *Client) DetectSpineMRIWithCallback(request *DetectSpineMRIRequest, callback func(response *DetectSpineMRIResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectSpineMRIResponse
		var err error
		defer close(result)
		response, err = client.DetectSpineMRI(request)
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

// DetectSpineMRIRequest is the request struct for api DetectSpineMRI
type DetectSpineMRIRequest struct {
	*requests.RpcRequest
	OrgName    string                   `position:"Body" name:"OrgName"`
	DataFormat string                   `position:"Body" name:"DataFormat"`
	URLList    *[]DetectSpineMRIURLList `position:"Body" name:"URLList"  type:"Repeated"`
	OrgId      string                   `position:"Body" name:"OrgId"`
}

// DetectSpineMRIURLList is a repeated param struct in DetectSpineMRIRequest
type DetectSpineMRIURLList struct {
	URL string `name:"URL"`
}

// DetectSpineMRIResponse is the response struct for api DetectSpineMRI
type DetectSpineMRIResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDetectSpineMRIRequest creates a request to invoke DetectSpineMRI API
func CreateDetectSpineMRIRequest() (request *DetectSpineMRIRequest) {
	request = &DetectSpineMRIRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "DetectSpineMRI", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectSpineMRIResponse creates a response to parse from DetectSpineMRI response
func CreateDetectSpineMRIResponse() (response *DetectSpineMRIResponse) {
	response = &DetectSpineMRIResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
