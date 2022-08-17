package email

import (
	"bytes"
	"embed"
	"html/template"
	"notifications/errors"
	"os"
	"time"
)

// Generic info for all notification methods
type EmailHtmlTemplate struct {
	PageTitle         string
	HeaderClass       string
	HeaderMsg         string
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

//go:embed template/email_template.html
var templateFile embed.FS

// Load template information fr
func (et *EmailHtmlTemplate) setEmailTemplate() error {
	var present bool

	et.PageTitle = "Tekton Pipeline Notification"

	if et.GitRepo, present = os.LookupEnv("GIT_REPO"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_REPO"}
	}

	if et.GitBranch, present = os.LookupEnv("GIT_BRANCH"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_BRANCH"}
	}

	if et.GitCommit, present = os.LookupEnv("GIT_COMMIT"); !present {
		return &errors.EnvVarNotDefined{Name: "GIT_COMMIT"}
	}

	if et.PusherName, present = os.LookupEnv("PUSHER_NAME"); !present {
		return &errors.EnvVarNotDefined{Name: "PUSHER_NAME"}
	}

	if et.PusherEmail, present = os.LookupEnv("PUSHER_EMAIL"); !present {
		return &errors.EnvVarNotDefined{Name: "PUSHER_EMAIL"}
	}

	if et.PipelinerunName, present = os.LookupEnv("PIPELINERUN_NAME"); !present {
		return &errors.EnvVarNotDefined{Name: "PIPELINERUN_NAME"}
	}

	if et.PipelinerunStatus, present = os.LookupEnv("PIPELINERUN_STATUS"); !present {
		return &errors.EnvVarNotDefined{Name: "PIPELINERUN_STATUS"}
	}

	if et.Namespace, present = os.LookupEnv("NAMESPACE"); !present {
		return &errors.EnvVarNotDefined{Name: "NAMESPACE"}
	}

	// current date format Ex: "Mon, 15 Aug 2022 15:06:44 -03"
	et.CurrentDate = time.Now().Format(time.RFC1123)

	return nil
}

// set template with commom env variables
func (et *EmailHtmlTemplate) ParseEmailTemplate(tplBuffer *bytes.Buffer) error {

	// set Email template info
	err := et.setEmailTemplate()
	if err != nil {
		return err
	}

	if et.PipelinerunStatus == "Completed" || et.PipelinerunStatus == "Succeeded" {
		et.HeaderClass = "success-header"
		et.HeaderMsg = "Success!"
	} else {
		et.HeaderClass = "failure-header"
		et.HeaderMsg = "Failure!"
	}

	tmpl := template.Must(template.ParseFS(templateFile, "template/email_template.html"))

	err = tmpl.Execute(tplBuffer, et)
	if err != nil {
		return err
	}

	return nil
}
