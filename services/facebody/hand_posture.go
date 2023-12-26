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

// HandPosture invokes the facebody.HandPosture API synchronously
func (client *Client) HandPosture(request *HandPostureRequest) (response *HandPostureResponse, err error) {
	response = CreateHandPostureResponse()
	err = client.DoAction(request, response)
	return
}

// HandPostureWithChan invokes the facebody.HandPosture API asynchronously
func (client *Client) HandPostureWithChan(request *HandPostureRequest) (<-chan *HandPostureResponse, <-chan error) {
	responseChan := make(chan *HandPostureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HandPosture(request)
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

// HandPostureWithCallback invokes the facebody.HandPosture API asynchronously
func (client *Client) HandPostureWithCallback(request *HandPostureRequest, callback func(response *HandPostureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HandPostureResponse
		var err error
		defer close(result)
		response, err = client.HandPosture(request)
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

// HandPostureRequest is the request struct for api HandPosture
type HandPostureRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	ImageURL           string           `position:"Body" name:"ImageURL"`
}

// HandPostureResponse is the response struct for api HandPosture
type HandPostureResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateHandPostureRequest creates a request to invoke HandPosture API
func CreateHandPostureRequest() (request *HandPostureRequest) {
	request = &HandPostureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "HandPosture", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateHandPostureResponse creates a response to parse from HandPosture response
func CreateHandPostureResponse() (response *HandPostureResponse) {
	response = &HandPostureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
