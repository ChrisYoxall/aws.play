# AWS Playground

Experimenting with and learning AWS



## Identity and Access Management

Using IAM Identity Center authentication seems to be recommended way of providing AWS credentials.

Refer: https://docs.aws.amazon.com/sdkref/latest/guide/access-sso.html

IAM Identity Center console: https://console.aws.amazon.com/singlesignon has link to AWS access portal where you can log on with user
created in IAM Identity Center. Can get short lived credentials here.

Doing 'aws configure sso' will result in a '~/.aws/config' file created and a cached token.


From the 'AWS security credentials' page https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html:

- Access key IDs beginning with AKIA are long-term access keys for an IAM user or an AWS account root user.
- Access key IDs beginning with ASIA are temporary credentials access keys that you create using AWS STS operations.
- When you make an API call using temporary security credentials, the call must include a session token,


This was a good page: https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html    For my Go SQS application ended up
creating a user and a long lived token.


https://ben11kehoe.medium.com/never-put-aws-temporary-credentials-in-env-vars-or-credentials-files-theres-a-better-way-25ec45b4d73e





## Regions

Sydney region: ap-southeast-2
