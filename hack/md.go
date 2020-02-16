package hack

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/go-logr/logr"
	"gitlab.com/golang-commonmark/markdown"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/kubectl/pkg/scheme"
)

func GetDeployment(logger logr.Logger, url string) (*appsv1.Deployment, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	md := markdown.New(markdown.XHTMLOutput(true))
	tokens := md.Parse(buf.Bytes())

	var deployment *appsv1.Deployment

	decode := scheme.Codecs.UniversalDeserializer().Decode

	for _, t := range tokens {
		switch block := t.(type) {
		case *markdown.Fence:
			deploymentYaml := block.Content
			obj, _, err := decode([]byte(deploymentYaml), nil, nil)
			if err != nil {
				logger.Info("Unmarshall of fence failed")
				continue
			}
			deployment = obj.(*appsv1.Deployment)
		case *markdown.CodeBlock:
			deploymentYaml := block.Content
			obj, _, err := decode([]byte(deploymentYaml), nil, nil)
			if err != nil {
				logger.Info("Unmarshall of code block failed")
				continue
			}
			deployment = obj.(*appsv1.Deployment)
		case *markdown.CodeInline:
			deploymentYaml := block.Content
			obj, _, err := decode([]byte(deploymentYaml), nil, nil)
			if err != nil {
				logger.Info("Unmarshall of code inline failed")
				continue
			}
			deployment = obj.(*appsv1.Deployment)
		default:
			continue
		}
	}
	if deployment == nil {
		return nil, errors.New("didn't find code block")
	}

	return deployment, nil
}
