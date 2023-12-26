package ddospro

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

// ModifyDomainTransmitRule invokes the ddospro.ModifyDomainTransmitRule API synchronously
// api document: https://help.aliyun.com/api/ddospro/modifydomaintransmitrule.html
func (client *Client) ModifyDomainTransmitRule(request *ModifyDomainTransmitRuleRequest) (response *ModifyDomainTransmitRuleResponse, err error) {
	response = CreateModifyDomainTransmitRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDomainTransmitRuleWithChan invokes the ddospro.ModifyDomainTransmitRule API asynchronously
// api document: https://help.aliyun.com/api/ddospro/modifydomaintransmitrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainTransmitRuleWithChan(request *ModifyDomainTransmitRuleRequest) (<-chan *ModifyDomainTransmitRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyDomainTransmitRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDomainTransmitRule(request)
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

// ModifyDomainTransmitRuleWithCallback invokes the ddospro.ModifyDomainTransmitRule API asynchronously
// api document: https://help.aliyun.com/api/ddospro/modifydomaintransmitrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainTransmitRuleWithCallback(request *ModifyDomainTransmitRuleRequest, callback func(response *ModifyDomainTransmitRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDomainTransmitRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyDomainTransmitRule(request)
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

// ModifyDomainTransmitRuleRequest is the request struct for api ModifyDomainTransmitRule
type ModifyDomainTransmitRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
	Ip              string           `position:"Query" name:"Ip"`
	Type            string           `position:"Query" name:"Type"`
	RealServer      *[]string        `position:"Query" name:"RealServer"  type:"Repeated"`
}

// ModifyDomainTransmitRuleResponse is the response struct for api ModifyDomainTransmitRule
type ModifyDomainTransmitRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDomainTransmitRuleRequest creates a request to invoke ModifyDomainTransmitRule API
func CreateModifyDomainTransmitRuleRequest() (request *ModifyDomainTransmitRuleRequest) {
	request = &ModifyDomainTransmitRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "ModifyDomainTransmitRule", "", "")
	return
}

// CreateModifyDomainTransmitRuleResponse creates a response to parse from ModifyDomainTransmitRule response
func CreateModifyDomainTransmitRuleResponse() (response *ModifyDomainTransmitRuleResponse) {
	response = &ModifyDomainTransmitRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
