package main

import (
	"github.com/modelcontextprotocol-servers-api/mcp-server/config"
	"github.com/modelcontextprotocol-servers-api/mcp-server/models"
	tools_mcp "github.com/modelcontextprotocol-servers-api/mcp-server/tools/mcp"
	tools_message "github.com/modelcontextprotocol-servers-api/mcp-server/tools/message"
	tools_sse "github.com/modelcontextprotocol-servers-api/mcp-server/tools/sse"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_mcp.CreateDelete_mcpTool(cfg),
		tools_mcp.CreateGet_mcpTool(cfg),
		tools_mcp.CreatePost_mcpTool(cfg),
		tools_message.CreatePost_messageTool(cfg),
		tools_sse.CreateGet_sseTool(cfg),
	}
}
