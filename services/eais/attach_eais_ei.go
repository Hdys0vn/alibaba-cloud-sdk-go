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

// AttachEaisEi invokes the eais.AttachEaisEi API synchronously
func (client *Client) AttachEaisEi(request *AttachEaisEiRequest) (response *AttachEaisEiResponse, err error) {
	response = CreateAttachEaisEiResponse()
	err = client.DoAction(request, response)
	return
}

// AttachEaisEiWithChan invokes the eais.AttachEaisEi API asynchronously
func (client *Client) AttachEaisEiWithChan(request *AttachEaisEiRequest) (<-chan *AttachEaisEiResponse, <-chan error) {
	responseChan := make(chan *AttachEaisEiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachEaisEi(request)
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

// AttachEaisEiWithCallback invokes the eais.AttachEaisEi API asynchronously
func (client *Client) AttachEaisEiWithCallback(request *AttachEaisEiRequest, callback func(response *AttachEaisEiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachEaisEiResponse
		var err error
		defer close(result)
		response, err = client.AttachEaisEi(request)
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

// AttachEaisEiRequest is the request struct for api AttachEaisEi
type AttachEaisEiRequest struct {
	*requests.RpcRequest
	EiInstanceType   string `position:"Query" name:"EiInstanceType"`
	ClientInstanceId string `position:"Query" name:"ClientInstanceId"`
	EiInstanceId     string `position:"Query" name:"EiInstanceId"`
}

// AttachEaisEiResponse is the response struct for api AttachEaisEi
type AttachEaisEiResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	EiInstanceId     string `json:"EiInstanceId" xml:"EiInstanceId"`
	ClientInstanceId string `json:"ClientInstanceId" xml:"ClientInstanceId"`
}

// CreateAttachEaisEiRequest creates a request to invoke AttachEaisEi API
func CreateAttachEaisEiRequest() (request *AttachEaisEiRequest) {
	request = &AttachEaisEiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "AttachEaisEi", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachEaisEiResponse creates a response to parse from AttachEaisEi response
func CreateAttachEaisEiResponse() (response *AttachEaisEiResponse) {
	response = &AttachEaisEiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
