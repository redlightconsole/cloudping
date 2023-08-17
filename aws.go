package cloudping

const AWSProvider = "aws"

func init() {
	AddTarget(
		NewRegionTarget(AWSProvider, "Africa (Cape Town)", "af-south-1", "dynamodb.af-south-1.amazonaws.com"),
	)
}
