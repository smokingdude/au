package pkg

type NpmLsResult struct {
	Version              string                     `json:"version"`
	Name                 string                     `json:"name"`
	ExtendedDependencies map[string]NpmLsDependency `json:"dependencies"`
}

type NpmLsDependency struct {
	Version    string `json:"version"`
	Resolved   string `json:"resolved"`
	Overridden bool   `json:"overridden"`
	Name       string `json:"name"`
	Integrity  string `json:"integrity"`
	Dev        bool   `json:"dev"`
	License    string `json:"license"`
	//Engines    map[string]string `json:"engines"` // FIXME can be array of strings or map[string]string
	//Funding    []struct { // FIXME can be an array or a single object
	//	Type string `json:"type"`
	//	Url  string `json:"url"`
	//} `json:"funding"`
	Id                   string                     `json:"_id"`
	Extraneous           bool                       `json:"extraneous"`
	Path                 string                     `json:"path"`
	Dependencies         map[string]string          `json:"_dependencies"`
	DevDependencies      map[string]string          `json:"devDependencies"`
	PeerDependencies     map[string]string          `json:"peerDependencies"`
	ExtendedDependencies map[string]NpmLsDependency `json:"dependencies"`
}
