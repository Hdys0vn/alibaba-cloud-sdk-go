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

// ListCards invokes the cc5g.ListCards API synchronously
func (client *Client) ListCards(request *ListCardsRequest) (response *ListCardsResponse, err error) {
	response = CreateListCardsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCardsWithChan invokes the cc5g.ListCards API asynchronously
func (client *Client) ListCardsWithChan(request *ListCardsRequest) (<-chan *ListCardsResponse, <-chan error) {
	responseChan := make(chan *ListCardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCards(request)
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

// ListCardsWithCallback invokes the cc5g.ListCards API asynchronously
func (client *Client) ListCardsWithCallback(request *ListCardsRequest, callback func(response *ListCardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCardsResponse
		var err error
		defer close(result)
		response, err = client.ListCards(request)
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

// ListCardsRequest is the request struct for api ListCards
type ListCardsRequest struct {
	*requests.RpcRequest
	IpAddress                string           `position:"Query" name:"IpAddress"`
	Iccids                   *[]string        `position:"Query" name:"Iccids"  type:"Repeated"`
	Iccid                    string           `position:"Query" name:"Iccid"`
	NextToken                string           `position:"Query" name:"NextToken"`
	Lock                     requests.Boolean `position:"Query" name:"Lock"`
	Msisdn                   string           `position:"Query" name:"Msisdn"`
	Apn                      string           `position:"Query" name:"Apn"`
	NetLinkId                string           `position:"Query" name:"NetLinkId"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
	Online                   requests.Boolean `position:"Query" name:"Online"`
	MaxResults               requests.Integer `position:"Query" name:"MaxResults"`
	Statuses                 *[]string        `position:"Query" name:"Statuses"  type:"Repeated"`
}

// ListCardsResponse is the response struct for api ListCards
type ListCardsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	MaxResults string `json:"MaxResults" xml:"MaxResults"`
	TotalCount string `json:"TotalCount" xml:"TotalCount"`
	Cards      []Card `json:"Cards" xml:"Cards"`
}

// CreateListCardsRequest creates a request to invoke ListCards API
func CreateListCardsRequest() (request *ListCardsRequest) {
	request = &ListCardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "ListCards", "fivegcc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListCardsResponse creates a response to parse from ListCards response
func CreateListCardsResponse() (response *ListCardsResponse) {
	response = &ListCardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
