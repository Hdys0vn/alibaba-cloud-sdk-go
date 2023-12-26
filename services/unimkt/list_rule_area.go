package unimkt

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

// ListRuleArea invokes the unimkt.ListRuleArea API synchronously
func (client *Client) ListRuleArea(request *ListRuleAreaRequest) (response *ListRuleAreaResponse, err error) {
	response = CreateListRuleAreaResponse()
	err = client.DoAction(request, response)
	return
}

// ListRuleAreaWithChan invokes the unimkt.ListRuleArea API asynchronously
func (client *Client) ListRuleAreaWithChan(request *ListRuleAreaRequest) (<-chan *ListRuleAreaResponse, <-chan error) {
	responseChan := make(chan *ListRuleAreaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRuleArea(request)
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

// ListRuleAreaWithCallback invokes the unimkt.ListRuleArea API asynchronously
func (client *Client) ListRuleAreaWithCallback(request *ListRuleAreaRequest, callback func(response *ListRuleAreaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRuleAreaResponse
		var err error
		defer close(result)
		response, err = client.ListRuleArea(request)
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

// ListRuleAreaRequest is the request struct for api ListRuleArea
type ListRuleAreaRequest struct {
	*requests.RpcRequest
	AdSlotType       string           `position:"Query" name:"AdSlotType"`
	RuleName         string           `position:"Query" name:"RuleName"`
	UserId           string           `position:"Query" name:"UserId"`
	OriginSiteUserId string           `position:"Query" name:"OriginSiteUserId"`
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	AppName          string           `position:"Query" name:"AppName"`
	TenantId         string           `position:"Query" name:"TenantId"`
	AdSlotId         string           `position:"Query" name:"AdSlotId"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	EndCreateTime    requests.Integer `position:"Query" name:"EndCreateTime"`
	Business         string           `position:"Query" name:"Business"`
	RuleType         string           `position:"Query" name:"RuleType"`
	MediaId          string           `position:"Query" name:"MediaId"`
	MediaStatus      string           `position:"Query" name:"MediaStatus"`
	Environment      string           `position:"Query" name:"Environment"`
	StartCreateTime  requests.Integer `position:"Query" name:"StartCreateTime"`
	UserSite         string           `position:"Query" name:"UserSite"`
	RuleId           string           `position:"Query" name:"RuleId"`
}

// ListRuleAreaResponse is the response struct for api ListRuleArea
type ListRuleAreaResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Success    bool   `json:"Success" xml:"Success"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Total      int64  `json:"Total" xml:"Total"`
	Model      []Rule `json:"Model" xml:"Model"`
}

// CreateListRuleAreaRequest creates a request to invoke ListRuleArea API
func CreateListRuleAreaRequest() (request *ListRuleAreaRequest) {
	request = &ListRuleAreaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "ListRuleArea", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRuleAreaResponse creates a response to parse from ListRuleArea response
func CreateListRuleAreaResponse() (response *ListRuleAreaResponse) {
	response = &ListRuleAreaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
