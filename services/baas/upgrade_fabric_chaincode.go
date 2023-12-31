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

// UpgradeFabricChaincode invokes the baas.UpgradeFabricChaincode API synchronously
// api document: https://help.aliyun.com/api/baas/upgradefabricchaincode.html
func (client *Client) UpgradeFabricChaincode(request *UpgradeFabricChaincodeRequest) (response *UpgradeFabricChaincodeResponse, err error) {
	response = CreateUpgradeFabricChaincodeResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeFabricChaincodeWithChan invokes the baas.UpgradeFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/upgradefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeFabricChaincodeWithChan(request *UpgradeFabricChaincodeRequest) (<-chan *UpgradeFabricChaincodeResponse, <-chan error) {
	responseChan := make(chan *UpgradeFabricChaincodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeFabricChaincode(request)
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

// UpgradeFabricChaincodeWithCallback invokes the baas.UpgradeFabricChaincode API asynchronously
// api document: https://help.aliyun.com/api/baas/upgradefabricchaincode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeFabricChaincodeWithCallback(request *UpgradeFabricChaincodeRequest, callback func(response *UpgradeFabricChaincodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeFabricChaincodeResponse
		var err error
		defer close(result)
		response, err = client.UpgradeFabricChaincode(request)
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

// UpgradeFabricChaincodeRequest is the request struct for api UpgradeFabricChaincode
type UpgradeFabricChaincodeRequest struct {
	*requests.RpcRequest
	EndorsePolicy    string `position:"Body" name:"EndorsePolicy"`
	OrganizationId   string `position:"Body" name:"OrganizationId"`
	ChaincodeId      string `position:"Body" name:"ChaincodeId"`
	CollectionConfig string `position:"Body" name:"CollectionConfig"`
	Location         string `position:"Body" name:"Location"`
}

// UpgradeFabricChaincodeResponse is the response struct for api UpgradeFabricChaincode
type UpgradeFabricChaincodeResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Success   bool                           `json:"Success" xml:"Success"`
	ErrorCode int                            `json:"ErrorCode" xml:"ErrorCode"`
	Result    ResultInUpgradeFabricChaincode `json:"Result" xml:"Result"`
}

// CreateUpgradeFabricChaincodeRequest creates a request to invoke UpgradeFabricChaincode API
func CreateUpgradeFabricChaincodeRequest() (request *UpgradeFabricChaincodeRequest) {
	request = &UpgradeFabricChaincodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "UpgradeFabricChaincode", "baas", "openAPI")
	return
}

// CreateUpgradeFabricChaincodeResponse creates a response to parse from UpgradeFabricChaincode response
func CreateUpgradeFabricChaincodeResponse() (response *UpgradeFabricChaincodeResponse) {
	response = &UpgradeFabricChaincodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
