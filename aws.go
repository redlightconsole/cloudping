package cloudping

const AWSProvider = "aws"

func init() {
	AddTarget(
		NewRegionTarget(AWSProvider, "Africa (Cape Town)", "af-south-1", "dynamodb.af-south-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Hong Kong)", "ap-east-1", "dynamodb.ap-east-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Tokyo)", "ap-northeast-1", "dynamodb.ap-northeast-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Seoul)", "ap-northeast-2", "dynamodb.ap-northeast-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Osaka)", "ap-northeast-3", "dynamodb.ap-northeast-3.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Mumbai)", "ap-south-1", "dynamodb.ap-south-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Hyderabad)", "ap-south-2", "dynamodb.ap-south-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Singapore)", "ap-southeast-1", "dynamodb.ap-southeast-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Sydney)", "ap-southeast-2", "dynamodb.ap-southeast-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Jakarta)", "ap-southeast-3", "dynamodb.ap-southeast-3.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Asia Pacific (Melbourne)", "ap-southeast-4", "dynamodb.ap-southeast-4.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Canada (Central)", "ca-central-1", "dynamodb.ca-central-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Frankfurt)", "eu-central-1", "dynamodb.eu-central-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Zurich)", "eu-central-2", "dynamodb.eu-central-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Stockholm)", "eu-north-1", "dynamodb.eu-north-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Milan)", "eu-south-1", "dynamodb.eu-south-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Spain)", "eu-south-2", "dynamodb.eu-south-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Ireland)", "eu-west-1", "dynamodb.eu-west-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (London)", "eu-west-2", "dynamodb.eu-west-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Europe (Paris)", "eu-west-3", "dynamodb.eu-west-3.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Middle East (UAE)", "me-central-1", "dynamodb.me-central-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "Middle East (Bahrain)", "me-south-1", "dynamodb.me-south-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "South America (Sao Paulo)", "sa-east-1", "dynamodb.sa-east-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "US East (N. Virginia)", "us-east-1", "dynamodb.us-east-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "US East (Ohio)", "us-east-2", "dynamodb.us-east-2.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "US West (N. California)", "us-west-1", "dynamodb.us-west-1.amazonaws.com", RequestTypeTCP),
		NewRegionTarget(AWSProvider, "US West (Oregon)", "us-west-2", "dynamodb.us-west-2.amazonaws.com", RequestTypeTCP),
	)
}
