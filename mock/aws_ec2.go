package mock

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/corysm1th/cloudtables/pkg"
)

// EC2Client mocks the AWS API EC2 endpoint.
type EC2Client struct{ ec2iface.EC2API }

// DescribeInstances returns a mock EC2 instance.
func (e *EC2Client) DescribeInstances(i *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	nameKey, nameValue, id, itype, key, privIP, pubIP, zone := "Name", "or2ccenk8p123", "i-cjdu38ok3jd839ck3", "r5d.24xlarge", "ops-production",
		"10.1.2.100", "192.0.2.100", "us-west-2c"
	place := ec2.Placement{AvailabilityZone: &zone}
	nameTag := ec2.Tag{Key: &nameKey, Value: &nameValue}
	i1 := ec2.Instance{
		InstanceId:       &id, // 17 char
		InstanceType:     &itype,
		KeyName:          &key,
		Tags:             []*ec2.Tag{&nameTag},
		PrivateIpAddress: &privIP,
		PublicIpAddress:  &pubIP,
		Placement:        &place,
	}
	res1 := ec2.Reservation{
		Instances: []*ec2.Instance{&i1},
	}
	r := ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{&res1},
	}
	return &r, nil
}

// DescribeAddresses returns a mock Elastic IP address
func (e *EC2Client) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	domain, inst, pubIP := "devops4.life", "i-cjdu38ok3jd839ck3", "192.0.2.100"
	addr := ec2.Address{Domain: &domain, InstanceId: &inst, PublicIp: &pubIP}
	r := ec2.DescribeAddressesOutput{Addresses: []*ec2.Address{&addr}}
	return &r, nil
}

// CreateEIPs returns an array of pointers to elastic IP objects.
func CreateEIPs() []cloudtables.EC2EIPObj {
	IDs := []string{"i-9pqa3in4f", "i-98vanu34l", "i-cjhnfo9w4e", "i-zxo98fhw"}
	domains := []string{"sql.fastfiats.local", "web.fastfiats.local", "redis.fastfiats.local", "rabbitmq.fastfiats.local"}
	publicIPs := []string{"192.0.2.10", "192.0.2.11", "192.0.2.12", "192.0.2.13"}

	EIPs := []cloudtables.EC2EIPObj{}
	for i, d := range domains {
		eip := cloudtables.EC2EIPObj{
			PublicIP:   publicIPs[i],
			Domain:     d,
			InstanceID: IDs[i],
			Account:    "test",
			Region:     "us-west-2",
		}
		EIPs = append(EIPs, eip)
	}
	return EIPs
}

// CreateEC2Instances returns an array of pointers to EC2 instance objects.
func CreateEC2Instances() []cloudtables.EC2InstObj {
	IDs := []string{"i-9pqa3in4f", "i-98vanu34l", "i-cjhnfo9w4e", "i-zxo98fhw"}
	names := []string{"sql", "web", "redis", "rabbitmq"}
	types := []string{"medium", "large", "small", "micro"}
	privateIPs := []string{"10.1.1.10", "10.1.1.11", "10.1.1.12", "10.1.1.13"}
	publicIPs := []string{"192.0.2.10", "192.0.2.11", "192.0.2.12", "192.0.2.13"}

	instances := []cloudtables.EC2InstObj{}
	for i, n := range names {
		ec2 := cloudtables.EC2InstObj{
			ID:               IDs[i],
			Name:             n,
			Type:             types[i],
			Key:              "ops",
			PrivateIP:        privateIPs[i],
			PublicIP:         publicIPs[i],
			Account:          "test",
			AvailabilityZone: "us-west-2a",
			Region:           "us-west-2",
		}
		instances = append(instances, ec2)
	}
	return instances
}
