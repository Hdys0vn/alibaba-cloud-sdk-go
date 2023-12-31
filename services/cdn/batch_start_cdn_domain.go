package cdn

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

// BatchStartCdnDomain invokes the cdn.BatchStartCdnDomain API synchronously
func (client *Client) BatchStartCdnDomain(request *BatchStartCdnDomainRequest) (response *BatchStartCdnDomainResponse, err error) {
	response = CreateBatchStartCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStartCdnDomainWithChan invokes the cdn.BatchStartCdnDomain API asynchronously
func (client *Client) BatchStartCdnDomainWithChan(request *BatchStartCdnDomainRequest) (<-chan *BatchStartCdnDomainResponse, <-chan error) {
	responseChan := make(chan *BatchStartCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStartCdnDomain(request)
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

// BatchStartCdnDomainWithCallback invokes the cdn.BatchStartCdnDomain API asynchronously
func (client *Client) BatchStartCdnDomainWithCallback(request *BatchStartCdnDomainRequest, callback func(response *BatchStartCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStartCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.BatchStartCdnDomain(request)
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

// BatchStartCdnDomainRequest is the request struct for api BatchStartCdnDomain
type BatchStartCdnDomainRequest struct {
	*requests.RpcRequest
	DomainNames   string           `position:"Query" name:"DomainNames"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchStartCdnDomainResponse is the response struct for api BatchStartCdnDomain
type BatchStartCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchStartCdnDomainRequest creates a request to invoke BatchStartCdnDomain API
func CreateBatchStartCdnDomainRequest() (request *BatchStartCdnDomainRequest) {
	request = &BatchStartCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "BatchStartCdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchStartCdnDomainResponse creates a response to parse from BatchStartCdnDomain response
func CreateBatchStartCdnDomainResponse() (response *BatchStartCdnDomainResponse) {
	response = &BatchStartCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
