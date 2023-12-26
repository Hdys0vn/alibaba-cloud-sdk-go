package mse

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

// DeleteGatewayDomain invokes the mse.DeleteGatewayDomain API synchronously
func (client *Client) DeleteGatewayDomain(request *DeleteGatewayDomainRequest) (response *DeleteGatewayDomainResponse, err error) {
	response = CreateDeleteGatewayDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayDomainWithChan invokes the mse.DeleteGatewayDomain API asynchronously
func (client *Client) DeleteGatewayDomainWithChan(request *DeleteGatewayDomainRequest) (<-chan *DeleteGatewayDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayDomain(request)
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

// DeleteGatewayDomainWithCallback invokes the mse.DeleteGatewayDomain API asynchronously
func (client *Client) DeleteGatewayDomainWithCallback(request *DeleteGatewayDomainRequest, callback func(response *DeleteGatewayDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayDomain(request)
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

// DeleteGatewayDomainRequest is the request struct for api DeleteGatewayDomain
type DeleteGatewayDomainRequest struct {
	*requests.RpcRequest
	MseSessionId    string `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string `position:"Query" name:"GatewayUniqueId"`
	Id              string `position:"Query" name:"Id"`
	AcceptLanguage  string `position:"Query" name:"AcceptLanguage"`
}

// DeleteGatewayDomainResponse is the response struct for api DeleteGatewayDomain
type DeleteGatewayDomainResponse struct {
	*responses.BaseResponse
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                    `json:"Message" xml:"Message"`
	Code           int                       `json:"Code" xml:"Code"`
	Success        bool                      `json:"Success" xml:"Success"`
	Data           DataInDeleteGatewayDomain `json:"Data" xml:"Data"`
}

// CreateDeleteGatewayDomainRequest creates a request to invoke DeleteGatewayDomain API
func CreateDeleteGatewayDomainRequest() (request *DeleteGatewayDomainRequest) {
	request = &DeleteGatewayDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteGatewayDomain", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayDomainResponse creates a response to parse from DeleteGatewayDomain response
func CreateDeleteGatewayDomainResponse() (response *DeleteGatewayDomainResponse) {
	response = &DeleteGatewayDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}