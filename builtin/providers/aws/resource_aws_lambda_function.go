package aws

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/awslabs/aws-sdk-go/aws"
	_ "github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/service/lambda"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsLambdaFunctionCreate,
		Read:   resourceAwsLambdaFunctionRead,
		Update: resourceAwsLambdaFunctionUpdate,
		Delete: resourceAwsLambdaFunctionDelete,

		Schema: map[string]*schema.Schema{
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"function_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"handler": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"memory_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"runtime": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Default:  "nodejs",
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
		},
	}
}

// readZip reads a zipfile and returns the zip data as a base64-encoded byte
// array so you can upload it to the Lambda service.
func readZip(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	// Base64 is 4/3 the original size of data, and we'll padd with +1 in case
	// integer math rounds down. This may be larger than we need but it won't
	// need to be resized.
	b64size := (len(data)/3 + 1) * 4
	target := make([]byte, b64size)
	base64.StdEncoding.Encode(target, data)
	return target, nil
}

// resourceAwsLambdaFunction maps to:
// CreateFunction in the API / SDK
func resourceAwsLambdaFunctionCreate(d *schema.ResourceData, meta interface{}) error {

	svc := lambda.New(nil)

	memory_size := d.Get("memory_size")
	memory_size_int := memory_size.(int)
	memory_size_int64 := int64(memory_size_int)
	memory_size_long := aws.Long(memory_size_int64)

	timeout := d.Get("timeout")
	timeout_int := timeout.(int)
	timeout_int64 := int64(timeout_int)
	timeout_long := aws.Long(timeout_int64)

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: []byte("blah"),
		},
		Description:  aws.String(d.Get("description").(string)),
		FunctionName: aws.String(d.Get("function_name").(string)),
		Handler:      aws.String(d.Get("handler").(string)),
		MemorySize:   memory_size_long,
		Role:         aws.String(d.Get("role").(string)),
		Runtime:      aws.String(d.Get("runtime").(string)),
		Timeout:      timeout_long,
	}

	resp, err := svc.CreateFunction(params)

	functionName := resp.FunctionName
	fmt.Println(functionName)
	fmt.Println(resp)

	// Do something with resp

	return err
}

// resourceAwsLambdaFunctionRead maps to:
// GetFunction in the API / SDK
func resourceAwsLambdaFunctionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceAwsLambdaFunction maps to:
// DeleteFunction in the API / SDK
func resourceAwsLambdaFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// resourceAwsLambdaFunctionUpdate maps to:
// UpdateFunctionCode in the API / SDK
func resourceAwsLambdaFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
