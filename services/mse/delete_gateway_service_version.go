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

// DeleteGatewayServiceVersion invokes the mse.DeleteGatewayServiceVersion API synchronously
func (client *Client) DeleteGatewayServiceVersion(request *DeleteGatewayServiceVersionRequest) (response *DeleteGatewayServiceVersionResponse, err error) {
	response = CreateDeleteGatewayServiceVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayServiceVersionWithChan invokes the mse.DeleteGatewayServiceVersion API asynchronously
func (client *Client) DeleteGatewayServiceVersionWithChan(request *DeleteGatewayServiceVersionRequest) (<-chan *DeleteGatewayServiceVersionResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayServiceVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayServiceVersion(request)
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

// DeleteGatewayServiceVersionWithCallback invokes the mse.DeleteGatewayServiceVersion API asynchronously
func (client *Client) DeleteGatewayServiceVersionWithCallback(request *DeleteGatewayServiceVersionRequest, callback func(response *DeleteGatewayServiceVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayServiceVersionResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayServiceVersion(request)
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

// DeleteGatewayServiceVersionRequest is the request struct for api DeleteGatewayServiceVersion
type DeleteGatewayServiceVersionRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	ServiceVersion  string           `position:"Query" name:"ServiceVersion"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	ServiceId       requests.Integer `position:"Query" name:"ServiceId"`
}

// DeleteGatewayServiceVersionResponse is the response struct for api DeleteGatewayServiceVersion
type DeleteGatewayServiceVersionResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateDeleteGatewayServiceVersionRequest creates a request to invoke DeleteGatewayServiceVersion API
func CreateDeleteGatewayServiceVersionRequest() (request *DeleteGatewayServiceVersionRequest) {
	request = &DeleteGatewayServiceVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteGatewayServiceVersion", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayServiceVersionResponse creates a response to parse from DeleteGatewayServiceVersion response
func CreateDeleteGatewayServiceVersionResponse() (response *DeleteGatewayServiceVersionResponse) {
	response = &DeleteGatewayServiceVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}