package r_kvstore

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

// SwitchInstanceProxy invokes the r_kvstore.SwitchInstanceProxy API synchronously
func (client *Client) SwitchInstanceProxy(request *SwitchInstanceProxyRequest) (response *SwitchInstanceProxyResponse, err error) {
	response = CreateSwitchInstanceProxyResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchInstanceProxyWithChan invokes the r_kvstore.SwitchInstanceProxy API asynchronously
func (client *Client) SwitchInstanceProxyWithChan(request *SwitchInstanceProxyRequest) (<-chan *SwitchInstanceProxyResponse, <-chan error) {
	responseChan := make(chan *SwitchInstanceProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchInstanceProxy(request)
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

// SwitchInstanceProxyWithCallback invokes the r_kvstore.SwitchInstanceProxy API asynchronously
func (client *Client) SwitchInstanceProxyWithCallback(request *SwitchInstanceProxyRequest, callback func(response *SwitchInstanceProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchInstanceProxyResponse
		var err error
		defer close(result)
		response, err = client.SwitchInstanceProxy(request)
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

// SwitchInstanceProxyRequest is the request struct for api SwitchInstanceProxy
type SwitchInstanceProxyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Product              string           `position:"Query" name:"Product"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Category             string           `position:"Query" name:"Category"`
}

// SwitchInstanceProxyResponse is the response struct for api SwitchInstanceProxy
type SwitchInstanceProxyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchInstanceProxyRequest creates a request to invoke SwitchInstanceProxy API
func CreateSwitchInstanceProxyRequest() (request *SwitchInstanceProxyRequest) {
	request = &SwitchInstanceProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "SwitchInstanceProxy", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchInstanceProxyResponse creates a response to parse from SwitchInstanceProxy response
func CreateSwitchInstanceProxyResponse() (response *SwitchInstanceProxyResponse) {
	response = &SwitchInstanceProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
