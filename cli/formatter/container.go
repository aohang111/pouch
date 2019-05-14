package formatter

import (
	"github.com/alibaba/pouch/apis/types"
	"github.com/alibaba/pouch/pkg/utils"
)

//add default header for pouch ps output
type ContainerHeader map[string]string
type ContainerContext map[string]string

func NewContainerHeader() ContainerHeader {
	containerHeader := ContainerHeader{
		"Name":    "Name",
		"ID":      "ID",
		"Status":  "Status",
		"Created": "Created",
		"Image":   "Image",
		"Runtime": "Runtime",
		"Command": "Command",
		"ImageID": "ImageID",
		//"Labels":  "Labels",
		//"Mounts":  "Mounts",
		"State": "State",
	}
	return containerHeader
}
func NewContainerContext(c *types.Container) ContainerContext {
	id := c.ID[:6]
	created, _ := utils.FormatTimeInterval(c.Created)
	containerContext := ContainerContext{
		"Name":    c.Names[0],
		"ID":      id,
		"Status":  c.Status,
		"Created": created + " ago",
		"Image":   c.Image,
		"Runtime": c.HostConfig.Runtime,
		"Command": c.Command,
		"ImageID": c.ImageID,
		//"Labels":  "Labels",
		//"Mounts":  "Mounts",
		"State": c.State,
	}
	return containerContext
}
