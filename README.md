# AWS Playground

Experimenting with and learning AWS



## AWS Identity and Access Management (IAM)

IAM Identity Center is a web service that provides access control and identity management for AWS resources. Note that in the portal there is both 'IAM Identity Center' and 'IAM'.

Refer: https://docs.aws.amazon.com/sdkref/latest/guide/access-sso.html

IAM Identity Center console: https://console.aws.amazon.com/singlesignon has link to AWS access portal where you can log on with user
created in IAM Identity Center. Can get short lived credentials here.

Doing 'aws configure sso' will result in a '~/.aws/config' file created and a cached token.


From the 'AWS security credentials' page https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html:

- Access key IDs beginning with AKIA are long-term access keys for an IAM user or an AWS account root user.
- Access key IDs beginning with ASIA are temporary credentials access keys that you create using AWS STS operations.
- When you make an API call using temporary security credentials, the call must include a session token,


This was a good page: https://docs.aws.amazon.com/IAM/latest/UserGuide/security-creds.html   For my Go SQS application ended up creating a user in IAM with a long lived token and configuring access on the queue.


https://ben11kehoe.medium.com/never-put-aws-temporary-credentials-in-env-vars-or-credentials-files-theres-a-better-way-25ec45b4d73e



## AWS Cognito

A managed service that provides user authentication, authorization, and user management. Cognito allows organizations to add user sign-up,
sign-in, and access control to web and mobile applications quickly. Cognito provides several authentication options, including social
identity providers, such as Google, Facebook, or Amazon, as well as enterprise identity providers, such as Active Directory or SAML-based
identity providers.

Cognito also provides several features that enable organizations to manage user identities, including user registration, user sign-in, and
password reset. Cognito enables organizations to customize the user experience by providing customizable sign-up and sign-in pages that match
their brand's look and feel.

Cognito integrates with IAM to provide role-based access control. Organizations can use IAM policies to control access to Cognito resources,
such as user pools and identity providers. Cognito also supports fine-grained access control using attribute-based access control (ABAC), where
an organization can define policies that control access to resources based on user attributes, such as location, job role, or group membership.



## Regions

Sydney region: ap-southeast-2
