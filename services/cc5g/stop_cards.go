package cc5g

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

// StopCards invokes the cc5g.StopCards API synchronously
func (client *Client) StopCards(request *StopCardsRequest) (response *StopCardsResponse, err error) {
	response = CreateStopCardsResponse()
	err = client.DoAction(request, response)
	return
}

// StopCardsWithChan invokes the cc5g.StopCards API asynchronously
func (client *Client) StopCardsWithChan(request *StopCardsRequest) (<-chan *StopCardsResponse, <-chan error) {
	responseChan := make(chan *StopCardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopCards(request)
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

// StopCardsWithCallback invokes the cc5g.StopCards API asynchronously
func (client *Client) StopCardsWithCallback(request *StopCardsRequest, callback func(response *StopCardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopCardsResponse
		var err error
		defer close(result)
		response, err = client.StopCards(request)
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

// StopCardsRequest is the request struct for api StopCards
type StopCardsRequest struct {
	*requests.RpcRequest
	Iccids      *[]string        `position:"Query" name:"Iccids"  type:"Repeated"`
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken string           `position:"Query" name:"ClientToken"`
}

// StopCardsResponse is the response struct for api StopCards
type StopCardsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopCardsRequest creates a request to invoke StopCards API
func CreateStopCardsRequest() (request *StopCardsRequest) {
	request = &StopCardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "StopCards", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopCardsResponse creates a response to parse from StopCards response
func CreateStopCardsResponse() (response *StopCardsResponse) {
	response = &StopCardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
