package jarvis

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

// CreateCdnSubscription invokes the jarvis.CreateCdnSubscription API synchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnsubscription.html
func (client *Client) CreateCdnSubscription(request *CreateCdnSubscriptionRequest) (response *CreateCdnSubscriptionResponse, err error) {
	response = CreateCreateCdnSubscriptionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCdnSubscriptionWithChan invokes the jarvis.CreateCdnSubscription API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnsubscription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCdnSubscriptionWithChan(request *CreateCdnSubscriptionRequest) (<-chan *CreateCdnSubscriptionResponse, <-chan error) {
	responseChan := make(chan *CreateCdnSubscriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCdnSubscription(request)
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

// CreateCdnSubscriptionWithCallback invokes the jarvis.CreateCdnSubscription API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnsubscription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCdnSubscriptionWithCallback(request *CreateCdnSubscriptionRequest, callback func(response *CreateCdnSubscriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCdnSubscriptionResponse
		var err error
		defer close(result)
		response, err = client.CreateCdnSubscription(request)
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

// CreateCdnSubscriptionRequest is the request struct for api CreateCdnSubscription
type CreateCdnSubscriptionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	CdnUidList      string           `position:"Query" name:"CdnUidList"`
	SourceCode      string           `position:"Query" name:"SourceCode"`
}

// CreateCdnSubscriptionResponse is the response struct for api CreateCdnSubscription
type CreateCdnSubscriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateCreateCdnSubscriptionRequest creates a request to invoke CreateCdnSubscription API
func CreateCreateCdnSubscriptionRequest() (request *CreateCdnSubscriptionRequest) {
	request = &CreateCdnSubscriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "CreateCdnSubscription", "jarvis", "openAPI")
	return
}

// CreateCreateCdnSubscriptionResponse creates a response to parse from CreateCdnSubscription response
func CreateCreateCdnSubscriptionResponse() (response *CreateCdnSubscriptionResponse) {
	response = &CreateCdnSubscriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
