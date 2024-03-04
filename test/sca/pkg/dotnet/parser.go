package dotnet

type Dependency struct {
	Name    string
	Version string
}

type Project struct {
	Dependencies []Dependency
}

//func Parse(path string) (*Project, error) {
//	data, err := os.ReadFile(path)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read file %q: %w", path, err)
//	}
//
//	raw := &rawProject{}
//
//	if err := xml.Unmarshal(data, raw); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal: %w", err)
//	}
//
//	return projectFromRaw(raw), nil
//}

type rawProject struct {
	ItemGroup []struct {
		PackageReference []struct {
			Include string `xml:"Include,attr"`
			Version string `xml:"Version,attr"`
		} `xml:"PackageReference"`
	} `xml:"ItemGroup"`
}

//func projectFromRaw(raw *rawProject) *Project {
//	p := &Project{}
//	for _, ig := range raw.ItemGroup {
//		for _, pr := range ig.PackageReference {
//			p.Dependencies = append(p.Dependencies, Dependency{Name: pr.Include, Version: pr.Version})
//		}
//	}
//
//	return p
//}
