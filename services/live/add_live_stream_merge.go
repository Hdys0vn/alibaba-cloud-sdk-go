package live

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

// AddLiveStreamMerge invokes the live.AddLiveStreamMerge API synchronously
func (client *Client) AddLiveStreamMerge(request *AddLiveStreamMergeRequest) (response *AddLiveStreamMergeResponse, err error) {
	response = CreateAddLiveStreamMergeResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveStreamMergeWithChan invokes the live.AddLiveStreamMerge API asynchronously
func (client *Client) AddLiveStreamMergeWithChan(request *AddLiveStreamMergeRequest) (<-chan *AddLiveStreamMergeResponse, <-chan error) {
	responseChan := make(chan *AddLiveStreamMergeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveStreamMerge(request)
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

// AddLiveStreamMergeWithCallback invokes the live.AddLiveStreamMerge API asynchronously
func (client *Client) AddLiveStreamMergeWithCallback(request *AddLiveStreamMergeRequest, callback func(response *AddLiveStreamMergeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveStreamMergeResponse
		var err error
		defer close(result)
		response, err = client.AddLiveStreamMerge(request)
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

// AddLiveStreamMergeRequest is the request struct for api AddLiveStreamMerge
type AddLiveStreamMergeRequest struct {
	*requests.RpcRequest
	InAppName2    string           `position:"Query" name:"InAppName2"`
	InAppName1    string           `position:"Query" name:"InAppName1"`
	StartTime     string           `position:"Query" name:"StartTime"`
	Protocol      string           `position:"Query" name:"Protocol"`
	AppName       string           `position:"Query" name:"AppName"`
	InStreamName2 string           `position:"Query" name:"InStreamName2"`
	StreamName    string           `position:"Query" name:"StreamName"`
	InStreamName1 string           `position:"Query" name:"InStreamName1"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// AddLiveStreamMergeResponse is the response struct for api AddLiveStreamMerge
type AddLiveStreamMergeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveStreamMergeRequest creates a request to invoke AddLiveStreamMerge API
func CreateAddLiveStreamMergeRequest() (request *AddLiveStreamMergeRequest) {
	request = &AddLiveStreamMergeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveStreamMerge", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLiveStreamMergeResponse creates a response to parse from AddLiveStreamMerge response
func CreateAddLiveStreamMergeResponse() (response *AddLiveStreamMergeResponse) {
	response = &AddLiveStreamMergeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
