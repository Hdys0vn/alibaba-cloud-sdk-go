package elasticsearch

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

// DescribeTemplates invokes the elasticsearch.DescribeTemplates API synchronously
func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
	response = CreateDescribeTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTemplatesWithChan invokes the elasticsearch.DescribeTemplates API asynchronously
func (client *Client) DescribeTemplatesWithChan(request *DescribeTemplatesRequest) (<-chan *DescribeTemplatesResponse, <-chan error) {
	responseChan := make(chan *DescribeTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTemplates(request)
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

// DescribeTemplatesWithCallback invokes the elasticsearch.DescribeTemplates API asynchronously
func (client *Client) DescribeTemplatesWithCallback(request *DescribeTemplatesRequest, callback func(response *DescribeTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTemplatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeTemplates(request)
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

// DescribeTemplatesRequest is the request struct for api DescribeTemplates
type DescribeTemplatesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
}

// DescribeTemplatesResponse is the response struct for api DescribeTemplates
type DescribeTemplatesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateDescribeTemplatesRequest creates a request to invoke DescribeTemplates API
func CreateDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DescribeTemplates", "/openapi/instances/[InstanceId]/templates", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeTemplatesResponse creates a response to parse from DescribeTemplates response
func CreateDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
