// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/cloudformation/iam_user_osdCcsAdmin.json
// templates/policies/4.7/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json
// templates/policies/4.7/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json
// templates/policies/4.7/openshift_image_registry_installer_cloud_credentials_policy.json
// templates/policies/4.7/openshift_ingress_operator_cloud_credentials_policy.json
// templates/policies/4.7/openshift_machine_api_aws_cloud_credentials_policy.json
// templates/policies/4.7/operator_iam_role_policy.json
// templates/policies/4.7/sts_installer_permission_policy.json
// templates/policies/4.7/sts_instance_controlplane_permission_policy.json
// templates/policies/4.7/sts_instance_worker_permission_policy.json
// templates/policies/4.7/sts_support_permission_policy.json
// templates/policies/4.8/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json
// templates/policies/4.8/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json
// templates/policies/4.8/openshift_image_registry_installer_cloud_credentials_policy.json
// templates/policies/4.8/openshift_ingress_operator_cloud_credentials_policy.json
// templates/policies/4.8/openshift_machine_api_aws_cloud_credentials_policy.json
// templates/policies/4.8/operator_iam_role_policy.json
// templates/policies/4.8/sts_installer_permission_policy.json
// templates/policies/4.8/sts_instance_controlplane_permission_policy.json
// templates/policies/4.8/sts_instance_worker_permission_policy.json
// templates/policies/4.8/sts_support_permission_policy.json
// templates/policies/4.9/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json
// templates/policies/4.9/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json
// templates/policies/4.9/openshift_image_registry_installer_cloud_credentials_policy.json
// templates/policies/4.9/openshift_ingress_operator_cloud_credentials_policy.json
// templates/policies/4.9/openshift_machine_api_aws_cloud_credentials_policy.json
// templates/policies/4.9/operator_iam_role_policy.json
// templates/policies/4.9/sts_installer_permission_policy.json
// templates/policies/4.9/sts_instance_controlplane_permission_policy.json
// templates/policies/4.9/sts_instance_worker_permission_policy.json
// templates/policies/4.9/sts_support_permission_policy.json
// templates/policies/osd_scp_policy.json
// templates/policies/sts_installer_trust_policy.json
// templates/policies/sts_instance_controlplane_trust_policy.json
// templates/policies/sts_instance_worker_trust_policy.json
// templates/policies/sts_support_trust_policy.json
package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesCloudformationIam_user_osdccsadminJson = []byte(`{
  "Resources": {
    "osdCcsAdmin": {
      "Type": "AWS::IAM::User",
      "Properties": {
        "ManagedPolicyArns": [
          "arn:aws:iam::aws:policy/AdministratorAccess"
        ],
        "UserName": "osdCcsAdmin"
      }
    }
  }
}
`)

func templatesCloudformationIam_user_osdccsadminJsonBytes() ([]byte, error) {
	return _templatesCloudformationIam_user_osdccsadminJson, nil
}

func templatesCloudformationIam_user_osdccsadminJson() (*asset, error) {
	bytes, err := templatesCloudformationIam_user_osdccsadminJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cloudformation/iam_user_osdCcsAdmin.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "iam:GetUser",
        "iam:GetUserPolicy",
        "iam:ListAccessKeys"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, nil
}

func templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteTags",
        "ec2:DeleteVolume",
        "ec2:DescribeInstances",
        "ec2:DescribeSnapshots",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVolumesModifications",
        "ec2:DetachVolume",
        "ec2:ModifyVolume"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, nil
}

func templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:CreateBucket",
        "s3:DeleteBucket",
        "s3:PutBucketTagging",
        "s3:GetBucketTagging",
        "s3:PutBucketPublicAccessBlock",
        "s3:GetBucketPublicAccessBlock",
        "s3:PutEncryptionConfiguration",
        "s3:GetEncryptionConfiguration",
        "s3:PutLifecycleConfiguration",
        "s3:GetLifecycleConfiguration",
        "s3:GetBucketLocation",
        "s3:ListBucket",
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject",
        "s3:ListBucketMultipartUploads",
        "s3:AbortMultipartUpload",
        "s3:ListMultipartUploadParts"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJson, nil
}

func templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/openshift_image_registry_installer_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "elasticloadbalancing:DescribeLoadBalancers",
        "route53:ListHostedZones",
        "route53:ChangeResourceRecordSets",
        "tag:GetResources"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJson, nil
}

func templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/openshift_ingress_operator_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateTags",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeVpcs",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "elasticloadbalancing:DescribeLoadBalancers",
        "elasticloadbalancing:DescribeTargetGroups",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:RegisterTargets",
        "iam:PassRole",
        "iam:CreateServiceLinkedRole"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:Decrypt",
        "kms:Encrypt",
        "kms:GenerateDataKey",
        "kms:GenerateDataKeyWithoutPlainText",
        "kms:DescribeKey"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:RevokeGrant",
        "kms:CreateGrant",
        "kms:ListGrants"
      ],
      "Resource": "*",
      "Condition": {
        "Bool": {
          "kms:GrantIsForAWSResource": true
        }
      }
    }
  ]
}
`)

func templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJson, nil
}

func templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/openshift_machine_api_aws_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Operator_iam_role_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "%{oidc_provider_arn}"
      },
      "Action": [
        "sts:AssumeRoleWithWebIdentity"
      ],
      "Condition": {
        "StringEquals": {
          "%{issuer_url}:aud": "openshift"
        }
      }
    }
  ]
}
`)

func templatesPolicies47Operator_iam_role_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Operator_iam_role_policyJson, nil
}

