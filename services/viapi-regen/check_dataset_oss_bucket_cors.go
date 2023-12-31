package viapi_regen

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

// CheckDatasetOssBucketCORS invokes the viapi_regen.CheckDatasetOssBucketCORS API synchronously
func (client *Client) CheckDatasetOssBucketCORS(request *CheckDatasetOssBucketCORSRequest) (response *CheckDatasetOssBucketCORSResponse, err error) {
	response = CreateCheckDatasetOssBucketCORSResponse()
	err = client.DoAction(request, response)
	return
}

// CheckDatasetOssBucketCORSWithChan invokes the viapi_regen.CheckDatasetOssBucketCORS API asynchronously
func (client *Client) CheckDatasetOssBucketCORSWithChan(request *CheckDatasetOssBucketCORSRequest) (<-chan *CheckDatasetOssBucketCORSResponse, <-chan error) {
	responseChan := make(chan *CheckDatasetOssBucketCORSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckDatasetOssBucketCORS(request)
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

// CheckDatasetOssBucketCORSWithCallback invokes the viapi_regen.CheckDatasetOssBucketCORS API asynchronously
func (client *Client) CheckDatasetOssBucketCORSWithCallback(request *CheckDatasetOssBucketCORSRequest, callback func(response *CheckDatasetOssBucketCORSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckDatasetOssBucketCORSResponse
		var err error
		defer close(result)
		response, err = client.CheckDatasetOssBucketCORS(request)
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

// CheckDatasetOssBucketCORSRequest is the request struct for api CheckDatasetOssBucketCORS
type CheckDatasetOssBucketCORSRequest struct {
	*requests.RpcRequest
	LabelsetId requests.Integer `position:"Body" name:"LabelsetId"`
}

// CheckDatasetOssBucketCORSResponse is the response struct for api CheckDatasetOssBucketCORS
type CheckDatasetOssBucketCORSResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCheckDatasetOssBucketCORSRequest creates a request to invoke CheckDatasetOssBucketCORS API
func CreateCheckDatasetOssBucketCORSRequest() (request *CheckDatasetOssBucketCORSRequest) {
	request = &CheckDatasetOssBucketCORSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "CheckDatasetOssBucketCORS", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckDatasetOssBucketCORSResponse creates a response to parse from CheckDatasetOssBucketCORS response
func CreateCheckDatasetOssBucketCORSResponse() (response *CheckDatasetOssBucketCORSResponse) {
	response = &CheckDatasetOssBucketCORSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
