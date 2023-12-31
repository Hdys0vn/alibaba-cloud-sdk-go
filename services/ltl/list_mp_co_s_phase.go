package ltl

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

// ListMPCoSPhase invokes the ltl.ListMPCoSPhase API synchronously
func (client *Client) ListMPCoSPhase(request *ListMPCoSPhaseRequest) (response *ListMPCoSPhaseResponse, err error) {
	response = CreateListMPCoSPhaseResponse()
	err = client.DoAction(request, response)
	return
}

// ListMPCoSPhaseWithChan invokes the ltl.ListMPCoSPhase API asynchronously
func (client *Client) ListMPCoSPhaseWithChan(request *ListMPCoSPhaseRequest) (<-chan *ListMPCoSPhaseResponse, <-chan error) {
	responseChan := make(chan *ListMPCoSPhaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMPCoSPhase(request)
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

// ListMPCoSPhaseWithCallback invokes the ltl.ListMPCoSPhase API asynchronously
func (client *Client) ListMPCoSPhaseWithCallback(request *ListMPCoSPhaseRequest, callback func(response *ListMPCoSPhaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMPCoSPhaseResponse
		var err error
		defer close(result)
		response, err = client.ListMPCoSPhase(request)
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

// ListMPCoSPhaseRequest is the request struct for api ListMPCoSPhase
type ListMPCoSPhaseRequest struct {
	*requests.RpcRequest
	Size         requests.Integer `position:"Query" name:"Size"`
	Num          requests.Integer `position:"Query" name:"Num"`
	Name         string           `position:"Query" name:"Name"`
	PhaseGroupId string           `position:"Query" name:"PhaseGroupId"`
	ApiVersion   string           `position:"Query" name:"ApiVersion"`
	BizChainId   string           `position:"Query" name:"BizChainId"`
}

// ListMPCoSPhaseResponse is the response struct for api ListMPCoSPhase
type ListMPCoSPhaseResponse struct {
	*responses.BaseResponse
	Code      int                  `json:"Code" xml:"Code"`
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Success   bool                 `json:"Success" xml:"Success"`
	Data      DataInListMPCoSPhase `json:"Data" xml:"Data"`
}

// CreateListMPCoSPhaseRequest creates a request to invoke ListMPCoSPhase API
func CreateListMPCoSPhaseRequest() (request *ListMPCoSPhaseRequest) {
	request = &ListMPCoSPhaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "ListMPCoSPhase", "", "")
	request.Method = requests.POST
	return
}

// CreateListMPCoSPhaseResponse creates a response to parse from ListMPCoSPhase response
func CreateListMPCoSPhaseResponse() (response *ListMPCoSPhaseResponse) {
	response = &ListMPCoSPhaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
