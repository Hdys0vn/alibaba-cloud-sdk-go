package cr_ee

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

// DeleteChartRelease invokes the cr.DeleteChartRelease API synchronously
// api document: https://help.aliyun.com/api/cr/deletechartrelease.html
func (client *Client) DeleteChartRelease(request *DeleteChartReleaseRequest) (response *DeleteChartReleaseResponse, err error) {
	response = CreateDeleteChartReleaseResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteChartReleaseWithChan invokes the cr.DeleteChartRelease API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartReleaseWithChan(request *DeleteChartReleaseRequest) (<-chan *DeleteChartReleaseResponse, <-chan error) {
	responseChan := make(chan *DeleteChartReleaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteChartRelease(request)
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

// DeleteChartReleaseWithCallback invokes the cr.DeleteChartRelease API asynchronously
// api document: https://help.aliyun.com/api/cr/deletechartrelease.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteChartReleaseWithCallback(request *DeleteChartReleaseRequest, callback func(response *DeleteChartReleaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteChartReleaseResponse
		var err error
		defer close(result)
		response, err = client.DeleteChartRelease(request)
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

// DeleteChartReleaseRequest is the request struct for api DeleteChartRelease
type DeleteChartReleaseRequest struct {
	*requests.RpcRequest
	Release           string `position:"Query" name:"Release"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	RepoName          string `position:"Query" name:"RepoName"`
	RepoNamespaceName string `position:"Query" name:"RepoNamespaceName"`
	Chart             string `position:"Query" name:"Chart"`
}

// DeleteChartReleaseResponse is the response struct for api DeleteChartRelease
type DeleteChartReleaseResponse struct {
	*responses.BaseResponse
	DeleteChartReleaseIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                        string `json:"Code" xml:"Code"`
	RequestId                   string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteChartReleaseRequest creates a request to invoke DeleteChartRelease API
func CreateDeleteChartReleaseRequest() (request *DeleteChartReleaseRequest) {
	request = &DeleteChartReleaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "DeleteChartRelease", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteChartReleaseResponse creates a response to parse from DeleteChartRelease response
func CreateDeleteChartReleaseResponse() (response *DeleteChartReleaseResponse) {
	response = &DeleteChartReleaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
