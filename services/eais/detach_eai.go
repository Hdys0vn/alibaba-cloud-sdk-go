package eais

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

// DetachEai invokes the eais.DetachEai API synchronously
func (client *Client) DetachEai(request *DetachEaiRequest) (response *DetachEaiResponse, err error) {
	response = CreateDetachEaiResponse()
	err = client.DoAction(request, response)
	return
}

// DetachEaiWithChan invokes the eais.DetachEai API asynchronously
func (client *Client) DetachEaiWithChan(request *DetachEaiRequest) (<-chan *DetachEaiResponse, <-chan error) {
	responseChan := make(chan *DetachEaiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachEai(request)
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

// DetachEaiWithCallback invokes the eais.DetachEai API asynchronously
func (client *Client) DetachEaiWithCallback(request *DetachEaiRequest, callback func(response *DetachEaiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachEaiResponse
		var err error
		defer close(result)
		response, err = client.DetachEai(request)
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

// DetachEaiRequest is the request struct for api DetachEai
type DetachEaiRequest struct {
	*requests.RpcRequest
	ElasticAcceleratedInstanceId string `position:"Query" name:"ElasticAcceleratedInstanceId"`
}

// DetachEaiResponse is the response struct for api DetachEai
type DetachEaiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachEaiRequest creates a request to invoke DetachEai API
func CreateDetachEaiRequest() (request *DetachEaiRequest) {
	request = &DetachEaiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "DetachEai", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachEaiResponse creates a response to parse from DetachEai response
func CreateDetachEaiResponse() (response *DetachEaiResponse) {
	response = &DetachEaiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}