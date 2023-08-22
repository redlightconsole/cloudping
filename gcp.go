package cloudping

const GCPProvider = "gcp"

func init() {
	// From https://github.com/GoogleCloudPlatform/gcping/blob/main/internal/config/endpoints.go#L69
	AddTarget(
		NewRegionTarget(GCPProvider, "Taiwan", "asia-east1", "asia-east1-5tkroniexa-de.a.run.app"),
		NewRegionTarget(GCPProvider, "Hong Kong", "asia-east2", "asia-east2-5tkroniexa-df.a.run.app"),
		NewRegionTarget(GCPProvider, "Tokyo", "asia-northeast1", "asia-northeast1-5tkroniexa-an.a.run.app"),
		NewRegionTarget(GCPProvider, "Osaka", "asia-northeast2", "asia-northeast2-5tkroniexa-dt.a.run.app"),
		NewRegionTarget(GCPProvider, "Seoul", "asia-northeast3", "asia-northeast3-5tkroniexa-du.a.run.app"),
		NewRegionTarget(GCPProvider, "Mumbai", "asia-south1", "asia-south1-5tkroniexa-el.a.run.app"),
		NewRegionTarget(GCPProvider, "Delhi", "asia-south2", "asia-south2-5tkroniexa-em.a.run.app"),
		NewRegionTarget(GCPProvider, "Singapore", "asia-southeast1", "asia-southeast1-5tkroniexa-as.a.run.app"),
		NewRegionTarget(GCPProvider, "Jakarta", "asia-southeast2", "asia-southeast2-5tkroniexa-et.a.run.app"),
		NewRegionTarget(GCPProvider, "Sydney", "australia-southeast1", "australia-southeast1-5tkroniexa-ts.a.run.app"),
		NewRegionTarget(GCPProvider, "Melbourne", "australia-southeast2", "australia-southeast2-5tkroniexa-km.a.run.app"),
		NewRegionTarget(GCPProvider, "Warsaw", "europe-central2", "europe-central2-5tkroniexa-lm.a.run.app"),
		NewRegionTarget(GCPProvider, "Finland", "europe-north1", "europe-north1-5tkroniexa-lz.a.run.app"),
		NewRegionTarget(GCPProvider, "Belgium", "europe-west1", "europe-west1-5tkroniexa-ew.a.run.app"),
		NewRegionTarget(GCPProvider, "London", "europe-west2", "europe-west2-5tkroniexa-nw.a.run.app"),
		NewRegionTarget(GCPProvider, "Frankfurt", "europe-west3", "europe-west3-5tkroniexa-ey.a.run.app"),
		NewRegionTarget(GCPProvider, "Netherlands", "europe-west4", "europe-west4-5tkroniexa-ez.a.run.app"),
		NewRegionTarget(GCPProvider, "Zurich", "europe-west6", "europe-west6-5tkroniexa-oa.a.run.app"),
		NewRegionTarget(GCPProvider, "Milan", "europe-west8", "europe-west8-5tkroniexa-oc.a.run.app"),
		NewRegionTarget(GCPProvider, "Paris", "europe-west9", "europe-west9-5tkroniexa-od.a.run.app"),
		NewRegionTarget(GCPProvider, "Berlin", "europe-west10", "europe-west10-5tkroniexa-oe.a.run.app"),
		NewRegionTarget(GCPProvider, "Madrid", "europe-southwest1", "europe-southwest1-5tkroniexa-no.a.run.app"),
		NewRegionTarget(GCPProvider, "Tel Aviv", "me-west1", "me-west1-5tkroniexa-zf.a.run.app"),
		NewRegionTarget(GCPProvider, "Montréal", "northamerica-northeast1", "northamerica-northeast1-5tkroniexa-nn.a.run.app"),
		NewRegionTarget(GCPProvider, "Toronto", "northamerica-northeast2", "northamerica-northeast2-5tkroniexa-pd.a.run.app"),
		NewRegionTarget(GCPProvider, "São Paulo", "southamerica-east1", "southamerica-east1-5tkroniexa-rj.a.run.app"),
		NewRegionTarget(GCPProvider, "Santiago", "southamerica-west1", "southamerica-west1-5tkroniexa-tl.a.run.app"),
		NewRegionTarget(GCPProvider, "Iowa", "us-central1", "us-central1-5tkroniexa-uc.a.run.app"),
		NewRegionTarget(GCPProvider, "South Carolina", "us-east1", "us-east1-5tkroniexa-ue.a.run.app"),
		NewRegionTarget(GCPProvider, "North Virginia", "us-east4", "us-east4-5tkroniexa-uk.a.run.app"),
		NewRegionTarget(GCPProvider, "Columbus", "us-east5", "us-east5-5tkroniexa-ul.a.run.app"),
		NewRegionTarget(GCPProvider, "Dallas", "us-south1", "us-south1-5tkroniexa-vp.a.run.app"),
		NewRegionTarget(GCPProvider, "Oregon", "us-west1", "us-west1-5tkroniexa-uw.a.run.app"),
		NewRegionTarget(GCPProvider, "Los Angeles", "us-west2", "us-west2-5tkroniexa-wl.a.run.app"),
		NewRegionTarget(GCPProvider, "Salt Lake City", "us-west3", "us-west3-5tkroniexa-wm.a.run.app"),
		NewRegionTarget(GCPProvider, "Las Vegas", "us-west4", "us-west4-5tkroniexa-wn.a.run.app"),
	)
}
