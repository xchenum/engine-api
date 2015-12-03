package lib

import "net/url"

// ContainerKill terminates the container process but does not remove the container from the docker host.
func (cli *Client) ContainerKill(containerID, signal string) error {
	var query url.Values
	query.Set("signal", signal)

	_, err := cli.POST("/containers/"+containerID+"/kill", query, nil, nil)
	return err
}
