# CloudTables

Cloud Resource Manager

Provides an easily searchable database of AWS resources across multiple accounts.

Feel free to submit feature requests for the new version.

On the roadmap:

* 80% or better test coverage
* Support for more resource types
* Support for more cloud providers

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

### TLS Configuration

TLS is configured through environment variables.

At a minimum, TLS certificates are required for transport security.

Environment variables, and their default values:

```sh
CERT_FILE="../tls/cert.pem"         # Server TLS certificate
KEY_FILE="../tls/cert-key.pem"      # Server TLS private key
CA_FILE="../tls/ca.pem"             # Server TLS trust chain
MUTUAL_AUTH=false                   # Enable (true) or Disable (default) client authentication
```

**Client Certificate Authentication**

CloudTables ships with support for client certificate authentication, aka mutual authentication.  To enable this feature, set `MUTUAL_AUTH=true` and use client certificates which are signed by the same authority that signed the the TLS server certificate configured with `CERT_FILE`.

For additional account management features, such as rotating certificates or revoking access, it's fairly easy to set up a Hashicorp Vault server, and use that as the certificate authority.  You can use one already deployed in your infrastructure, or a dedicated deployment running in a sidekick container.

**Self Signed Certificates**

Included is a Makefile which will generate self-signed certificates if you want them.  The server certificate expires after 5 years.

```sh
make cert/server
```

The client certificate expires after 6 months, at which point you'll need to generate and distribute a new one to your users.

```sh
make cert/client
```

Install the generated `cloudtables_user.p12` file in your web browser.
