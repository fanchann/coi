package parser

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/fanchann/coi/internal/types"
)

func ExtractInterfaceFunctions(filePath string) (*types.Interfaces, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(content)

	packageNameRegex := regexp.MustCompile(`^package\s+(\w+)`)
	packageNameMatches := packageNameRegex.FindStringSubmatch(contentStr)
	if len(packageNameMatches) < 2 || packageNameMatches[1] == "" {
		return nil, types.ErrPackageNameNotFound
	}
	packageName := packageNameMatches[1]

	newContentStr := extractCOIComment(contentStr)
	// tag `// coi` not found
	if newContentStr == "" {
		return nil, types.ErrTagCoiNotFound
	}

	interfaceNameRegex := regexp.MustCompile(`type\s+(\w+)\s+interface`)
	interfaceNameMatches := interfaceNameRegex.FindStringSubmatch(newContentStr)
	if len(interfaceNameMatches) < 2 {
		return nil, types.ErrInterfaceNotFound
	}
	interfaceName := interfaceNameMatches[1]

	result := &types.Interfaces{
		PackageName: packageName,
		IName:       interfaceName,
		IMethod:     parseInterfaceFunctions(newContentStr),
	}

	return result, nil
}

func parseInterfaceFunctions(str string) []string {
	var functions []string

	re := regexp.MustCompile(`(?:\n|^)\t*(\w+\(.*?\))(?:\s+\w+(?:\s*\.\s*\w+)*|\[\w+\])*(?:,|$)`)
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "type") {
			continue
		}
		if strings.HasPrefix(line, "}") {
			break
		}
		if len(line) > 0 {
			matches := re.FindAllString(line, -1)
			for _, match := range matches {
				functions = append(functions, match)
			}
		}
	}
	return functions
}

func extractCOIComment(s string) string {
	coiRegex := regexp.MustCompile(`(?s)// coi.*`)
	match := coiRegex.FindString(s)
	if match != "" {
		return match
	}
	return ""
}
