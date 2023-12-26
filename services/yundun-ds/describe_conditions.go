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

// DescribeConditions invokes the yundun_ds.DescribeConditions API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeconditions.html
func (client *Client) DescribeConditions(request *DescribeConditionsRequest) (response *DescribeConditionsResponse, err error) {
	response = CreateDescribeConditionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConditionsWithChan invokes the yundun_ds.DescribeConditions API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeconditions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConditionsWithChan(request *DescribeConditionsRequest) (<-chan *DescribeConditionsResponse, <-chan error) {
	responseChan := make(chan *DescribeConditionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConditions(request)
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

// DescribeConditionsWithCallback invokes the yundun_ds.DescribeConditions API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describeconditions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConditionsWithCallback(request *DescribeConditionsRequest, callback func(response *DescribeConditionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConditionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeConditions(request)
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

// DescribeConditionsRequest is the request struct for api DescribeConditions
type DescribeConditionsRequest struct {
	*requests.RpcRequest
	ProductCode string           `position:"Query" name:"ProductCode"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	SearchType  requests.Integer `position:"Query" name:"SearchType"`
	Lang        string           `position:"Query" name:"Lang"`
	QueryType   requests.Integer `position:"Query" name:"QueryType"`
}

// DescribeConditionsResponse is the response struct for api DescribeConditions
type DescribeConditionsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Condition Condition `json:"Condition" xml:"Condition"`
}

// CreateDescribeConditionsRequest creates a request to invoke DescribeConditions API
func CreateDescribeConditionsRequest() (request *DescribeConditionsRequest) {
	request = &DescribeConditionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeConditions", "sddp", "openAPI")
	return
}

// CreateDescribeConditionsResponse creates a response to parse from DescribeConditions response
func CreateDescribeConditionsResponse() (response *DescribeConditionsResponse) {
	response = &DescribeConditionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}