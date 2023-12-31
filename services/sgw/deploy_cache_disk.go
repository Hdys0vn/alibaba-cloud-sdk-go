package sgw

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

// DeployCacheDisk invokes the sgw.DeployCacheDisk API synchronously
func (client *Client) DeployCacheDisk(request *DeployCacheDiskRequest) (response *DeployCacheDiskResponse, err error) {
	response = CreateDeployCacheDiskResponse()
	err = client.DoAction(request, response)
	return
}

// DeployCacheDiskWithChan invokes the sgw.DeployCacheDisk API asynchronously
func (client *Client) DeployCacheDiskWithChan(request *DeployCacheDiskRequest) (<-chan *DeployCacheDiskResponse, <-chan error) {
	responseChan := make(chan *DeployCacheDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployCacheDisk(request)
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

// DeployCacheDiskWithCallback invokes the sgw.DeployCacheDisk API asynchronously
func (client *Client) DeployCacheDiskWithCallback(request *DeployCacheDiskRequest, callback func(response *DeployCacheDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployCacheDiskResponse
		var err error
		defer close(result)
		response, err = client.DeployCacheDisk(request)
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

// DeployCacheDiskRequest is the request struct for api DeployCacheDisk
type DeployCacheDiskRequest struct {
	*requests.RpcRequest
	SizeInGB      requests.Integer `position:"Query" name:"SizeInGB"`
	CacheConfig   string           `position:"Query" name:"CacheConfig"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DiskCategory  string           `position:"Query" name:"DiskCategory"`
	GatewayId     string           `position:"Query" name:"GatewayId"`
}

// DeployCacheDiskResponse is the response struct for api DeployCacheDisk
type DeployCacheDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateDeployCacheDiskRequest creates a request to invoke DeployCacheDisk API
func CreateDeployCacheDiskRequest() (request *DeployCacheDiskRequest) {
	request = &DeployCacheDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DeployCacheDisk", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployCacheDiskResponse creates a response to parse from DeployCacheDisk response
func CreateDeployCacheDiskResponse() (response *DeployCacheDiskResponse) {
	response = &DeployCacheDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
