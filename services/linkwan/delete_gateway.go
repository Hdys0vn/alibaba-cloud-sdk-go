package linkwan

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

// DeleteGateway invokes the linkwan.DeleteGateway API synchronously
func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (response *DeleteGatewayResponse, err error) {
	response = CreateDeleteGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayWithChan invokes the linkwan.DeleteGateway API asynchronously
func (client *Client) DeleteGatewayWithChan(request *DeleteGatewayRequest) (<-chan *DeleteGatewayResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGateway(request)
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

// DeleteGatewayWithCallback invokes the linkwan.DeleteGateway API asynchronously
func (client *Client) DeleteGatewayWithCallback(request *DeleteGatewayRequest, callback func(response *DeleteGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayResponse
		var err error
		defer close(result)
		response, err = client.DeleteGateway(request)
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

// DeleteGatewayRequest is the request struct for api DeleteGateway
type DeleteGatewayRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	GwEui         string `position:"Query" name:"GwEui"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// DeleteGatewayResponse is the response struct for api DeleteGateway
type DeleteGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteGatewayRequest creates a request to invoke DeleteGateway API
func CreateDeleteGatewayRequest() (request *DeleteGatewayRequest) {
	request = &DeleteGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "DeleteGateway", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayResponse creates a response to parse from DeleteGateway response
func CreateDeleteGatewayResponse() (response *DeleteGatewayResponse) {
	response = &DeleteGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
