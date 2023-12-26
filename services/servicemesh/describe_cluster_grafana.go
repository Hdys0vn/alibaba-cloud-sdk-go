package servicemesh

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

// DescribeClusterGrafana invokes the servicemesh.DescribeClusterGrafana API synchronously
func (client *Client) DescribeClusterGrafana(request *DescribeClusterGrafanaRequest) (response *DescribeClusterGrafanaResponse, err error) {
	response = CreateDescribeClusterGrafanaResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterGrafanaWithChan invokes the servicemesh.DescribeClusterGrafana API asynchronously
func (client *Client) DescribeClusterGrafanaWithChan(request *DescribeClusterGrafanaRequest) (<-chan *DescribeClusterGrafanaResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterGrafanaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterGrafana(request)
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

// DescribeClusterGrafanaWithCallback invokes the servicemesh.DescribeClusterGrafana API asynchronously
func (client *Client) DescribeClusterGrafanaWithCallback(request *DescribeClusterGrafanaRequest, callback func(response *DescribeClusterGrafanaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterGrafanaResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterGrafana(request)
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

// DescribeClusterGrafanaRequest is the request struct for api DescribeClusterGrafana
type DescribeClusterGrafanaRequest struct {
	*requests.RpcRequest
	K8sClusterId  string `position:"Query" name:"K8sClusterId"`
	ServiceMeshId string `position:"Query" name:"ServiceMeshId"`
}

// DescribeClusterGrafanaResponse is the response struct for api DescribeClusterGrafana
type DescribeClusterGrafanaResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	Dashboards []Dashboard `json:"Dashboards" xml:"Dashboards"`
}

// CreateDescribeClusterGrafanaRequest creates a request to invoke DescribeClusterGrafana API
func CreateDescribeClusterGrafanaRequest() (request *DescribeClusterGrafanaRequest) {
	request = &DescribeClusterGrafanaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "DescribeClusterGrafana", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterGrafanaResponse creates a response to parse from DescribeClusterGrafana response
func CreateDescribeClusterGrafanaResponse() (response *DescribeClusterGrafanaResponse) {
	response = &DescribeClusterGrafanaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}