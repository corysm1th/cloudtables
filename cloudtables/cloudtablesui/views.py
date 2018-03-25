import boto.ec2, boto.s3, ConfigParser, os
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
	#awsregions = ['us-west-1', 'us-west-2', 'us-east-1']
	awsregions = ec2.regions()
	awscreds.read(['credentials', os.path.expanduser('~/.aws/credentials')])

	for section in awscreds.sections():
		for region in awsregions:
			ec2conn = boto.ec2.connect_to_region(region, profile_name = section)
			syncec2(ec2conn, section, region)
			synceip(ec2conn, section, region)

		s3conn = boto.s3.connect_to_region(region, profile_name = section)
		syncs3(s3conn, section)

	return redirect('/')

