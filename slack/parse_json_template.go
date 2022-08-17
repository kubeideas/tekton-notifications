package slack

import (
	"bytes"
	"embed"
	"html/template"
	"notifications/errors"
	"os"
	"time"
)

// Generic info for all notification methods
type JsonMsgTemplate struct {
	Color             string
	GitRepo           string
	GitBranch         string
	GitCommit         string
	PusherName        string
	PusherEmail       string
	PipelinerunName   string
	PipelinerunStatus string
	Namespace         string
	CurrentDate       string
}

//go:embed template/slack_template.json
var templateFile embed.FS

// Load template information
func (j *JsonMsgTemplate) setEmailTemplate() error {
	var present bool

	if j.GitRepo, present = os.LookupEnv("GIT_REPO"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_REPO"}
	}

	if j.GitBranch, present = os.LookupEnv("GIT_BRANCH"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_BRANCH"}
	}

	if j.GitCommit, present = os.LookupEnv("GIT_COMMIT"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_COMMIT"}
	}

	if j.PusherName, present = os.LookupEnv("PUSHER_NAME"); !present {
		return &errors.EnvVarNotDefined{Name: "PUSHER_NAME"}
	}

	if j.PusherEmail, present = os.LookupEnv("PUSHER_EMAIL"); !present {
		return &errors.EnvVarNotDefined{Name: "PUSHER_EMAIL"}
	}

	if j.PipelinerunName, present = os.LookupEnv("PIPELINERUN_NAME"); !present {
		return &errors.EnvVarNotDefined{Name: "PIPELINERUN_NAME"}
	}

	if j.PipelinerunStatus, present = os.LookupEnv("PIPELINERUN_STATUS"); !present {
		return &errors.EnvVarNotDefined{Name: "PIPELINERUN_STATUS"}
	}

	if j.Namespace, present = os.LookupEnv("NAMESPACE"); !present {
		return &errors.EnvVarNotDefined{Name: "NAMESPACE"}
	}

	// current date format Ex: "Mon, 15 Aug 2022 15:06:44 -03"
	j.CurrentDate = time.Now().Format(time.RFC1123)

	return nil
}

// set template with commom env variables
func (j *JsonMsgTemplate) ParseJsonMsgTemplate(tplBuffer *bytes.Buffer) error {

	// set Email template info
	err := j.setEmailTemplate()
	if err != nil {
		return err
	}

	if j.PipelinerunStatus == "Completed" || j.PipelinerunStatus == "Succeeded" {
		j.Color = "good"
	} else {
		j.Color = "danger"
	}

	tmpl := template.Must(template.ParseFS(templateFile, "template/slack_template.json"))
	err = tmpl.Execute(tplBuffer, j)
	if err != nil {
		return err
	}

	return nil
}
