package mse

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

// UpdateGatewayRouteCORS invokes the mse.UpdateGatewayRouteCORS API synchronously
func (client *Client) UpdateGatewayRouteCORS(request *UpdateGatewayRouteCORSRequest) (response *UpdateGatewayRouteCORSResponse, err error) {
	response = CreateUpdateGatewayRouteCORSResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayRouteCORSWithChan invokes the mse.UpdateGatewayRouteCORS API asynchronously
func (client *Client) UpdateGatewayRouteCORSWithChan(request *UpdateGatewayRouteCORSRequest) (<-chan *UpdateGatewayRouteCORSResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayRouteCORSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayRouteCORS(request)
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

// UpdateGatewayRouteCORSWithCallback invokes the mse.UpdateGatewayRouteCORS API asynchronously
func (client *Client) UpdateGatewayRouteCORSWithCallback(request *UpdateGatewayRouteCORSRequest, callback func(response *UpdateGatewayRouteCORSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayRouteCORSResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayRouteCORS(request)
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

// UpdateGatewayRouteCORSRequest is the request struct for api UpdateGatewayRouteCORS
type UpdateGatewayRouteCORSRequest struct {
	*requests.RpcRequest
	MseSessionId    string                         `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string                         `position:"Query" name:"GatewayUniqueId"`
	Id              requests.Integer               `position:"Query" name:"Id"`
	GatewayId       requests.Integer               `position:"Query" name:"GatewayId"`
	CorsJSON        UpdateGatewayRouteCORSCorsJSON `position:"Query" name:"CorsJSON"  type:"Struct"`
	AcceptLanguage  string                         `position:"Query" name:"AcceptLanguage"`
}

// UpdateGatewayRouteCORSCorsJSON is a repeated param struct in UpdateGatewayRouteCORSRequest
type UpdateGatewayRouteCORSCorsJSON struct {
	AllowCredentials string `name:"AllowCredentials"`
	AllowOrigins     string `name:"AllowOrigins"`
	AllowMethods     string `name:"AllowMethods"`
	AllowHeaders     string `name:"AllowHeaders"`
	ExposeHeaders    string `name:"ExposeHeaders"`
	TimeUnit         string `name:"TimeUnit"`
	UnitNum          string `name:"UnitNum"`
	Status           string `name:"Status"`
}

// UpdateGatewayRouteCORSResponse is the response struct for api UpdateGatewayRouteCORS
type UpdateGatewayRouteCORSResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayRouteCORSRequest creates a request to invoke UpdateGatewayRouteCORS API
func CreateUpdateGatewayRouteCORSRequest() (request *UpdateGatewayRouteCORSRequest) {
	request = &UpdateGatewayRouteCORSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayRouteCORS", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayRouteCORSResponse creates a response to parse from UpdateGatewayRouteCORS response
func CreateUpdateGatewayRouteCORSResponse() (response *UpdateGatewayRouteCORSResponse) {
	response = &UpdateGatewayRouteCORSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
