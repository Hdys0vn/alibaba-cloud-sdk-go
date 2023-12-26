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

// ModifyACLRule invokes the smartag.ModifyACLRule API synchronously
func (client *Client) ModifyACLRule(request *ModifyACLRuleRequest) (response *ModifyACLRuleResponse, err error) {
	response = CreateModifyACLRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyACLRuleWithChan invokes the smartag.ModifyACLRule API asynchronously
func (client *Client) ModifyACLRuleWithChan(request *ModifyACLRuleRequest) (<-chan *ModifyACLRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyACLRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyACLRule(request)
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

// ModifyACLRuleWithCallback invokes the smartag.ModifyACLRule API asynchronously
func (client *Client) ModifyACLRuleWithCallback(request *ModifyACLRuleRequest, callback func(response *ModifyACLRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyACLRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyACLRule(request)
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

// ModifyACLRuleRequest is the request struct for api ModifyACLRule
type ModifyACLRuleRequest struct {
	*requests.RpcRequest
	DpiGroupIds          *[]string        `position:"Query" name:"DpiGroupIds"  type:"Repeated"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange      string           `position:"Query" name:"SourcePortRange"`
	SourceCidr           string           `position:"Query" name:"SourceCidr"`
	Description          string           `position:"Query" name:"Description"`
	Type                 string           `position:"Query" name:"Type"`
	DestCidr             string           `position:"Query" name:"DestCidr"`
	DpiSignatureIds      *[]string        `position:"Query" name:"DpiSignatureIds"  type:"Repeated"`
	Direction            string           `position:"Query" name:"Direction"`
	Policy               string           `position:"Query" name:"Policy"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	AcrId                string           `position:"Query" name:"AcrId"`
	DestPortRange        string           `position:"Query" name:"DestPortRange"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyACLRuleResponse is the response struct for api ModifyACLRule
type ModifyACLRuleResponse struct {
	*responses.BaseResponse
	Policy          string                         `json:"Policy" xml:"Policy"`
	Description     string                         `json:"Description" xml:"Description"`
	RequestId       string                         `json:"RequestId" xml:"RequestId"`
	SourcePortRange string                         `json:"SourcePortRange" xml:"SourcePortRange"`
	SourceCidr      string                         `json:"SourceCidr" xml:"SourceCidr"`
	Priority        int                            `json:"Priority" xml:"Priority"`
	AclId           string                         `json:"AclId" xml:"AclId"`
	AcrId           string                         `json:"AcrId" xml:"AcrId"`
	DestPortRange   string                         `json:"DestPortRange" xml:"DestPortRange"`
	Direction       string                         `json:"Direction" xml:"Direction"`
	Name            string                         `json:"Name" xml:"Name"`
	GmtCreate       int64                          `json:"GmtCreate" xml:"GmtCreate"`
	DestCidr        string                         `json:"DestCidr" xml:"DestCidr"`
	IpProtocol      string                         `json:"IpProtocol" xml:"IpProtocol"`
	DpiGroupIds     DpiGroupIdsInModifyACLRule     `json:"DpiGroupIds" xml:"DpiGroupIds"`
	DpiSignatureIds DpiSignatureIdsInModifyACLRule `json:"DpiSignatureIds" xml:"DpiSignatureIds"`
}

// CreateModifyACLRuleRequest creates a request to invoke ModifyACLRule API
func CreateModifyACLRuleRequest() (request *ModifyACLRuleRequest) {
	request = &ModifyACLRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyACLRule", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyACLRuleResponse creates a response to parse from ModifyACLRule response
func CreateModifyACLRuleResponse() (response *ModifyACLRuleResponse) {
	response = &ModifyACLRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