func templatesPolicies47Operator_iam_role_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Operator_iam_role_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/operator_iam_role_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Sts_installer_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "autoscaling:DescribeAutoScalingGroups",
                "ec2:AllocateAddress",
                "ec2:AssociateAddress",
                "ec2:AssociateDhcpOptions",
                "ec2:AssociateRouteTable",
                "ec2:AttachInternetGateway",
                "ec2:AttachNetworkInterface",
                "ec2:AuthorizeSecurityGroupEgress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CopyImage",
                "ec2:CreateDhcpOptions",
                "ec2:CreateInternetGateway",
                "ec2:CreateNatGateway",
                "ec2:CreateNetworkInterface",
                "ec2:CreateRoute",
                "ec2:CreateRouteTable",
                "ec2:CreateSecurityGroup",
                "ec2:CreateSubnet",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:CreateVpc",
                "ec2:CreateVpcEndpoint",
                "ec2:DeleteDhcpOptions",
                "ec2:DeleteInternetGateway",
                "ec2:DeleteNatGateway",
                "ec2:DeleteNetworkInterface",
                "ec2:DeleteRoute",
                "ec2:DeleteRouteTable",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteSnapshot",
                "ec2:DeleteSubnet",
                "ec2:DeleteTags",
                "ec2:DeleteVolume",
                "ec2:DeleteVpc",
                "ec2:DeleteVpcEndpoints",
                "ec2:DeregisterImage",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceCreditSpecifications",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePrefixLists",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstancesOfferings",
                "ec2:DescribeRouteTables",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeVolumes",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcs",
                "ec2:DetachInternetGateway",
                "ec2:DisassociateRouteTable",
                "ec2:GetEbsDefaultKmsKeyId",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyNetworkInterfaceAttribute",
                "ec2:ModifySubnetAttribute",
                "ec2:ModifyVpcAttribute",
                "ec2:ReleaseAddress",
                "ec2:ReplaceRouteTableAssociation",
                "ec2:RevokeSecurityGroupEgress",
                "ec2:RevokeSecurityGroupIngress",
                "ec2:RunInstances",
                "ec2:StartInstances",
                "ec2:StopInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "iam:AddRoleToInstanceProfile",
                "iam:CreateInstanceProfile",
                "iam:DeleteInstanceProfile",
                "iam:GetInstanceProfile",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:GetUser",
                "iam:ListAttachedRolePolicies",
                "iam:ListInstanceProfiles",
                "iam:ListInstanceProfilesForRole",
                "iam:ListRolePolicies",
                "iam:ListRoles",
                "iam:ListUserPolicies",
                "iam:ListUsers",
                "iam:PassRole",
                "iam:RemoveRoleFromInstanceProfile",
                "iam:SimulatePrincipalPolicy",
                "iam:TagRole",
                "iam:UntagRole",
                "route53:ChangeResourceRecordSets",
                "route53:ChangeTagsForResource",
                "route53:CreateHostedZone",
                "route53:DeleteHostedZone",
                "route53:GetChange",
                "route53:GetHostedZone",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "route53:ListTagsForResource",
                "route53:UpdateHostedZoneComment",
                "s3:CreateBucket",
                "s3:DeleteBucket",
                "s3:DeleteObject",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLocation",
                "s3:GetBucketLogging",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetObject",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:GetObjectVersion",
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
                "s3:ListBucketVersions",
                "s3:PutBucketAcl",
                "s3:PutBucketTagging",
                "s3:PutEncryptionConfiguration",
                "s3:PutObject",
                "s3:PutObjectAcl",
                "s3:PutObjectTagging",
                "sts:AssumeRole",
                "sts:AssumeRoleWithWebIdentity",
                "sts:GetCallerIdentity",
                "tag:GetResources",
                "tag:UntagResources",
                "ec2:CreateVpcEndpointServiceConfiguration",
                "ec2:DeleteVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:ModifyVpcEndpointServicePermissions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies47Sts_installer_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Sts_installer_permission_policyJson, nil
}

func templatesPolicies47Sts_installer_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Sts_installer_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/sts_installer_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Sts_instance_controlplane_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AttachVolume",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteVolume",
                "ec2:Describe*",
                "ec2:DetachVolume",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyVolume",
                "ec2:RevokeSecurityGroupIngress",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerPolicy",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DeleteListener",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteLoadBalancerListeners",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:Describe*",
                "elasticloadbalancing:DetachLoadBalancerFromSubnets",
                "elasticloadbalancing:ModifyListener",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "kms:DescribeKey"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies47Sts_instance_controlplane_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Sts_instance_controlplane_permission_policyJson, nil
}

func templatesPolicies47Sts_instance_controlplane_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Sts_instance_controlplane_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/sts_instance_controlplane_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Sts_instance_worker_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeInstances",
                "ec2:DescribeRegions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies47Sts_instance_worker_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Sts_instance_worker_permission_policyJson, nil
}

func templatesPolicies47Sts_instance_worker_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Sts_instance_worker_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/sts_instance_worker_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies47Sts_support_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudtrail:DescribeTrails",
                "cloudtrail:LookupEvents",
                "cloudwatch:GetMetricData",
                "cloudwatch:GetMetricStatistics",
                "cloudwatch:ListMetrics",
                "ec2:CopySnapshot",
                "ec2:CreateSnapshot",
                "ec2:CreateSnapshots",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAddressesAttribute",
                "ec2:DescribeAggregateIdFormat",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeByoipCidrs",
                "ec2:DescribeCapacityReservations",
                "ec2:DescribeCarrierGateways",
                "ec2:DescribeClassicLinkInstances",
                "ec2:DescribeClientVpnAuthorizationRules",
                "ec2:DescribeClientVpnConnections",
                "ec2:DescribeClientVpnEndpoints",
                "ec2:DescribeClientVpnRoutes",
                "ec2:DescribeClientVpnTargetNetworks",
                "ec2:DescribeCoipPools",
                "ec2:DescribeCustomerGateways",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeEgressOnlyInternetGateways",
                "ec2:DescribeIamInstanceProfileAssociations",
                "ec2:DescribeIdFormat",
                "ec2:DescribeIdentityIdFormat",
                "ec2:DescribeImageAttribute",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypeOfferings",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInstances",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeIpv6Pools",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeLaunchTemplates",
                "ec2:DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
                "ec2:DescribeLocalGatewayRouteTableVpcAssociations",
                "ec2:DescribeLocalGatewayRouteTables",
                "ec2:DescribeLocalGatewayVirtualInterfaceGroups",
                "ec2:DescribeLocalGatewayVirtualInterfaces",
                "ec2:DescribeLocalGateways",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePlacementGroups",
                "ec2:DescribePrefixLists",
                "ec2:DescribePrincipalIdFormat",
                "ec2:DescribePublicIpv4Pools",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstances",
                "ec2:DescribeRouteTables",
                "ec2:DescribeScheduledInstances",
                "ec2:DescribeSecurityGroupReferences",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSnapshotAttribute",
                "ec2:DescribeSnapshots",
                "ec2:DescribeSpotFleetInstances",
                "ec2:DescribeStaleSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeTransitGatewayAttachments",
                "ec2:DescribeTransitGatewayConnectPeers",
                "ec2:DescribeTransitGatewayConnects",
                "ec2:DescribeTransitGatewayMulticastDomains",
                "ec2:DescribeTransitGatewayPeeringAttachments",
                "ec2:DescribeTransitGatewayRouteTables",
                "ec2:DescribeTransitGatewayVpcAttachments",
                "ec2:DescribeTransitGateways",
                "ec2:DescribeVolumeAttribute",
                "ec2:DescribeVolumeStatus",
                "ec2:DescribeVolumes",
                "ec2:DescribeVolumesModifications",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpointConnectionNotifications",
                "ec2:DescribeVpcEndpointConnections",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcPeeringConnections",
                "ec2:DescribeVpcs",
                "ec2:DescribeVpnConnections",
                "ec2:DescribeVpnGateways",
                "ec2:GetAssociatedIpv6PoolCidrs",
                "ec2:GetTransitGatewayAttachmentPropagations",
                "ec2:GetTransitGatewayMulticastDomainAssociations",
                "ec2:GetTransitGatewayPrefixListReferences",
                "ec2:GetTransitGatewayRouteTableAssociations",
                "ec2:GetTransitGatewayRouteTablePropagations",
                "ec2:RebootInstances",
                "ec2:SearchLocalGatewayRoutes",
                "ec2:SearchTransitGatewayMulticastGroups",
                "ec2:SearchTransitGatewayRoutes",
                "ec2:StartInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DescribeAccountLimits",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListenerCertificates",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerPolicies",
                "elasticloadbalancing:DescribeLoadBalancerPolicyTypes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeRules",
                "elasticloadbalancing:DescribeSSLPolicies",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "route53:GetHostedZone",
                "route53:GetHostedZoneCount",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "s3:GetBucketTagging",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:ListAllMyBuckets"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "s3:ListBucket",
            "Resource": [
                "arn:aws:s3:::managed-velero*",
                "arn:aws:s3:::*image-registry*"
            ]
        }
    ]
}
`)

func templatesPolicies47Sts_support_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies47Sts_support_permission_policyJson, nil
}

func templatesPolicies47Sts_support_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies47Sts_support_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.7/sts_support_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "iam:GetUser",
        "iam:GetUserPolicy",
        "iam:ListAccessKeys"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, nil
}

func templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteTags",
        "ec2:DeleteVolume",
        "ec2:DescribeInstances",
        "ec2:DescribeSnapshots",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVolumesModifications",
        "ec2:DetachVolume",
        "ec2:ModifyVolume"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, nil
}

func templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:CreateBucket",
        "s3:DeleteBucket",
        "s3:PutBucketTagging",
        "s3:GetBucketTagging",
        "s3:PutBucketPublicAccessBlock",
        "s3:GetBucketPublicAccessBlock",
        "s3:PutEncryptionConfiguration",
        "s3:GetEncryptionConfiguration",
        "s3:PutLifecycleConfiguration",
        "s3:GetLifecycleConfiguration",
        "s3:GetBucketLocation",
        "s3:ListBucket",
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject",
        "s3:ListBucketMultipartUploads",
        "s3:AbortMultipartUpload",
        "s3:ListMultipartUploadParts"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJson, nil
}

func templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/openshift_image_registry_installer_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "elasticloadbalancing:DescribeLoadBalancers",
        "route53:ListHostedZones",
        "route53:ChangeResourceRecordSets",
        "tag:GetResources"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJson, nil
}

func templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/openshift_ingress_operator_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateTags",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeVpcs",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "elasticloadbalancing:DescribeLoadBalancers",
        "elasticloadbalancing:DescribeTargetGroups",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:RegisterTargets",
        "elasticloadbalancing:DeregisterTargets",
        "iam:PassRole",
        "iam:CreateServiceLinkedRole"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:Decrypt",
        "kms:Encrypt",
        "kms:GenerateDataKey",
        "kms:GenerateDataKeyWithoutPlainText",
        "kms:DescribeKey"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:RevokeGrant",
        "kms:CreateGrant",
        "kms:ListGrants"
      ],
      "Resource": "*",
      "Condition": {
        "Bool": {
          "kms:GrantIsForAWSResource": true
        }
      }
    }
  ]
}
`)

func templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJson, nil
}

func templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/openshift_machine_api_aws_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Operator_iam_role_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "%{oidc_provider_arn}"
      },
      "Action": [
        "sts:AssumeRoleWithWebIdentity"
      ],
      "Condition": {
        "StringEquals": {
          "%{issuer_url}:sub": [ "%{service_accounts}" ]
        }
      }
    }
  ]
}
`)

func templatesPolicies48Operator_iam_role_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Operator_iam_role_policyJson, nil
}

func templatesPolicies48Operator_iam_role_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Operator_iam_role_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/operator_iam_role_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Sts_installer_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "autoscaling:DescribeAutoScalingGroups",
                "ec2:AllocateAddress",
                "ec2:AssociateAddress",
                "ec2:AssociateDhcpOptions",
                "ec2:AssociateRouteTable",
                "ec2:AttachInternetGateway",
                "ec2:AttachNetworkInterface",
                "ec2:AuthorizeSecurityGroupEgress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CopyImage",
                "ec2:CreateDhcpOptions",
                "ec2:CreateInternetGateway",
                "ec2:CreateNatGateway",
                "ec2:CreateNetworkInterface",
                "ec2:CreateRoute",
                "ec2:CreateRouteTable",
                "ec2:CreateSecurityGroup",
                "ec2:CreateSubnet",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:CreateVpc",
                "ec2:CreateVpcEndpoint",
                "ec2:DeleteDhcpOptions",
                "ec2:DeleteInternetGateway",
                "ec2:DeleteNatGateway",
                "ec2:DeleteNetworkInterface",
                "ec2:DeleteRoute",
                "ec2:DeleteRouteTable",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteSnapshot",
                "ec2:DeleteSubnet",
                "ec2:DeleteTags",
                "ec2:DeleteVolume",
                "ec2:DeleteVpc",
                "ec2:DeleteVpcEndpoints",
                "ec2:DeregisterImage",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceCreditSpecifications",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePrefixLists",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstancesOfferings",
                "ec2:DescribeRouteTables",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeVolumes",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcs",
                "ec2:DetachInternetGateway",
                "ec2:DisassociateRouteTable",
                "ec2:GetEbsDefaultKmsKeyId",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyNetworkInterfaceAttribute",
                "ec2:ModifySubnetAttribute",
                "ec2:ModifyVpcAttribute",
                "ec2:ReleaseAddress",
                "ec2:ReplaceRouteTableAssociation",
                "ec2:RevokeSecurityGroupEgress",
                "ec2:RevokeSecurityGroupIngress",
                "ec2:RunInstances",
                "ec2:StartInstances",
                "ec2:StopInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "iam:AddRoleToInstanceProfile",
                "iam:CreateInstanceProfile",
                "iam:DeleteInstanceProfile",
                "iam:GetInstanceProfile",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:GetUser",
                "iam:ListAttachedRolePolicies",
                "iam:ListInstanceProfiles",
                "iam:ListInstanceProfilesForRole",
                "iam:ListRolePolicies",
                "iam:ListRoles",
                "iam:ListUserPolicies",
                "iam:ListUsers",
                "iam:PassRole",
                "iam:RemoveRoleFromInstanceProfile",
                "iam:SimulatePrincipalPolicy",
                "iam:TagRole",
                "iam:UntagRole",
                "route53:ChangeResourceRecordSets",
                "route53:ChangeTagsForResource",
                "route53:CreateHostedZone",
                "route53:DeleteHostedZone",
                "route53:GetChange",
                "route53:GetHostedZone",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "route53:ListTagsForResource",
                "route53:UpdateHostedZoneComment",
                "s3:CreateBucket",
                "s3:DeleteBucket",
                "s3:DeleteObject",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLocation",
                "s3:GetBucketLogging",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetObject",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:GetObjectVersion",
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
                "s3:ListBucketVersions",
                "s3:PutBucketAcl",
                "s3:PutBucketTagging",
                "s3:PutEncryptionConfiguration",
                "s3:PutObject",
                "s3:PutObjectAcl",
                "s3:PutObjectTagging",
                "sts:AssumeRole",
                "sts:AssumeRoleWithWebIdentity",
                "sts:GetCallerIdentity",
                "tag:GetResources",
                "tag:UntagResources",
                "ec2:CreateVpcEndpointServiceConfiguration",
                "ec2:DeleteVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:ModifyVpcEndpointServicePermissions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies48Sts_installer_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Sts_installer_permission_policyJson, nil
}

func templatesPolicies48Sts_installer_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Sts_installer_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/sts_installer_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Sts_instance_controlplane_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AttachVolume",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteVolume",
                "ec2:Describe*",
                "ec2:DetachVolume",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyVolume",
                "ec2:RevokeSecurityGroupIngress",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerPolicy",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DeleteListener",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteLoadBalancerListeners",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:Describe*",
                "elasticloadbalancing:DetachLoadBalancerFromSubnets",
                "elasticloadbalancing:ModifyListener",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "kms:DescribeKey"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies48Sts_instance_controlplane_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Sts_instance_controlplane_permission_policyJson, nil
}

func templatesPolicies48Sts_instance_controlplane_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Sts_instance_controlplane_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/sts_instance_controlplane_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Sts_instance_worker_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeInstances",
                "ec2:DescribeRegions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies48Sts_instance_worker_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Sts_instance_worker_permission_policyJson, nil
}

func templatesPolicies48Sts_instance_worker_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Sts_instance_worker_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/sts_instance_worker_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies48Sts_support_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudtrail:DescribeTrails",
                "cloudtrail:LookupEvents",
                "cloudwatch:GetMetricData",
                "cloudwatch:GetMetricStatistics",
                "cloudwatch:ListMetrics",
                "ec2:CopySnapshot",
                "ec2:CreateSnapshot",
                "ec2:CreateSnapshots",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAddressesAttribute",
                "ec2:DescribeAggregateIdFormat",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeByoipCidrs",
                "ec2:DescribeCapacityReservations",
                "ec2:DescribeCarrierGateways",
                "ec2:DescribeClassicLinkInstances",
                "ec2:DescribeClientVpnAuthorizationRules",
                "ec2:DescribeClientVpnConnections",
                "ec2:DescribeClientVpnEndpoints",
                "ec2:DescribeClientVpnRoutes",
                "ec2:DescribeClientVpnTargetNetworks",
                "ec2:DescribeCoipPools",
                "ec2:DescribeCustomerGateways",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeEgressOnlyInternetGateways",
                "ec2:DescribeIamInstanceProfileAssociations",
                "ec2:DescribeIdFormat",
                "ec2:DescribeIdentityIdFormat",
                "ec2:DescribeImageAttribute",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypeOfferings",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInstances",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeIpv6Pools",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeLaunchTemplates",
                "ec2:DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
                "ec2:DescribeLocalGatewayRouteTableVpcAssociations",
                "ec2:DescribeLocalGatewayRouteTables",
                "ec2:DescribeLocalGatewayVirtualInterfaceGroups",
                "ec2:DescribeLocalGatewayVirtualInterfaces",
                "ec2:DescribeLocalGateways",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePlacementGroups",
                "ec2:DescribePrefixLists",
                "ec2:DescribePrincipalIdFormat",
                "ec2:DescribePublicIpv4Pools",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstances",
                "ec2:DescribeRouteTables",
                "ec2:DescribeScheduledInstances",
                "ec2:DescribeSecurityGroupReferences",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSnapshotAttribute",
                "ec2:DescribeSnapshots",
                "ec2:DescribeSpotFleetInstances",
                "ec2:DescribeStaleSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeTransitGatewayAttachments",
                "ec2:DescribeTransitGatewayConnectPeers",
                "ec2:DescribeTransitGatewayConnects",
                "ec2:DescribeTransitGatewayMulticastDomains",
                "ec2:DescribeTransitGatewayPeeringAttachments",
                "ec2:DescribeTransitGatewayRouteTables",
                "ec2:DescribeTransitGatewayVpcAttachments",
                "ec2:DescribeTransitGateways",
                "ec2:DescribeVolumeAttribute",
                "ec2:DescribeVolumeStatus",
                "ec2:DescribeVolumes",
                "ec2:DescribeVolumesModifications",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpointConnectionNotifications",
                "ec2:DescribeVpcEndpointConnections",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcPeeringConnections",
                "ec2:DescribeVpcs",
                "ec2:DescribeVpnConnections",
                "ec2:DescribeVpnGateways",
                "ec2:GetAssociatedIpv6PoolCidrs",
                "ec2:GetTransitGatewayAttachmentPropagations",
                "ec2:GetTransitGatewayMulticastDomainAssociations",
                "ec2:GetTransitGatewayPrefixListReferences",
                "ec2:GetTransitGatewayRouteTableAssociations",
                "ec2:GetTransitGatewayRouteTablePropagations",
                "ec2:RebootInstances",
                "ec2:SearchLocalGatewayRoutes",
                "ec2:SearchTransitGatewayMulticastGroups",
                "ec2:SearchTransitGatewayRoutes",
                "ec2:StartInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DescribeAccountLimits",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListenerCertificates",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerPolicies",
                "elasticloadbalancing:DescribeLoadBalancerPolicyTypes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeRules",
                "elasticloadbalancing:DescribeSSLPolicies",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "route53:GetHostedZone",
                "route53:GetHostedZoneCount",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "s3:GetBucketTagging",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:ListAllMyBuckets"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "s3:ListBucket",
            "Resource": [
                "arn:aws:s3:::managed-velero*",
                "arn:aws:s3:::*image-registry*"
            ]
        }
    ]
}
`)

