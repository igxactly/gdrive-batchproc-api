package appmain

import (
	"controllers"
	"github.com/astaxie/beegae"
)

func init() {
	beegae.Router("/jobqueue", &controllers.JobqueueController{})
	beegae.Run()
}
