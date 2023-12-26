package cloudfw

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

// DescribeOutgoingDestinationIP invokes the cloudfw.DescribeOutgoingDestinationIP API synchronously
func (client *Client) DescribeOutgoingDestinationIP(request *DescribeOutgoingDestinationIPRequest) (response *DescribeOutgoingDestinationIPResponse, err error) {
	response = CreateDescribeOutgoingDestinationIPResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOutgoingDestinationIPWithChan invokes the cloudfw.DescribeOutgoingDestinationIP API asynchronously
func (client *Client) DescribeOutgoingDestinationIPWithChan(request *DescribeOutgoingDestinationIPRequest) (<-chan *DescribeOutgoingDestinationIPResponse, <-chan error) {
	responseChan := make(chan *DescribeOutgoingDestinationIPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOutgoingDestinationIP(request)
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

// DescribeOutgoingDestinationIPWithCallback invokes the cloudfw.DescribeOutgoingDestinationIP API asynchronously
func (client *Client) DescribeOutgoingDestinationIPWithCallback(request *DescribeOutgoingDestinationIPRequest, callback func(response *DescribeOutgoingDestinationIPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOutgoingDestinationIPResponse
		var err error
		defer close(result)
		response, err = client.DescribeOutgoingDestinationIP(request)
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

// DescribeOutgoingDestinationIPRequest is the request struct for api DescribeOutgoingDestinationIP
type DescribeOutgoingDestinationIPRequest struct {
	*requests.RpcRequest
	SecuritySuggest string `position:"Query" name:"SecuritySuggest"`
	AclCoverage     string `position:"Query" name:"AclCoverage"`
	PublicIP        string `position:"Query" name:"PublicIP"`
	StartTime       string `position:"Query" name:"StartTime"`
	AclStatus       string `position:"Query" name:"AclStatus"`
	PageSize        string `position:"Query" name:"PageSize"`
	DstIP           string `position:"Query" name:"DstIP"`
	Lang            string `position:"Query" name:"Lang"`
	CategoryId      string `position:"Query" name:"CategoryId"`
	Order           string `position:"Query" name:"Order"`
	TagId           string `position:"Query" name:"TagId"`
	PrivateIP       string `position:"Query" name:"PrivateIP"`
	EndTime         string `position:"Query" name:"EndTime"`
	CurrentPage     string `position:"Query" name:"CurrentPage"`
	Sort            string `position:"Query" name:"Sort"`
	ApplicationName string `position:"Query" name:"ApplicationName"`
	Port            string `position:"Query" name:"Port"`
	TagIdNew        string `position:"Query" name:"TagIdNew"`
}

// DescribeOutgoingDestinationIPResponse is the response struct for api DescribeOutgoingDestinationIP
type DescribeOutgoingDestinationIPResponse struct {
	*responses.BaseResponse
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	DstIPList  []DstIPData `json:"DstIPList" xml:"DstIPList"`
}

// CreateDescribeOutgoingDestinationIPRequest creates a request to invoke DescribeOutgoingDestinationIP API
func CreateDescribeOutgoingDestinationIPRequest() (request *DescribeOutgoingDestinationIPRequest) {
	request = &DescribeOutgoingDestinationIPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeOutgoingDestinationIP", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOutgoingDestinationIPResponse creates a response to parse from DescribeOutgoingDestinationIP response
func CreateDescribeOutgoingDestinationIPResponse() (response *DescribeOutgoingDestinationIPResponse) {
	response = &DescribeOutgoingDestinationIPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}