import boto.ec2, boto.s3, ConfigParser, os
from ast import literal_eval

from ctsync import syncec2, synceip, syncs3
from .models import ec2, eip, s3
from django.shortcuts import render, redirect

# Create your views here.
def index(request):
	ec2instances = ec2.objects.all()
	s3buckets = s3.objects.all()
	eips = eip.objects.all()
	return render(request, 'index.html', {'ec2instances':ec2instances, 's3buckets':s3buckets, 'eips':eips})

def sync(request):
	awscreds = ConfigParser.ConfigParser()
	#awsregions = ['us-west-1', 'us-west-2', 'us-east-1', 'ap-northeast-1', 'ap-southeast-2', 'sa-east-1', 'ap-northeast-2', 'us-east-2', 'ap-southeast-1', 'ca-central-1', 'ap-south-1', 'eu-central-1', 'eu-west-1', 'eu-west-2']
	ec2regions = literal_eval(os.environ.get('EC2_REGIONS'))
	awscreds.read(['credentials', os.path.expanduser('~/.aws/credentials')])

	for section in awscreds.sections():
		for region in ec2regions:
			ec2conn = boto.ec2.connect_to_region(region, profile_name = section)
			syncec2(ec2conn, section, region)
			synceip(ec2conn, section, region)
		
		s3conn = boto.s3.connect_to_region(region, profile_name = section)
		syncs3(s3conn, section)

	return redirect('/')