func templatesPolicies48Sts_support_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies48Sts_support_permission_policyJson, nil
}

func templatesPolicies48Sts_support_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies48Sts_support_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.8/sts_support_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "iam:GetUser",
        "iam:GetUserPolicy",
        "iam:ListAccessKeys"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, nil
}

func templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:AttachVolume",
        "ec2:CreateSnapshot",
        "ec2:CreateTags",
        "ec2:CreateVolume",
        "ec2:DeleteSnapshot",
        "ec2:DeleteTags",
        "ec2:DeleteVolume",
        "ec2:DescribeInstances",
        "ec2:DescribeSnapshots",
        "ec2:DescribeTags",
        "ec2:DescribeVolumes",
        "ec2:DescribeVolumesModifications",
        "ec2:DetachVolume",
        "ec2:ModifyVolume"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, nil
}

func templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:CreateBucket",
        "s3:DeleteBucket",
        "s3:PutBucketTagging",
        "s3:GetBucketTagging",
        "s3:PutBucketPublicAccessBlock",
        "s3:GetBucketPublicAccessBlock",
        "s3:PutEncryptionConfiguration",
        "s3:GetEncryptionConfiguration",
        "s3:PutLifecycleConfiguration",
        "s3:GetLifecycleConfiguration",
        "s3:GetBucketLocation",
        "s3:ListBucket",
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject",
        "s3:ListBucketMultipartUploads",
        "s3:AbortMultipartUpload",
        "s3:ListMultipartUploadParts"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJson, nil
}

func templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/openshift_image_registry_installer_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "elasticloadbalancing:DescribeLoadBalancers",
        "route53:ListHostedZones",
        "route53:ChangeResourceRecordSets",
        "tag:GetResources"
      ],
      "Resource": "*"
    }
  ]
}
`)

func templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJson, nil
}

func templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/openshift_ingress_operator_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ec2:CreateTags",
        "ec2:DescribeAvailabilityZones",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeImages",
        "ec2:DescribeInstances",
        "ec2:DescribeInternetGateways",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeVpcs",
        "ec2:RunInstances",
        "ec2:TerminateInstances",
        "elasticloadbalancing:DescribeLoadBalancers",
        "elasticloadbalancing:DescribeTargetGroups",
        "elasticloadbalancing:DescribeTargetHealth",
        "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
        "elasticloadbalancing:RegisterTargets",
        "elasticloadbalancing:DeregisterTargets",
        "iam:PassRole",
        "iam:CreateServiceLinkedRole"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:Decrypt",
        "kms:Encrypt",
        "kms:GenerateDataKey",
        "kms:GenerateDataKeyWithoutPlainText",
        "kms:DescribeKey"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "kms:RevokeGrant",
        "kms:CreateGrant",
        "kms:ListGrants"
      ],
      "Resource": "*",
      "Condition": {
        "Bool": {
          "kms:GrantIsForAWSResource": true
        }
      }
    }
  ]
}
`)

func templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJson, nil
}

func templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/openshift_machine_api_aws_cloud_credentials_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Operator_iam_role_policyJson = []byte(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "%{oidc_provider_arn}"
      },
      "Action": [
        "sts:AssumeRoleWithWebIdentity"
      ],
      "Condition": {
        "StringEquals": {
          "%{issuer_url}:sub": [ "%{service_accounts}" ]
        }
      }
    }
  ]
}
`)

func templatesPolicies49Operator_iam_role_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Operator_iam_role_policyJson, nil
}

func templatesPolicies49Operator_iam_role_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Operator_iam_role_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/operator_iam_role_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Sts_installer_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "autoscaling:DescribeAutoScalingGroups",
                "ec2:AllocateAddress",
                "ec2:AssociateAddress",
                "ec2:AssociateDhcpOptions",
                "ec2:AssociateRouteTable",
                "ec2:AttachInternetGateway",
                "ec2:AttachNetworkInterface",
                "ec2:AuthorizeSecurityGroupEgress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CopyImage",
                "ec2:CreateDhcpOptions",
                "ec2:CreateInternetGateway",
                "ec2:CreateNatGateway",
                "ec2:CreateNetworkInterface",
                "ec2:CreateRoute",
                "ec2:CreateRouteTable",
                "ec2:CreateSecurityGroup",
                "ec2:CreateSubnet",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:CreateVpc",
                "ec2:CreateVpcEndpoint",
                "ec2:DeleteDhcpOptions",
                "ec2:DeleteInternetGateway",
                "ec2:DeleteNatGateway",
                "ec2:DeleteNetworkInterface",
                "ec2:DeleteRoute",
                "ec2:DeleteRouteTable",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteSnapshot",
                "ec2:DeleteSubnet",
                "ec2:DeleteTags",
                "ec2:DeleteVolume",
                "ec2:DeleteVpc",
                "ec2:DeleteVpcEndpoints",
                "ec2:DeregisterImage",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceCreditSpecifications",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePrefixLists",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstancesOfferings",
                "ec2:DescribeRouteTables",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeVolumes",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcs",
                "ec2:DetachInternetGateway",
                "ec2:DisassociateRouteTable",
                "ec2:GetEbsDefaultKmsKeyId",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyNetworkInterfaceAttribute",
                "ec2:ModifySubnetAttribute",
                "ec2:ModifyVpcAttribute",
                "ec2:ReleaseAddress",
                "ec2:ReplaceRouteTableAssociation",
                "ec2:RevokeSecurityGroupEgress",
                "ec2:RevokeSecurityGroupIngress",
                "ec2:RunInstances",
                "ec2:StartInstances",
                "ec2:StopInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "iam:AddRoleToInstanceProfile",
                "iam:CreateInstanceProfile",
                "iam:DeleteInstanceProfile",
                "iam:GetInstanceProfile",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:GetUser",
                "iam:ListAttachedRolePolicies",
                "iam:ListInstanceProfiles",
                "iam:ListInstanceProfilesForRole",
                "iam:ListRolePolicies",
                "iam:ListRoles",
                "iam:ListUserPolicies",
                "iam:ListUsers",
                "iam:PassRole",
                "iam:RemoveRoleFromInstanceProfile",
                "iam:SimulatePrincipalPolicy",
                "iam:TagRole",
                "iam:UntagRole",
                "route53:ChangeResourceRecordSets",
                "route53:ChangeTagsForResource",
                "route53:CreateHostedZone",
                "route53:DeleteHostedZone",
                "route53:GetChange",
                "route53:GetHostedZone",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "route53:ListTagsForResource",
                "route53:UpdateHostedZoneComment",
                "s3:CreateBucket",
                "s3:DeleteBucket",
                "s3:DeleteObject",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLocation",
                "s3:GetBucketLogging",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetObject",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:GetObjectVersion",
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
                "s3:ListBucketVersions",
                "s3:PutBucketAcl",
                "s3:PutBucketTagging",
                "s3:PutEncryptionConfiguration",
                "s3:PutObject",
                "s3:PutObjectAcl",
                "s3:PutObjectTagging",
                "sts:AssumeRole",
                "sts:AssumeRoleWithWebIdentity",
                "sts:GetCallerIdentity",
                "tag:GetResources",
                "tag:UntagResources",
                "ec2:CreateVpcEndpointServiceConfiguration",
                "ec2:DeleteVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:ModifyVpcEndpointServicePermissions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies49Sts_installer_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Sts_installer_permission_policyJson, nil
}

func templatesPolicies49Sts_installer_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Sts_installer_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/sts_installer_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Sts_instance_controlplane_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:AttachVolume",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:CreateVolume",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteVolume",
                "ec2:Describe*",
                "ec2:DetachVolume",
                "ec2:ModifyInstanceAttribute",
                "ec2:ModifyVolume",
                "ec2:RevokeSecurityGroupIngress",
                "elasticloadbalancing:AddTags",
                "elasticloadbalancing:AttachLoadBalancerToSubnets",
                "elasticloadbalancing:ApplySecurityGroupsToLoadBalancer",
                "elasticloadbalancing:CreateListener",
                "elasticloadbalancing:CreateLoadBalancer",
                "elasticloadbalancing:CreateLoadBalancerPolicy",
                "elasticloadbalancing:CreateLoadBalancerListeners",
                "elasticloadbalancing:CreateTargetGroup",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DeleteListener",
                "elasticloadbalancing:DeleteLoadBalancer",
                "elasticloadbalancing:DeleteLoadBalancerListeners",
                "elasticloadbalancing:DeleteTargetGroup",
                "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
                "elasticloadbalancing:DeregisterTargets",
                "elasticloadbalancing:Describe*",
                "elasticloadbalancing:DetachLoadBalancerFromSubnets",
                "elasticloadbalancing:ModifyListener",
                "elasticloadbalancing:ModifyLoadBalancerAttributes",
                "elasticloadbalancing:ModifyTargetGroup",
                "elasticloadbalancing:ModifyTargetGroupAttributes",
                "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
                "elasticloadbalancing:RegisterTargets",
                "elasticloadbalancing:SetLoadBalancerPoliciesForBackendServer",
                "elasticloadbalancing:SetLoadBalancerPoliciesOfListener",
                "kms:DescribeKey"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies49Sts_instance_controlplane_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Sts_instance_controlplane_permission_policyJson, nil
}

func templatesPolicies49Sts_instance_controlplane_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Sts_instance_controlplane_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/sts_instance_controlplane_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Sts_instance_worker_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:DescribeInstances",
                "ec2:DescribeRegions"
            ],
            "Resource": "*"
        }
    ]
}
`)

