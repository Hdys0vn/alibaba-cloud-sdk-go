package jarvis_public

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

// DescribeAttackedIp invokes the jarvis_public.DescribeAttackedIp API synchronously
// api document: https://help.aliyun.com/api/jarvis-public/describeattackedip.html
func (client *Client) DescribeAttackedIp(request *DescribeAttackedIpRequest) (response *DescribeAttackedIpResponse, err error) {
	response = CreateDescribeAttackedIpResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAttackedIpWithChan invokes the jarvis_public.DescribeAttackedIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis-public/describeattackedip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAttackedIpWithChan(request *DescribeAttackedIpRequest) (<-chan *DescribeAttackedIpResponse, <-chan error) {
	responseChan := make(chan *DescribeAttackedIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAttackedIp(request)
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

// DescribeAttackedIpWithCallback invokes the jarvis_public.DescribeAttackedIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis-public/describeattackedip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAttackedIpWithCallback(request *DescribeAttackedIpRequest, callback func(response *DescribeAttackedIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAttackedIpResponse
		var err error
		defer close(result)
		response, err = client.DescribeAttackedIp(request)
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

// DescribeAttackedIpRequest is the request struct for api DescribeAttackedIp
type DescribeAttackedIpRequest struct {
	*requests.RpcRequest
	SourceIp     string           `position:"Query" name:"SourceIp"`
	ServerIpList string           `position:"Query" name:"ServerIpList"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	Lang         string           `position:"Query" name:"Lang"`
	Region       string           `position:"Query" name:"Region"`
	ProductType  string           `position:"Query" name:"ProductType"`
}

// DescribeAttackedIpResponse is the response struct for api DescribeAttackedIp
type DescribeAttackedIpResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Module    string   `json:"Module" xml:"Module"`
	IpList    []string `json:"IpList" xml:"IpList"`
}

// CreateDescribeAttackedIpRequest creates a request to invoke DescribeAttackedIp API
func CreateDescribeAttackedIpRequest() (request *DescribeAttackedIpRequest) {
	request = &DescribeAttackedIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis-public", "2018-06-21", "DescribeAttackedIp", "jarvis-public", "openAPI")
	return
}

// CreateDescribeAttackedIpResponse creates a response to parse from DescribeAttackedIp response
func CreateDescribeAttackedIpResponse() (response *DescribeAttackedIpResponse) {
	response = &DescribeAttackedIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
