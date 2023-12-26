package aegis

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

// DescribeSasLeftCondition invokes the aegis.DescribeSasLeftCondition API synchronously
// api document: https://help.aliyun.com/api/aegis/describesasleftcondition.html
func (client *Client) DescribeSasLeftCondition(request *DescribeSasLeftConditionRequest) (response *DescribeSasLeftConditionResponse, err error) {
	response = CreateDescribeSasLeftConditionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSasLeftConditionWithChan invokes the aegis.DescribeSasLeftCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesasleftcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSasLeftConditionWithChan(request *DescribeSasLeftConditionRequest) (<-chan *DescribeSasLeftConditionResponse, <-chan error) {
	responseChan := make(chan *DescribeSasLeftConditionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSasLeftCondition(request)
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

// DescribeSasLeftConditionWithCallback invokes the aegis.DescribeSasLeftCondition API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesasleftcondition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSasLeftConditionWithCallback(request *DescribeSasLeftConditionRequest, callback func(response *DescribeSasLeftConditionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSasLeftConditionResponse
		var err error
		defer close(result)
		response, err = client.DescribeSasLeftCondition(request)
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

// DescribeSasLeftConditionRequest is the request struct for api DescribeSasLeftCondition
type DescribeSasLeftConditionRequest struct {
	*requests.RpcRequest
	SourceIp         string `position:"Query" name:"SourceIp"`
	ConditionType    string `position:"Query" name:"ConditionType"`
	Lang             string `position:"Query" name:"Lang"`
	FilterConditions string `position:"Query" name:"FilterConditions"`
}

// DescribeSasLeftConditionResponse is the response struct for api DescribeSasLeftCondition
type DescribeSasLeftConditionResponse struct {
	*responses.BaseResponse
	RequestId     string      `json:"RequestId" xml:"RequestId"`
	ConditionList []Condition `json:"ConditionList" xml:"ConditionList"`
}

// CreateDescribeSasLeftConditionRequest creates a request to invoke DescribeSasLeftCondition API
func CreateDescribeSasLeftConditionRequest() (request *DescribeSasLeftConditionRequest) {
	request = &DescribeSasLeftConditionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSasLeftCondition", "vipaegis", "openAPI")
	return
}

// CreateDescribeSasLeftConditionResponse creates a response to parse from DescribeSasLeftCondition response
func CreateDescribeSasLeftConditionResponse() (response *DescribeSasLeftConditionResponse) {
	response = &DescribeSasLeftConditionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
