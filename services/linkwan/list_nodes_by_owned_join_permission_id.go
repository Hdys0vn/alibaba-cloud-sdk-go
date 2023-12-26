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

// ListNodesByOwnedJoinPermissionId invokes the linkwan.ListNodesByOwnedJoinPermissionId API synchronously
func (client *Client) ListNodesByOwnedJoinPermissionId(request *ListNodesByOwnedJoinPermissionIdRequest) (response *ListNodesByOwnedJoinPermissionIdResponse, err error) {
	response = CreateListNodesByOwnedJoinPermissionIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodesByOwnedJoinPermissionIdWithChan invokes the linkwan.ListNodesByOwnedJoinPermissionId API asynchronously
func (client *Client) ListNodesByOwnedJoinPermissionIdWithChan(request *ListNodesByOwnedJoinPermissionIdRequest) (<-chan *ListNodesByOwnedJoinPermissionIdResponse, <-chan error) {
	responseChan := make(chan *ListNodesByOwnedJoinPermissionIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodesByOwnedJoinPermissionId(request)
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

// ListNodesByOwnedJoinPermissionIdWithCallback invokes the linkwan.ListNodesByOwnedJoinPermissionId API asynchronously
func (client *Client) ListNodesByOwnedJoinPermissionIdWithCallback(request *ListNodesByOwnedJoinPermissionIdRequest, callback func(response *ListNodesByOwnedJoinPermissionIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodesByOwnedJoinPermissionIdResponse
		var err error
		defer close(result)
		response, err = client.ListNodesByOwnedJoinPermissionId(request)
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

// ListNodesByOwnedJoinPermissionIdRequest is the request struct for api ListNodesByOwnedJoinPermissionId
type ListNodesByOwnedJoinPermissionIdRequest struct {
	*requests.RpcRequest
	JoinPermissionId string           `position:"Query" name:"JoinPermissionId"`
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	FuzzyDevEui      string           `position:"Query" name:"FuzzyDevEui"`
	Limit            requests.Integer `position:"Query" name:"Limit"`
	Offset           requests.Integer `position:"Query" name:"Offset"`
	Ascending        requests.Boolean `position:"Query" name:"Ascending"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
	SortingField     string           `position:"Query" name:"SortingField"`
}

// ListNodesByOwnedJoinPermissionIdResponse is the response struct for api ListNodesByOwnedJoinPermissionId
type ListNodesByOwnedJoinPermissionIdResponse struct {
	*responses.BaseResponse
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	Success   bool                                   `json:"Success" xml:"Success"`
	Data      DataInListNodesByOwnedJoinPermissionId `json:"Data" xml:"Data"`
}

// CreateListNodesByOwnedJoinPermissionIdRequest creates a request to invoke ListNodesByOwnedJoinPermissionId API
func CreateListNodesByOwnedJoinPermissionIdRequest() (request *ListNodesByOwnedJoinPermissionIdRequest) {
	request = &ListNodesByOwnedJoinPermissionIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodesByOwnedJoinPermissionId", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodesByOwnedJoinPermissionIdResponse creates a response to parse from ListNodesByOwnedJoinPermissionId response
func CreateListNodesByOwnedJoinPermissionIdResponse() (response *ListNodesByOwnedJoinPermissionIdResponse) {
	response = &ListNodesByOwnedJoinPermissionIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}