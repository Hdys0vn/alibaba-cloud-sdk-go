package baas

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

// DeleteFabricChaincode invokes the baas.DeleteFabricChaincode API synchronously
// api document: https://help.aliyun.com/api/baas/deletefabricchaincode.html
func (client *Client) DeleteFabricChaincode(request *DeleteFabricChaincodeRequest) (response *DeleteFabricChaincodeResponse, err error) {
	response = CreateDeleteFabricChaincodeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFabricChaincodeWithChan invokes the baas.DeleteFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/deletefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFabricChaincodeWithChan(request *DeleteFabricChaincodeRequest) (<-chan *DeleteFabricChaincodeResponse, <-chan error) {
	responseChan := make(chan *DeleteFabricChaincodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFabricChaincode(request)
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

// DeleteFabricChaincodeWithCallback invokes the baas.DeleteFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/deletefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFabricChaincodeWithCallback(request *DeleteFabricChaincodeRequest, callback func(response *DeleteFabricChaincodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFabricChaincodeResponse
		var err error
		defer close(result)
		response, err = client.DeleteFabricChaincode(request)
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

// DeleteFabricChaincodeRequest is the request struct for api DeleteFabricChaincode
type DeleteFabricChaincodeRequest struct {
	*requests.RpcRequest
	ChaincodeId string `position:"Body" name:"ChaincodeId"`
}

// DeleteFabricChaincodeResponse is the response struct for api DeleteFabricChaincode
type DeleteFabricChaincodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeleteFabricChaincodeRequest creates a request to invoke DeleteFabricChaincode API
func CreateDeleteFabricChaincodeRequest() (request *DeleteFabricChaincodeRequest) {
	request = &DeleteFabricChaincodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DeleteFabricChaincode", "baas", "openAPI")
	return
}

// CreateDeleteFabricChaincodeResponse creates a response to parse from DeleteFabricChaincode response
func CreateDeleteFabricChaincodeResponse() (response *DeleteFabricChaincodeResponse) {
	response = &DeleteFabricChaincodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
