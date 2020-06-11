package iotmakerCloudAws

type AwsStorageRegionsData struct {
	Name      string
	Region    string
	Continent string
	Endpoint  []string
	Protocol  []string
}

type AwsStorageRegion int

func (el AwsStorageRegion) String() string {
	return awsStorageRegions[el].Region
}

func (el AwsStorageRegion) Get() AwsStorageRegionsData {
	return awsStorageRegions[el]
}

const (
	KAwsStorageRegionsUsEast2 AwsStorageRegion = iota
	KAwsStorageRegionsUsEast1
	KAwsStorageRegionsUsWest1
	KAwsStorageRegionsUsWest2
	KAwsStorageRegionsAfSouth1
	KAwsStorageRegionsApEast1
	KAwsStorageRegionsApSouth1
	KAwsStorageRegionsApNortheast3
	KAwsStorageRegionsApNortheast2
	KAwsStorageRegionsApSoutheast1
	KAwsStorageRegionsApSoutheast2
	KAwsStorageRegionsApNortheast1
	KAwsStorageRegionsCaCentral1
	KAwsStorageRegionsCnNorth1
	KAwsStorageRegionsCnNorthwest1
	KAwsStorageRegionsEuCentral1
	KAwsStorageRegionsEuWest1
	KAwsStorageRegionsEuWest2
	KAwsStorageRegionsEuSouth1
	KAwsStorageRegionsEuWest3
	KAwsStorageRegionsEuNorth1
	KAwsStorageRegionsMeSouth1
	KAwsStorageRegionsSaEast1
	KAwsStorageRegionsUsGovEast1
	KAwsStorageRegionsUsGovWest1
)

var awsStorageRegions = [...]AwsStorageRegionsData{
	{
		Name:      "Leste dos EUA (Ohio)",
		Region:    "us-east-2",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-east-2.amazonaws.com",
			"rds-fips.us-east-2.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "Leste dos EUA (Norte da Virgínia)",
		Region:    "us-east-1",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-east-1.amazonaws.com",
			"rds-fips.us-east-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "Oeste dos EUA (Norte da Califórnia)",
		Region:    "us-west-1",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-west-1.amazonaws.com",
			"rds-fips.us-west-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "Oeste dos EUA (Oregon)",
		Region:    "us-west-2",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-west-2.amazonaws.com",
			"rds-fips.us-west-2.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "África (Cidade do Cabo)",
		Region:    "af-south-1",
		Continent: "Africa",
		Endpoint: []string{
			"rds.af-south-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Hong Kong)",
		Region:    "ap-east-1",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-east-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Mumbai)",
		Region:    "ap-south-1",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-south-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Osaka – Local)",
		Region:    "ap-northeast-3",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-northeast-3.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Seul)",
		Region:    "ap-northeast-2",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-northeast-2.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Cingapura)",
		Region:    "ap-southeast-1",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-southeast-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Sydney)",
		Region:    "ap-southeast-2",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-southeast-2.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Ásia-Pacífico (Tóquio)",
		Region:    "ap-northeast-1",
		Continent: "Asia",
		Endpoint: []string{
			"rds.ap-northeast-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Canadá (Central)",
		Region:    "ca-central-1",
		Continent: "Canada",
		Endpoint: []string{
			"rds.ca-central-1.amazonaws.com",
			"rds-fips.ca-central-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "China (Pequim)",
		Region:    "cn-north-1",
		Continent: "China",
		Endpoint: []string{
			"rds.cn-north-1.amazonaws.com.cn",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "China (Ningxia)",
		Region:    "cn-northwest-1",
		Continent: "China",
		Endpoint: []string{
			"rds.cn-northwest-1.amazonaws.com.cn",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Frankfurt)",
		Region:    "eu-central-1",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-central-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Irlanda)",
		Region:    "eu-west-1",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-west-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Londres)",
		Region:    "eu-west-2",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-west-2.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Milão)",
		Region:    "eu-south-1",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-south-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Paris)",
		Region:    "eu-west-3",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-west-3.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Europa (Estocolmo)",
		Region:    "eu-north-1",
		Continent: "Europa",
		Endpoint: []string{
			"rds.eu-north-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "Oriente Médio (Bahrein)",
		Region:    "me-south-1",
		Continent: "Asia",
		Endpoint: []string{
			"rds.me-south-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "América do Sul (São Paulo)",
		Region:    "sa-east-1",
		Continent: "South America",
		Endpoint: []string{
			"rds.sa-east-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
		},
	},
	{
		Name:      "AWS GovCloud (Leste dos EUA)",
		Region:    "us-gov-east-1",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-gov-east-1.amazonaws.com",
			"rds.us-gov-east-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
	{
		Name:      "AWS GovCloud (EU)",
		Region:    "us-gov-west-1",
		Continent: "EUA",
		Endpoint: []string{
			"rds.us-gov-west-1.amazonaws.com",
			"rds.us-gov-west-1.amazonaws.com",
		},
		Protocol: []string{
			"https",
			"https",
		},
	},
}
