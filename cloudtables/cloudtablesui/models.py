from __future__ import unicode_literals

from django.db import models

# Create your models here.

class ec2(models.Model):
	name = models.SlugField(null=True, default="Sure, Not")
	ec2id = models.SlugField(unique=True)
	ec2type = models.SlugField()
	keyname = models.SlugField()
	pubip = models.GenericIPAddressField(null=True)
	privip = models.GenericIPAddressField(null=True)
	awsacct = models.SlugField()
	availzone = models.SlugField()
	ec2region = models.SlugField()

class eip(models.Model):
	publicip = models.GenericIPAddressField(unique=True)
	domain = models.SlugField()
	instanceid = models.SlugField(null=True, default="unassigned")
	awsacct = models.SlugField()
	region = models.SlugField()

class s3(models.Model):
	bucket = models.SlugField(unique=True)
	awsacct = models.SlugField()
