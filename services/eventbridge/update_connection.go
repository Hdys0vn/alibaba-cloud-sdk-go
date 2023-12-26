package eventbridge

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

// UpdateConnection invokes the eventbridge.UpdateConnection API synchronously
func (client *Client) UpdateConnection(request *UpdateConnectionRequest) (response *UpdateConnectionResponse, err error) {
	response = CreateUpdateConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConnectionWithChan invokes the eventbridge.UpdateConnection API asynchronously
func (client *Client) UpdateConnectionWithChan(request *UpdateConnectionRequest) (<-chan *UpdateConnectionResponse, <-chan error) {
	responseChan := make(chan *UpdateConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConnection(request)
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

// UpdateConnectionWithCallback invokes the eventbridge.UpdateConnection API asynchronously
func (client *Client) UpdateConnectionWithCallback(request *UpdateConnectionRequest, callback func(response *UpdateConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConnectionResponse
		var err error
		defer close(result)
		response, err = client.UpdateConnection(request)
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

// UpdateConnectionRequest is the request struct for api UpdateConnection
type UpdateConnectionRequest struct {
	*requests.RpcRequest
	ConnectionName    string                            `position:"Query" name:"ConnectionName"`
	Description       string                            `position:"Query" name:"Description"`
	NetworkParameters UpdateConnectionNetworkParameters `position:"Query" name:"NetworkParameters"  type:"Struct"`
	AuthParameters    UpdateConnectionAuthParameters    `position:"Query" name:"AuthParameters"  type:"Struct"`
}

// UpdateConnectionNetworkParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionNetworkParameters struct {
	VpcId           string `name:"VpcId"`
	SecurityGroupId string `name:"SecurityGroupId"`
	NetworkType     string `name:"NetworkType"`
	VswitcheId      string `name:"VswitcheId"`
}

// UpdateConnectionAuthParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParameters struct {
	BasicAuthParameters      UpdateConnectionAuthParametersBasicAuthParameters      `name:"BasicAuthParameters" type:"Struct"`
	ApiKeyAuthParameters     UpdateConnectionAuthParametersApiKeyAuthParameters     `name:"ApiKeyAuthParameters" type:"Struct"`
	AuthorizationType        string                                                 `name:"AuthorizationType"`
	InvocationHttpParameters UpdateConnectionAuthParametersInvocationHttpParameters `name:"InvocationHttpParameters" type:"Struct"`
	OAuthParameters          UpdateConnectionAuthParametersOAuthParameters          `name:"OAuthParameters" type:"Struct"`
}

// UpdateConnectionAuthParametersBasicAuthParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersBasicAuthParameters struct {
	Password string `name:"Password"`
	Username string `name:"Username"`
}

// UpdateConnectionAuthParametersApiKeyAuthParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersApiKeyAuthParameters struct {
	ApiKeyName  string `name:"ApiKeyName"`
	ApiKeyValue string `name:"ApiKeyValue"`
}

// UpdateConnectionAuthParametersInvocationHttpParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersInvocationHttpParameters struct {
	BodyParameters        *[]UpdateConnectionAuthParametersInvocationHttpParametersBodyParametersItem        `name:"BodyParameters" type:"Repeated"`
	HeaderParameters      *[]UpdateConnectionAuthParametersInvocationHttpParametersHeaderParametersItem      `name:"HeaderParameters" type:"Repeated"`
	QueryStringParameters *[]UpdateConnectionAuthParametersInvocationHttpParametersQueryStringParametersItem `name:"QueryStringParameters" type:"Repeated"`
}

// UpdateConnectionAuthParametersOAuthParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParameters struct {
	ClientParameters      UpdateConnectionAuthParametersOAuthParametersClientParameters    `name:"ClientParameters" type:"Struct"`
	AuthorizationEndpoint string                                                           `name:"AuthorizationEndpoint"`
	HttpMethod            string                                                           `name:"HttpMethod"`
	OAuthHttpParameters   UpdateConnectionAuthParametersOAuthParametersOAuthHttpParameters `name:"OAuthHttpParameters" type:"Struct"`
}

// UpdateConnectionAuthParametersInvocationHttpParametersBodyParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersInvocationHttpParametersBodyParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionAuthParametersInvocationHttpParametersHeaderParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersInvocationHttpParametersHeaderParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionAuthParametersInvocationHttpParametersQueryStringParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersInvocationHttpParametersQueryStringParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionAuthParametersOAuthParametersClientParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParametersClientParameters struct {
	ClientID     string `name:"ClientID"`
	ClientSecret string `name:"ClientSecret"`
}

// UpdateConnectionAuthParametersOAuthParametersOAuthHttpParameters is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParametersOAuthHttpParameters struct {
	BodyParameters        *[]UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersItem        `name:"BodyParameters" type:"Repeated"`
	HeaderParameters      *[]UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersItem      `name:"HeaderParameters" type:"Repeated"`
	QueryStringParameters *[]UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersItem `name:"QueryStringParameters" type:"Repeated"`
}

// UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersBodyParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersHeaderParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersItem is a repeated param struct in UpdateConnectionRequest
type UpdateConnectionAuthParametersOAuthParametersOAuthHttpParametersQueryStringParametersItem struct {
	IsValueSecret string `name:"IsValueSecret"`
	Value         string `name:"Value"`
	Key           string `name:"Key"`
}

// UpdateConnectionResponse is the response struct for api UpdateConnection
type UpdateConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateUpdateConnectionRequest creates a request to invoke UpdateConnection API
func CreateUpdateConnectionRequest() (request *UpdateConnectionRequest) {
	request = &UpdateConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "UpdateConnection", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateConnectionResponse creates a response to parse from UpdateConnection response
func CreateUpdateConnectionResponse() (response *UpdateConnectionResponse) {
	response = &UpdateConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}