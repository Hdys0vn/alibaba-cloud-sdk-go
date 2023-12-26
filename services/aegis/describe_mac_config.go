package aegis

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

// DescribeMacConfig invokes the aegis.DescribeMacConfig API synchronously
// api document: https://help.aliyun.com/api/aegis/describemacconfig.html
func (client *Client) DescribeMacConfig(request *DescribeMacConfigRequest) (response *DescribeMacConfigResponse, err error) {
	response = CreateDescribeMacConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMacConfigWithChan invokes the aegis.DescribeMacConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describemacconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMacConfigWithChan(request *DescribeMacConfigRequest) (<-chan *DescribeMacConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeMacConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMacConfig(request)
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

// DescribeMacConfigWithCallback invokes the aegis.DescribeMacConfig API asynchronously
// api document: https://help.aliyun.com/api/aegis/describemacconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMacConfigWithCallback(request *DescribeMacConfigRequest, callback func(response *DescribeMacConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMacConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeMacConfig(request)
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

// DescribeMacConfigRequest is the request struct for api DescribeMacConfig
type DescribeMacConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Types           string           `position:"Query" name:"Types"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Extern          string           `position:"Query" name:"Extern"`
	Type            string           `position:"Query" name:"Type"`
	Config          string           `position:"Query" name:"Config"`
	Target          string           `position:"Query" name:"Target"`
}

// DescribeMacConfigResponse is the response struct for api DescribeMacConfig
type DescribeMacConfigResponse struct {
	*responses.BaseResponse
	RequestId   string   `json:"RequestId" xml:"RequestId"`
	Count       int      `json:"Count" xml:"Count"`
	PageSize    int      `json:"PageSize" xml:"PageSize"`
	CurrentPage int      `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int      `json:"TotalCount" xml:"TotalCount"`
	Configs     []Config `json:"Configs" xml:"Configs"`
}

// CreateDescribeMacConfigRequest creates a request to invoke DescribeMacConfig API
func CreateDescribeMacConfigRequest() (request *DescribeMacConfigRequest) {
	request = &DescribeMacConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeMacConfig", "vipaegis", "openAPI")
	return
}

// CreateDescribeMacConfigResponse creates a response to parse from DescribeMacConfig response
func CreateDescribeMacConfigResponse() (response *DescribeMacConfigResponse) {
	response = &DescribeMacConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
