# CloudTables

Feel free to submit feature requests for the new version.

On the roadmap:

* 80% or better test coverage
* Support for more resource types
* Support for more cloud providers

# CloudTables

Cloud Resource Manager

Provides an easily searchable database of AWS resources across multiple accounts.

## Installation

Start by setting up your AWS credentials file.

```sh
mkdir $HOME/.aws
cat <<EOF > $HOME/.aws/credentials
[aws_acct_1]
aws_access_key_id=50M353C237!D
aws_secret_access_key=50M353C237P@55W02D

[aws_acct_2]
aws_access_key_id=50M353C237!D
aws_secret_access_key=50M353C237P@55W02D
EOF
```

In the square brackets of the credentials file is the name of the AWS account.  This will show up in the "Account" column in CloudTables for every resource in that account, so pick something meaningful.

The front end will be exposed on port 443 of the host, and the inital sync should kick off automatically.

### UI Certificates

TODO

### API Certificates

TODO