func templatesPolicies49Sts_instance_worker_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Sts_instance_worker_permission_policyJson, nil
}

func templatesPolicies49Sts_instance_worker_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Sts_instance_worker_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/sts_instance_worker_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPolicies49Sts_support_permission_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudtrail:DescribeTrails",
                "cloudtrail:LookupEvents",
                "cloudwatch:GetMetricData",
                "cloudwatch:GetMetricStatistics",
                "cloudwatch:ListMetrics",
                "ec2:CopySnapshot",
                "ec2:CreateSnapshot",
                "ec2:CreateSnapshots",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeAddresses",
                "ec2:DescribeAddressesAttribute",
                "ec2:DescribeAggregateIdFormat",
                "ec2:DescribeAvailabilityZones",
                "ec2:DescribeByoipCidrs",
                "ec2:DescribeCapacityReservations",
                "ec2:DescribeCarrierGateways",
                "ec2:DescribeClassicLinkInstances",
                "ec2:DescribeClientVpnAuthorizationRules",
                "ec2:DescribeClientVpnConnections",
                "ec2:DescribeClientVpnEndpoints",
                "ec2:DescribeClientVpnRoutes",
                "ec2:DescribeClientVpnTargetNetworks",
                "ec2:DescribeCoipPools",
                "ec2:DescribeCustomerGateways",
                "ec2:DescribeDhcpOptions",
                "ec2:DescribeEgressOnlyInternetGateways",
                "ec2:DescribeIamInstanceProfileAssociations",
                "ec2:DescribeIdFormat",
                "ec2:DescribeIdentityIdFormat",
                "ec2:DescribeImageAttribute",
                "ec2:DescribeImages",
                "ec2:DescribeInstanceAttribute",
                "ec2:DescribeInstanceStatus",
                "ec2:DescribeInstanceTypeOfferings",
                "ec2:DescribeInstanceTypes",
                "ec2:DescribeInstances",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeIpv6Pools",
                "ec2:DescribeKeyPairs",
                "ec2:DescribeLaunchTemplates",
                "ec2:DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations",
                "ec2:DescribeLocalGatewayRouteTableVpcAssociations",
                "ec2:DescribeLocalGatewayRouteTables",
                "ec2:DescribeLocalGatewayVirtualInterfaceGroups",
                "ec2:DescribeLocalGatewayVirtualInterfaces",
                "ec2:DescribeLocalGateways",
                "ec2:DescribeNatGateways",
                "ec2:DescribeNetworkAcls",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribePlacementGroups",
                "ec2:DescribePrefixLists",
                "ec2:DescribePrincipalIdFormat",
                "ec2:DescribePublicIpv4Pools",
                "ec2:DescribeRegions",
                "ec2:DescribeReservedInstances",
                "ec2:DescribeRouteTables",
                "ec2:DescribeScheduledInstances",
                "ec2:DescribeSecurityGroupReferences",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSnapshotAttribute",
                "ec2:DescribeSnapshots",
                "ec2:DescribeSpotFleetInstances",
                "ec2:DescribeStaleSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeTags",
                "ec2:DescribeTransitGatewayAttachments",
                "ec2:DescribeTransitGatewayConnectPeers",
                "ec2:DescribeTransitGatewayConnects",
                "ec2:DescribeTransitGatewayMulticastDomains",
                "ec2:DescribeTransitGatewayPeeringAttachments",
                "ec2:DescribeTransitGatewayRouteTables",
                "ec2:DescribeTransitGatewayVpcAttachments",
                "ec2:DescribeTransitGateways",
                "ec2:DescribeVolumeAttribute",
                "ec2:DescribeVolumeStatus",
                "ec2:DescribeVolumes",
                "ec2:DescribeVolumesModifications",
                "ec2:DescribeVpcAttribute",
                "ec2:DescribeVpcClassicLink",
                "ec2:DescribeVpcClassicLinkDnsSupport",
                "ec2:DescribeVpcEndpointConnectionNotifications",
                "ec2:DescribeVpcEndpointConnections",
                "ec2:DescribeVpcEndpointServiceConfigurations",
                "ec2:DescribeVpcEndpointServicePermissions",
                "ec2:DescribeVpcEndpointServices",
                "ec2:DescribeVpcEndpoints",
                "ec2:DescribeVpcPeeringConnections",
                "ec2:DescribeVpcs",
                "ec2:DescribeVpnConnections",
                "ec2:DescribeVpnGateways",
                "ec2:GetAssociatedIpv6PoolCidrs",
                "ec2:GetTransitGatewayAttachmentPropagations",
                "ec2:GetTransitGatewayMulticastDomainAssociations",
                "ec2:GetTransitGatewayPrefixListReferences",
                "ec2:GetTransitGatewayRouteTableAssociations",
                "ec2:GetTransitGatewayRouteTablePropagations",
                "ec2:RebootInstances",
                "ec2:SearchLocalGatewayRoutes",
                "ec2:SearchTransitGatewayMulticastGroups",
                "ec2:SearchTransitGatewayRoutes",
                "ec2:StartInstances",
                "ec2:TerminateInstances",
                "elasticloadbalancing:ConfigureHealthCheck",
                "elasticloadbalancing:DescribeAccountLimits",
                "elasticloadbalancing:DescribeInstanceHealth",
                "elasticloadbalancing:DescribeListenerCertificates",
                "elasticloadbalancing:DescribeListeners",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerAttributes",
                "elasticloadbalancing:DescribeLoadBalancerPolicies",
                "elasticloadbalancing:DescribeLoadBalancerPolicyTypes",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeRules",
                "elasticloadbalancing:DescribeSSLPolicies",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTags",
                "elasticloadbalancing:DescribeTargetGroupAttributes",
                "elasticloadbalancing:DescribeTargetGroups",
                "elasticloadbalancing:DescribeTargetHealth",
                "route53:GetHostedZone",
                "route53:GetHostedZoneCount",
                "route53:ListHostedZones",
                "route53:ListHostedZonesByName",
                "route53:ListResourceRecordSets",
                "s3:GetBucketTagging",
                "s3:GetObjectAcl",
                "s3:GetObjectTagging",
                "s3:ListAllMyBuckets"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": "s3:ListBucket",
            "Resource": [
                "arn:aws:s3:::managed-velero*",
                "arn:aws:s3:::*image-registry*"
            ]
        }
    ]
}
`)

func templatesPolicies49Sts_support_permission_policyJsonBytes() ([]byte, error) {
	return _templatesPolicies49Sts_support_permission_policyJson, nil
}

func templatesPolicies49Sts_support_permission_policyJson() (*asset, error) {
	bytes, err := templatesPolicies49Sts_support_permission_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/4.9/sts_support_permission_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPoliciesOsd_scp_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Id": "OSD SCP Policy Document",
    "Statement": [
        {
            "Sid": "AllowEC2",
            "Effect": "Allow",
            "Action": [
                "ec2:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowAutoscaling",
            "Effect": "Allow",
            "Action": [
                "autoscaling:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowS3",
            "Effect": "Allow",
            "Action": [
                "s3:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowIAM",
            "Effect": "Allow",
            "Action": [
                "iam:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowELB",
            "Effect": "Allow",
            "Action": [
                "elasticloadbalancing:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowCloudWatch",
            "Effect": "Allow",
            "Action": [
                "cloudwatch:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowCloudWatchEvents",
            "Effect": "Allow",
            "Action": [
                "events:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowCloudWatchLogs",
            "Effect": "Allow",
            "Action": [
                "logs:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowSupport",
            "Effect": "Allow",
            "Action": [
                "support:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowKMS",
            "Effect": "Allow",
            "Action": [
                "kms:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowSTS",
            "Effect": "Allow",
            "Action": [
                "sts:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowTag",
            "Effect": "Allow",
            "Action": [
                "tag:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowRoute53",
            "Effect": "Allow",
            "Action": [
                "route53:*"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "AllowServiceQuotas",
            "Effect": "Allow",
            "Action": [
                "servicequotas:ListServices",
                "servicequotas:GetRequestedServiceQuotaChange",
                "servicequotas:GetServiceQuota",
                "servicequotas:RequestServiceQuotaIncrease",
                "servicequotas:ListServiceQuotas"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
}
`)

