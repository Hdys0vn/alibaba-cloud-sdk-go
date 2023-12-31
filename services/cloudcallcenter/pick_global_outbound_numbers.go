package cloudcallcenter

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

// PickGlobalOutboundNumbers invokes the cloudcallcenter.PickGlobalOutboundNumbers API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pickglobaloutboundnumbers.html
func (client *Client) PickGlobalOutboundNumbers(request *PickGlobalOutboundNumbersRequest) (response *PickGlobalOutboundNumbersResponse, err error) {
	response = CreatePickGlobalOutboundNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// PickGlobalOutboundNumbersWithChan invokes the cloudcallcenter.PickGlobalOutboundNumbers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pickglobaloutboundnumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PickGlobalOutboundNumbersWithChan(request *PickGlobalOutboundNumbersRequest) (<-chan *PickGlobalOutboundNumbersResponse, <-chan error) {
	responseChan := make(chan *PickGlobalOutboundNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PickGlobalOutboundNumbers(request)
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

// PickGlobalOutboundNumbersWithCallback invokes the cloudcallcenter.PickGlobalOutboundNumbers API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/pickglobaloutboundnumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PickGlobalOutboundNumbersWithCallback(request *PickGlobalOutboundNumbersRequest, callback func(response *PickGlobalOutboundNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PickGlobalOutboundNumbersResponse
		var err error
		defer close(result)
		response, err = client.PickGlobalOutboundNumbers(request)
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

// PickGlobalOutboundNumbersRequest is the request struct for api PickGlobalOutboundNumbers
type PickGlobalOutboundNumbersRequest struct {
	*requests.RpcRequest
	IsVirtual    requests.Boolean `position:"Query" name:"IsVirtual"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId *[]string        `position:"Query" name:"SkillGroupId"  type:"Repeated"`
	Count        requests.Integer `position:"Query" name:"Count"`
	CalleeNumber string           `position:"Query" name:"CalleeNumber"`
}

// PickGlobalOutboundNumbersResponse is the response struct for api PickGlobalOutboundNumbers
type PickGlobalOutboundNumbersResponse struct {
	*responses.BaseResponse
	RequestId       string                                     `json:"RequestId" xml:"RequestId"`
	Success         bool                                       `json:"Success" xml:"Success"`
	Code            string                                     `json:"Code" xml:"Code"`
	Message         string                                     `json:"Message" xml:"Message"`
	HttpStatusCode  int                                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DialNumberPairs DialNumberPairsInPickGlobalOutboundNumbers `json:"DialNumberPairs" xml:"DialNumberPairs"`
}

// CreatePickGlobalOutboundNumbersRequest creates a request to invoke PickGlobalOutboundNumbers API
func CreatePickGlobalOutboundNumbersRequest() (request *PickGlobalOutboundNumbersRequest) {
	request = &PickGlobalOutboundNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "PickGlobalOutboundNumbers", "", "")
	request.Method = requests.POST
	return
}

// CreatePickGlobalOutboundNumbersResponse creates a response to parse from PickGlobalOutboundNumbers response
func CreatePickGlobalOutboundNumbersResponse() (response *PickGlobalOutboundNumbersResponse) {
	response = &PickGlobalOutboundNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
