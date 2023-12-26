package edas

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

// DeleteEcu invokes the edas.DeleteEcu API synchronously
func (client *Client) DeleteEcu(request *DeleteEcuRequest) (response *DeleteEcuResponse, err error) {
	response = CreateDeleteEcuResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEcuWithChan invokes the edas.DeleteEcu API asynchronously
func (client *Client) DeleteEcuWithChan(request *DeleteEcuRequest) (<-chan *DeleteEcuResponse, <-chan error) {
	responseChan := make(chan *DeleteEcuResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEcu(request)
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

// DeleteEcuWithCallback invokes the edas.DeleteEcu API asynchronously
func (client *Client) DeleteEcuWithCallback(request *DeleteEcuRequest, callback func(response *DeleteEcuResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEcuResponse
		var err error
		defer close(result)
		response, err = client.DeleteEcu(request)
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

// DeleteEcuRequest is the request struct for api DeleteEcu
type DeleteEcuRequest struct {
	*requests.RoaRequest
	EcuId string `position:"Query" name:"EcuId"`
}

// DeleteEcuResponse is the response struct for api DeleteEcu
type DeleteEcuResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEcuRequest creates a request to invoke DeleteEcu API
func CreateDeleteEcuRequest() (request *DeleteEcuRequest) {
	request = &DeleteEcuRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeleteEcu", "/pop/v5/resource/delete_ecu", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEcuResponse creates a response to parse from DeleteEcu response
func CreateDeleteEcuResponse() (response *DeleteEcuResponse) {
	response = &DeleteEcuResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
