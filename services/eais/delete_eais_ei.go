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

// DeleteEaisEi invokes the eais.DeleteEaisEi API synchronously
func (client *Client) DeleteEaisEi(request *DeleteEaisEiRequest) (response *DeleteEaisEiResponse, err error) {
	response = CreateDeleteEaisEiResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEaisEiWithChan invokes the eais.DeleteEaisEi API asynchronously
func (client *Client) DeleteEaisEiWithChan(request *DeleteEaisEiRequest) (<-chan *DeleteEaisEiResponse, <-chan error) {
	responseChan := make(chan *DeleteEaisEiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEaisEi(request)
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

// DeleteEaisEiWithCallback invokes the eais.DeleteEaisEi API asynchronously
func (client *Client) DeleteEaisEiWithCallback(request *DeleteEaisEiRequest, callback func(response *DeleteEaisEiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEaisEiResponse
		var err error
		defer close(result)
		response, err = client.DeleteEaisEi(request)
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

// DeleteEaisEiRequest is the request struct for api DeleteEaisEi
type DeleteEaisEiRequest struct {
	*requests.RpcRequest
	EiInstanceId string           `position:"Query" name:"EiInstanceId"`
	Force        requests.Boolean `position:"Query" name:"Force"`
}

// DeleteEaisEiResponse is the response struct for api DeleteEaisEi
type DeleteEaisEiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEaisEiRequest creates a request to invoke DeleteEaisEi API
func CreateDeleteEaisEiRequest() (request *DeleteEaisEiRequest) {
	request = &DeleteEaisEiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eais", "2019-06-24", "DeleteEaisEi", "eais", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEaisEiResponse creates a response to parse from DeleteEaisEi response
func CreateDeleteEaisEiResponse() (response *DeleteEaisEiResponse) {
	response = &DeleteEaisEiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}