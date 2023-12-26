package polardb

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

// DescribeDBClusterServerlessConf invokes the polardb.DescribeDBClusterServerlessConf API synchronously
func (client *Client) DescribeDBClusterServerlessConf(request *DescribeDBClusterServerlessConfRequest) (response *DescribeDBClusterServerlessConfResponse, err error) {
	response = CreateDescribeDBClusterServerlessConfResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterServerlessConfWithChan invokes the polardb.DescribeDBClusterServerlessConf API asynchronously
func (client *Client) DescribeDBClusterServerlessConfWithChan(request *DescribeDBClusterServerlessConfRequest) (<-chan *DescribeDBClusterServerlessConfResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterServerlessConfResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterServerlessConf(request)
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

// DescribeDBClusterServerlessConfWithCallback invokes the polardb.DescribeDBClusterServerlessConf API asynchronously
func (client *Client) DescribeDBClusterServerlessConfWithCallback(request *DescribeDBClusterServerlessConfRequest, callback func(response *DescribeDBClusterServerlessConfResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterServerlessConfResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterServerlessConf(request)
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

// DescribeDBClusterServerlessConfRequest is the request struct for api DescribeDBClusterServerlessConf
type DescribeDBClusterServerlessConfRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RequestType          string           `position:"Query" name:"RequestType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterServerlessConfResponse is the response struct for api DescribeDBClusterServerlessConf
type DescribeDBClusterServerlessConfResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	ScaleMin              string `json:"ScaleMin" xml:"ScaleMin"`
	ScaleMax              string `json:"ScaleMax" xml:"ScaleMax"`
	ScaleRoNumMin         string `json:"ScaleRoNumMin" xml:"ScaleRoNumMin"`
	ScaleRoNumMax         string `json:"ScaleRoNumMax" xml:"ScaleRoNumMax"`
	AllowShutDown         string `json:"AllowShutDown" xml:"AllowShutDown"`
	SecondsUntilAutoPause string `json:"SecondsUntilAutoPause" xml:"SecondsUntilAutoPause"`
	DBClusterId           string `json:"DBClusterId" xml:"DBClusterId"`
	ScaleApRoNumMin       string `json:"ScaleApRoNumMin" xml:"ScaleApRoNumMin"`
	ScaleApRoNumMax       string `json:"ScaleApRoNumMax" xml:"ScaleApRoNumMax"`
	Switchs               string `json:"Switchs" xml:"Switchs"`
	DBMaxscaleId          string `json:"DBMaxscaleId" xml:"DBMaxscaleId"`
}

// CreateDescribeDBClusterServerlessConfRequest creates a request to invoke DescribeDBClusterServerlessConf API
func CreateDescribeDBClusterServerlessConfRequest() (request *DescribeDBClusterServerlessConfRequest) {
	request = &DescribeDBClusterServerlessConfRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterServerlessConf", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterServerlessConfResponse creates a response to parse from DescribeDBClusterServerlessConf response
func CreateDescribeDBClusterServerlessConfResponse() (response *DescribeDBClusterServerlessConfResponse) {
	response = &DescribeDBClusterServerlessConfResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
