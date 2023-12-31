package mts

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

// ListMedia invokes the mts.ListMedia API synchronously
func (client *Client) ListMedia(request *ListMediaRequest) (response *ListMediaResponse, err error) {
	response = CreateListMediaResponse()
	err = client.DoAction(request, response)
	return
}

// ListMediaWithChan invokes the mts.ListMedia API asynchronously
func (client *Client) ListMediaWithChan(request *ListMediaRequest) (<-chan *ListMediaResponse, <-chan error) {
	responseChan := make(chan *ListMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMedia(request)
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

// ListMediaWithCallback invokes the mts.ListMedia API asynchronously
func (client *Client) ListMediaWithCallback(request *ListMediaRequest, callback func(response *ListMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMediaResponse
		var err error
		defer close(result)
		response, err = client.ListMedia(request)
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

// ListMediaRequest is the request struct for api ListMedia
type ListMediaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	From                 string           `position:"Query" name:"From"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      requests.Integer `position:"Query" name:"MaximumPageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	To                   string           `position:"Query" name:"To"`
}

// ListMediaResponse is the response struct for api ListMedia
type ListMediaResponse struct {
	*responses.BaseResponse
	RequestId     string               `json:"RequestId" xml:"RequestId"`
	NextPageToken string               `json:"NextPageToken" xml:"NextPageToken"`
	MediaList     MediaListInListMedia `json:"MediaList" xml:"MediaList"`
}

// CreateListMediaRequest creates a request to invoke ListMedia API
func CreateListMediaRequest() (request *ListMediaRequest) {
	request = &ListMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListMedia", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListMediaResponse creates a response to parse from ListMedia response
func CreateListMediaResponse() (response *ListMediaResponse) {
	response = &ListMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
