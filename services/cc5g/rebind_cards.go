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

// RebindCards invokes the cc5g.RebindCards API synchronously
func (client *Client) RebindCards(request *RebindCardsRequest) (response *RebindCardsResponse, err error) {
	response = CreateRebindCardsResponse()
	err = client.DoAction(request, response)
	return
}

// RebindCardsWithChan invokes the cc5g.RebindCards API asynchronously
func (client *Client) RebindCardsWithChan(request *RebindCardsRequest) (<-chan *RebindCardsResponse, <-chan error) {
	responseChan := make(chan *RebindCardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RebindCards(request)
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

// RebindCardsWithCallback invokes the cc5g.RebindCards API asynchronously
func (client *Client) RebindCardsWithCallback(request *RebindCardsRequest, callback func(response *RebindCardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RebindCardsResponse
		var err error
		defer close(result)
		response, err = client.RebindCards(request)
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

// RebindCardsRequest is the request struct for api RebindCards
type RebindCardsRequest struct {
	*requests.RpcRequest
	Iccids      *[]string        `position:"Query" name:"Iccids"  type:"Repeated"`
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken string           `position:"Query" name:"ClientToken"`
}

// RebindCardsResponse is the response struct for api RebindCards
type RebindCardsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRebindCardsRequest creates a request to invoke RebindCards API
func CreateRebindCardsRequest() (request *RebindCardsRequest) {
	request = &RebindCardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "RebindCards", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRebindCardsResponse creates a response to parse from RebindCards response
func CreateRebindCardsResponse() (response *RebindCardsResponse) {
	response = &RebindCardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
