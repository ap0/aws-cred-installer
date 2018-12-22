# aws-cred-installer

Parses the `~/.aws/credentials` file and writes environment variable export statements to the shell.

## Usage

Your AWS `credentials` file will look something like this:

```
[dev]
aws_secret_access_key = XXffdsdfsdf...
aws_access_key_id = ASDASDDDSD

[prod]
aws_secret_access_key = YY&&SDFSDFSFd.s..
aws_access_key_id = ADJJKFDKJFJD
```

Calling `aws-cred-installer prod` will output the following:

```
export AWS_SECRET_ACCESS_KEY=XXffdsdfsdf...
export AWS_ACCESS_KEY_ID=ASDASDDDSD
```

To evaluate this and write this to your environment, you can run:

```
`aws-cred-installer prod`
```

And it will load those environment variables to your shell.