package eflo

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

// ListSubnets invokes the eflo.ListSubnets API synchronously
func (client *Client) ListSubnets(request *ListSubnetsRequest) (response *ListSubnetsResponse, err error) {
	response = CreateListSubnetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSubnetsWithChan invokes the eflo.ListSubnets API asynchronously
func (client *Client) ListSubnetsWithChan(request *ListSubnetsRequest) (<-chan *ListSubnetsResponse, <-chan error) {
	responseChan := make(chan *ListSubnetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSubnets(request)
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

// ListSubnetsWithCallback invokes the eflo.ListSubnets API asynchronously
func (client *Client) ListSubnetsWithCallback(request *ListSubnetsRequest, callback func(response *ListSubnetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSubnetsResponse
		var err error
		defer close(result)
		response, err = client.ListSubnets(request)
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

// ListSubnetsRequest is the request struct for api ListSubnets
type ListSubnetsRequest struct {
	*requests.RpcRequest
	Type            string            `position:"Body" name:"Type"`
	PageNumber      requests.Integer  `position:"Body" name:"PageNumber"`
	ResourceGroupId string            `position:"Body" name:"ResourceGroupId"`
	PageSize        requests.Integer  `position:"Body" name:"PageSize"`
	Tag             *[]ListSubnetsTag `position:"Body" name:"Tag"  type:"Repeated"`
	SubnetId        string            `position:"Body" name:"SubnetId"`
	VpdId           string            `position:"Body" name:"VpdId"`
	EnablePage      requests.Boolean  `position:"Body" name:"EnablePage"`
	ZoneId          string            `position:"Body" name:"ZoneId"`
	SubnetName      string            `position:"Body" name:"SubnetName"`
	Status          string            `position:"Body" name:"Status"`
}

// ListSubnetsTag is a repeated param struct in ListSubnetsRequest
type ListSubnetsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListSubnetsResponse is the response struct for api ListSubnets
type ListSubnetsResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateListSubnetsRequest creates a request to invoke ListSubnets API
func CreateListSubnetsRequest() (request *ListSubnetsRequest) {
	request = &ListSubnetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListSubnets", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSubnetsResponse creates a response to parse from ListSubnets response
func CreateListSubnetsResponse() (response *ListSubnetsResponse) {
	response = &ListSubnetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
