package dcdn

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

// BatchSetDcdnIpaDomainConfigs invokes the dcdn.BatchSetDcdnIpaDomainConfigs API synchronously
func (client *Client) BatchSetDcdnIpaDomainConfigs(request *BatchSetDcdnIpaDomainConfigsRequest) (response *BatchSetDcdnIpaDomainConfigsResponse, err error) {
	response = CreateBatchSetDcdnIpaDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSetDcdnIpaDomainConfigsWithChan invokes the dcdn.BatchSetDcdnIpaDomainConfigs API asynchronously
func (client *Client) BatchSetDcdnIpaDomainConfigsWithChan(request *BatchSetDcdnIpaDomainConfigsRequest) (<-chan *BatchSetDcdnIpaDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchSetDcdnIpaDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSetDcdnIpaDomainConfigs(request)
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

// BatchSetDcdnIpaDomainConfigsWithCallback invokes the dcdn.BatchSetDcdnIpaDomainConfigs API asynchronously
func (client *Client) BatchSetDcdnIpaDomainConfigsWithCallback(request *BatchSetDcdnIpaDomainConfigsRequest, callback func(response *BatchSetDcdnIpaDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSetDcdnIpaDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchSetDcdnIpaDomainConfigs(request)
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

// BatchSetDcdnIpaDomainConfigsRequest is the request struct for api BatchSetDcdnIpaDomainConfigs
type BatchSetDcdnIpaDomainConfigsRequest struct {
	*requests.RpcRequest
	Functions     string           `position:"Query" name:"Functions"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchSetDcdnIpaDomainConfigsResponse is the response struct for api BatchSetDcdnIpaDomainConfigs
type BatchSetDcdnIpaDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchSetDcdnIpaDomainConfigsRequest creates a request to invoke BatchSetDcdnIpaDomainConfigs API
func CreateBatchSetDcdnIpaDomainConfigsRequest() (request *BatchSetDcdnIpaDomainConfigsRequest) {
	request = &BatchSetDcdnIpaDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "BatchSetDcdnIpaDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchSetDcdnIpaDomainConfigsResponse creates a response to parse from BatchSetDcdnIpaDomainConfigs response
func CreateBatchSetDcdnIpaDomainConfigsResponse() (response *BatchSetDcdnIpaDomainConfigsResponse) {
	response = &BatchSetDcdnIpaDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
