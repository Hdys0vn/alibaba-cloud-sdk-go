package smartag

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

// DiagnoseSmartAccessGateway invokes the smartag.DiagnoseSmartAccessGateway API synchronously
func (client *Client) DiagnoseSmartAccessGateway(request *DiagnoseSmartAccessGatewayRequest) (response *DiagnoseSmartAccessGatewayResponse, err error) {
	response = CreateDiagnoseSmartAccessGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DiagnoseSmartAccessGatewayWithChan invokes the smartag.DiagnoseSmartAccessGateway API asynchronously
func (client *Client) DiagnoseSmartAccessGatewayWithChan(request *DiagnoseSmartAccessGatewayRequest) (<-chan *DiagnoseSmartAccessGatewayResponse, <-chan error) {
	responseChan := make(chan *DiagnoseSmartAccessGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DiagnoseSmartAccessGateway(request)
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

// DiagnoseSmartAccessGatewayWithCallback invokes the smartag.DiagnoseSmartAccessGateway API asynchronously
func (client *Client) DiagnoseSmartAccessGatewayWithCallback(request *DiagnoseSmartAccessGatewayRequest, callback func(response *DiagnoseSmartAccessGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DiagnoseSmartAccessGatewayResponse
		var err error
		defer close(result)
		response, err = client.DiagnoseSmartAccessGateway(request)
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

// DiagnoseSmartAccessGatewayRequest is the request struct for api DiagnoseSmartAccessGateway
type DiagnoseSmartAccessGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DiagnoseSmartAccessGatewayResponse is the response struct for api DiagnoseSmartAccessGateway
type DiagnoseSmartAccessGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDiagnoseSmartAccessGatewayRequest creates a request to invoke DiagnoseSmartAccessGateway API
func CreateDiagnoseSmartAccessGatewayRequest() (request *DiagnoseSmartAccessGatewayRequest) {
	request = &DiagnoseSmartAccessGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DiagnoseSmartAccessGateway", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDiagnoseSmartAccessGatewayResponse creates a response to parse from DiagnoseSmartAccessGateway response
func CreateDiagnoseSmartAccessGatewayResponse() (response *DiagnoseSmartAccessGatewayResponse) {
	response = &DiagnoseSmartAccessGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