func templatesPoliciesOsd_scp_policyJsonBytes() ([]byte, error) {
	return _templatesPoliciesOsd_scp_policyJson, nil
}

func templatesPoliciesOsd_scp_policyJson() (*asset, error) {
	bytes, err := templatesPoliciesOsd_scp_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/osd_scp_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPoliciesSts_installer_trust_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": [
                    "arn:aws:iam::%{aws_account_id}:role/RH-Managed-OpenShift-Installer"
                ]
            },
            "Action": [
                "sts:AssumeRole"
            ]
        }
    ]
}
`)

func templatesPoliciesSts_installer_trust_policyJsonBytes() ([]byte, error) {
	return _templatesPoliciesSts_installer_trust_policyJson, nil
}

func templatesPoliciesSts_installer_trust_policyJson() (*asset, error) {
	bytes, err := templatesPoliciesSts_installer_trust_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/sts_installer_trust_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPoliciesSts_instance_controlplane_trust_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": [
                    "ec2.amazonaws.com"
                ]
            },
            "Action": [
                "sts:AssumeRole"
            ]
        }
    ]
}
`)

func templatesPoliciesSts_instance_controlplane_trust_policyJsonBytes() ([]byte, error) {
	return _templatesPoliciesSts_instance_controlplane_trust_policyJson, nil
}

