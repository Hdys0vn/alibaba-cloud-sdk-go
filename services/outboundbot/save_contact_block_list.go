package outboundbot

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

// SaveContactBlockList invokes the outboundbot.SaveContactBlockList API synchronously
func (client *Client) SaveContactBlockList(request *SaveContactBlockListRequest) (response *SaveContactBlockListResponse, err error) {
	response = CreateSaveContactBlockListResponse()
	err = client.DoAction(request, response)
	return
}

// SaveContactBlockListWithChan invokes the outboundbot.SaveContactBlockList API asynchronously
func (client *Client) SaveContactBlockListWithChan(request *SaveContactBlockListRequest) (<-chan *SaveContactBlockListResponse, <-chan error) {
	responseChan := make(chan *SaveContactBlockListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveContactBlockList(request)
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

// SaveContactBlockListWithCallback invokes the outboundbot.SaveContactBlockList API asynchronously
func (client *Client) SaveContactBlockListWithCallback(request *SaveContactBlockListRequest, callback func(response *SaveContactBlockListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveContactBlockListResponse
		var err error
		defer close(result)
		response, err = client.SaveContactBlockList(request)
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

// SaveContactBlockListRequest is the request struct for api SaveContactBlockList
type SaveContactBlockListRequest struct {
	*requests.RpcRequest
	ContactBlockListList  *[]string `position:"Query" name:"ContactBlockListList"  type:"Repeated"`
	InstanceId            string    `position:"Query" name:"InstanceId"`
	ContactBlockListsJson string    `position:"Query" name:"ContactBlockListsJson"`
}

// SaveContactBlockListResponse is the response struct for api SaveContactBlockList
type SaveContactBlockListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	AffectedRows   int    `json:"AffectedRows" xml:"AffectedRows"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateSaveContactBlockListRequest creates a request to invoke SaveContactBlockList API
func CreateSaveContactBlockListRequest() (request *SaveContactBlockListRequest) {
	request = &SaveContactBlockListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SaveContactBlockList", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveContactBlockListResponse creates a response to parse from SaveContactBlockList response
func CreateSaveContactBlockListResponse() (response *SaveContactBlockListResponse) {
	response = &SaveContactBlockListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
