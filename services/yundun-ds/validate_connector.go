package yundun_ds

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

// ValidateConnector invokes the yundun_ds.ValidateConnector API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/validateconnector.html
func (client *Client) ValidateConnector(request *ValidateConnectorRequest) (response *ValidateConnectorResponse, err error) {
	response = CreateValidateConnectorResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateConnectorWithChan invokes the yundun_ds.ValidateConnector API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/validateconnector.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ValidateConnectorWithChan(request *ValidateConnectorRequest) (<-chan *ValidateConnectorResponse, <-chan error) {
	responseChan := make(chan *ValidateConnectorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateConnector(request)
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

// ValidateConnectorWithCallback invokes the yundun_ds.ValidateConnector API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/validateconnector.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ValidateConnectorWithCallback(request *ValidateConnectorRequest, callback func(response *ValidateConnectorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateConnectorResponse
		var err error
		defer close(result)
		response, err = client.ValidateConnector(request)
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

// ValidateConnectorRequest is the request struct for api ValidateConnector
type ValidateConnectorRequest struct {
	*requests.RpcRequest
	Password        string           `position:"Query" name:"Password"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Connector       string           `position:"Query" name:"Connector"`
	Lang            string           `position:"Query" name:"Lang"`
	ResourceType    requests.Integer `position:"Query" name:"ResourceType"`
	ServiceRegionId string           `position:"Query" name:"ServiceRegionId"`
	ParentId        string           `position:"Query" name:"ParentId"`
	UserName        string           `position:"Query" name:"UserName"`
}

// ValidateConnectorResponse is the response struct for api ValidateConnector
type ValidateConnectorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Connected bool   `json:"Connected" xml:"Connected"`
}

// CreateValidateConnectorRequest creates a request to invoke ValidateConnector API
func CreateValidateConnectorRequest() (request *ValidateConnectorRequest) {
	request = &ValidateConnectorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "ValidateConnector", "sddp", "openAPI")
	return
}

// CreateValidateConnectorResponse creates a response to parse from ValidateConnector response
func CreateValidateConnectorResponse() (response *ValidateConnectorResponse) {
	response = &ValidateConnectorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