func templatesPoliciesSts_instance_controlplane_trust_policyJson() (*asset, error) {
	bytes, err := templatesPoliciesSts_instance_controlplane_trust_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/sts_instance_controlplane_trust_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPoliciesSts_instance_worker_trust_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Service": [
                    "ec2.amazonaws.com"
                ]
            },
            "Action": [
                "sts:AssumeRole"
            ]
        }
    ]
}
`)

func templatesPoliciesSts_instance_worker_trust_policyJsonBytes() ([]byte, error) {
	return _templatesPoliciesSts_instance_worker_trust_policyJson, nil
}

func templatesPoliciesSts_instance_worker_trust_policyJson() (*asset, error) {
	bytes, err := templatesPoliciesSts_instance_worker_trust_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/sts_instance_worker_trust_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPoliciesSts_support_trust_policyJson = []byte(`{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": [
                    "arn:aws:iam::%{aws_account_id}:role/RH-Technical-Support-Access"
                ]
            },
            "Action": [
                "sts:AssumeRole"
            ]
        }
    ]
}
`)

func templatesPoliciesSts_support_trust_policyJsonBytes() ([]byte, error) {
	return _templatesPoliciesSts_support_trust_policyJson, nil
}

func templatesPoliciesSts_support_trust_policyJson() (*asset, error) {
	bytes, err := templatesPoliciesSts_support_trust_policyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/policies/sts_support_trust_policy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/cloudformation/iam_user_osdCcsAdmin.json":                                                            templatesCloudformationIam_user_osdccsadminJson,
	"templates/policies/4.7/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson,
	"templates/policies/4.7/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson,
	"templates/policies/4.7/openshift_image_registry_installer_cloud_credentials_policy.json":                       templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJson,
	"templates/policies/4.7/openshift_ingress_operator_cloud_credentials_policy.json":                               templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJson,
	"templates/policies/4.7/openshift_machine_api_aws_cloud_credentials_policy.json":                                templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJson,
	"templates/policies/4.7/operator_iam_role_policy.json":                                                          templatesPolicies47Operator_iam_role_policyJson,
	"templates/policies/4.7/sts_installer_permission_policy.json":                                                   templatesPolicies47Sts_installer_permission_policyJson,
	"templates/policies/4.7/sts_instance_controlplane_permission_policy.json":                                       templatesPolicies47Sts_instance_controlplane_permission_policyJson,
	"templates/policies/4.7/sts_instance_worker_permission_policy.json":                                             templatesPolicies47Sts_instance_worker_permission_policyJson,
	"templates/policies/4.7/sts_support_permission_policy.json":                                                     templatesPolicies47Sts_support_permission_policyJson,
	"templates/policies/4.8/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson,
	"templates/policies/4.8/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson,
	"templates/policies/4.8/openshift_image_registry_installer_cloud_credentials_policy.json":                       templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJson,
	"templates/policies/4.8/openshift_ingress_operator_cloud_credentials_policy.json":                               templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJson,
	"templates/policies/4.8/openshift_machine_api_aws_cloud_credentials_policy.json":                                templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJson,
	"templates/policies/4.8/operator_iam_role_policy.json":                                                          templatesPolicies48Operator_iam_role_policyJson,
	"templates/policies/4.8/sts_installer_permission_policy.json":                                                   templatesPolicies48Sts_installer_permission_policyJson,
	"templates/policies/4.8/sts_instance_controlplane_permission_policy.json":                                       templatesPolicies48Sts_instance_controlplane_permission_policyJson,
	"templates/policies/4.8/sts_instance_worker_permission_policy.json":                                             templatesPolicies48Sts_instance_worker_permission_policyJson,
	"templates/policies/4.8/sts_support_permission_policy.json":                                                     templatesPolicies48Sts_support_permission_policyJson,
	"templates/policies/4.9/openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson,
	"templates/policies/4.9/openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson,
	"templates/policies/4.9/openshift_image_registry_installer_cloud_credentials_policy.json":                       templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJson,
	"templates/policies/4.9/openshift_ingress_operator_cloud_credentials_policy.json":                               templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJson,
	"templates/policies/4.9/openshift_machine_api_aws_cloud_credentials_policy.json":                                templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJson,
	"templates/policies/4.9/operator_iam_role_policy.json":                                                          templatesPolicies49Operator_iam_role_policyJson,
	"templates/policies/4.9/sts_installer_permission_policy.json":                                                   templatesPolicies49Sts_installer_permission_policyJson,
	"templates/policies/4.9/sts_instance_controlplane_permission_policy.json":                                       templatesPolicies49Sts_instance_controlplane_permission_policyJson,
	"templates/policies/4.9/sts_instance_worker_permission_policy.json":                                             templatesPolicies49Sts_instance_worker_permission_policyJson,
	"templates/policies/4.9/sts_support_permission_policy.json":                                                     templatesPolicies49Sts_support_permission_policyJson,
	"templates/policies/osd_scp_policy.json":                                                                        templatesPoliciesOsd_scp_policyJson,
	"templates/policies/sts_installer_trust_policy.json":                                                            templatesPoliciesSts_installer_trust_policyJson,
	"templates/policies/sts_instance_controlplane_trust_policy.json":                                                templatesPoliciesSts_instance_controlplane_trust_policyJson,
	"templates/policies/sts_instance_worker_trust_policy.json":                                                      templatesPoliciesSts_instance_worker_trust_policyJson,
	"templates/policies/sts_support_trust_policy.json":                                                              templatesPoliciesSts_support_trust_policyJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"cloudformation": &bintree{nil, map[string]*bintree{
			"iam_user_osdCcsAdmin.json": &bintree{templatesCloudformationIam_user_osdccsadminJson, map[string]*bintree{}},
		}},
		"policies": &bintree{nil, map[string]*bintree{
			"4.7": &bintree{nil, map[string]*bintree{
				"openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": &bintree{templatesPolicies47Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, map[string]*bintree{}},
				"openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        &bintree{templatesPolicies47Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_image_registry_installer_cloud_credentials_policy.json":                       &bintree{templatesPolicies47Openshift_image_registry_installer_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_ingress_operator_cloud_credentials_policy.json":                               &bintree{templatesPolicies47Openshift_ingress_operator_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_machine_api_aws_cloud_credentials_policy.json":                                &bintree{templatesPolicies47Openshift_machine_api_aws_cloud_credentials_policyJson, map[string]*bintree{}},
				"operator_iam_role_policy.json":                                                          &bintree{templatesPolicies47Operator_iam_role_policyJson, map[string]*bintree{}},
				"sts_installer_permission_policy.json":                                                   &bintree{templatesPolicies47Sts_installer_permission_policyJson, map[string]*bintree{}},
				"sts_instance_controlplane_permission_policy.json":                                       &bintree{templatesPolicies47Sts_instance_controlplane_permission_policyJson, map[string]*bintree{}},
				"sts_instance_worker_permission_policy.json":                                             &bintree{templatesPolicies47Sts_instance_worker_permission_policyJson, map[string]*bintree{}},
				"sts_support_permission_policy.json":                                                     &bintree{templatesPolicies47Sts_support_permission_policyJson, map[string]*bintree{}},
			}},
			"4.8": &bintree{nil, map[string]*bintree{
				"openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": &bintree{templatesPolicies48Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, map[string]*bintree{}},
				"openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        &bintree{templatesPolicies48Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_image_registry_installer_cloud_credentials_policy.json":                       &bintree{templatesPolicies48Openshift_image_registry_installer_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_ingress_operator_cloud_credentials_policy.json":                               &bintree{templatesPolicies48Openshift_ingress_operator_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_machine_api_aws_cloud_credentials_policy.json":                                &bintree{templatesPolicies48Openshift_machine_api_aws_cloud_credentials_policyJson, map[string]*bintree{}},
				"operator_iam_role_policy.json":                                                          &bintree{templatesPolicies48Operator_iam_role_policyJson, map[string]*bintree{}},
				"sts_installer_permission_policy.json":                                                   &bintree{templatesPolicies48Sts_installer_permission_policyJson, map[string]*bintree{}},
				"sts_instance_controlplane_permission_policy.json":                                       &bintree{templatesPolicies48Sts_instance_controlplane_permission_policyJson, map[string]*bintree{}},
				"sts_instance_worker_permission_policy.json":                                             &bintree{templatesPolicies48Sts_instance_worker_permission_policyJson, map[string]*bintree{}},
				"sts_support_permission_policy.json":                                                     &bintree{templatesPolicies48Sts_support_permission_policyJson, map[string]*bintree{}},
			}},
			"4.9": &bintree{nil, map[string]*bintree{
				"openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policy.json": &bintree{templatesPolicies49Openshift_cloud_credential_operator_cloud_credential_operator_iam_ro_creds_policyJson, map[string]*bintree{}},
				"openshift_cluster_csi_drivers_ebs_cloud_credentials_policy.json":                        &bintree{templatesPolicies49Openshift_cluster_csi_drivers_ebs_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_image_registry_installer_cloud_credentials_policy.json":                       &bintree{templatesPolicies49Openshift_image_registry_installer_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_ingress_operator_cloud_credentials_policy.json":                               &bintree{templatesPolicies49Openshift_ingress_operator_cloud_credentials_policyJson, map[string]*bintree{}},
				"openshift_machine_api_aws_cloud_credentials_policy.json":                                &bintree{templatesPolicies49Openshift_machine_api_aws_cloud_credentials_policyJson, map[string]*bintree{}},
				"operator_iam_role_policy.json":                                                          &bintree{templatesPolicies49Operator_iam_role_policyJson, map[string]*bintree{}},
				"sts_installer_permission_policy.json":                                                   &bintree{templatesPolicies49Sts_installer_permission_policyJson, map[string]*bintree{}},
				"sts_instance_controlplane_permission_policy.json":                                       &bintree{templatesPolicies49Sts_instance_controlplane_permission_policyJson, map[string]*bintree{}},
				"sts_instance_worker_permission_policy.json":                                             &bintree{templatesPolicies49Sts_instance_worker_permission_policyJson, map[string]*bintree{}},
				"sts_support_permission_policy.json":                                                     &bintree{templatesPolicies49Sts_support_permission_policyJson, map[string]*bintree{}},
			}},
			"osd_scp_policy.json":                         &bintree{templatesPoliciesOsd_scp_policyJson, map[string]*bintree{}},
			"sts_installer_trust_policy.json":             &bintree{templatesPoliciesSts_installer_trust_policyJson, map[string]*bintree{}},
			"sts_instance_controlplane_trust_policy.json": &bintree{templatesPoliciesSts_instance_controlplane_trust_policyJson, map[string]*bintree{}},
			"sts_instance_worker_trust_policy.json":       &bintree{templatesPoliciesSts_instance_worker_trust_policyJson, map[string]*bintree{}},
			"sts_support_trust_policy.json":               &bintree{templatesPoliciesSts_support_trust_policyJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
