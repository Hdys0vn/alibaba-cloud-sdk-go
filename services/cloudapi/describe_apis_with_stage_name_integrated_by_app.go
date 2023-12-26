package cloudapi

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

// DescribeApisWithStageNameIntegratedByApp invokes the cloudapi.DescribeApisWithStageNameIntegratedByApp API synchronously
func (client *Client) DescribeApisWithStageNameIntegratedByApp(request *DescribeApisWithStageNameIntegratedByAppRequest) (response *DescribeApisWithStageNameIntegratedByAppResponse, err error) {
	response = CreateDescribeApisWithStageNameIntegratedByAppResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApisWithStageNameIntegratedByAppWithChan invokes the cloudapi.DescribeApisWithStageNameIntegratedByApp API asynchronously
func (client *Client) DescribeApisWithStageNameIntegratedByAppWithChan(request *DescribeApisWithStageNameIntegratedByAppRequest) (<-chan *DescribeApisWithStageNameIntegratedByAppResponse, <-chan error) {
	responseChan := make(chan *DescribeApisWithStageNameIntegratedByAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApisWithStageNameIntegratedByApp(request)
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

// DescribeApisWithStageNameIntegratedByAppWithCallback invokes the cloudapi.DescribeApisWithStageNameIntegratedByApp API asynchronously
func (client *Client) DescribeApisWithStageNameIntegratedByAppWithCallback(request *DescribeApisWithStageNameIntegratedByAppRequest, callback func(response *DescribeApisWithStageNameIntegratedByAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApisWithStageNameIntegratedByAppResponse
		var err error
		defer close(result)
		response, err = client.DescribeApisWithStageNameIntegratedByApp(request)
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

// DescribeApisWithStageNameIntegratedByAppRequest is the request struct for api DescribeApisWithStageNameIntegratedByApp
type DescribeApisWithStageNameIntegratedByAppRequest struct {
	*requests.RpcRequest
	Method        string           `position:"Query" name:"Method"`
	Description   string           `position:"Query" name:"Description"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	Path          string           `position:"Query" name:"Path"`
	ApiName       string           `position:"Query" name:"ApiName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	AppId         requests.Integer `position:"Query" name:"AppId"`
	ApiUid        string           `position:"Query" name:"ApiUid"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeApisWithStageNameIntegratedByAppResponse is the response struct for api DescribeApisWithStageNameIntegratedByApp
type DescribeApisWithStageNameIntegratedByAppResponse struct {
	*responses.BaseResponse
	PageNumber          int                                                           `json:"PageNumber" xml:"PageNumber"`
	RequestId           string                                                        `json:"RequestId" xml:"RequestId"`
	PageSize            int                                                           `json:"PageSize" xml:"PageSize"`
	TotalCount          int                                                           `json:"TotalCount" xml:"TotalCount"`
	AppApiRelationInfos AppApiRelationInfosInDescribeApisWithStageNameIntegratedByApp `json:"AppApiRelationInfos" xml:"AppApiRelationInfos"`
}

// CreateDescribeApisWithStageNameIntegratedByAppRequest creates a request to invoke DescribeApisWithStageNameIntegratedByApp API
func CreateDescribeApisWithStageNameIntegratedByAppRequest() (request *DescribeApisWithStageNameIntegratedByAppRequest) {
	request = &DescribeApisWithStageNameIntegratedByAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisWithStageNameIntegratedByApp", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeApisWithStageNameIntegratedByAppResponse creates a response to parse from DescribeApisWithStageNameIntegratedByApp response
func CreateDescribeApisWithStageNameIntegratedByAppResponse() (response *DescribeApisWithStageNameIntegratedByAppResponse) {
	response = &DescribeApisWithStageNameIntegratedByAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
