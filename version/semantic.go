package version

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (s *SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
}

func (s *SemanticVersion) IncrementMajor() {
	s.major += 1
}

func (s *SemanticVersion) IncrementMinor() {
	s.minor += 1
}

func (s *SemanticVersion) IncrementPatch() {
	s.patch += 1
}
