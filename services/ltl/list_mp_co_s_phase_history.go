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

// ListMPCoSPhaseHistory invokes the ltl.ListMPCoSPhaseHistory API synchronously
func (client *Client) ListMPCoSPhaseHistory(request *ListMPCoSPhaseHistoryRequest) (response *ListMPCoSPhaseHistoryResponse, err error) {
	response = CreateListMPCoSPhaseHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// ListMPCoSPhaseHistoryWithChan invokes the ltl.ListMPCoSPhaseHistory API asynchronously
func (client *Client) ListMPCoSPhaseHistoryWithChan(request *ListMPCoSPhaseHistoryRequest) (<-chan *ListMPCoSPhaseHistoryResponse, <-chan error) {
	responseChan := make(chan *ListMPCoSPhaseHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMPCoSPhaseHistory(request)
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

// ListMPCoSPhaseHistoryWithCallback invokes the ltl.ListMPCoSPhaseHistory API asynchronously
func (client *Client) ListMPCoSPhaseHistoryWithCallback(request *ListMPCoSPhaseHistoryRequest, callback func(response *ListMPCoSPhaseHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMPCoSPhaseHistoryResponse
		var err error
		defer close(result)
		response, err = client.ListMPCoSPhaseHistory(request)
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

// ListMPCoSPhaseHistoryRequest is the request struct for api ListMPCoSPhaseHistory
type ListMPCoSPhaseHistoryRequest struct {
	*requests.RpcRequest
	Num          requests.Integer `position:"Query" name:"Num"`
	PhaseId      string           `position:"Query" name:"PhaseId"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	ApiVersion   string           `position:"Query" name:"ApiVersion"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	BizChainId   string           `position:"Query" name:"BizChainId"`
	DataKey      string           `position:"Query" name:"DataKey"`
	Size         requests.Integer `position:"Query" name:"Size"`
	PhaseGroupId string           `position:"Query" name:"PhaseGroupId"`
}

// ListMPCoSPhaseHistoryResponse is the response struct for api ListMPCoSPhaseHistory
type ListMPCoSPhaseHistoryResponse struct {
	*responses.BaseResponse
	Code      int                         `json:"Code" xml:"Code"`
	Message   string                      `json:"Message" xml:"Message"`
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Success   bool                        `json:"Success" xml:"Success"`
	Data      DataInListMPCoSPhaseHistory `json:"Data" xml:"Data"`
}

// CreateListMPCoSPhaseHistoryRequest creates a request to invoke ListMPCoSPhaseHistory API
func CreateListMPCoSPhaseHistoryRequest() (request *ListMPCoSPhaseHistoryRequest) {
	request = &ListMPCoSPhaseHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "ListMPCoSPhaseHistory", "", "")
	request.Method = requests.POST
	return
}

// CreateListMPCoSPhaseHistoryResponse creates a response to parse from ListMPCoSPhaseHistory response
func CreateListMPCoSPhaseHistoryResponse() (response *ListMPCoSPhaseHistoryResponse) {
	response = &ListMPCoSPhaseHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
