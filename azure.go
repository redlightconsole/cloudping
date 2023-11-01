package cloudping

const AzureProvider = "azure"

func init() {
	// From https://github.com/richorama/AzureSpeedTest2/blob/master/lib/locations.js
	AddTarget(
		NewRegionTarget(AzureProvider, "West US 3", "westus3", "azurespeedtestwestus3.z1.web.core.windows.net"),
		NewRegionTarget(AzureProvider, "Qatar Central", "qc", "speedtestqc.z1.web.core.windows.net"),
		NewRegionTarget(AzureProvider, "Italy North", "itn", "speedtestitn.z38.web.core.windows.net"),
		NewRegionTarget(AzureProvider, "South Africa North", "san", "speedtestsan.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West Europe", "we", "speedtestwe.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Southeast Asia", "sea", "speedtestsea.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "East Asia", "ea", "speedtestea.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "North Central US", "nsus", "speedtestnsus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "North Europe", "ne", "speedtestne.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "South Central US", "scus", "speedtestscus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West US", "wus", "speedtestwus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "East US", "eus", "speedtesteus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Japan East", "jpe", "speedtestjpe.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Japan West", "jpw", "speedtestjpw.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Central US", "cus", "speedtestcus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "East US 2", "eus2", "speedtesteus2.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Australia Southeast", "ozse", "speedtestozse.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Australia East", "oze", "speedtestoze.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West UK", "ukw", "speedtestukw.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "South UK", "uks", "speedtestuks.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Canada Central", "cac", "speedtestcac.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Canada East", "cae", "speedtestcae.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West US 2", "westus2", "speedtestwestus2.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West India", "westindia", "speedtestwestindia.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "South India", "eastindia", "speedtesteastindia.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Central India", "centralindia", "speedtestcentralindia.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Korea Central", "koreacentral", "speedtestkoreacentral.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Korea South", "koreasouth", "speedtestkoreasouth.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "West Central US", "westcentralus", "speedtestwestcentralus.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "France Central", "frc", "speedtestfrc.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "UAE North", "uaen", "speedtestuaen.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Germany North", "den", "speedtestden.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Switzerland North", "chn", "speedtestchn.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Switzerland West", "chw", "speedtestchw.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Norway East", "east", "azspeednoeast.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Brazil", "nea", "speedtestnea.blob.core.windows.net"),
		NewRegionTarget(AzureProvider, "Sweden Central", "esc", "speedtestesc.blob.core.windows.net"),
	)
}
