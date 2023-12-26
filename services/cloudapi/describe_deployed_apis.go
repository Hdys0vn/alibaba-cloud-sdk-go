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

// DescribeDeployedApis invokes the cloudapi.DescribeDeployedApis API synchronously
func (client *Client) DescribeDeployedApis(request *DescribeDeployedApisRequest) (response *DescribeDeployedApisResponse, err error) {
	response = CreateDescribeDeployedApisResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeployedApisWithChan invokes the cloudapi.DescribeDeployedApis API asynchronously
func (client *Client) DescribeDeployedApisWithChan(request *DescribeDeployedApisRequest) (<-chan *DescribeDeployedApisResponse, <-chan error) {
	responseChan := make(chan *DescribeDeployedApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeployedApis(request)
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

// DescribeDeployedApisWithCallback invokes the cloudapi.DescribeDeployedApis API asynchronously
func (client *Client) DescribeDeployedApisWithCallback(request *DescribeDeployedApisRequest, callback func(response *DescribeDeployedApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeployedApisResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeployedApis(request)
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

// DescribeDeployedApisRequest is the request struct for api DescribeDeployedApis
type DescribeDeployedApisRequest struct {
	*requests.RpcRequest
	StageName     string                     `position:"Query" name:"StageName"`
	PageNumber    requests.Integer           `position:"Query" name:"PageNumber"`
	SecurityToken string                     `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer           `position:"Query" name:"PageSize"`
	Tag           *[]DescribeDeployedApisTag `position:"Query" name:"Tag"  type:"Repeated"`
	ApiMethod     string                     `position:"Query" name:"ApiMethod"`
	GroupId       string                     `position:"Query" name:"GroupId"`
	ApiPath       string                     `position:"Query" name:"ApiPath"`
	EnableTagAuth requests.Boolean           `position:"Query" name:"EnableTagAuth"`
	ApiName       string                     `position:"Query" name:"ApiName"`
	ApiId         string                     `position:"Query" name:"ApiId"`
}

// DescribeDeployedApisTag is a repeated param struct in DescribeDeployedApisRequest
type DescribeDeployedApisTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDeployedApisResponse is the response struct for api DescribeDeployedApis
type DescribeDeployedApisResponse struct {
	*responses.BaseResponse
	PageNumber   int          `json:"PageNumber" xml:"PageNumber"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	TotalCount   int          `json:"TotalCount" xml:"TotalCount"`
	DeployedApis DeployedApis `json:"DeployedApis" xml:"DeployedApis"`
}

// CreateDescribeDeployedApisRequest creates a request to invoke DescribeDeployedApis API
func CreateDescribeDeployedApisRequest() (request *DescribeDeployedApisRequest) {
	request = &DescribeDeployedApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDeployedApis", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDeployedApisResponse creates a response to parse from DescribeDeployedApis response
func CreateDescribeDeployedApisResponse() (response *DescribeDeployedApisResponse) {
	response = &DescribeDeployedApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}