package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tangerinefrog/HopView/internal/models"
	"tangerinefrog/HopView/internal/network"

	"github.com/gin-gonic/gin"
)

func tracerouteHandler(c *gin.Context) {
	target := c.Query("target")
	if target == "" {
		c.String(http.StatusBadRequest, "Missing 'target' query param")
		return
	}

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.String(http.StatusInternalServerError, "Streaming unsupported")
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "text/event-stream")
	c.Header("Connection", "keep-alive")

	nodesChan := make(chan *models.Node)
	ctx := c.Request.Context()

	go network.TraceRoute(ctx, target, nodesChan)
	for {
		select {
		case <-ctx.Done():
			fmt.Fprint(c.Writer, "event: done\ndata: {}\n\n")
			flusher.Flush()
			return
		case node, ok := <-nodesChan:
			if !ok {
				fmt.Fprint(c.Writer, "event: done\ndata:{}\n\n")
				flusher.Flush()
				return
			}
			nodeBytes, err := json.Marshal(node)
			if err != nil {
				c.String(http.StatusInternalServerError, "Error parsing node to JSON format")
				return
			}

			_, err = fmt.Fprintf(c.Writer, "data: %s\n\n", string(nodeBytes))
			if err != nil {
				c.String(http.StatusInternalServerError, "Error writing node data to buffer")
				return
			}

			flusher.Flush()
		}
	}
}
