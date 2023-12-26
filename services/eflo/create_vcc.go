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

// CreateVcc invokes the eflo.CreateVcc API synchronously
func (client *Client) CreateVcc(request *CreateVccRequest) (response *CreateVccResponse, err error) {
	response = CreateCreateVccResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVccWithChan invokes the eflo.CreateVcc API asynchronously
func (client *Client) CreateVccWithChan(request *CreateVccRequest) (<-chan *CreateVccResponse, <-chan error) {
	responseChan := make(chan *CreateVccResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVcc(request)
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

// CreateVccWithCallback invokes the eflo.CreateVcc API asynchronously
func (client *Client) CreateVccWithCallback(request *CreateVccRequest, callback func(response *CreateVccResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVccResponse
		var err error
		defer close(result)
		response, err = client.CreateVcc(request)
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

// CreateVccRequest is the request struct for api CreateVcc
type CreateVccRequest struct {
	*requests.RpcRequest
	BgpCidr            string           `position:"Body" name:"BgpCidr"`
	CenId              string           `position:"Body" name:"CenId"`
	Description        string           `position:"Body" name:"Description"`
	CenOwnerId         string           `position:"Body" name:"CenOwnerId"`
	AccessCouldService requests.Boolean `position:"Body" name:"AccessCouldService"`
	ResourceGroupId    string           `position:"Body" name:"ResourceGroupId"`
	VccName            string           `position:"Body" name:"VccName"`
	Tag                *[]CreateVccTag  `position:"Body" name:"Tag"  type:"Repeated"`
	VccId              string           `position:"Body" name:"VccId"`
	ConnectionType     string           `position:"Body" name:"ConnectionType"`
	Bandwidth          requests.Integer `position:"Body" name:"Bandwidth"`
	VSwitchId          string           `position:"Body" name:"VSwitchId"`
	VpdId              string           `position:"Body" name:"VpdId"`
	VpcId              string           `position:"Body" name:"VpcId"`
	ZoneId             string           `position:"Body" name:"ZoneId"`
}

// CreateVccTag is a repeated param struct in CreateVccRequest
type CreateVccTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateVccResponse is the response struct for api CreateVcc
type CreateVccResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateCreateVccRequest creates a request to invoke CreateVcc API
func CreateCreateVccRequest() (request *CreateVccRequest) {
	request = &CreateVccRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "CreateVcc", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVccResponse creates a response to parse from CreateVcc response
func CreateCreateVccResponse() (response *CreateVccResponse) {
	response = &CreateVccResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}