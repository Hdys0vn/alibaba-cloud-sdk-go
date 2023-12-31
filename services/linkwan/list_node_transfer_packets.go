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

// ListNodeTransferPackets invokes the linkwan.ListNodeTransferPackets API synchronously
func (client *Client) ListNodeTransferPackets(request *ListNodeTransferPacketsRequest) (response *ListNodeTransferPacketsResponse, err error) {
	response = CreateListNodeTransferPacketsResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeTransferPacketsWithChan invokes the linkwan.ListNodeTransferPackets API asynchronously
func (client *Client) ListNodeTransferPacketsWithChan(request *ListNodeTransferPacketsRequest) (<-chan *ListNodeTransferPacketsResponse, <-chan error) {
	responseChan := make(chan *ListNodeTransferPacketsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeTransferPackets(request)
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

// ListNodeTransferPacketsWithCallback invokes the linkwan.ListNodeTransferPackets API asynchronously
func (client *Client) ListNodeTransferPacketsWithCallback(request *ListNodeTransferPacketsRequest, callback func(response *ListNodeTransferPacketsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeTransferPacketsResponse
		var err error
		defer close(result)
		response, err = client.ListNodeTransferPackets(request)
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

// ListNodeTransferPacketsRequest is the request struct for api ListNodeTransferPackets
type ListNodeTransferPacketsRequest struct {
	*requests.RpcRequest
	EndMillis    requests.Integer `position:"Query" name:"EndMillis"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	GwEui        string           `position:"Query" name:"GwEui"`
	Ascending    requests.Boolean `position:"Query" name:"Ascending"`
	DevEui       string           `position:"Query" name:"DevEui"`
	ApiProduct   string           `position:"Body" name:"ApiProduct"`
	ApiRevision  string           `position:"Body" name:"ApiRevision"`
	Category     string           `position:"Query" name:"Category"`
	BeginMillis  requests.Integer `position:"Query" name:"BeginMillis"`
	SortingField string           `position:"Query" name:"SortingField"`
}

// ListNodeTransferPacketsResponse is the response struct for api ListNodeTransferPackets
type ListNodeTransferPacketsResponse struct {
	*responses.BaseResponse
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Success   bool                          `json:"Success" xml:"Success"`
	Data      DataInListNodeTransferPackets `json:"Data" xml:"Data"`
}

// CreateListNodeTransferPacketsRequest creates a request to invoke ListNodeTransferPackets API
func CreateListNodeTransferPacketsRequest() (request *ListNodeTransferPacketsRequest) {
	request = &ListNodeTransferPacketsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodeTransferPackets", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodeTransferPacketsResponse creates a response to parse from ListNodeTransferPackets response
func CreateListNodeTransferPacketsResponse() (response *ListNodeTransferPacketsResponse) {
	response = &ListNodeTransferPacketsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
