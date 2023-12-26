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

// DescribeIndexTemplate invokes the elasticsearch.DescribeIndexTemplate API synchronously
func (client *Client) DescribeIndexTemplate(request *DescribeIndexTemplateRequest) (response *DescribeIndexTemplateResponse, err error) {
	response = CreateDescribeIndexTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIndexTemplateWithChan invokes the elasticsearch.DescribeIndexTemplate API asynchronously
func (client *Client) DescribeIndexTemplateWithChan(request *DescribeIndexTemplateRequest) (<-chan *DescribeIndexTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeIndexTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIndexTemplate(request)
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

// DescribeIndexTemplateWithCallback invokes the elasticsearch.DescribeIndexTemplate API asynchronously
func (client *Client) DescribeIndexTemplateWithCallback(request *DescribeIndexTemplateRequest, callback func(response *DescribeIndexTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIndexTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeIndexTemplate(request)
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

// DescribeIndexTemplateRequest is the request struct for api DescribeIndexTemplate
type DescribeIndexTemplateRequest struct {
	*requests.RoaRequest
	InstanceId    string `position:"Path" name:"InstanceId"`
	IndexTemplate string `position:"Path" name:"IndexTemplate"`
}

// DescribeIndexTemplateResponse is the response struct for api DescribeIndexTemplate
type DescribeIndexTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeIndexTemplateRequest creates a request to invoke DescribeIndexTemplate API
func CreateDescribeIndexTemplateRequest() (request *DescribeIndexTemplateRequest) {
	request = &DescribeIndexTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "DescribeIndexTemplate", "/openapi/instances/[InstanceId]/index-templates/[IndexTemplate]", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeIndexTemplateResponse creates a response to parse from DescribeIndexTemplate response
func CreateDescribeIndexTemplateResponse() (response *DescribeIndexTemplateResponse) {
	response = &DescribeIndexTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
