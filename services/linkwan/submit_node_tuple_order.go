package linkwan

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

// SubmitNodeTupleOrder invokes the linkwan.SubmitNodeTupleOrder API synchronously
func (client *Client) SubmitNodeTupleOrder(request *SubmitNodeTupleOrderRequest) (response *SubmitNodeTupleOrderResponse, err error) {
	response = CreateSubmitNodeTupleOrderResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitNodeTupleOrderWithChan invokes the linkwan.SubmitNodeTupleOrder API asynchronously
func (client *Client) SubmitNodeTupleOrderWithChan(request *SubmitNodeTupleOrderRequest) (<-chan *SubmitNodeTupleOrderResponse, <-chan error) {
	responseChan := make(chan *SubmitNodeTupleOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitNodeTupleOrder(request)
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

// SubmitNodeTupleOrderWithCallback invokes the linkwan.SubmitNodeTupleOrder API asynchronously
func (client *Client) SubmitNodeTupleOrderWithCallback(request *SubmitNodeTupleOrderRequest, callback func(response *SubmitNodeTupleOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitNodeTupleOrderResponse
		var err error
		defer close(result)
		response, err = client.SubmitNodeTupleOrder(request)
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

// SubmitNodeTupleOrderRequest is the request struct for api SubmitNodeTupleOrder
type SubmitNodeTupleOrderRequest struct {
	*requests.RpcRequest
	LoraVersion   string           `position:"Query" name:"LoraVersion"`
	TupleType     string           `position:"Query" name:"TupleType"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	RequiredCount requests.Integer `position:"Query" name:"RequiredCount"`
}

// SubmitNodeTupleOrderResponse is the response struct for api SubmitNodeTupleOrder
type SubmitNodeTupleOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSubmitNodeTupleOrderRequest creates a request to invoke SubmitNodeTupleOrder API
func CreateSubmitNodeTupleOrderRequest() (request *SubmitNodeTupleOrderRequest) {
	request = &SubmitNodeTupleOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "SubmitNodeTupleOrder", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitNodeTupleOrderResponse creates a response to parse from SubmitNodeTupleOrder response
func CreateSubmitNodeTupleOrderResponse() (response *SubmitNodeTupleOrderResponse) {
	response = &SubmitNodeTupleOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
