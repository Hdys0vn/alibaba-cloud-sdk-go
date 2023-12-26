package alb

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

// CreateRules invokes the alb.CreateRules API synchronously
func (client *Client) CreateRules(request *CreateRulesRequest) (response *CreateRulesResponse, err error) {
	response = CreateCreateRulesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRulesWithChan invokes the alb.CreateRules API asynchronously
func (client *Client) CreateRulesWithChan(request *CreateRulesRequest) (<-chan *CreateRulesResponse, <-chan error) {
	responseChan := make(chan *CreateRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRules(request)
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

// CreateRulesWithCallback invokes the alb.CreateRules API asynchronously
func (client *Client) CreateRulesWithCallback(request *CreateRulesRequest, callback func(response *CreateRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRulesResponse
		var err error
		defer close(result)
		response, err = client.CreateRules(request)
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

// CreateRulesRequest is the request struct for api CreateRules
type CreateRulesRequest struct {
	*requests.RpcRequest
	ClientToken string              `position:"Query" name:"ClientToken"`
	Rules       *[]CreateRulesRules `position:"Query" name:"Rules"  type:"Repeated"`
	ListenerId  string              `position:"Query" name:"ListenerId"`
	DryRun      requests.Boolean    `position:"Query" name:"DryRun"`
}

// CreateRulesRules is a repeated param struct in CreateRulesRequest
type CreateRulesRules struct {
	RuleConditions *[]CreateRulesRulesRuleConditionsItem `name:"RuleConditions" type:"Repeated"`
	RuleName       string                                `name:"RuleName"`
	Priority       string                                `name:"Priority"`
	RuleActions    *[]CreateRulesRulesRuleActionsItem    `name:"RuleActions" type:"Repeated"`
	Direction      string                                `name:"Direction"`
}

// CreateRulesRulesRuleConditionsItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItem struct {
	MethodConfig             CreateRulesRulesRuleConditionsItemMethodConfig             `name:"MethodConfig" type:"Struct"`
	SourceIpConfig           CreateRulesRulesRuleConditionsItemSourceIpConfig           `name:"SourceIpConfig" type:"Struct"`
	HostConfig               CreateRulesRulesRuleConditionsItemHostConfig               `name:"HostConfig" type:"Struct"`
	QueryStringConfig        CreateRulesRulesRuleConditionsItemQueryStringConfig        `name:"QueryStringConfig" type:"Struct"`
	ResponseStatusCodeConfig CreateRulesRulesRuleConditionsItemResponseStatusCodeConfig `name:"ResponseStatusCodeConfig" type:"Struct"`
	PathConfig               CreateRulesRulesRuleConditionsItemPathConfig               `name:"PathConfig" type:"Struct"`
	CookieConfig             CreateRulesRulesRuleConditionsItemCookieConfig             `name:"CookieConfig" type:"Struct"`
	Type                     string                                                     `name:"Type"`
	HeaderConfig             CreateRulesRulesRuleConditionsItemHeaderConfig             `name:"HeaderConfig" type:"Struct"`
	ResponseHeaderConfig     CreateRulesRulesRuleConditionsItemResponseHeaderConfig     `name:"ResponseHeaderConfig" type:"Struct"`
}

// CreateRulesRulesRuleActionsItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItem struct {
	FixedResponseConfig CreateRulesRulesRuleActionsItemFixedResponseConfig `name:"FixedResponseConfig" type:"Struct"`
	TrafficMirrorConfig CreateRulesRulesRuleActionsItemTrafficMirrorConfig `name:"TrafficMirrorConfig" type:"Struct"`
	ForwardGroupConfig  CreateRulesRulesRuleActionsItemForwardGroupConfig  `name:"ForwardGroupConfig" type:"Struct"`
	RemoveHeaderConfig  CreateRulesRulesRuleActionsItemRemoveHeaderConfig  `name:"RemoveHeaderConfig" type:"Struct"`
	InsertHeaderConfig  CreateRulesRulesRuleActionsItemInsertHeaderConfig  `name:"InsertHeaderConfig" type:"Struct"`
	TrafficLimitConfig  CreateRulesRulesRuleActionsItemTrafficLimitConfig  `name:"TrafficLimitConfig" type:"Struct"`
	CorsConfig          CreateRulesRulesRuleActionsItemCorsConfig          `name:"CorsConfig" type:"Struct"`
	RedirectConfig      CreateRulesRulesRuleActionsItemRedirectConfig      `name:"RedirectConfig" type:"Struct"`
	Type                string                                             `name:"Type"`
	Order               string                                             `name:"Order"`
	RewriteConfig       CreateRulesRulesRuleActionsItemRewriteConfig       `name:"RewriteConfig" type:"Struct"`
}

// CreateRulesRulesRuleConditionsItemMethodConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemMethodConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemSourceIpConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemSourceIpConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemHostConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemHostConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemQueryStringConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemQueryStringConfig struct {
	Values *[]CreateRulesRulesRuleConditionsItemQueryStringConfigValuesItem `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemResponseStatusCodeConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemResponseStatusCodeConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemPathConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemPathConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemCookieConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemCookieConfig struct {
	Values *[]CreateRulesRulesRuleConditionsItemCookieConfigValuesItem `name:"Values" type:"Repeated"`
}

// CreateRulesRulesRuleConditionsItemHeaderConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemHeaderConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
	Key    string    `name:"Key"`
}

// CreateRulesRulesRuleConditionsItemResponseHeaderConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemResponseHeaderConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
	Key    string    `name:"Key"`
}

// CreateRulesRulesRuleActionsItemFixedResponseConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemFixedResponseConfig struct {
	HttpCode    string `name:"HttpCode"`
	Content     string `name:"Content"`
	ContentType string `name:"ContentType"`
}

// CreateRulesRulesRuleActionsItemTrafficMirrorConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemTrafficMirrorConfig struct {
	MirrorGroupConfig CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfig `name:"MirrorGroupConfig" type:"Struct"`
	TargetType        string                                                              `name:"TargetType"`
}

// CreateRulesRulesRuleActionsItemForwardGroupConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemForwardGroupConfig struct {
	ServerGroupStickySession CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupStickySession `name:"ServerGroupStickySession" type:"Struct"`
	ServerGroupTuples        *[]CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupTuplesItem `name:"ServerGroupTuples" type:"Repeated"`
}

// CreateRulesRulesRuleActionsItemRemoveHeaderConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemRemoveHeaderConfig struct {
	Key string `name:"Key"`
}

// CreateRulesRulesRuleActionsItemInsertHeaderConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemInsertHeaderConfig struct {
	ValueType    string `name:"ValueType"`
	CoverEnabled string `name:"CoverEnabled"`
	Value        string `name:"Value"`
	Key          string `name:"Key"`
}

// CreateRulesRulesRuleActionsItemTrafficLimitConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemTrafficLimitConfig struct {
	QPS      string `name:"QPS"`
	PerIpQps string `name:"PerIpQps"`
}

// CreateRulesRulesRuleActionsItemCorsConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemCorsConfig struct {
	AllowCredentials string    `name:"AllowCredentials"`
	AllowOrigin      *[]string `name:"AllowOrigin" type:"Repeated"`
	MaxAge           string    `name:"MaxAge"`
	AllowMethods     *[]string `name:"AllowMethods" type:"Repeated"`
	AllowHeaders     *[]string `name:"AllowHeaders" type:"Repeated"`
	ExposeHeaders    *[]string `name:"ExposeHeaders" type:"Repeated"`
}

// CreateRulesRulesRuleActionsItemRedirectConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemRedirectConfig struct {
	Path     string `name:"Path"`
	Protocol string `name:"Protocol"`
	Port     string `name:"Port"`
	Query    string `name:"Query"`
	Host     string `name:"Host"`
	HttpCode string `name:"HttpCode"`
}

// CreateRulesRulesRuleActionsItemRewriteConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemRewriteConfig struct {
	Path  string `name:"Path"`
	Query string `name:"Query"`
	Host  string `name:"Host"`
}

// CreateRulesRulesRuleConditionsItemQueryStringConfigValuesItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemQueryStringConfigValuesItem struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateRulesRulesRuleConditionsItemCookieConfigValuesItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleConditionsItemCookieConfigValuesItem struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfig is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples *[]CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem `name:"ServerGroupTuples" type:"Repeated"`
}

// CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupStickySession is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupStickySession struct {
	Enabled string `name:"Enabled"`
	Timeout string `name:"Timeout"`
}

// CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupTuplesItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemForwardGroupConfigServerGroupTuplesItem struct {
	ServerGroupId string `name:"ServerGroupId"`
	Weight        string `name:"Weight"`
}

// CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem is a repeated param struct in CreateRulesRequest
type CreateRulesRulesRuleActionsItemTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem struct {
	ServerGroupId string `name:"ServerGroupId"`
}

// CreateRulesResponse is the response struct for api CreateRules
type CreateRulesResponse struct {
	*responses.BaseResponse
	JobId     string   `json:"JobId" xml:"JobId"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	RuleIds   []RuleId `json:"RuleIds" xml:"RuleIds"`
}

// CreateCreateRulesRequest creates a request to invoke CreateRules API
func CreateCreateRulesRequest() (request *CreateRulesRequest) {
	request = &CreateRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "CreateRules", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRulesResponse creates a response to parse from CreateRules response
func CreateCreateRulesResponse() (response *CreateRulesResponse) {
	response = &CreateRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
