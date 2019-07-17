package repo_link

import (
	"fmt"

	"gopkg.in/errgo.v1"

	"github.com/Scalingo/cli/config"
)

func ManualDeploy(app, branch string) error {
	if app == "" {
		return errgo.New("no app defined")
	}

	c, err := config.ScalingoClient()
	if err != nil {
		return errgo.Notef(err, "fail to get Scalingo client")
	}

	err = c.ScmRepoLinkManualDeploy(app, branch)
	if err != nil {
		return errgo.Mask(err)
	}

	fmt.Printf("Manual deployment triggered for app '%s' on branch '%s'.\n", app, branch)
	return nil
}
